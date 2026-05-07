# 0005 Client Storage

## Status

Accepted

## Context

The frontend needs lightweight persistence for API base URL, recent jobs, and
user preferences.

## Decision

Use `localStorage` for preferences and recent job IDs. Do not store uploaded
audio in the browser in v1.

## Consequences

State is simple and local to the device. Cross-device sync is out of scope.

## Alternatives Considered

IndexedDB and OPFS were considered for local audio caches but rejected for v1
because processing happens backend-side.

