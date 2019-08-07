package kodorlnc

// Copyright Steinwurf ApS 2018.
// Distributed under the "STEINWURF RESEARCH LICENSE 1.0".
// See accompanying file LICENSE.rst or
// http://www.steinwurf.com/licensing

/*
#cgo CFLAGS: -I../kodo-rlnc-c
#cgo LDFLAGS: -L../kodo-rlnc-c -lkodo_rlnc_c_static -lkodo_rlnc -lfifi -lcpuid
#include <encoder.h>
*/
import "C"
import "runtime"

// Encoder is used for encoding data
type Encoder struct {
	mEncoder C.krlnc_encoder_t
}

// NewEncoder creates a new encoder object.
// @param finiteField The finite field that should be used.
// @param symbols The number of symbols in a coding block
// @param symbolSize The size of a symbol in bytes
// @return Pointer to a new encoder instance.
func NewEncoder(
	finiteField int32, symbols uint32, symbolSize uint32) *Encoder {
	encoder := new(Encoder)
	encoder.mEncoder = C.krlnc_create_encoder(
		C.int32_t(finiteField), C.uint32_t(symbols), C.uint32_t(symbolSize))

	runtime.SetFinalizer(encoder, deleteEncoder)
	return encoder
}

// deleteEncoder deallocates and release the memory consumed by an encoder
// @param encoder The encoder which should be deallocated
func deleteEncoder(encoder *Encoder) {
	C.krlnc_delete_encoder(encoder.mEncoder)
}

// Reset resets the encoder and ensure that the object is in a clean
// state.
// @param encoder The encoder which should be reset
func (encoder *Encoder) Reset() {
	C.krlnc_reset_encoder(encoder.mEncoder)
}

// SetCodingVectorFormat sets the coding vector format
// @param encoder The encoder which should be configured
// @param formatId The selected coding vector format
func (encoder *Encoder) SetCodingVectorFormat(formatID int32) {
	C.krlnc_encoder_set_coding_vector_format(encoder.mEncoder, C.int32_t(formatID))
}

// SymbolSize returns the symbol size of the encoder.
// @param encoder The encoder to check
// @return The size of a symbol in bytes
func (encoder *Encoder) SymbolSize() uint32 {
	return uint32(C.krlnc_encoder_symbol_size(encoder.mEncoder))
}

// Symbols returns the number of symbols in a block.
// @param encoder The encoder to check
// @return The number of symbols
func (encoder *Encoder) Symbols() uint32 {
	return uint32(C.krlnc_encoder_symbols(encoder.mEncoder))
}

// BlockSize returns the block size of an encoder.
// @param encoder The encoder to query.
// @return The block size, i.e. the total size in bytes that this encoder
//         operates on.
func (encoder *Encoder) BlockSize() uint32 {
	return uint32(C.krlnc_encoder_block_size(encoder.mEncoder))
}

// SetSymbolStorage specifies the source data for a given symbol.
// @param encoder The encoder which will encode the symbol
// @param data The buffer containing the data to be encoded
// @param index The index of the symbol in the coding block
func (encoder *Encoder) SetSymbolStorage(data *[]uint8, index uint32) {
	C.krlnc_encoder_set_symbol_storage(
		encoder.mEncoder, (*C.uint8_t)(&(*data)[0]), C.uint32_t(index))
}

// SetSymbolsStorage specifies the source data for all symbols. This will
// specify all symbols.
// @param encoder The encoder which will encode the data
// @param data The buffer containing the data to be encoded
func (encoder *Encoder) SetSymbolsStorage(data *[]uint8) {
	C.krlnc_encoder_set_symbols_storage(
		encoder.mEncoder, (*C.uint8_t)(&(*data)[0]))
}

// MaxPayloadSize returns the maximum possible payload size of an encoder,
// which is the maximum size of a generated payload.
// @param encoder The encoder to query.
// @return The payload size in bytes
func (encoder *Encoder) MaxPayloadSize() uint32 {
	return uint32(C.krlnc_encoder_max_payload_size(encoder.mEncoder))
}

// ProducePayload produces a payload representing a single encoded symbol in the
// provided buffer.
// @param encoder The encoder to use.
// @param payload The buffer which should contain the symbol.
// @return The total bytes used from the payload buffer
func (encoder *Encoder) ProducePayload(payload *[]uint8) uint32 {
	return uint32(C.krlnc_encoder_produce_payload(
		encoder.mEncoder, (*C.uint8_t)(&(*payload)[0])))
}

// Rank returns the rank of an encoder that indicates how many symbols are
//  available for encoding.
// @param encoder The encoder to query
// @return The rank of the encoder
func (encoder *Encoder) Rank() uint32 {
	return uint32(C.krlnc_encoder_rank(encoder.mEncoder))
}

// IsSystematicOn returns whether the encoder is in the systematic mode, i.e.
// if it will initially send the original source symbols with a simple header.
// @param encoder The encoder
// @return Non-zero if the encoder is in the systematic mode, otherwise 0
func (encoder *Encoder) IsSystematicOn() bool {
	return C.krlnc_encoder_is_systematic_on(encoder.mEncoder) != 0
}

// SetSystematicOn switches the systematic encoding on
// @param encoder The encoder
func (encoder *Encoder) SetSystematicOn() {
	C.krlnc_encoder_set_systematic_on(encoder.mEncoder)
}

// SetSystematicOff switches the systematic encoding off
// @param encoder The encoder
func (encoder *Encoder) SetSystematicOff() {
	C.krlnc_encoder_set_systematic_off(encoder.mEncoder)
}

// InSystematicPhase returns whether the encoder is in the systematic phase,
// i.e. there is a systematic packet to send
// @param encoder The encoder
// @return true if the encoder is in the systematic phase, otherwise false
func (encoder *Encoder) InSystematicPhase() bool {
	return C.krlnc_encoder_in_systematic_phase(encoder.mEncoder) != 0
}

// CoefficientVectorSize returns the size of the coefficient vector.
// @param encoder The encoder to check
// @return The size of the coefficient vector in bytes
func (encoder *Encoder) CoefficientVectorSize() uint32 {
	return uint32(C.krlnc_encoder_coefficient_vector_size(encoder.mEncoder))
}

// ProduceSymbol write an encoded symbol according to the provided symbol coefficients.
// @param encoder The encoder to use.
// @param symbol_data The destination buffer for the encoded symbol
// @param coefficients The desired coding coefficients that should
//        be used to calculate the encoded symbol.
// @return The number of bytes used.
func (encoder *Encoder) ProduceSymbol(
	symbolData *[]uint8, coefficients *[]uint8) uint32 {
	return uint32(C.krlnc_encoder_produce_symbol(
		encoder.mEncoder,
		(*C.uint8_t)(&(*symbolData)[0]),
		(*C.uint8_t)(&(*coefficients)[0])))
}

// ConsumeSystematicSymbol write a systematic/decoded symbol that corresponds
// to the provided symbol index.
// @param encoder The encoder to use.
// @param symbol_data The destination of the decoded source symbol.
// @param index The index of this decoded symbol in the data block.
// @return The number of bytes used.
func (encoder *Encoder) ConsumeSystematicSymbol(
	symbolData *[]uint8, index uint32) uint32 {
	return uint32(C.krlnc_encoder_produce_systematic_symbol(
		encoder.mEncoder,
		(*C.uint8_t)(&(*symbolData)[0]),
		C.uint32_t(index)))
}

// SetSeed sets the seed of the coefficient generator.
// @param encoder The encoder to use
// @param seedValue The seed value for the generator.
func (encoder *Encoder) SetSeed(seedValue uint32) {
	C.krlnc_encoder_set_seed(encoder.mEncoder, C.uint32_t(seedValue))
}

// Generate fills the input buffer with symbol coefficients used for either
// encoding or decoding a symbol.
// @param encoder The encoder to use.
// @param coefficients Pointer to the memory where the coefficients should
//        be stored. The coefficient buffer should have at least
//        CoefficientVectorSize() capacity.
func (encoder *Encoder) Generate(coefficients *[]uint8) {
	C.krlnc_encoder_generate(
		encoder.mEncoder, (*C.uint8_t)(&(*coefficients)[0]))
}

// GeneratePartial generates a "partial" coding vector that will only contain
// non-zero coefficients for the symbols that are currently defined.
// This allows encoding before defining all original source symbols,
// i.e. on-the-fly encoding.
// @param encoder The encoder to use.
// @param coefficients Pointer to the memory where the coefficients should
//        be stored. The coefficient buffer should have at least
//        CoefficientVectorSize() capacity.
func (encoder *Encoder) GeneratePartial(coefficients *[]uint8) {
	C.krlnc_encoder_generate_partial(
		encoder.mEncoder, (*C.uint8_t)(&(*coefficients)[0]))
}

// Density returns the density of the generated coding vector coefficients.
// @param encoder The encoder to query
// @return The coding vector density as a float
func (encoder *Encoder) Density() float32 {
	return float32(C.krlnc_encoder_density(encoder.mEncoder))
}

// SetDensity sets the density of the generated coding vector coefficients.
// @param encoder The encoder to use
// @param density The density value (0.0 < density <= 1.0)
func (encoder *Encoder) SetDensity(density float32) {
	C.krlnc_encoder_set_density(encoder.mEncoder, C.float(density))
}

// SetLogStdout enables the default log function of the encoder, which prints to the standard output.
// @param encoder The encoder to use
func (encoder *Encoder) SetLogStdout() {
	C.krlnc_encoder_set_log_stdout(encoder.mEncoder)
}

// SetLogOff disables the log function of the encoder.
// @param encoder The encoder to use
func (encoder *Encoder) SetLogOff() {
	C.krlnc_encoder_set_log_off(encoder.mEncoder)
}
