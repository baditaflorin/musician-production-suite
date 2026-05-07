# 0011 Logging

## Status

Accepted

## Context

The backend needs structured logs and the frontend should avoid noisy production
console output.

## Decision

Use Go `slog` JSON logs on stdout in the backend. The frontend logs only
developer diagnostics in development.

## Consequences

Container logs are machine-readable and compatible with common collectors.

## Alternatives Considered

Text logs were rejected because production parsing is weaker.
