# 0006 WASM Modules

## Status

Accepted

## Context

Some audio and notation libraries have WASM builds, but v1 uses a runtime
backend.

## Decision

Do not ship WASM modules in v1. Lazy client-side previews may be added later.

## Consequences

The initial payload remains small and GitHub Pages needs no COOP/COEP
workaround.

## Alternatives Considered

Running Essentia or Verovio in the browser was deferred because it would not
remove the need for backend Demucs, LilyPond, FFmpeg, SoX, and VST hosting.

