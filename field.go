package kodorlnc

// Copyright Steinwurf ApS 2018.
// Distributed under the "STEINWURF RESEARCH LICENSE 1.0".
// See accompanying file LICENSE.rst or
// http://www.steinwurf.com/licensing

/*
#cgo CFLAGS: -I../kodo-rlnc-c
#include <kodo_rlnc_c.h>
*/
import "C"

/// Enum specifying the available finite fields
const (
	Binary   = C.krlnc_binary
	Binary4  = C.krlnc_binary4
	Binary8  = C.krlnc_binary8
	Binary16 = C.krlnc_binary16
)
