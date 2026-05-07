# 0004 API Contract

## Status

Accepted

## Context

Mode C requires a stable contract between Pages and the Docker backend.

## Decision

Expose REST/JSON under `/api/v1` and document it in `api/openapi.yaml`.
Frontend configuration provides the API base URL at build time.

## Consequences

The backend can be deployed independently. The frontend must handle backend
unavailability and CORS failures clearly.

## Alternatives Considered

GraphQL was rejected as unnecessary for a small job/artifact API.
