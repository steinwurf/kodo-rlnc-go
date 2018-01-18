package kodorlnc

// Copyright Steinwurf ApS 2018.
// Distributed under the "STEINWURF RESEARCH LICENSE 1.0".
// See accompanying file LICENSE.rst or
// http://www.steinwurf.com/licensing

/*
#cgo CFLAGS: -I../kodo_rlnc_c
#cgo LDFLAGS: -L../kodo_rlnc_c -lboost_system -lboost_chrono -lboost_timer -lboost_iostreams -lboost_filesystem -lcpuid -lboost_thread -lfifi -lkodo_rlnc -lkodo_core_nocode -lkodo_rlnc_c_static
#include "kodo_rlnc_c.h"
*/
import "C"

/// Enum specifying the available finite fields
const (
	Binary  = C.kodo_rlnc_binary
	Binary4 = C.kodo_rlnc_binary4
	Binary8 = C.kodo_rlnc_binary8
)
