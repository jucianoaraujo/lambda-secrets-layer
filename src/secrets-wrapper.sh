#!/bin/bash

# 1. Execute the Go binary and evaluate its export string to set local environment variables
eval $(/opt/retrieve-secrets)

# 2. Shift the arguments and execute the native Lambda runtime bootstrap command
exec "$@"
