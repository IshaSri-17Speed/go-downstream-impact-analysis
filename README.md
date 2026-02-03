# Go Downstream Impact Analysis Tool

This repository explores a Go-based tool for evaluating the real-world
downstream impact of changes made to Go libraries by running builds and
test suites of dependent projects.

The goal is to move beyond API-level compatibility checks
and understand whether a change actually breaks downstream consumers in
practice.

This work is inspired by similar tooling in other ecosystems (e.g. Rust)
and is being developed as part of a proposed [LFX Mentorship
(Term 01 – 2026) under the CNCF](https://mentorship.lfx.linuxfoundation.org/project/5f537fc2-548b-487a-99ed-c61f7e8bcd47) / [OpenTelemetry](https://opentelemetry.io/) umbrella.



## Problem Context

Go libraries such as [OpenTelemetry](https://github.com/open-telemetry/opentelemetry-go) form part of the dependency graph of a
large number of projects. Changes to these libraries
propagate transitively through module dependencies and can affect build
and test behavior in downstream projects.

Current change-evaluation approaches in the Go ecosystem primarily rely
on API-level analysis tools, such as [apidiff](https://pkg.go.dev/golang.org/x/exp/apidiff), which
detect changes to exported symbols and type signatures between versions.
While useful, these approaches are limited to interface-level comparison
and do not capture behavioral changes, indirect dependency interactions,
or assumptions encoded in downstream build pipelines and test suites.

As a result, API compatibility does not necessarily imply downstream
build or test compatibility. Evaluating the impact of a change therefore
requires empirical validation against downstream projects rather than
static interface comparison alone.

This project focuses on enabling such validation by executing and
comparing downstream build and test results across versions.



## Planned Capabilities

The tool is designed around three core capabilities:

1. **Dependency Discovery**  
   Identify downstream Go modules that depend on a given module using
   reliable data sources such as [deps.dev](https://docs.deps.dev/api/v3/) and [pkg.go.dev](https://pkg.go.dev/).

2. **Downstream Test Execution**  
   Safely clone and execute test suites of downstream projects for
   different versions of a module in isolated environments.

3. **Impact Comparison**  
   Compare results before and after a change and produce a summary that
   highlights newly introduced failures and unaffected dependents.



## Current Status

This repository currently contains:

- Initial problem analysis and design notes
- An early prototype for downstream dependency discovery using [deps.dev](https://deps.dev/)

The implementation will be incrementally expanded during the mentorship.



## Repository Structure

```text
├── cmd/            # CLI entry points 
├── docs/           # Design and problem understanding documents
├── internal/       # Internal packages 
├── main.go         # Prototype entry point
└── README.md

```

## Disclaimer

This is an early-stage prototype intended for experimentation and design
validation. APIs and behavior are expected to evolve.
