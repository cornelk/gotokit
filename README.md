## gotokit - An opinionated standard library for Golang microservices

[![Build status](https://github.com/cornelk/gotokit/actions/workflows/go.yaml/badge.svg?branch=main)](https://github.com/cornelk/gotokit/actions)
[![go.dev reference](https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white&style=flat-square)](https://pkg.go.dev/github.com/cornelk/gotokit)
[![Go Report Card](https://goreportcard.com/badge/github.com/cornelk/gotokit)](https://goreportcard.com/report/github.com/cornelk/gotokit)
[![codecov](https://codecov.io/gh/cornelk/gotokit/branch/main/graph/badge.svg?token=NS5UY28V3A)](https://codecov.io/gh/cornelk/gotokit)

## Project layout

    ├─ config           config file reading from environment variables
    ├─ database         PostgreSQL client and struct scanner
    ├─ env              test/dev/staging/prod environment defines
    ├─ envfile          loads environment variables from env files
    ├─ jsonutils        additional helpers for JSON processing 
    ├─ log              fast, structured, optimized logging based on zap
    ├─ multierror       provides appendable errors
