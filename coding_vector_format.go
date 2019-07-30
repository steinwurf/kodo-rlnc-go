package kodorlnc

// Copyright Steinwurf ApS 2018.
// Distributed under the "STEINWURF RESEARCH LICENSE 1.0".
// See accompanying file LICENSE.rst or
// http://www.steinwurf.com/licensing

/*
#cgo CFLAGS: -I../kodo-rlnc-c
#include <common.h>
*/
import "C"

/// Enum specifying the available coding vector formats fields
const (
	FullVector = C.krlnc_full_vector
	Seed       = C.krlnc_seed
	SparseSeed = C.krlnc_sparse_seed
)
