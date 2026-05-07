# Contributing

## Local Setup

```sh
make install-hooks
make dev
```

## Commit Style

Use Conventional Commits:

```text
feat: add audio upload
fix: handle missing ffmpeg
docs: document deployment
```

## Checks

Run the same checks as the hooks:

```sh
make fmt
make lint
make test
make build
make smoke
```

