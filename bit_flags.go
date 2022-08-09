package bit_flags

import "math/bits"

type uints interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
}

// BitFlags stores boolean flags
//
// Use uintN type to store up to N flags, e.g. uint8 to <=8 flags
type BitFlags[T uints] struct {
	flags T
}

func New[T uints]() *BitFlags[T] {
	return new(BitFlags[T])
}

// Set sets the stored flags from the specified
func (f *BitFlags[T]) Set(flags T) {
	f.flags = flags
}

// Get returns the stored flags
func (f BitFlags[T]) Get() T {
	return f.flags
}

// Add adds the specified flags to the stored
func (f *BitFlags[T]) Add(flags T) {
	f.flags |= flags
}

// Remove deletes the specified flags from the stored
func (f *BitFlags[T]) Remove(flags T) {
	f.flags ^= flags
}

// Reset sets the stored flags to zero
func (f *BitFlags[T]) Reset() {
	f.flags = 0
}

// Has checks if all the specified flags exists in the stored
func (f BitFlags[T]) Has(flags T) bool {
	return f.flags&flags == flags
}

// List returns a list of the stored flags
func (f BitFlags[T]) List() []T {
	result := make([]T, 0, bits.Len64(uint64(f.flags)))
	for i := 0; i < cap(result); i++ {
		pow := T(1 << i)
		if f.flags&pow == pow {
			result = append(result, pow)
		}
	}
	return result
}
