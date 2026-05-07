# 0010 GitHub Pages Publishing

## Status

Accepted

## Context

The live Pages URL is a first-class deliverable and must work from the start.

## Decision

Publish from the `main` branch `docs/` directory. Vite emits the production
frontend directly to `docs/`, with hashed assets and a `404.html` SPA fallback.

## Consequences

`docs/` is intentionally committed and must not be gitignored. Build artifacts
for Pages live alongside documentation.

## Alternatives Considered

A `gh-pages` branch was rejected because it complicates local hooks and commit
visibility.
