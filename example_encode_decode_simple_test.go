package kodorlnc_test

// Copyright Steinwurf ApS 2018.
// Distributed under the "STEINWURF RESEARCH LICENSE 1.0".
// See accompanying file LICENSE.rst or
// http://www.steinwurf.com/licensing

import (
	"fmt"
	"math/rand"
	"time"

	. "github.com/steinwurf/kodo-rlnc-go"
)

func Example_encodeDecodeSimple() {
	// Seed random number generator to produce different results every time
	rand.Seed(time.Now().UTC().UnixNano())

	// Set the number of symbols (i.e. the generation size in RLNC
	// terminology) and the size of a symbol in bytes
	var symbols, symbolSize uint32 = 10, 100

	// Initialization of encoder and decoder
	encoder := NewEncoder(Binary8, symbols, symbolSize)
	decoder := NewDecoder(Binary8, symbols, symbolSize)

	// Allocate some storage for a "payload" the payload is what we would
	// eventually send over a network
	payload := make([]uint8, encoder.MaxPayloadSize())

	// Allocate some data to encode. In this case we make a buffer
	// with the same size as the encoder's block size (the max.
	// amount a single encoder can encode)
	dataIn := make([]uint8, encoder.BlockSize())

	// Just for fun - fill the data with random data
	for i := range dataIn {
		dataIn[i] = uint8(rand.Uint32())
	}

	// Assign the data buffer to the encoder so that we may start
	// to produce encoded symbols from it
	encoder.SetSymbolsStorage(&dataIn)

	// Set the storage for the decoder
	dataOut := make([]uint8, len(dataIn))
	decoder.SetSymbolsStorage(&dataOut)

	// Set systematic off
	encoder.SetSystematicOff()

	for !decoder.IsComplete() {

		// Encode the packet into the payload buffer
		encoder.ProducePayload(&payload)
		// Pass that packet to the decoder
		decoder.ConsumePayload(&payload)
	}

	// Check if we properly decoded the data
	for i, v := range dataIn {
		if v != dataOut[i] {
			fmt.Println("Unexpected failure to decode")
			fmt.Println("Please file a bug report :)")
			return
		}
	}
	fmt.Println("Data decoded correctly")
	// Output: Data decoded correctly
}
