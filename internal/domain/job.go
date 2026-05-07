package domain

import "time"

type JobStatus string

const (
	JobQueued    JobStatus = "queued"
	JobRunning   JobStatus = "running"
	JobSucceeded JobStatus = "succeeded"
	JobFailed    JobStatus = "failed"
)

type ChordSegment struct {
	Start float64 `json:"start"`
	End   float64 `json:"end"`
	Chord string  `json:"chord"`
}

type Artifact struct {
	Name      string `json:"name"`
	Kind      string `json:"kind"`
	URL       string `json:"url"`
	SizeBytes int64  `json:"sizeBytes"`
	Path      string `json:"-"`
}

type Job struct {
	ID        string         `json:"id"`
	Filename  string         `json:"filename"`
	Status    JobStatus      `json:"status"`
	Progress  int            `json:"progress"`
	Message   string         `json:"message"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	BPM       *float64       `json:"bpm"`
	Key       *string        `json:"key"`
	Chords    []ChordSegment `json:"chords"`
	Artifacts []Artifact     `json:"artifacts"`
	Error     *string        `json:"error"`
	InputPath string         `json:"-"`
	WorkDir   string         `json:"-"`
}
