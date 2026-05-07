# 0003 Frontend Tooling

## Status

Accepted

## Context

The frontend needs a polished upload, job progress, and artifact download
experience with strong typing.

## Decision

Use React, TypeScript strict mode, Vite, Tailwind CSS, Zod, TanStack Query,
Vitest, and Playwright.

## Consequences

The app builds quickly to static files in `docs/` and can be served by GitHub
Pages.

## Alternatives Considered

Plain TypeScript was rejected because the job UI benefits from component state.
Next.js was rejected because static Pages publishing is simpler with Vite.

