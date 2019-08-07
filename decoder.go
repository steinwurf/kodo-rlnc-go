package kodorlnc

// Copyright Steinwurf ApS 2018.
// Distributed under the "STEINWURF RESEARCH LICENSE 1.0".
// See accompanying file LICENSE.rst or
// http://www.steinwurf.com/licensing

/*
#cgo CFLAGS: -I../kodo-rlnc-c
#cgo LDFLAGS: -L../kodo-rlnc-c -lkodo_rlnc_c_static -lkodo_rlnc -lfifi -lcpuid
#include <decoder.h>
*/
import "C"
import "runtime"

// Decoder is used for encoding data
type Decoder struct {
	mDecoder C.krlnc_decoder_t
}

// NewDecoder Creates a new decoder object.
// @param finiteField The finite field that should be used.
// @param symbols The number of symbols in a coding block
// @param symbolSize The size of a symbol in bytes
// @return Pointer to a new decoder instance.
func NewDecoder(finiteField int32, symbols uint32, symbolSize uint32) *Decoder {
	decoder := new(Decoder)
	decoder.mDecoder = C.krlnc_create_decoder(
		C.int32_t(finiteField), C.uint32_t(symbols), C.uint32_t(symbolSize))
	runtime.SetFinalizer(decoder, deleteDecoder)
	return decoder
}

// deleteDecoder deallocates and release the memory consumed by an decoder
// @param decoder The decoder which should be deallocated
func deleteDecoder(decoder *Decoder) {
	C.krlnc_delete_decoder(decoder.mDecoder)
}

// Reset resets the decoder and ensure that the object is in a clean state.
// @param decoder The decoder which should be reset
func (decoder *Decoder) Reset() {
	C.krlnc_reset_decoder(decoder.mDecoder)
}

// Symbols returns the number of symbols in a block (i.e. the generation size).
// @param decoder The decoder to check
// @return The number of symbols
func (decoder *Decoder) Symbols() uint32 {
	return uint32(C.krlnc_decoder_symbols(decoder.mDecoder))
}

// SymbolSize returns the symbol size of the decoder.
// @param decoder The decoder to check
// @return The size of a symbol in bytes
func (decoder *Decoder) SymbolSize() uint32 {
	return uint32(C.krlnc_decoder_symbol_size(decoder.mDecoder))
}

// BlockSize returns the block size of an decoder.
// @param decoder The decoder to query.
// @return The block size, i.e. the total size in bytes that this decoder
//         operates on.
func (decoder *Decoder) BlockSize() uint32 {
	return uint32(C.krlnc_decoder_block_size(decoder.mDecoder))
}

// SetSymbolStorage specifies the data buffer where the decoder should store a
// given symbol.
// @param decoder The decoder which will decode the symbol
// @param data The buffer that should contain the decoded symbol
// @param index The index of the symbol in the coding block
func (decoder *Decoder) SetSymbolStorage(data *[]uint8, index uint32) {
	C.krlnc_decoder_set_symbol_storage(
		decoder.mDecoder, (*C.uint8_t)(&(*data)[0]), C.uint32_t(index))
}

// SetSymbolsStorage specifies the data buffer where the decoder should store
// the decoded symbols.
// This will specify the storage for all symbols.
// @param decoder The decoder which will decode the data
// @param data The buffer that should contain the decoded symbols
func (decoder *Decoder) SetSymbolsStorage(data *[]uint8) {
	C.krlnc_decoder_set_symbols_storage(
		decoder.mDecoder, (*C.uint8_t)(&(*data)[0]))
}

// MaxPayloadSize returns the maximum payload size of an decoder,
// which is the expected size of an incoming payload.
// @param decoder The decoder to query.
// @return The payload size in bytes
func (decoder *Decoder) MaxPayloadSize() uint32 {
	return uint32(C.krlnc_decoder_max_payload_size(decoder.mDecoder))
}

// ConsumePayload consumes an encoded symbol stored in the payload buffer
// @param decoder The decoder to use.
// @param payload The buffer storing the payload of an encoded symbol.
//        The payload buffer may be changed by this operation,
//        so it cannot be reused. If the payload is needed at several places,
//        make sure to keep a copy of the original payload.
func (decoder *Decoder) ConsumePayload(payload *[]uint8) {
	C.krlnc_decoder_consume_payload(
		decoder.mDecoder, (*C.uint8_t)(&(*payload)[0]))
}

// ProducePayload produces a recoded symbol in the provided payload buffer.
// @param decoder The decoder to use.
// @param payload The buffer which should contain the recoded symbol.
// @return The total bytes used from the payload buffer
func (decoder *Decoder) ProducePayload(payload *[]uint8) uint32 {
	return uint32(C.krlnc_decoder_produce_payload(
		decoder.mDecoder, (*C.uint8_t)(&(*payload)[0])))
}

// IsComplete checks whether decoding is complete.
// @param decoder The decoder to query
// @return true if the decoding is complete, otherwise false
func (decoder *Decoder) IsComplete() bool {
	return C.krlnc_decoder_is_complete(decoder.mDecoder) != 0
}

// IsPartiallyComplete checks whether decoding is partially complete.
// This means that some symbols in the decoder are fully decoded.
// You can use the IsSymbolDecoded() function to determine which symbols.
// @param decoder The decoder to query
// @return true if the decoding is partially complete, otherwise false
func (decoder *Decoder) IsPartiallyComplete() bool {
	return C.krlnc_decoder_is_partially_complete(decoder.mDecoder) != 0
}

// Rank returns the rank of a decoder indicates how many symbols have
// been decoded or partially decoded.
// @param decoder The decoder to query
// @return The rank of the decoder
func (decoder *Decoder) Rank() uint32 {
	return uint32(C.krlnc_decoder_rank(decoder.mDecoder))
}

// SymbolsDecoded returns the number of decoded (i.e. fully decoded) symbols.
// @param decoder The decoder to query
// @return The number of decoded symbols at the decoder
func (decoder *Decoder) SymbolsDecoded() uint32 {
	return uint32(C.krlnc_decoder_symbols_decoded(decoder.mDecoder))
}

// IsSymbolMissing indicates whether a symbol is missing at a decoder.
// @param decoder The decoder to query
// @param index Index of the symbol whose state should be checked
// @return true if the symbol is missing, otherwise false
func (decoder *Decoder) IsSymbolMissing(index uint32) bool {
	return C.krlnc_decoder_is_symbol_missing(
		decoder.mDecoder, C.uint32_t(index)) != 0
}

// IsSymbolPartiallyDecoded indicates whether a symbol has been partially
// decoded at a decoder.
// @param decoder The decoder to query
// @param index Index of the symbol whose state should be checked
// @return true if the symbol has been partially decoded, otherwise false
func (decoder *Decoder) IsSymbolPartiallyDecoded(index uint32) bool {
	return C.krlnc_decoder_is_symbol_partially_decoded(
		decoder.mDecoder, C.uint32_t(index)) != 0
}

// IsSymbolDecoded indicates whether a symbol is available in an decoded
// (i.e. fully decoded)
// form at the decoder.
// @param decoder The decoder to query
// @param index Index of the symbol whose state should be checked
// @return true if the symbol is decoded, otherwise false
func (decoder *Decoder) IsSymbolDecoded(index uint32) bool {
	return C.krlnc_decoder_is_symbol_decoded(
		decoder.mDecoder, C.uint32_t(index)) != 0
}

// IsSymbolPivot indicates if a symbol is partially or fully decoded.
// A symbol with a pivot element is defined in the coding matrix of a decoder.
// @param decoder The decoder to query
// @param index Index of the symbol whose state should be checked
// @return true if the symbol is defined, otherwise false
func (decoder *Decoder) IsSymbolPivot(index uint32) bool {
	return C.krlnc_decoder_is_symbol_pivot(
		decoder.mDecoder, C.uint32_t(index)) != 0
}

// IsStatusUpdaterEnabled returns whether the symbol status updater is enabled
// or not.
// The status updater can be used to accurately track the status of each
// symbol during the decoding process (this can impact the performance).
// The default state is OFF.
// @param decoder The decoder to query
// @return true if the symbol status updater is enabled, otherwise false
func (decoder *Decoder) IsStatusUpdaterEnabled() bool {
	return C.krlnc_decoder_is_status_updater_enabled(decoder.mDecoder) != 0
}

// SetStatusUpdaterOn enable the status updater so that a full update is
// performed every time a symbol is read.
// @param decoder The decoder to modify
func (decoder *Decoder) SetStatusUpdaterOn() {
	C.krlnc_decoder_set_status_updater_on(decoder.mDecoder)
}

// SetStatusUpdaterOff disable the status updater.
// @param decoder The decoder to modify
func (decoder *Decoder) SetStatusUpdaterOff() {
	C.krlnc_decoder_set_status_updater_off(decoder.mDecoder)
}

// UpdateSymbolStatus force a manual update on the symbol status so that all
// symbols that are currently considered partially decoded will labelled as
// decoded if their coding vector only has a single non-zero coefficient
// (which is 1).
// @param decoder The decoder to update
func (decoder *Decoder) UpdateSymbolStatus() {
	C.krlnc_decoder_update_symbol_status(decoder.mDecoder)
}

// CoefficientVectorSize returns the size of the coefficient vector.
// @param decoder The decoder to check
// @return The size of the coefficient vector in bytes
func (decoder *Decoder) CoefficientVectorSize() uint32 {
	return uint32(C.krlnc_decoder_coefficient_vector_size(decoder.mDecoder))
}

// ConsumeSymbol reads and decodes an encoded symbol according to the provided
// coding coefficients.
// @param decoder The decoder to use.
// @param symbolData The encoded symbol
// @param coefficients The coding coefficients that were used to calculate the
//                     encoded symbol
func (decoder *Decoder) ConsumeSymbol(
	symbolData *[]uint8, coefficients *[]uint8) {
	C.krlnc_decoder_consume_symbol(
		decoder.mDecoder,
		(*C.uint8_t)(&(*symbolData)[0]),
		(*C.uint8_t)(&(*coefficients)[0]))
}

// ConsumeSystematicSymbol Reads and decodes a systematic/decoded symbol with
// the corresponding symbol index.
// @param decoder The decoder to use.
// @param symbolData The systematic source symbol.
// @param index The index of this decoded symbol in the data block.
func (decoder *Decoder) ConsumeSystematicSymbol(
	symbolData *[]uint8, index uint32) {
	C.krlnc_decoder_consume_systematic_symbol(
		decoder.mDecoder,
		(*C.uint8_t)(&(*symbolData)[0]),
		C.uint32_t(index))
}

// SetSeed sets the seed of the coefficient generator.
// @param decoder The decoder to use
// @param seedValue The seed value for the generator.
func (decoder *Decoder) SetSeed(seedValue uint32) {
	C.krlnc_decoder_set_seed(decoder.mDecoder, C.uint32_t(seedValue))
}

// Generate fills the input buffer with symbol coefficients used for either
// encoding or decoding a symbol.
// @param decoder The decoder to use.
// @param coefficients Pointer to the memory where the coefficients should
//        be stored. The coefficient buffer should have at least
//        CoefficientVectorSize() capacity.
func (decoder *Decoder) Generate(coefficients *[]uint8) {
	C.krlnc_decoder_generate(
		decoder.mDecoder, (*C.uint8_t)(&(*coefficients)[0]))
}

// GeneratePartial generates a "partial" coding vector that will only contain
// non-zero coefficients for the symbols that are currently defined.
// This allows encoding before defining all original source symbols,
// i.e. on-the-fly encoding.
// @param decoder The decoder to use.
// @param coefficients Pointer to the memory where the coefficients should
//        be stored. The coefficient buffer should have at least
//        CoefficientVectorSize() capacity.
func (decoder *Decoder) GeneratePartial(coefficients *[]uint8) {
	C.krlnc_decoder_generate_partial(
		decoder.mDecoder, (*C.uint8_t)(&(*coefficients)[0]))
}

// SetLogStdout enables the default log function of the decoder, which prints to the standard output.
// @param decoder The decoder to use
func (decoder *Decoder) SetLogStdout() {
	C.krlnc_decoder_set_log_stdout(decoder.mDecoder)
}

// SetLogOff disables the log function of the decoder.
// @param decoder The decoder to use
func (decoder *Decoder) SetLogOff() {
	C.krlnc_decoder_set_log_off(decoder.mDecoder)
}
