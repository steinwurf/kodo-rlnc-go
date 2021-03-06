package kodorlnc_test

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	. "github.com/steinwurf/kodo-rlnc-go"
)

func Benchmark(b *testing.B) {
	// Seed random number generator to produce different results every time
	rand.Seed(time.Now().UTC().UnixNano())

	symbols := uint32(25)
	symbolSize := uint32(1600)
	fields := map[string]int32{
		"Binary":   Binary,
		"Binary4":  Binary4,
		"Binary8":  Binary8,
		"Binary16": Binary16,
	}

	for name, field := range fields {
		encoder := NewEncoder(field, symbols, symbolSize)

		// Allocate some data to encode. In this case we make a buffer
		// with the same size as the encoder's block size (the max.
		// amount a single encoder can encode)
		blockSize := encoder.Symbols() * encoder.SymbolSize()
		dataIn := make([]uint8, blockSize)

		// Just for fun - fill the data with random data
		for i := range dataIn {
			dataIn[i] = uint8(rand.Uint32())
		}

		var payloads [][]uint8
		b.Run(name+"Encode", func(b *testing.B) { payloads = encodeData(b, encoder, &dataIn) })

		dataOut := make([]uint8, len(dataIn))
		decoder := NewDecoder(field, symbols, symbolSize)
		b.Run(name+"Decode", func(b *testing.B) { decodeData(b, decoder, &dataOut, &payloads) })

		var success bool = true
		// Check if we properly decoded the data
		for i, v := range dataIn {
			if v != dataOut[i] {
				success = false
				break
			}
		}

		if success == true {
			fmt.Println("Data decoded correctly")
		} else {
			fmt.Println("Decoding failed")
			b.FailNow()
		}
	}
}

func encodeData(
	b *testing.B, encoder *Encoder, dataIn *[]uint8) [][]uint8 {

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		encoder.Reset()

		benchpayload := make([]uint8, encoder.MaxPayloadSize())

		// We measure pure coding, so we always turn off the systematic mode
		encoder.SetSystematicOff()

		// Copy the input data to the encoder
		encoder.SetSymbolsStorage(dataIn)

		for i := 0; i < int(encoder.Symbols())*2; i++ {
			encoder.ProducePayload(&benchpayload)
		}
	}

	// Generate an ample number of coded symbols (considering kodo_binary)
	payloadCount := 2 * int(encoder.Symbols())

	// The generated payloads will be stored for the decoder
	payloads := make([][]uint8, payloadCount)
	for i := range payloads {
		payloads[i] = make([]uint8, encoder.MaxPayloadSize())
	}
	for i := 0; i < len(payloads); i++ {
		encoder.ProducePayload(&payloads[i])
	}

	return payloads
}

func decodeData(
	b *testing.B,
	decoder *Decoder,
	dataOut *[]uint8,
	payloads *[][]uint8) {

	f := func(decoder *Decoder, payloads [][]uint8) {
		for _, payload := range payloads {
			if decoder.IsComplete() {
				break
			}
			decoder.ConsumePayload(&payload)
		}
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		decoder.Reset()
		// Set the storage for the decoder
		decoder.SetSymbolsStorage(dataOut)
		f(decoder, *payloads)
	}
}
