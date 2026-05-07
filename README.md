# musician-production-suite

https://baditaflorin.github.io/musician-production-suite/

Browser UI plus audio backend for stems, cleanup, BPM/key/chords, transcription,
MIDI, sheet music PDF, and mixdown.

## Quickstart

```sh
npm install
go mod download
make install-hooks
make dev
```

Frontend: http://localhost:5173

Backend: http://localhost:8080

## What It Does

Upload an audio file, run the production pipeline, and download generated
artifacts: analysis JSON, cleaned audio, stems, MIDI, MusicXML, PDF score, and a
mixdown. The backend uses native tools when installed and creates deterministic
fallback artifacts when optional engines are unavailable so the workflow remains
testable.

## Architecture

The frontend is built with TypeScript, React, Vite, and Tailwind CSS. It is
published from `docs/` for GitHub Pages. The backend is a Go API packaged as a
Docker image and intended to run separately behind nginx.

See:

- https://github.com/baditaflorin/musician-production-suite/blob/main/docs/architecture.md
- https://github.com/baditaflorin/musician-production-suite/blob/main/docs/adr/0001-deployment-mode.md
- https://github.com/baditaflorin/musician-production-suite/blob/main/deploy/README.md
