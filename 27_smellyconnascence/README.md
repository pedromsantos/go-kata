# Connascence Violations Kata

## Overview

This is a **verification fixture, not a practice exercise**. Each folder
contains a small, self-contained example of exactly one Connascence type,
translated directly from *Agile Technical Practices Distilled*'s
Connascence chapter worked examples. Its purpose is to give a known-answer
set for [jev-review](https://github.com/pedromsantos/jev-review)'s upcoming
Connascence rules to be verified against before those rules are implemented.

`CoV` (Connascence of Value) is intentionally not represented -- already
covered by jev-review's existing checks. `CoMT` (Connascence of Manual
Task) is intentionally not represented either -- it's connascence with an
external, undocumented manual step no code review can see.

This is the Go translation of the reference kata in ts-kata's
`26_SmellyConnascence`; equivalent kata exist for Java, Python, C#, and C.

## What's here

| File | Type | Why |
|---|---|---|
| `Position/notification_system.go` | Connascence of Position | three same-typed `string` parameters carry meaning only through argument order |
| `Meaning/transport_selector.go` | Connascence of Meaning | `"1"`/`"2"`/`"3"`/`"4"` mean bike/car/train/bus only by an unstated, shared convention |
| `Algorithm/checksum_calculator.go` | Connascence of Algorithm | the checksum computation (`sum % 10`) is duplicated across two methods instead of extracted once |
| `ExecutionOrder/receipt_sender.go` | Connascence of Execution Order | `Archive()` is only correct after `SendToCustomer()`, but nothing enforces that order |
| `Timing/background_job_runner.go` | Connascence of Timing | waits a fixed, arbitrary delay instead of the job's actual completion |
| `Identity/global_counter.go` + `Identity/counter_consumer.go` | Connascence of Identity | every consumer's correctness depends on sharing this exact package-level instance |

## Run

```sh
go build ./27_smellyconnascence/...
go vet ./27_smellyconnascence/...
```
