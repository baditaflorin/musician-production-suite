# 0015 Deployment Topology

## Status

Accepted

## Context

The frontend and backend deploy independently.

## Decision

Serve the frontend from GitHub Pages. Run the backend Docker container behind
nginx on host port `25342`, with optional Prometheus profile.

## Consequences

Static hosting remains simple. Backend operators must provision a Linux server
with enough CPU, RAM, and disk for audio jobs.

## Alternatives Considered

Serving the frontend from the Go binary was rejected by the Pages-first
constraint.

