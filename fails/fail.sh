#!/bin/sh
set -euo pipefail

if [[ $# < 1 ]]; then
    echo "Usage: $0 password" >&2
    exit 1
fi
