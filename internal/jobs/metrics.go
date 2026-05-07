package jobs

import "github.com/prometheus/client_golang/prometheus"

var (
	jobsCreated = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "mps_jobs_created_total",
		Help: "Total production jobs created.",
	})
	jobsCompleted = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "mps_jobs_completed_total",
		Help: "Total production jobs completed.",
	})
	jobsFailed = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "mps_jobs_failed_total",
		Help: "Total production jobs failed.",
	})
	jobDuration = prometheus.NewHistogram(prometheus.HistogramOpts{
		Name:    "mps_job_duration_seconds",
		Help:    "Production job duration in seconds.",
		Buckets: prometheus.DefBuckets,
	})
	artifactsCreated = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "mps_artifacts_created_total",
		Help: "Total artifacts created.",
	})
)

func init() {
	prometheus.MustRegister(jobsCreated, jobsCompleted, jobsFailed, jobDuration, artifactsCreated)
}

func RecordCreated() {
	jobsCreated.Inc()
}
