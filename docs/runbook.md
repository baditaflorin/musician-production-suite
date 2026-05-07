# Runbook

## Health

```sh
curl http://localhost:8080/healthz
curl http://localhost:8080/readyz
```

## Logs

Docker logs:

```sh
docker compose -f deploy/docker-compose.yml logs -f app
```

## Metrics

Prometheus metrics are exposed at `/metrics` and should be blocked from public
internet access by nginx.

## Sizing

Minimum for fallback mode: 2 CPU, 2 GB RAM, 10 GB disk.

Recommended for Demucs/CREPE workloads: 8 CPU, 16 GB RAM, 100 GB disk, and a
GPU-capable variant when available.
