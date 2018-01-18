package kodorlnc

// Copyright Steinwurf ApS 2018.
// Distributed under the "STEINWURF RESEARCH LICENSE 1.0".
// See accompanying file LICENSE.rst or
// http://www.steinwurf.com/licensing

/*
#cgo CFLAGS: -I../kodo-rlnc-c
#cgo LDFLAGS: -L../kodo-rlnc-c -lkodo_rlnc_c_static -lkodo_rlnc -lfifi -lcpuid
#include <kodo_rlnc_c.h>
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
func NewDecoderFactory(
	finiteField int32, symbols uint32, symbolSize uint32) *DecoderFactory {
	factory := new(DecoderFactory)
	factory.mFactory = C.kodo_rlnc_decoder_factory_construct(
		C.int32_t(finiteField), C.uint32_t(symbols), C.uint32_t(symbolSize))
	return factory
}

// Destruct deallocatesthe memory consumed by a factory
// @param factory The factory which should be deallocated
func (factory *DecoderFactory) Destruct() {
	C.kodo_rlnc_decoder_factory_destruct(factory.mFactory)
}

// Symbols returns the number of symbols in a block
// @param factory The factory to query
// @return the number of symbols in a block
func (factory *DecoderFactory) Symbols() uint32 {
	return uint32(C.kodo_rlnc_decoder_factory_symbols(factory.mFactory))
}

// SymbolSize returns the symbol size in bytes
// @param factory The factory to query
// @return the symbol size in bytes
func (factory *DecoderFactory) SymbolSize() uint32 {
	return uint32(C.kodo_rlnc_decoder_factory_symbol_size(factory.mFactory))
}

// SetSymbols sets the number of symbols
// @param factory The factory which should be configured
// @param symbols the number of symbols
func (factory *DecoderFactory) SetSymbols(symbols uint32) {
	C.kodo_rlnc_decoder_factory_set_symbols(
		factory.mFactory, C.uint32_t(symbols))
}

// SetSymbolSize sets the symbol size
// @param factory The factory which should be configured
// @param the symbol size in bytes
func (factory *DecoderFactory) SetSymbolSize(symbolSize uint32) {
	C.kodo_rlnc_decoder_factory_set_symbol_size(
		factory.mFactory, C.uint32_t(symbolSize))
}

// Build builds the actual decoder
// @param factory The decoder factory which should be used to build the decoder
// @return pointer to an instantiation of an decoder
func (factory *DecoderFactory) Build() *Decoder {
	decoder := new(Decoder)
	decoder.mDecoder = C.kodo_rlnc_decoder_factory_build(factory.mFactory)
	return decoder
}
