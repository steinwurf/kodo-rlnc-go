package kodorlnc_test

import (
	"testing"

	. "github.com/steinwurf/kodo-rlnc-go"
	. "gopkg.in/check.v1"
)

// Hook up gocheck into the "go test" runner.
func Test(t *testing.T) { TestingT(t) }

type MySuite struct{}

var _ = Suite(&MySuite{})

func (s *MySuite) TestDecoderFactory(c *C) {
	var symbols uint32 = 50
	var symbolSize uint32 = 750
	decoderFactory := NewDecoderFactory(Binary8, symbols, symbolSize)

	c.Assert(symbols, Equals, decoderFactory.Symbols())
	c.Assert(symbolSize, Equals, decoderFactory.SymbolSize())

	var newSymbols uint32 = 25
	decoderFactory.SetSymbols(newSymbols)
	c.Assert(newSymbols, Equals, decoderFactory.Symbols())

	var newSymbolSize uint32 = 300
	decoderFactory.SetSymbolSize(newSymbolSize)
	c.Assert(newSymbolSize, Equals, decoderFactory.SymbolSize())
}

func (s *MySuite) TestDecoder(c *C) {
	var symbols uint32 = 50
	var symbolSize uint32 = 750
	decoderFactory := NewDecoderFactory(Binary4, symbols, symbolSize)
	decoder := decoderFactory.Build()
	c.Assert(symbols, Equals, decoder.Symbols())
	c.Assert(symbolSize, Equals, decoder.SymbolSize())
	c.Assert(false, Equals, decoder.IsComplete())
	c.Assert(uint32(0), Equals, decoder.Rank())
	c.Assert(symbols*symbolSize, Equals, decoder.BlockSize())
	c.Assert(symbolSize+uint32(34), Equals, decoder.PayloadSize())
}

func (s *MySuite) TestEncoderFactory(c *C) {
	var symbols uint32 = 50
	var symbolSize uint32 = 750
	encoderFactory := NewEncoderFactory(Binary8, symbols, symbolSize)

	c.Assert(symbols, Equals, encoderFactory.Symbols())
	c.Assert(symbolSize, Equals, encoderFactory.SymbolSize())

	var newSymbols uint32 = 25
	encoderFactory.SetSymbols(newSymbols)

	c.Assert(newSymbols, Equals, encoderFactory.Symbols())

	var newSymbolSize uint32 = 300
	encoderFactory.SetSymbolSize(newSymbolSize)
	c.Assert(newSymbolSize, Equals, encoderFactory.SymbolSize())
}

func (s *MySuite) TestEncoder(c *C) {
	var symbols uint32 = 50
	var symbolSize uint32 = 750
	encoderFactory := NewEncoderFactory(Binary4, symbols, symbolSize)
	encoder := encoderFactory.Build()
	c.Assert(symbols, Equals, encoder.Symbols())
	c.Assert(symbolSize, Equals, encoder.SymbolSize())
	c.Assert(symbols*symbolSize, Equals, encoder.BlockSize())
	c.Assert(symbolSize+uint32(34), Equals, encoder.PayloadSize())

	c.Assert(true, Equals, encoder.IsSystematicOn())
	encoder.SetSystematicOff()
	c.Assert(false, Equals, encoder.IsSystematicOn())
	encoder.SetSystematicOn()
	c.Assert(true, Equals, encoder.IsSystematicOn())
}
