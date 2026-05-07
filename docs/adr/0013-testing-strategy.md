# 0013 Testing Strategy

## Status

Accepted

## Context

The product needs confidence across frontend logic, backend handlers, and the
happy path.

## Decision

Use Vitest for frontend tests, Go unit tests for backend packages, optional
integration tests under `test/integration`, and a shell smoke test that builds,
serves, and verifies the app and API.

## Consequences

Checks remain local and compatible with git hooks.

## Alternatives Considered

GitHub Actions were rejected by project constraint.
