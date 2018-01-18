package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"

	"gitlab.com/steinwurf/kodo-go/src/kodo"
)

func main() {
	// Seed random number generator to produce different results every time
	rand.Seed(time.Now().UTC().UnixNano())

	algorithms := map[string]interface{}{
		"full_vector":        kodo.FullVector,
		"on_the_fly":         kodo.OnTheFly,
		"sliding_window":     kodo.SlidingWindow,
		"sparse_full_vector": kodo.SparseFullVector,
		"seed":               kodo.Seed,
		"sparse_seed":        kodo.SparseSeed,
		"perpetual":          kodo.Perpetual,
		"fulcrum":            kodo.Fulcrum,
		"reed_solomon":       kodo.ReedSolomon}

	fields := map[string]interface{}{
		"binary":  kodo.Binary,
		"binary4": kodo.Binary4,
		"binary8": kodo.Binary8}

	var algorithm = flag.String(
		"algorithm", "full_vector", "The algorithm to use")
	var field = flag.String("field", "binary8", "The finite field to use")
	var symbols = flag.Int("symbols", 16, "The number of symbols")
	var symbol_size = flag.Int("symbol_size", 1600, "The size of each symbol")

	flag.Parse()

	fmt.Printf("Algorithm: %s / Finite field: %s\n", *algorithm, *field)
	fmt.Printf("Symbols: %d / Symbol size: %d\n", *symbols, *symbol_size)

	success, encoding_rate, decoding_rate := run_coding_test(
		algorithms[*algorithm], fields[*field], uint32(*symbols),
		uint32(*symbol_size))

	fmt.Println("Encoding rate: ", encoding_rate, "MB/s")
	fmt.Println("Decoding rate: ", decoding_rate, "MB/s")

	if success == true {
		fmt.Println("Data decoded correctly")
	} else {
		fmt.Println("Decoding failed")
	}

}

func run_coding_test(algorithm interface{}, field interface{}, symbols uint32,
	symbol_size uint32) (bool, float64, float64) {

	a := algorithm.(int)
	f := field.(int)

	start := time.Now()
	encoder_factory := kodo.NewEncoderFactory(int32(a), int32(f), symbols,
		symbol_size)

	decoder_factory := kodo.NewDecoderFactory(kodo.FullVector,
		kodo.Binary8, symbols, symbol_size)

	encoder := encoder_factory.Build()
	decoder := decoder_factory.Build()
	setup_time := time.Since(start)

	// We measure pure coding, so we always turn off the systematic mode
	encoder.SetSystematicOff()

	// Allocate some data to encode. In this case we make a buffer
	// with the same size as the encoder's block size (the max.
	// amount a single encoder can encode)
	data_in := make([]uint8, encoder.BlockSize())

	// Set the storage for the decoder
	data_out := make([]uint8, len(data_in))

	// Just for fun - fill the data with random data
	for i, _ := range data_in {
		data_in[i] = uint8(rand.Uint32())
	}

	// Generate an ample number of coded symbols (considering kodo_binary)
	payload_count := 2 * symbols

	// The generated payloads will be stored in this slice
	payloads := make([][]uint8, payload_count)
	for i := range payloads {
		payloads[i] = make([]uint8, encoder.PayloadSize())
	}

	// Start the encoding timer
	start = time.Now()

	// Copy the input data to the encoder
	encoder.SetConstSymbols(&data_in[0], encoder.BlockSize())

	// Generate coded symbols with the encoder
	for i := uint32(0); i < payload_count; i++ {
		encoder.WritePayload(&payloads[i][0])
	}
	encoding_time := time.Since(start)

	// Calculate the encoding rate in megabytes / seconds
	encoded_bytes := payload_count * symbol_size
	encoding_rate := float64(encoded_bytes) / encoding_time.Seconds() / 1000000

	// Start the decoding timer
	start = time.Now()

	decoder.SetMutableSymbols(&data_out[0], decoder.BlockSize())

	// Feed the coded symbols to the decoder
	for _, v := range payloads {
		if decoder.IsComplete() {
			break
		}
		decoder.ReadPayload(&v[0])
	}

	// Calculate the decoding time
	decoding_time := time.Since(start)

	// Calculate the decoding rate in megabytes / seconds
	decoded_bytes := symbols * symbol_size
	decoding_rate := float64(decoded_bytes) / decoding_time.Seconds() / 1000000

	var success bool = true
	// Check if we properly decoded the data
	for i, v := range data_in {
		if v != data_out[i] {
			success = false
			break
		}
	}

	fmt.Println("Setup time: ", setup_time.Seconds()*1000000, "microsec")
	fmt.Println("Encoding time: ", encoding_time.Seconds()*1000000, "microsec")
	fmt.Println("Decoding time: ", decoding_time.Seconds()*1000000, "microsec")

	return success, encoding_rate, decoding_rate

}
