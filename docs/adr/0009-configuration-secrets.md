# 0009 Configuration And Secrets

## Status

Accepted

## Context

The frontend must never contain secrets and the backend should be configured by
environment.

## Decision

Use environment variables with `MPS_` prefixes and document placeholders in
`.env.example`. No secrets are required for v1.

## Consequences

Deployment is portable across local, Docker Compose, and hosted servers.

## Alternatives Considered

Config files were rejected for deployment secrets because they are easier to
commit accidentally.

