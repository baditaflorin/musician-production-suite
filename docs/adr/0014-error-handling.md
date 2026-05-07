# 0014 Error Handling

## Status

Accepted

## Context

Audio processing can fail for many environmental reasons.

## Decision

Return structured JSON errors from the API, persist job errors in job state, and
use `internal/utils.HandleErrorOrLogWithMessages` for top-level command errors.

## Consequences

The frontend can show actionable failures and backend logs stay consistent.

## Alternatives Considered

Panics and opaque text errors were rejected.
