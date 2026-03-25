#!/usr/bin/env bash

#===============================================================
# Description: Create a new solution and the necessary files
# Author: eidnoat
# Date: 2026-03-25
# Version: 0.1
#===============================================================
set -euo pipefail

if [[ $# -lt 1 ]]; then
  echo "err: params not enough" >&2
  echo "usage: $0 <question_title>" >&2
  exit 1
fi

question_title="${1//[ -]/_}"

if [[ -d "$question_title" ]]; then
  echo "err: directory '$question_title' already exists!" >&2
  exit 1
fi

mkdir "$question_title"
touch "$question_title"/{solution.go,README.md}

echo "✅ Solution workspace created: $question_title/"