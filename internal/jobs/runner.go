package jobs

import (
	"context"
	"log/slog"
	"time"

	"github.com/example/musician-production-suite/internal/domain"
	"github.com/example/musician-production-suite/pkg/audio"
)

type Runner struct {
	store    *Store
	pipeline *audio.Pipeline
	logger   *slog.Logger
}

func NewRunner(store *Store, pipeline *audio.Pipeline, logger *slog.Logger) *Runner {
	return &Runner{store: store, pipeline: pipeline, logger: logger}
}

func (r *Runner) Start(job *domain.Job) {
	go r.run(job.ID)
}

func (r *Runner) run(id string) {
	start := time.Now()
	job, err := r.store.Update(id, func(j *domain.Job) {
		j.Status = domain.JobRunning
		j.Progress = 3
		j.Message = "Job runner started"
	})
	if err != nil {
		r.logger.Error("job_start_failed", "job_id", id, "error", err)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Hour)
	defer cancel()
	err = r.pipeline.Process(ctx, job, func(progress int, message string, mutate func(*domain.Job)) {
		_, updateErr := r.store.Update(id, func(j *domain.Job) {
			j.Progress = progress
			j.Message = message
			if mutate != nil {
				mutate(j)
			}
		})
		if updateErr != nil {
			r.logger.Error("job_update_failed", "job_id", id, "error", updateErr)
		}
	})

	if err != nil {
		msg := err.Error()
		_, _ = r.store.Update(id, func(j *domain.Job) {
			j.Status = domain.JobFailed
			j.Error = &msg
			j.Message = "Processing failed"
		})
		jobsFailed.Inc()
		r.logger.Error("job_failed", "job_id", id, "duration", time.Since(start), "error", err)
		return
	}

	final, err := r.store.Update(id, func(j *domain.Job) {
		j.Status = domain.JobSucceeded
		j.Progress = 100
		j.Message = "Production suite artifacts are ready"
	})
	if err == nil {
		artifactsCreated.Add(float64(len(final.Artifacts)))
	}
	jobsCompleted.Inc()
	jobDuration.Observe(time.Since(start).Seconds())
	r.logger.Info("job_completed", "job_id", id, "duration", time.Since(start))
}
