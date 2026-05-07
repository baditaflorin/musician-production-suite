#!/bin/sh
set -eu

VERSION="${VERSION:-v0.1.0}"
git diff --quiet
git tag "$VERSION"
printf 'Created tag %s\n' "$VERSION"

