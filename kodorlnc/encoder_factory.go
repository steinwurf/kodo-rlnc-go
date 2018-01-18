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

// EncoderFactory builds Encoders
type EncoderFactory struct {
	mFactory C.kodo_rlnc_encoder_factory_t
}

// NewEncoderFactory builds a new encoder factory
// @param finiteField The finite field that should be used by the encoder.
// @param symbols The maximum number of symbols supported by encoders
//        built with this factory.
// @param symbolSize The maximum symbol size in bytes supported by
//        encoders built using the returned factory
// @return A new factory capable of building encoders using the
//         selected parameters.
func NewEncoderFactory(finiteField int32, symbols uint32, symbolSize uint32) EncoderFactory {
	factory := EncoderFactory
	factory.mFactory = C.kodo_rlnc_encoder_factory_construct(
		finiteField, symbols, symbolSize)
	return factory
}

// EncoderFactoryDestruct deallocates and release the memory consumed by a factory
// @param factory The factory which should be deallocated
func EncoderFactoryDestruct(factory *EncoderFactory) {
	return C.kodo_rlnc_encoder_factory_destruct(factory.mFactory)
}

// EncoderFactorySymbols returns the number of symbols in a block
// @param factory The factory to query
// @return the number of symbols in a block
func EncoderFactorySymbols(factory *EncoderFactory) uint32 {
	return uint32(C.kodo_rlnc_encoder_factory_symbols(factory.mFactory))
}

// EncoderFactorySymbolSize returns the symbol size in bytes
// @param factory The factory to query
// @return the symbol size in bytes
func EncoderFactorySymbolSize(factory *EncoderFactory) uint32 {
	return uint32(C.kodo_rlnc_encoder_factory_symbolSize(factory.mFactory))
}

// EncoderFactorySetSymbols sets the number of symbols
// @param factory The factory which should be configured
// @param symbols the number of symbols
func EncoderFactorySetSymbols(factory *EncoderFactory, symbols uint32) {
	C.kodo_rlnc_encoder_factory_set_symbols(factory.mFactory, symbols)
}

// EncoderFactorySetSymbolSize sets the symbol size
// @param factory The factory which should be configured
// @param the symbol size in bytes
func EncoderFactorySetSymbolSize(factory *EncoderFactory, symbolSize uint32) {
	C.kodo_rlnc_encoder_factory_set_symbolSize(factory.mFactory, symbolSize)
}

// EncoderFactoryBuild builds the actual encoder
// @param factory The encoder factory which should be used to build the encoder
// @return pointer to an instantiation of an encoder
func EncoderFactoryBuild(factory *EncoderFactory) *Encoder {
	encoder := new(Encoder)
	encoder.m_encoder = C.kodo_rlnc_encoder_factory_build(factory.mFactory)
	return encoder
}
