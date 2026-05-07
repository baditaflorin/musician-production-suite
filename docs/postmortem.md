# Postmortem

## What Was Built

Mode C scaffold and v1 implementation for a GitHub Pages frontend plus Docker
backend audio processing API.

## Was Mode C Correct?

Yes. The requested suite depends on native tools, PyTorch models, long-running
jobs, and generated media artifacts. Mode A and Mode B cannot cover user-uploaded
runtime processing for v1.

## What Worked

The split keeps the public UI static while allowing production audio tools to
run server-side.

## What Did Not

The local environment may not have every external audio engine installed, so the
backend includes deterministic fallback artifacts for development and smoke
tests.

## Surprises

The breadth of the requested stack makes adapter boundaries more important than
any single engine choice.

## Accepted Tech Debt

VST hosting is represented as a pipeline stage placeholder until a specific
server-safe host is selected.

## Next Improvements

1. Add real Demucs and CREPE adapters with model cache management.
2. Add editable transcription review before score export.
3. Add authenticated job cleanup and retention controls.

## Time

Initial scaffold and implementation were completed as a compact v1 pass.
