# 0008 Go Backend Layout

## Status

Accepted

## Context

The backend must be maintainable, testable, and easy to containerize.

## Decision

Use `cmd/server`, `internal/api`, `internal/config`, `internal/jobs`,
`internal/domain`, `internal/utils`, `pkg/audio`, `api`, `configs`, `scripts`,
and `test`.

## Consequences

Internal packages remain private to the backend and the public package boundary
is small.

## Alternatives Considered

A single package was rejected because process orchestration and HTTP concerns
would mix.
