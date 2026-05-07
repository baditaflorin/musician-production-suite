package api

import (
	"encoding/json"
	"errors"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/baditaflorin/musician-production-suite/internal/domain"
	"github.com/baditaflorin/musician-production-suite/internal/jobs"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

type Handler struct {
	deps Dependencies
}

func NewHandler(deps Dependencies) *Handler {
	return &Handler{deps: deps}
}

func (h *Handler) Health(w http.ResponseWriter, _ *http.Request) {
	writeJSON(w, http.StatusOK, map[string]string{"status": "ok", "version": h.deps.Version})
}

func (h *Handler) Ready(w http.ResponseWriter, _ *http.Request) {
	if _, err := os.Stat(h.deps.Config.StorageDir); err != nil {
		writeError(w, http.StatusServiceUnavailable, "storage unavailable")
		return
	}
	writeJSON(w, http.StatusOK, map[string]string{"status": "ready"})
}

func (h *Handler) CreateJob(w http.ResponseWriter, r *http.Request) {
	r.Body = http.MaxBytesReader(w, r.Body, h.deps.Config.MaxUploadMB*1024*1024)
	if err := r.ParseMultipartForm(h.deps.Config.MaxUploadMB * 1024 * 1024); err != nil {
		writeError(w, http.StatusBadRequest, "invalid upload")
		return
	}
	file, header, err := r.FormFile("audio")
	if err != nil {
		writeError(w, http.StatusBadRequest, "missing audio field")
		return
	}
	defer file.Close()

	id := uuid.NewString()
	workDir := filepath.Join(h.deps.Store.Root(), id)
	if err := os.MkdirAll(workDir, 0o755); err != nil {
		writeError(w, http.StatusInternalServerError, "create job directory failed")
		return
	}
	inputPath := filepath.Join(workDir, "input"+filepath.Ext(header.Filename))
	dst, err := os.Create(inputPath)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "create input failed")
		return
	}
	defer dst.Close()
	if _, err := dst.ReadFrom(file); err != nil {
		writeError(w, http.StatusInternalServerError, "save input failed")
		return
	}

	now := time.Now().UTC()
	job := &domain.Job{
		ID:        id,
		Filename:  header.Filename,
		Status:    domain.JobQueued,
		Progress:  0,
		Message:   "Queued",
		CreatedAt: now,
		UpdatedAt: now,
		Chords:    []domain.ChordSegment{},
		Artifacts: []domain.Artifact{},
		InputPath: inputPath,
		WorkDir:   workDir,
	}
	if err := h.deps.Store.Create(job); err != nil {
		writeError(w, http.StatusInternalServerError, "create job failed")
		return
	}
	jobs.RecordCreated()
	h.deps.Runner.Start(job)
	writeJSON(w, http.StatusCreated, job)
}

func (h *Handler) GetJob(w http.ResponseWriter, r *http.Request) {
	job, err := h.deps.Store.Get(chi.URLParam(r, "id"))
	if err != nil {
		if errors.Is(err, jobs.ErrNotFound) {
			writeError(w, http.StatusNotFound, "job not found")
			return
		}
		writeError(w, http.StatusInternalServerError, "get job failed")
		return
	}
	writeJSON(w, http.StatusOK, job)
}

func (h *Handler) GetArtifact(w http.ResponseWriter, r *http.Request) {
	path, err := h.deps.Store.ArtifactPath(chi.URLParam(r, "id"), chi.URLParam(r, "*"))
	if err != nil {
		writeError(w, http.StatusNotFound, "artifact not found")
		return
	}
	http.ServeFile(w, r, path)
}

func writeJSON(w http.ResponseWriter, status int, value any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(value)
}

func writeError(w http.ResponseWriter, status int, message string) {
	writeJSON(w, status, map[string]string{"error": message})
}
