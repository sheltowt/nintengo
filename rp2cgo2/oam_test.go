package rp2cgo2

import (
	"testing"
)

func TestDisableReads(t *testing.T) {
	oam := NewOAM()

	oam.Buffer.Store(0, 0xff)

	if oam.Buffer.Fetch(0) != 0xff {
		t.Error("Memory is not 0xff")
	}

	oam.Buffer.Store(0, 0x00)

	if oam.Buffer.Fetch(0) != 0x00 {
		t.Error("Memory is not 0x00")
	}

	oam.Buffer.DisableReads()

	if oam.Buffer.Fetch(0) != 0xff {
		t.Error("Memory is not 0xff")
	}

	oam.Buffer.EnableReads()

	if oam.Buffer.Fetch(0) != 0x00 {
		t.Error("Memory is not 0x00")
	}
}

func TestDisableWrites(t *testing.T) {
	oam := NewOAM()

	oam.Buffer.Store(0, 0xff)

	if oam.Buffer.Fetch(0) != 0xff {
		t.Error("Memory is not 0xff")
	}

	oam.Buffer.Store(0, 0x00)

	if oam.Buffer.Fetch(0) != 0x00 {
		t.Error("Memory is not 0x00")
	}

	oam.Buffer.DisableWrites()

	oam.Buffer.Store(0, 0xff)

	if oam.Buffer.Fetch(0) != 0x00 {
		t.Error("Memory is not 0x00")
	}

	oam.Buffer.EnableWrites()

	oam.Buffer.Store(0, 0xff)

	if oam.Buffer.Fetch(0) != 0xff {
		t.Error("Memory is not 0xff")
	}
}

func TestClearOAMBuffer(t *testing.T) {
	oam := NewOAM()

	for i := uint16(0); i < 256; i++ {
		oam.BasicMemory.Store(i, 0x00)
	}

	for i := uint16(0); i < 32; i++ {
		oam.Buffer.Store(i, 0x00)
	}

	scanline := uint16(0)
	cycle := uint16(0)
	size := uint16(8)

	for cycle = 1; cycle <= 64; cycle++ {
		oam.SpriteEvaluation(scanline, cycle, size)
	}

	actual := uint8(0)
	expected := uint8(0xff)

	for i := uint16(0); i < 32; i++ {
		actual = oam.Buffer.Fetch(i)

		if actual != expected {
			t.Errorf("Memory is %02X not %02X\n", actual, expected)
		}
	}
}

func TestSpriteEvaluation1(t *testing.T) {
	oam := NewOAM()

	scanline := uint16(0)
	cycle := uint16(0)
	size := uint16(8)

	buf := []uint8{
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,

		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,

		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,

		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
	}

	for i := uint16(0); i < 32; i++ {
		oam.Buffer.Store(i, 0xff)
	}

	for i, e := range buf {
		oam.Store(uint16(i), e)
	}

	scanline = 0

	for cycle = 65; cycle <= 256; cycle++ {
		oam.SpriteEvaluation(scanline, cycle, size)
	}

	buf = []uint8{
		0x00, 0x00, 0x00, 0x01,
		0x00, 0x00, 0x00, 0x01,
		0x00, 0x00, 0x00, 0x01,
		0x00, 0x00, 0x00, 0x01,

		0x00, 0x00, 0x00, 0x01,
		0x00, 0x00, 0x00, 0x01,
		0x00, 0x00, 0x00, 0x01,
		0x00, 0x00, 0x00, 0x01,
	}

	for i, e := range buf {
		if oam.Buffer.Fetch(uint16(i)) != e {
			t.Errorf("Memory %v is %02X not %02X\n", i, oam.Buffer.Fetch(uint16(i)), e)
		}
	}
}

func TestSpriteEvaluation2(t *testing.T) {
	oam := NewOAM()

	scanline := uint16(0)
	cycle := uint16(0)
	size := uint16(8)

	buf := []uint8{
		0x00, 0x00, 0x00, 0x00,
		0xff, 0x00, 0x00, 0x00,
		0x00, 0x01, 0x00, 0x00,
		0xff, 0x00, 0x00, 0x00,

		0xff, 0x00, 0x00, 0x00,
		0x00, 0x02, 0x00, 0x00,
		0xff, 0x00, 0x00, 0x00,
		0x00, 0x03, 0x00, 0x00,

		0x00, 0x04, 0x00, 0x00,
		0xff, 0x00, 0x00, 0x00,
		0x00, 0x05, 0x00, 0x00,
		0xff, 0x00, 0x00, 0x00,

		0xff, 0x00, 0x00, 0x00,
		0x00, 0x06, 0x00, 0x00,
		0xff, 0x00, 0x00, 0x00,
		0x00, 0x07, 0x00, 0x00,
	}

	for i := uint16(0); i < 32; i++ {
		oam.Buffer.Store(i, 0xff)
	}

	for i, e := range buf {
		oam.Store(uint16(i), e)
	}

	scanline = 7

	for cycle = 65; cycle <= 256; cycle++ {
		oam.SpriteEvaluation(scanline, cycle, size)
	}

	buf = []uint8{
		0x00, 0x00, 0x00, 0x01,
		0x00, 0x01, 0x00, 0x01,
		0x00, 0x02, 0x00, 0x01,
		0x00, 0x03, 0x00, 0x01,

		0x00, 0x04, 0x00, 0x01,
		0x00, 0x05, 0x00, 0x01,
		0x00, 0x06, 0x00, 0x01,
		0x00, 0x07, 0x00, 0x01,
	}

	for i, e := range buf {
		if oam.Buffer.Fetch(uint16(i)) != e {
			t.Errorf("Memory %v is %02X not %02X\n", i, oam.Buffer.Fetch(uint16(i)), e)
		}
	}
}

func TestSpriteEvaluation3(t *testing.T) {
	oam := NewOAM()

	scanline := uint16(0)
	cycle := uint16(0)
	size := uint16(8)

	buf := []uint8{
		0x02, 0x00, 0x00, 0x00,
		0xff, 0x00, 0x00, 0x00,
		0x03, 0x01, 0x01, 0x01,
		0xff, 0x00, 0x00, 0x00,

		0xff, 0x00, 0x00, 0x00,
		0x04, 0x02, 0x02, 0x02,
		0xff, 0x00, 0x00, 0x00,
		0x05, 0x03, 0x03, 0x03,

		0x06, 0x04, 0x04, 0x04,
		0xff, 0x00, 0x00, 0x00,
		0x07, 0x05, 0x05, 0x05,
		0xff, 0x00, 0x00, 0x00,

		0xff, 0x00, 0x00, 0x00,
		0x08, 0x06, 0x06, 0x06,
		0xff, 0x00, 0x00, 0x00,
		0x09, 0x07, 0x07, 0x07,
	}

	for i := uint16(0); i < 32; i++ {
		oam.Buffer.Store(i, 0xff)
	}

	for i, e := range buf {
		oam.Store(uint16(i), e)
	}

	scanline = 9

	for cycle = 65; cycle <= 256; cycle++ {
		oam.SpriteEvaluation(scanline, cycle, size)
	}

	buf = []uint8{
		0x02, 0x00, 0x00, 0x01,
		0x03, 0x01, 0x01, 0x02,
		0x04, 0x02, 0x02, 0x03,
		0x05, 0x03, 0x03, 0x04,

		0x06, 0x04, 0x04, 0x05,
		0x07, 0x05, 0x05, 0x06,
		0x08, 0x06, 0x06, 0x07,
		0x09, 0x07, 0x07, 0x08,
	}

	for i, e := range buf {
		if oam.Buffer.Fetch(uint16(i)) != e {
			t.Errorf("Memory %v is %02X not %02X\n", i, oam.Buffer.Fetch(uint16(i)), e)
		}
	}
}

func TestSpriteEvaluation4(t *testing.T) {
	oam := NewOAM()

	scanline := uint16(0)
	cycle := uint16(0)
	size := uint16(8)

	buf := []uint8{
		0x02, 0x00, 0x00, 0x00,
		0xff, 0x00, 0x00, 0x00,
		0x03, 0x01, 0x01, 0x01,
		0xff, 0x00, 0x00, 0x00,

		0xff, 0x00, 0x00, 0x00,
		0x04, 0x02, 0x02, 0x02,
		0xff, 0x00, 0x00, 0x00,
		0x05, 0x03, 0x03, 0x03,

		0x06, 0x04, 0x04, 0x04,
		0xff, 0x00, 0x00, 0x00,
		0x07, 0x05, 0x05, 0x05,
		0xff, 0x00, 0x00, 0x00,

		0xff, 0x00, 0x00, 0x00,
		0x08, 0x06, 0x06, 0x06,
		0x09, 0x07, 0x07, 0x07,
		0xff, 0x00, 0x00, 0x00,
		// sprites to be skipped
		0x02, 0x00, 0x00, 0x00,
		0xff, 0x00, 0x00, 0x00,
		0x03, 0x00, 0x00, 0x00,
		0xff, 0x00, 0x00, 0x00,

		0xff, 0x00, 0x00, 0x00,
		0x04, 0x00, 0x00, 0x00,
		0xff, 0x00, 0x00, 0x00,
		0x05, 0x00, 0x00, 0x00,

		0x06, 0x00, 0x00, 0x00,
		0xff, 0x00, 0x00, 0x00,
		0x07, 0x00, 0x00, 0x00,
		0xff, 0x00, 0x00, 0x00,

		0xff, 0x00, 0x00, 0x00,
		0x08, 0x00, 0x00, 0x00,
		0x09, 0x00, 0x00, 0x00,
		0xff, 0x00, 0x00, 0x00,
	}

	for i := uint16(0); i < 32; i++ {
		oam.Buffer.Store(i, 0xff)
	}

	for i, e := range buf {
		oam.Store(uint16(i), e)
	}

	scanline = 9

	for cycle = 65; cycle <= 256; cycle++ {
		oam.SpriteEvaluation(scanline, cycle, size)
	}

	buf = []uint8{
		0x02, 0x00, 0x00, 0x01,
		0x03, 0x01, 0x01, 0x02,
		0x04, 0x02, 0x02, 0x03,
		0x05, 0x03, 0x03, 0x04,

		0x06, 0x04, 0x04, 0x05,
		0x07, 0x05, 0x05, 0x06,
		0x08, 0x06, 0x06, 0x07,
		0x09, 0x07, 0x07, 0x08,
	}

	for i, e := range buf {
		if oam.Buffer.Fetch(uint16(i)) != e {
			t.Errorf("Memory %v is %02X not %02X\n", i, oam.Buffer.Fetch(uint16(i)), e)
		}
	}
}

func TestSpriteOverflowClear(t *testing.T) {
	oam := NewOAM()

	scanline := uint16(0)
	cycle := uint16(0)
	size := uint16(8)

	buf := []uint8{
		0x02, 0x00, 0x00, 0x00,
		0xff, 0x00, 0x00, 0x00,
		0x03, 0x01, 0x01, 0x01,
		0xff, 0x00, 0x00, 0x00,

		0xff, 0x00, 0x00, 0x00,
		0x04, 0x02, 0x02, 0x02,
		0xff, 0x00, 0x00, 0x00,
		0x05, 0x03, 0x03, 0x03,

		0x06, 0x04, 0x04, 0x04,
		0xff, 0x00, 0x00, 0x00,
		0x07, 0x05, 0x05, 0x05,
		0xff, 0x00, 0x00, 0x00,

		0xff, 0x00, 0x00, 0x00,
		0x08, 0x06, 0x06, 0x06,
		0xff, 0x00, 0x00, 0x00,
		0x09, 0x07, 0x07, 0x07,

		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
	}

	for i := uint16(0); i < 32; i++ {
		oam.Buffer.Store(i, 0xff)
	}

	for i, e := range buf {
		oam.Store(uint16(i), e)
	}

	scanline = 9

	for cycle = 65; cycle <= 256; cycle++ {
		spriteOverflow := oam.SpriteEvaluation(scanline, cycle, size)

		if spriteOverflow {
			t.Error("SpriteOverflow set")
		}
	}
}

func TestSpriteOverflowSet1(t *testing.T) {
	oam := NewOAM()

	scanline := uint16(0)
	cycle := uint16(0)
	size := uint16(8)

	buf := []uint8{
		0x02, 0x00, 0x00, 0x00,
		0xff, 0x00, 0x00, 0x00,
		0x03, 0x01, 0x01, 0x01,
		0xff, 0x00, 0x00, 0x00,

		0xff, 0x00, 0x00, 0x00,
		0x04, 0x02, 0x02, 0x02,
		0xff, 0x00, 0x00, 0x00,
		0x05, 0x03, 0x03, 0x03,

		0x06, 0x04, 0x04, 0x04,
		0xff, 0x00, 0x00, 0x00,
		0x07, 0x05, 0x05, 0x05,
		0xff, 0x00, 0x00, 0x00,

		0xff, 0x00, 0x00, 0x00,
		0x08, 0x06, 0x06, 0x06,
		0xff, 0x00, 0x00, 0x00,
		0x09, 0x07, 0x07, 0x07,

		0x09, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
	}

	for i := uint16(0); i < 32; i++ {
		oam.Buffer.Store(i, 0xff)
	}

	for i, e := range buf {
		oam.Store(uint16(i), e)
	}

	scanline = 9
	spriteOverflow := false

	for cycle = 65; cycle <= 256; cycle++ {
		spriteOverflow = spriteOverflow || oam.SpriteEvaluation(scanline, cycle, size)
	}

	if !spriteOverflow {
		t.Error("SpriteOverflow not set")
	}
}

func TestSpriteOverflowSet2(t *testing.T) {
	oam := NewOAM()

	scanline := uint16(0)
	cycle := uint16(0)
	size := uint16(8)

	buf := []uint8{
		0x02, 0x00, 0x00, 0x00,
		0xff, 0x00, 0x00, 0x00,
		0x03, 0x01, 0x01, 0x01,
		0xff, 0x00, 0x00, 0x00,

		0xff, 0x00, 0x00, 0x00,
		0x04, 0x02, 0x02, 0x02,
		0xff, 0x00, 0x00, 0x00,
		0x05, 0x03, 0x03, 0x03,

		0x06, 0x04, 0x04, 0x04,
		0xff, 0x00, 0x00, 0x00,
		0x07, 0x05, 0x05, 0x05,
		0xff, 0x00, 0x00, 0x00,

		0xff, 0x00, 0x00, 0x00,
		0x08, 0x06, 0x06, 0x06,
		0xff, 0x00, 0x00, 0x00,
		0x09, 0x07, 0x07, 0x07,

		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x09, 0x00,
		0x00, 0x00, 0x00, 0x00,
	}

	for i := uint16(0); i < 32; i++ {
		oam.Buffer.Store(i, 0xff)
	}

	for i, e := range buf {
		oam.Store(uint16(i), e)
	}

	scanline = 9
	spriteOverflow := false

	for cycle = 65; cycle <= 256; cycle++ {
		spriteOverflow = spriteOverflow || oam.SpriteEvaluation(scanline, cycle, size)
	}

	if !spriteOverflow {
		t.Error("SpriteOverflow not set")
	}
}

func TestOAMSprite(t *testing.T) {
	oam := NewOAM()

	address := uint16(0)

	address = 0x0000

	oam.Buffer.Store(address+0, 0x00)
	oam.Buffer.Store(address+1, 0x01)
	oam.Buffer.Store(address+2, 0x02)
	oam.Buffer.Store(address+3, 0x03)

	if oam.Sprite(0) != 0x00010203 {
		t.Error("Sprite is not 0x00010203")
	}

	address = 0x0004

	oam.Buffer.Store(address+0, 0x04)
	oam.Buffer.Store(address+1, 0x05)
	oam.Buffer.Store(address+2, 0x06)
	oam.Buffer.Store(address+3, 0x07)

	if oam.Sprite(1) != 0x04050607 {
		t.Error("Sprite is not 0x04050607")
	}

	address = 0x001c

	oam.Buffer.Store(address+0, 0xfc)
	oam.Buffer.Store(address+1, 0xfd)
	oam.Buffer.Store(address+2, 0xfe)
	oam.Buffer.Store(address+3, 0xff)

	if oam.Sprite(7) != 0xfcfdfeff {
		t.Error("Sprite is not 0xfcfdfeff")
	}

}
