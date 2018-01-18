package main

import (
	"fmt"
	"math/rand"
	"time"

	"gitlab.com/steinwurf/kodo-go/src/kodo"
)

func main() {
	// Seed random number generator to produce different results every time
	rand.Seed(time.Now().UTC().UnixNano())

	// Set the number of symbols (i.e. the generation size in RLNC
	// terminology) and the size of a symbol in bytes
	var symbols, symbol_size uint32 = 10, 100

	// Initilization of encoder and decoder
	encoder_factory := kodo.NewEncoderFactory(kodo.FullVector,
		kodo.Binary8, symbols, symbol_size)
	decoder_factory := kodo.NewDecoderFactory(kodo.FullVector,
		kodo.Binary8, symbols, symbol_size)

	// These lines show the API to clean the memory used by the factories
	defer kodo.DeleteEncoderFactory(encoder_factory)
	defer kodo.DeleteDecoderFactory(decoder_factory)

	encoder := encoder_factory.Build()
	decoder := decoder_factory.Build()

	// These lines show the API to clean the memory used by the coders
	defer kodo.DeleteEncoder(encoder)
	defer kodo.DeleteDecoder(decoder)

	// Allocate some storage for a "payload" the payload is what we would
	// eventually send over a network
	payload := make([]uint8, encoder.PayloadSize())

	// Allocate some data to encode. In this case we make a buffer
	// with the same size as the encoder's block size (the max.
	// amount a single encoder can encode)
	data_in := make([]uint8, encoder.BlockSize())

	// Just for fun - fill the data with random data
	for i, _ := range data_in {
		data_in[i] = uint8(rand.Uint32())
	}

	// Assign the data buffer to the encoder so that we may start
	// to produce encoded symbols from it
	encoder.SetConstSymbols(&data_in[0], symbols*symbol_size)

	// Set the storage for the decoder
	data_out := make([]uint8, len(data_in))
	decoder.SetMutableSymbols(&data_out[0], decoder.BlockSize())

	// Set systematic off
	encoder.SetSystematicOff()

	for !decoder.IsComplete() {

		// Encode the packet into the payload buffer
		encoder.WritePayload(&payload[0])
		// Pass that packet to the decoder
		decoder.ReadPayload(&payload[0])
	}

	// Check if we properly decoded the data
	for i, v := range data_in {
		if v != data_out[i] {
			fmt.Println("Unexpected failure to decode")
			fmt.Println("Please file a bug report :)")
			return
		}
	}
	fmt.Println("Data decoded correctly")
}
