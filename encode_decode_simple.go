package kodorlnc

import (
	"fmt"
	"math/rand"
	"time"
)

// ExampleEncodeDecodeSimple is a simple example of using kodorlnc
func ExampleEncodeDecodeSimple() {
	// Seed random number generator to produce different results every time
	rand.Seed(time.Now().UTC().UnixNano())

	// Set the number of symbols (i.e. the generation size in RLNC
	// terminology) and the size of a symbol in bytes
	var symbols, symbolSize uint32 = 10, 100

	// Initialization of encoder and decoder
	encoderFactory := NewEncoderFactory(Binary8, symbols, symbolSize)
	decoderFactory := NewDecoderFactory(Binary8, symbols, symbolSize)

	// These lines show the API to clean the memory used by the factories
	defer encoderFactory.Destruct()
	defer decoderFactory.Destruct()

	encoder := encoderFactory.Build()
	decoder := decoderFactory.Build()

	// These lines show the API to clean the memory used by the coders
	defer encoder.Destruct()
	defer decoder.Destruct()

	// Allocate some storage for a "payload" the payload is what we would
	// eventually send over a network
	payload := make([]uint8, encoder.PayloadSize())

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
	encoder.SetConstSymbols(&dataIn[0], symbols*symbolSize)

	// Set the storage for the decoder
	dataOut := make([]uint8, len(dataIn))
	decoder.SetMutableSymbols(&dataOut[0], decoder.BlockSize())

	// Set systematic off
	encoder.SetSystematicOff()

	for !decoder.IsComplete() {

		// Encode the packet into the payload buffer
		encoder.WritePayload(&payload[0])
		// Pass that packet to the decoder
		decoder.ReadPayload(&payload[0])
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
}
