#!/bin/bash
set -e

make migrate

exec "$@"