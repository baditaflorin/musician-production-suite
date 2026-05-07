# 0016 Local Git Hooks

## Status

Accepted

## Context

No GitHub Actions are allowed, so checks must run locally.

## Decision

Use plain `.githooks/` wired by `core.hooksPath` and Makefile wrapper targets.

## Consequences

Developers must run `make install-hooks` after cloning.

## Alternatives Considered

Lefthook was considered, but plain hooks reduce dependency surface.

