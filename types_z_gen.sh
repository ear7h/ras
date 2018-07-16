#!/bin/sh
go run types_gen/fixed/*.go > types_fixed.go
go run types_gen/arrays/*.go > types_arrays.go