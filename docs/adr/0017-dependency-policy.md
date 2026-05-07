# 0017 Dependency Policy

## Status

Accepted

## Context

The system should use production-ready libraries and avoid custom engines.

## Decision

Use established libraries for routing, metrics, validation, frontend state, and
build tooling. The backend delegates media work to proven command-line tools
when installed.

## Consequences

The v1 implementation is reliable and replaceable at adapter boundaries.

## Alternatives Considered

Hand-rolled parsers, audio codecs, and ML models were rejected.

