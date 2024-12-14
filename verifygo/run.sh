#!/bin/bash

# Load environment variables from .env
source .env

# Run the Go program with the phone number argument
go run verify.go -phone="$1"
