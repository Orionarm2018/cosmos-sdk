package types

import (
	"math/big"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFromInt64(t *testing.T) {
	for n := 0; n < 20; n++ {
		r := rand.Int63()
		assert.Equal(t, r, NewInt(r).Int64())
	}
}

func TestInt(t *testing.T) {
	// Max Int = 2^255-1 = 5.789e+76
	// Min Int = -(2^255-1) = -5.789e+76
	i1, ok := NewIntWithDecimal(1, 76)
	assert.True(t, ok)
	i2, ok := NewIntWithDecimal(2, 76)
	assert.True(t, ok)
	i3, ok := NewIntWithDecimal(3, 76)
	assert.True(t, ok)
	_, ok = NewIntWithDecimal(6, 76)
	assert.False(t, ok)
	_, ok = NewIntWithDecimal(9, 80)
	assert.False(t, ok)

	// Overflow check
	assert.NotPanics(t, func() { i1.Add(i1) })
	assert.NotPanics(t, func() { i2.Add(i2) })
	assert.Panics(t, func() { i3.Add(i3) })

	assert.NotPanics(t, func() { i1.Sub(i1.Neg()) })
	assert.NotPanics(t, func() { i2.Sub(i2.Neg()) })
	assert.Panics(t, func() { i3.Sub(i3.Neg()) })

	assert.Panics(t, func() { i1.Mul(i1) })
	assert.Panics(t, func() { i2.Mul(i2) })
	assert.Panics(t, func() { i3.Mul(i3) })

	assert.Panics(t, func() { i1.Neg().Mul(i1.Neg()) })
	assert.Panics(t, func() { i2.Neg().Mul(i2.Neg()) })
	assert.Panics(t, func() { i3.Neg().Mul(i3.Neg()) })

	// Underflow check
	i3n := i3.Neg()
	assert.NotPanics(t, func() { i3n.Sub(i1) })
	assert.NotPanics(t, func() { i3n.Sub(i2) })
	assert.Panics(t, func() { i3n.Sub(i3) })

	assert.NotPanics(t, func() { i3n.Add(i1.Neg()) })
	assert.NotPanics(t, func() { i3n.Add(i2.Neg()) })
	assert.Panics(t, func() { i3n.Add(i3.Neg()) })

	assert.Panics(t, func() { i1.Mul(i1.Neg()) })
	assert.Panics(t, func() { i2.Mul(i2.Neg()) })
	assert.Panics(t, func() { i3.Mul(i3.Neg()) })

	// Bound check
	intmax := NewIntFromBigInt(new(big.Int).Sub(new(big.Int).Exp(big.NewInt(2), big.NewInt(255), nil), big.NewInt(1)))
	intmin := intmax.Neg()
	assert.NotPanics(t, func() { intmax.Add(ZeroInt()) })
	assert.NotPanics(t, func() { intmin.Sub(ZeroInt()) })
	assert.Panics(t, func() { intmax.Add(OneInt()) })
	assert.Panics(t, func() { intmin.Sub(OneInt()) })
}

func TestUint(t *testing.T) {
	// Max Uint = 1.15e+77
	// Min Uint = 0
	i1, ok := NewUintWithDecimal(5, 76)
	assert.True(t, ok)
	i2, ok := NewUintWithDecimal(10, 76)
	assert.True(t, ok)
	i3, ok := NewUintWithDecimal(11, 76)
	assert.True(t, ok)
	_, ok = NewUintWithDecimal(12, 76)
	assert.False(t, ok)
	_, ok = NewUintWithDecimal(1, 80)
	assert.False(t, ok)

	// Overflow check
	assert.NotPanics(t, func() { i1.Add(i1) })
	assert.Panics(t, func() { i2.Add(i2) })
	assert.Panics(t, func() { i3.Add(i3) })

	assert.Panics(t, func() { i1.Mul(i1) })
	assert.Panics(t, func() { i2.Mul(i2) })
	assert.Panics(t, func() { i3.Mul(i3) })

	// Underflow check
	assert.NotPanics(t, func() { i2.Sub(i1) })
	assert.NotPanics(t, func() { i2.Sub(i2) })
	assert.Panics(t, func() { i2.Sub(i3) })

	// Bound check
	uintmax := NewUintFromBigInt(new(big.Int).Sub(new(big.Int).Exp(big.NewInt(2), big.NewInt(256), nil), big.NewInt(1)))
	uintmin := NewUint(0)
	assert.NotPanics(t, func() { uintmax.Add(ZeroUint()) })
	assert.NotPanics(t, func() { uintmin.Sub(ZeroUint()) })
	assert.Panics(t, func() { uintmax.Add(OneUint()) })
	assert.Panics(t, func() { uintmin.Sub(OneUint()) })

}
