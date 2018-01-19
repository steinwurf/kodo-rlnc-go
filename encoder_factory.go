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
func NewEncoderFactory(
	finiteField int32, symbols uint32, symbolSize uint32) *EncoderFactory {
	factory := new(EncoderFactory)
	factory.mFactory = C.kodo_rlnc_encoder_factory_construct(
		C.int32_t(finiteField), C.uint32_t(symbols), C.uint32_t(symbolSize))
	return factory
}

// Destruct deallocates the memory consumed by a factory
// @param factory The factory which should be deallocated
func (factory *EncoderFactory) Destruct() {
	C.kodo_rlnc_encoder_factory_destruct(factory.mFactory)
}

// Symbols returns the number of symbols in a block
// @param factory The factory to query
// @return the number of symbols in a block
func (factory *EncoderFactory) Symbols() uint32 {
	return uint32(C.kodo_rlnc_encoder_factory_symbols(factory.mFactory))
}

// SymbolSize returns the symbol size in bytes
// @param factory The factory to query
// @return the symbol size in bytes
func (factory *EncoderFactory) SymbolSize() uint32 {
	return uint32(C.kodo_rlnc_encoder_factory_symbol_size(factory.mFactory))
}

// SetSymbols sets the number of symbols
// @param factory The factory which should be configured
// @param symbols the number of symbols
func (factory *EncoderFactory) SetSymbols(symbols uint32) {
	C.kodo_rlnc_encoder_factory_set_symbols(
		factory.mFactory, C.uint32_t(symbols))
}

// SetSymbolSize sets the symbol size
// @param factory The factory which should be configured
// @param the symbol size in bytes
func (factory *EncoderFactory) SetSymbolSize(symbolSize uint32) {
	C.kodo_rlnc_encoder_factory_set_symbol_size(
		factory.mFactory, C.uint32_t(symbolSize))
}

// Build builds the actual encoder
// @param factory The encoder factory which should be used to build the encoder
// @return pointer to an instantiation of an encoder
func (factory *EncoderFactory) Build() *Encoder {
	encoder := new(Encoder)
	encoder.mEncoder = C.kodo_rlnc_encoder_factory_build(factory.mFactory)
	return encoder
}