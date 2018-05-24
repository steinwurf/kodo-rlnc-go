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

// Decoder is used for encoding data
type Decoder struct {
	mDecoder *C.krlnc_decoder_t
}

// deleteDecoder deallocates and release the memory consumed by an decoder
// @param decoder The decoder which should be deallocated
func deleteDecoder(decoder *Decoder) {
	C.krlnc_delete_decoder(decoder.mDecoder)
}

// PayloadSize returns the payload size of an decoder, which is the
// size of a generated payload.
// @param decoder The decoder to query.
// @return The payload size in bytes
func (decoder *Decoder) PayloadSize() uint32 {
	return uint32(C.krlnc_decoder_payload_size(decoder.mDecoder))
}

// ReadPayload reads symbol from the given payload buffer.
// @param decoder The decoder to use.
// @param payload The buffer storing the payload of an encoded symbol.
//        The payload buffer may be changed by this operation,
//        so it cannot be reused. If the payload is needed at several places,
//        make sure to keep a copy of the original payload.
func (decoder *Decoder) ReadPayload(payload *[]uint8) {
	C.krlnc_decoder_read_payload(
		decoder.mDecoder, (*C.uint8_t)(&(*payload)[0]))
}

// BlockSize returns the block size of an decoder.
// @param decoder The decoder to query.
// @return The block size, i.e. the total size in bytes that this decoder
//         operates on.
func (decoder *Decoder) BlockSize() uint32 {
	return uint32(C.krlnc_decoder_block_size(decoder.mDecoder))
}

// SymbolSize returns the symbol size of the decoder.
// @param decoder The decoder to check
// @return The size of a symbol in bytes
func (decoder *Decoder) SymbolSize() uint32 {
	return uint32(C.krlnc_decoder_symbol_size(decoder.mDecoder))
}

// Symbols returns the number of symbols in a block.
// @param decoder The decoder to check
// @return The number of symbols
func (decoder *Decoder) Symbols() uint32 {
	return uint32(C.krlnc_decoder_symbols(decoder.mDecoder))
}

// SetMutableSymbols specifies the data buffer where the decoder
// should store the decoded symbols.
// This will specify the storage for all symbols.
// @param decoder The decoder which will decode the data
// @param data The buffer that should contain the decoded symbols
// @param size The size of the buffer to be decoded
func (decoder *Decoder) SetMutableSymbols(data *[]uint8) {
	C.krlnc_decoder_set_mutable_symbols(
		decoder.mDecoder, (*C.uint8_t)(&(*data)[0]), C.uint32_t(len(*data)))
}

// IsComplete checks whether decoding is complete.
// @param decoder The decoder to query
// @return true if the decoding is complete, otherwise false
func (decoder *Decoder) IsComplete() bool {
	return C.krlnc_decoder_is_complete(decoder.mDecoder) != 0
}

// Rank returns the rank of a decoder indicates how many symbols have
// been decoded or partially decoded.
// @param decoder The decoder to query
// @return The rank of the decoder
func (decoder *Decoder) Rank() uint32 {
	return uint32(C.krlnc_decoder_rank(decoder.mDecoder))
}
