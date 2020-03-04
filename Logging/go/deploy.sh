#!/usr/bin/env bash

rm -rf go.mod
go mod init gologging.go

gcloud app deploy app.yaml  --quiet

