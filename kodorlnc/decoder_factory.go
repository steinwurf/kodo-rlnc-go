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

// DecoderFactory builds Decoders
type DecoderFactory struct {
	mFactory C.kodo_rlnc_decoder_factory_t
}

// NewDecoderFactory builds a new decoder factory
// @param finiteField The finite field that should be used by the decoder.
// @param symbols The maximum number of symbols supported by decoders
//        built with this factory.
// @param symbolSize The maximum symbol size in bytes supported by
//        decoders built using the returned factory
// @return A new factory capable of building decoders using the
//         selected parameters.
func NewDecoderFactory(finiteField int32, symbols uint32, symbolSize uint32) DecoderFactory {
	factory := DecoderFactory
	factory.mFactory = C.kodo_rlnc_decoder_factory_construct(
		finiteField, symbols, symbolSize)
	return factory
}

// DecoderFactoryDestruct deallocates and release the memory consumed by a factory
// @param factory The factory which should be deallocated
func DecoderFactoryDestruct(factory *DecoderFactory) {
	C.kodo_rlnc_decoder_factory_destruct(factory.mFactory)
}

// DecoderFactorySymbols returns the number of symbols in a block
// @param factory The factory to query
// @return the number of symbols in a block
func DecoderFactorySymbols(factory *DecoderFactory) uint32 {
	return uint32(C.kodo_rlnc_decoder_factory_symbols(factory.mFactory))
}

// DecoderFactorySymbolSize returns the symbol size in bytes
// @param factory The factory to query
// @return the symbol size in bytes
func DecoderFactorySymbolSize(factory *DecoderFactory) uint32 {
	return uint32(C.kodo_rlnc_decoder_factory_symbolSize(factory.mFactory))
}

// DecoderFactorySetSymbols sets the number of symbols
// @param factory The factory which should be configured
// @param symbols the number of symbols
func DecoderFactorySetSymbols(factory *DecoderFactory, symbols uint32) {
	C.kodo_rlnc_decoder_factory_set_symbols(factory.mFactory, symbols)
}

// DecoderFactorySetSymbolSize sets the symbol size
// @param factory The factory which should be configured
// @param the symbol size in bytes
func DecoderFactorySetSymbolSize(factory *DecoderFactory, symbolSize uint32) {
	C.kodo_rlnc_decoder_factory_set_symbolSize(factory.mFactory, symbolSize)
}

// DecoderFactoryBuild builds the actual decoder
// @param factory The decoder factory which should be used to build the decoder
// @return pointer to an instantiation of an decoder
func DecoderFactoryBuild(factory *DecoderFactory) *Decoder {
	decoder := new(Decoder)
	decoder.m_decoder = C.kodo_rlnc_decoder_factory_build(factory.mFactory)
	return decoder
}
