package bit_flags

import (
	"testing"

	"github.com/stretchr/testify/require"
)

const (
	testProperty1 = 1 << iota
	testProperty2
	testProperty3
	testProperty4
	testProperty5
	testProperty6
)

func TestBitProperties_Set(t *testing.T) {
	bf := New[uint8]()
	bf.Set(testProperty1 | testProperty5)
	require.Equal(t, uint8(17), bf.Get())
}

func TestBitProperties_Add(t *testing.T) {
	bf := New[uint8]()
	bf.Set(testProperty1 | testProperty5)
	bf.Add(testProperty2 | testProperty3)
	require.Equal(t, uint8(23), bf.Get())
}

func TestBitProperties_Remove(t *testing.T) {
	bf := New[uint8]()
	bf.Set(testProperty1 | testProperty5 | testProperty4 | testProperty6)
	bf.Remove(testProperty5)
	require.Equal(t, uint8(41), bf.Get())
	require.Equal(t, false, bf.Has(testProperty5|testProperty1))
	require.Equal(t, true, bf.Has(testProperty1))
}

func TestBitProperties_Reset(t *testing.T) {
	bf := New[uint8]()
	bf.Set(testProperty1 | testProperty5)
	bf.Reset()
	require.Equal(t, uint8(0), bf.Get())
}

func TestBitProperties_Has(t *testing.T) {
	bf := New[uint8]()
	bf.Set(testProperty1 | testProperty5 | testProperty4 | testProperty6)
	require.Equal(t, true, bf.Has(testProperty5|testProperty1))
	require.Equal(t, false, bf.Has(testProperty2))
}

func TestBitProperties_List(t *testing.T) {
	bf := New[uint8]()
	bf.Set(testProperty1 | testProperty5 | testProperty4 | testProperty6)
	exp := []uint8{testProperty1, testProperty4, testProperty5, testProperty6}
	require.Equal(t, exp, bf.List())
}
