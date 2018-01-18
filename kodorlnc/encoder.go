package kodorlnc

// Copyright Steinwurf ApS 2018.
// Distributed under the "STEINWURF RESEARCH LICENSE 1.0".
// See accompanying file LICENSE.rst or
// http://www.steinwurf.com/licensing

/*
#include "kodo_rlnc_c.h"
*/
import "C"

// Encoder is used for encoding data
type Encoder struct {
	mEncoder C.kodo_rlnc_encoder_t
}

// EncoderDestruct deallocates and release the memory consumed by an encoder
// @param encoder The encoder which should be deallocated
func EncoderDestruct(encoder *Encoder) {
	C.kodo_rlnc_encoder_destruct(encoder.mEncoder)
}

// EncoderPayloadSize returns the payload size of an encoder, which is the
// size of a generated payload.
// @param encoder The encoder to query.
// @return The payload size in bytes
func EncoderPayloadSize(encoder *Encoder) uint32 {
	return uint32(C.kodo_rlnc_encoder_payload_size(encoder.mEncoder))
}

// EncoderWritePayload writes a symbol into the provided payload buffer.
// @param encoder The encoder to use.
// @param payload The buffer which should contain the symbol.
// @return The total bytes used from the payload buffer
func EncoderWritePayload(encoder *Encoder, payload *uint8) uint32 {
	return uint32(C.kodo_rlnc_encoder_write_payload(encoder.mEncoder, (*C.uint8_t)(payload)))
}

// EncoderBlockSize returns the block size of an encoder.
// @param encoder The encoder to query.
// @return The block size, i.e. the total size in bytes that this encoder
//         operates on.
func EncoderBlockSize(encoder *Encoder) uint32 {
	return uint32(C.kodo_rlnc_encoder_block_size(encoder.mEncoder))
}

// EncoderSymbolSize returns the symbol size of the encoder.
// @param encoder The encoder to check
// @return The size of a symbol in bytes
func EncoderSymbolSize(encoder *Encoder) uint32 {
	return uint32(C.kodo_rlnc_encoder_symbol_size(encoder.mEncoder))
}

// EncoderSymbols returns the number of symbols in a block (i.e. the generation size).
// @param encoder The encoder to check
// @return The number of symbols
func EncoderSymbols(encoder *Encoder) uint32 {
	return uint32(C.kodo_rlnc_encoder_symbols(encoder.mEncoder))
}

// EncoderSetConstSymbols specifies the source data for all symbols.
// This will specify all symbols.
// @param encoder The encoder which will encode the data
// @param data The buffer containing the data to be encoded
// @param size The size of the buffer to be encoded
func EncoderSetConstSymbols(encoder *Encoder, data *uint8, size uint32) {
	C.kodo_rlnc_encoder_set_const_symbols(encoder.mEncoder, (*C.uint8_t)(data), size)
}

// IsSystematicOn returns whether the encoder is in the systematic mode, i.e.
// if it will initially send the original source symbols with a simple header.
// @param encoder The encoder
// @return Non-zero if the encoder is in the systematic mode, otherwise 0
func IsSystematicOn(encoder *Encoder) bool {
	return C.kodo_rlnc_is_systematic_on(encoder.mEncoder) != 0
}

// SetSystematicOn switches the systematic encoding on
// @param encoder The encoder
func SetSystematicOn(encoder *Encoder) {
	C.kodo_rlnc_set_systematic_on(encoder.mEncoder)
}

// EncoderSetSystematicOff switches the systematic encoding off
// @param encoder The encoder
func EncoderSetSystematicOff(encoder *Encoder) {
	C.kodo_rlnc_encoder_set_systematic_off(encoder.mEncoder)
}
