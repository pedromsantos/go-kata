# SOLID Violations Kata

## Overview

This is a **verification fixture, not a practice exercise**. Each folder
contains a small, self-contained example of exactly one SOLID principle
violation, translated directly from *Agile Technical Practices Distilled*'s
SOLID chapter worked examples (Car/`Save`, `CarEngineStatusReportController`,
`Chef`/`Oven`/`Microwave`, `IAmACar`, `Kitchen`/`MicrowaveOven`). Its purpose
is to give static-analysis/AI code-review tooling (specifically
[jev-review](https://github.com/pedromsantos/jev-review)) a known-answer set
to check its SOLID rules against -- every file's violation is deliberate and
documented below, not hidden.

This is the Go translation of the reference kata in ts-kata's
`25_SmellySolid`; equivalent kata exist for Java, Python, C#, and C too.

Go has no classes or inheritance: the `Lsp` and `Isp` folders below use Go
interfaces and structs in place of the TypeScript reference's abstract base
class and subclasses -- an abstract class becomes an `interface`, and a
subclass overriding a method becomes a `struct` implementing that interface.

## What's here

| File | Violates | Why |
|---|---|---|
| `Srp/car.go` | SRP | `Save()` mixes a persistence concern into a struct otherwise about domain behaviour (mileage/travel) |
| `Ocp/*.go` | OCP (and DIP) | every new report format needs a new method on the controller, and it constructs its concrete views directly instead of receiving them injected |
| `Lsp/microwave.go` | LSP | `Cook()` returns an error instead of honouring the `Oven` interface's contract |
| `Lsp/chef.go` | -- | not itself a violation, but its type-assertion on `*Microwave` is the client-code tell of `Microwave`'s LSP violation |
| `Isp/i_am_a_car.go` | ISP | bundles `RefillGasoline`/`RefillElectricity`, capabilities no single car supports both of |
| `Isp/electric_car.go` | -- | the forced implementer: errors on the gasoline method it can't honestly support |
| `Dip/kitchen.go` | DIP (and OCP) | constructs `MicrowaveOven` directly; can't work with any other oven without being edited |
| `Dip/microwave_oven.go` | DIP | constructs `MicrowaveGenerator` directly instead of receiving it injected |

## Run

```sh
go build ./26_smellysolid/...
go vet ./26_smellysolid/...
```
