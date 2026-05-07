#!/bin/sh
set -eu

npm run build
go build -o bin/mps-server ./cmd/server

PORT="${PORT:-18080}"
MPS_API_ADDR=":${PORT}" MPS_STORAGE_DIR=./tmp/smoke-jobs ./bin/mps-server >/tmp/mps-smoke.log 2>&1 &
SERVER_PID=$!
trap 'kill ${SERVER_PID} >/dev/null 2>&1 || true' EXIT

tries=0
until curl -fsS "http://127.0.0.1:${PORT}/healthz" >/dev/null; do
  tries=$((tries + 1))
  if [ "$tries" -gt 30 ]; then
    cat /tmp/mps-smoke.log
    exit 1
  fi
  sleep 0.2
done

curl -fsS "http://127.0.0.1:${PORT}/readyz" >/dev/null
curl -fsS "http://127.0.0.1:${PORT}/metrics" >/dev/null
test -f docs/index.html
grep -q "Musician Production Suite" docs/index.html

printf 'smoke ok\n'

