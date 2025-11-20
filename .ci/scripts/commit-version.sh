#!/usr/bin/env bash
set -euo pipefail

git add app/app.go
git commit -m "chore: release version ${RELEASE_VERSION}"
git push origin HEAD:master
