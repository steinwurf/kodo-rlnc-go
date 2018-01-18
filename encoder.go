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

// Encoder is used for encoding data
type Encoder struct {
	mEncoder C.kodo_rlnc_encoder_t
}

// Destruct deallocates and release the memory consumed by an encoder
// @param encoder The encoder which should be deallocated
func (encoder *Encoder) Destruct() {
	C.kodo_rlnc_encoder_destruct(encoder.mEncoder)
}

// PayloadSize returns the payload size of an encoder, which is the
// size of a generated payload.
// @param encoder The encoder to query.
// @return The payload size in bytes
func (encoder *Encoder) PayloadSize() uint32 {
	return uint32(C.kodo_rlnc_encoder_payload_size(encoder.mEncoder))
}

// WritePayload writes a symbol into the provided payload buffer.
// @param encoder The encoder to use.
// @param payload The buffer which should contain the symbol.
// @return The total bytes used from the payload buffer
func (encoder *Encoder) WritePayload(payload *uint8) uint32 {
	return uint32(C.kodo_rlnc_encoder_write_payload(
		encoder.mEncoder, (*C.uint8_t)(payload)))
}

// BlockSize returns the block size of an encoder.
// @param encoder The encoder to query.
// @return The block size, i.e. the total size in bytes that this encoder
//         operates on.
func (encoder *Encoder) BlockSize() uint32 {
	return uint32(C.kodo_rlnc_encoder_block_size(encoder.mEncoder))
}

// SymbolSize returns the symbol size of the encoder.
// @param encoder The encoder to check
// @return The size of a symbol in bytes
func (encoder *Encoder) SymbolSize() uint32 {
	return uint32(C.kodo_rlnc_encoder_symbol_size(encoder.mEncoder))
}

// Symbols returns the number of symbols in a block.
// @param encoder The encoder to check
// @return The number of symbols
func (encoder *Encoder) Symbols() uint32 {
	return uint32(C.kodo_rlnc_encoder_symbols(encoder.mEncoder))
}

// SetConstSymbols specifies the source data for all symbols.
// This will specify all symbols.
// @param encoder The encoder which will encode the data
// @param data The buffer containing the data to be encoded
// @param size The size of the buffer to be encoded
func (encoder *Encoder) SetConstSymbols(data *uint8, size uint32) {
	C.kodo_rlnc_encoder_set_const_symbols(
		encoder.mEncoder, (*C.uint8_t)(data), C.uint32_t(size))
}

// IsSystematicOn returns whether the encoder is in the systematic mode, i.e.
// if it will initially send the original source symbols with a simple header.
// @param encoder The encoder
// @return Non-zero if the encoder is in the systematic mode, otherwise 0
func (encoder *Encoder) IsSystematicOn() bool {
	return C.kodo_rlnc_is_systematic_on(encoder.mEncoder) != 0
}

// SetSystematicOn switches the systematic encoding on
// @param encoder The encoder
func (encoder *Encoder) SetSystematicOn() {
	C.kodo_rlnc_set_systematic_on(encoder.mEncoder)
}

// SetSystematicOff switches the systematic encoding off
// @param encoder The encoder
func (encoder *Encoder) SetSystematicOff() {
	C.kodo_rlnc_encoder_set_systematic_off(encoder.mEncoder)
}
