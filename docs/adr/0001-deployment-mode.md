# 0001 Deployment Mode

## Status

Accepted

## Context

The product must run a complete music production pipeline: source separation,
denoising, analysis, pitch tracking, score generation, MIDI export, PDF
rendering, and mixdown. The named engines include native C/C++ tools, PyTorch
models, and command-line processors.

## Decision

Use Mode C: GitHub Pages frontend plus Docker backend.

The frontend is static and published from `docs/`. The backend is a Go API that
accepts uploads, manages jobs, invokes native audio tools, and serves generated
artifacts.

## Consequences

GitHub Pages remains the public UI surface. Heavy processing runs in Docker, not
in the browser. Deployments require both Pages and a server capable of running
the backend container.

## Alternatives Considered

Mode A was rejected because PyTorch, Demucs, LilyPond, FFmpeg, SoX, RNNoise, and
VST hosting are not practical as a complete browser-only v1. Mode B was rejected
because users upload their own audio and need runtime processing.

