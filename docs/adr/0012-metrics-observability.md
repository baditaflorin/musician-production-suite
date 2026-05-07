# 0012 Metrics And Observability

## Status

Accepted

## Context

Mode C requires operational visibility for job processing.

## Decision

Expose `/metrics` with Prometheus HTTP and domain metrics: created jobs,
completed jobs, failed jobs, job duration, and artifact count.

## Consequences

Prometheus can scrape the API container. Nginx blocks public access to metrics.

## Alternatives Considered

Client analytics are out of scope and disabled by default.
