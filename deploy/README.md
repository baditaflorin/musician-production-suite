# Backend Deployment

## Prerequisites

- Linux amd64 server
- Docker and Docker Compose
- DNS pointing to the server
- TLS certificates in `/etc/letsencrypt`

## First Deploy

```sh
cd deploy
cp ../.env.example .env
docker compose pull
docker compose up -d
```

The public backend is exposed through nginx on host port `25342`.

## TLS

Replace `example.com` in `deploy/nginx/nginx.conf` with the real host. Obtain
certificates with certbot on the host and mount `/etc/letsencrypt` read-only.

## Rollback

```sh
docker compose pull app
docker compose up -d app
```

Pin `app.image` to a previous semver tag for deterministic rollback.

## Logs

```sh
docker compose logs -f app
docker compose logs -f nginx
```

## Backups

Back up the `mps_jobs` Docker volume if job artifacts must be retained.
