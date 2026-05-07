# 0002 Architecture Overview

## Status

Accepted

## Context

The app needs a static user interface and an asynchronous processing backend.

## Decision

Use four boundaries: frontend, API, job runner, and audio tool adapters. The API
owns HTTP contracts. The job runner owns lifecycle and storage. Tool adapters own
process execution and artifact creation.

## Consequences

Audio engines can be added without changing handlers. The frontend talks only to
the API contract.

## Alternatives Considered

A monolithic script runner was rejected because it would be hard to test and
hard to expose safely over HTTP.
