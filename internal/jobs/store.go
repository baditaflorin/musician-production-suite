package jobs

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/baditaflorin/musician-production-suite/internal/domain"
)

var ErrNotFound = errors.New("job not found")

type Store struct {
	root string
	mu   sync.RWMutex
	jobs map[string]*domain.Job
}

func NewFileStore(root string) (*Store, error) {
	if err := os.MkdirAll(root, 0o755); err != nil {
		return nil, err
	}
	return &Store{root: root, jobs: make(map[string]*domain.Job)}, nil
}

func (s *Store) Root() string {
	return s.root
}

func (s *Store) Create(job *domain.Job) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.jobs[job.ID] = clone(job)
	return s.persist(job)
}

func (s *Store) Get(id string) (*domain.Job, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	job, ok := s.jobs[id]
	if !ok {
		return nil, ErrNotFound
	}
	return clone(job), nil
}

func (s *Store) Update(id string, mutate func(*domain.Job)) (*domain.Job, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	job, ok := s.jobs[id]
	if !ok {
		return nil, ErrNotFound
	}
	mutate(job)
	job.UpdatedAt = time.Now().UTC()
	if err := s.refreshArtifacts(job); err != nil {
		return nil, err
	}
	if err := s.persist(job); err != nil {
		return nil, err
	}
	return clone(job), nil
}

func (s *Store) ArtifactPath(id string, name string) (string, error) {
	job, err := s.Get(id)
	if err != nil {
		return "", err
	}
	clean := filepath.Clean(name)
	path := filepath.Join(job.WorkDir, clean)
	rel, err := filepath.Rel(job.WorkDir, path)
	if err != nil || rel == ".." || strings.HasPrefix(rel, ".."+string(os.PathSeparator)) || filepath.IsAbs(rel) {
		return "", ErrNotFound
	}
	if _, err := os.Stat(path); err != nil {
		return "", ErrNotFound
	}
	return path, nil
}

func (s *Store) refreshArtifacts(job *domain.Job) error {
	artifacts := make([]domain.Artifact, 0)
	err := filepath.WalkDir(job.WorkDir, func(path string, entry os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		base := filepath.Base(path)
		if entry.IsDir() || base == "job.json" || strings.HasPrefix(base, "input.") {
			return nil
		}
		info, err := entry.Info()
		if err != nil {
			return err
		}
		rel, err := filepath.Rel(job.WorkDir, path)
		if err != nil {
			return err
		}
		artifacts = append(artifacts, domain.Artifact{
			Name:      filepath.Base(path),
			Kind:      kindFor(path),
			URL:       "/api/v1/jobs/" + job.ID + "/artifacts/" + rel,
			SizeBytes: info.Size(),
			Path:      path,
		})
		return nil
	})
	if err != nil {
		return err
	}
	job.Artifacts = artifacts
	return nil
}

func (s *Store) persist(job *domain.Job) error {
	data, err := json.MarshalIndent(job, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(filepath.Join(job.WorkDir, "job.json"), data, 0o644)
}

func clone(job *domain.Job) *domain.Job {
	copied := *job
	copied.Chords = append([]domain.ChordSegment(nil), job.Chords...)
	copied.Artifacts = append([]domain.Artifact(nil), job.Artifacts...)
	return &copied
}

func kindFor(path string) string {
	switch filepath.Ext(path) {
	case ".json":
		return "analysis"
	case ".wav", ".mp3", ".flac", ".ogg", ".aiff":
		return "audio"
	case ".mid", ".midi":
		return "midi"
	case ".musicxml", ".xml":
		return "score"
	case ".pdf":
		return "sheet-music"
	default:
		return "artifact"
	}
}
