package api

import (
	"bytes"
	"io"
	"log/slog"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/baditaflorin/musician-production-suite/internal/config"
	"github.com/baditaflorin/musician-production-suite/internal/jobs"
	"github.com/baditaflorin/musician-production-suite/pkg/audio"
	"github.com/stretchr/testify/require"
)

func TestCreateJobRequiresAudio(t *testing.T) {
	t.Parallel()
	dir := t.TempDir()
	store, err := jobs.NewFileStore(dir)
	require.NoError(t, err)
	logger := slog.New(slog.NewTextHandler(io.Discard, nil))
	runner := jobs.NewRunner(store, audio.NewPipeline(logger), logger)
	router := NewRouter(Dependencies{
		Config: config.Config{
			APIAddr:            ":0",
			StorageDir:         dir,
			CORSAllowedOrigins: []string{"*"},
			MaxUploadMB:        10,
		},
		Store:  store,
		Runner: runner,
		Logger: logger,
	})

	req := httptest.NewRequest(http.MethodPost, "/api/v1/jobs", bytes.NewReader(nil))
	res := httptest.NewRecorder()
	router.ServeHTTP(res, req)

	require.Equal(t, http.StatusBadRequest, res.Code)
}

func TestCreateJobAcceptsAudio(t *testing.T) {
	t.Parallel()
	dir := t.TempDir()
	store, err := jobs.NewFileStore(dir)
	require.NoError(t, err)
	logger := slog.New(slog.NewTextHandler(io.Discard, nil))
	runner := jobs.NewRunner(store, audio.NewPipeline(logger), logger)
	router := NewRouter(Dependencies{
		Config: config.Config{
			APIAddr:            ":0",
			StorageDir:         dir,
			CORSAllowedOrigins: []string{"*"},
			MaxUploadMB:        10,
		},
		Store:  store,
		Runner: runner,
		Logger: logger,
	})

	var body bytes.Buffer
	writer := multipart.NewWriter(&body)
	part, err := writer.CreateFormFile("audio", "clip.wav")
	require.NoError(t, err)
	_, err = part.Write([]byte("audio"))
	require.NoError(t, err)
	require.NoError(t, writer.Close())

	req := httptest.NewRequest(http.MethodPost, "/api/v1/jobs", &body)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	res := httptest.NewRecorder()
	router.ServeHTTP(res, req)

	require.Equal(t, http.StatusCreated, res.Code)
	entries, err := os.ReadDir(dir)
	require.NoError(t, err)
	require.Len(t, entries, 1)
}
