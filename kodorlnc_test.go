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

func (s *MySuite) TestDecoder(c *C) {
	var symbols uint32 = 50
	var symbolSize uint32 = 750
	decoder := NewDecoder(Binary8, symbols, symbolSize)

	c.Assert(symbols, Equals, decoder.Symbols())
	c.Assert(symbolSize, Equals, decoder.SymbolSize())
	c.Assert(symbolSize*symbols, Equals, decoder.BlockSize())
	c.Assert(int(symbolSize+53), Equals, int(decoder.MaxPayloadSize()))
	c.Assert(false, Equals, decoder.IsComplete())
	c.Assert(false, Equals, decoder.IsPartiallyComplete())
	c.Assert(uint32(0), Equals, decoder.Rank())
	c.Assert(uint32(0), Equals, decoder.SymbolsDecoded())

	c.Assert(true, Equals, decoder.IsSymbolMissing(0))
	c.Assert(false, Equals, decoder.IsSymbolPartiallyDecoded(0))
	c.Assert(false, Equals, decoder.IsSymbolDecoded(0))
	c.Assert(false, Equals, decoder.IsSymbolPivot(0))

	c.Assert(false, Equals, decoder.IsStatusUpdaterEnabled())
	decoder.SetStatusUpdaterOn()
	c.Assert(true, Equals, decoder.IsStatusUpdaterEnabled())
	decoder.SetStatusUpdaterOff()
	c.Assert(false, Equals, decoder.IsStatusUpdaterEnabled())
	decoder.UpdateSymbolStatus()

	c.Assert(50, Equals, int(decoder.CoefficientVectorSize()))

	decoder.SetLogStdout()
	decoder.SetLogOff()

	decoder.Reset()

	c.Assert(symbols, Equals, decoder.Symbols())
	c.Assert(symbolSize, Equals, decoder.SymbolSize())
}

func (s *MySuite) TestEncoder(c *C) {
	var symbols uint32 = 50
	var symbolSize uint32 = 750
	encoder := NewEncoder(Binary4, symbols, symbolSize)
	encoder.SetCodingVectorFormat(FullVector)
	c.Assert(symbols, Equals, encoder.Symbols())
	c.Assert(symbolSize, Equals, encoder.SymbolSize())
	c.Assert(symbols*symbolSize, Equals, encoder.BlockSize())
	c.Assert(int(symbolSize+28), Equals, int(encoder.MaxPayloadSize()))
	c.Assert(int(0), Equals, int(encoder.Rank()))

	c.Assert(false, Equals, encoder.InSystematicPhase())
	c.Assert(true, Equals, encoder.IsSystematicOn())
	encoder.SetSystematicOff()
	c.Assert(false, Equals, encoder.IsSystematicOn())
	encoder.SetSystematicOn()
	c.Assert(true, Equals, encoder.IsSystematicOn())

	c.Assert(25, Equals, int(encoder.CoefficientVectorSize()))

	c.Assert(float32(0.9375), Equals, encoder.Density())
	encoder.SetDensity(0.4)
	c.Assert(float32(0.4), Equals, encoder.Density())

	encoder.SetLogStdout()
	encoder.SetLogOff()

	encoder.Reset()

	c.Assert(symbols, Equals, encoder.Symbols())
	c.Assert(symbolSize, Equals, encoder.SymbolSize())
}
