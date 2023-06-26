#!/usr/bin/sh

GOOS=windows GOARCH=amd64 make
GOOS=windows GOARCH=386 make
GOOS=linux GOARCH=amd64 make
GOOS=linux GOARCH=386 make