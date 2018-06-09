package types

import (
	"math/big"
)

// integer defines internal wrapper for big.Int
type integer struct {
	i big.Int `json:"int"`
}

func newInteger(n int64) integer {
	return integer{*big.NewInt(n)}
}

func (i integer) bigInt() *big.Int {
	return new(big.Int).Set(&(i.i))
}

func newIntegerFromBigInt(i *big.Int) integer {
	return integer{*i}
}

func newIntegerFromString(s string) (res integer, ok bool) {
	i, ok := new(big.Int).SetString(s, 0)
	if !ok {
		return
	}
	return integer{*i}, ok
}

func newIntegerWithDecimal(n int64, dec int) (res integer) {
	if dec < 0 {
		return
	}
	exp := new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(dec)), nil)
	i := new(big.Int)
	return integer{*(i.Mul(big.NewInt(n), exp))}
}

func zeroInt() integer { return integer{*big.NewInt(0)} }
func oneInt() integer  { return integer{*big.NewInt(1)} }

func (i integer) toInt64() int64 {
	if !(&(i.i)).IsInt64() {
		panic("Out of bound in Int64()")
	}
	return (&(i.i)).Int64()
}

func (i integer) toUint64() uint64 {
	if !(&(i.i)).IsUint64() {
		panic("Out of bound in Int64()")
	}
	return (&(i.i)).Uint64()
}

func (i integer) isZero() bool { return (&(i.i)).Sign() == 0 }

func (i integer) sign() int { return (&(i.i)).Sign() }

func (i integer) equal(i2 integer) bool { return (&(i.i)).Cmp(&(i2.i)) == 0 }

func (i integer) gt(i2 integer) bool { return (&(i.i)).Cmp(&(i2.i)) == 1 }

func (i integer) lt(i2 integer) bool { return (&(i.i)).Cmp(&(i2.i)) == -1 }

func (i integer) add(i2 integer) integer { return integer{*new(big.Int).Add(&(i.i), &(i2.i))} }

func (i integer) sub(i2 integer) integer { return integer{*new(big.Int).Sub(&(i.i), &(i2.i))} }

func (i integer) mul(i2 integer) integer { return integer{*new(big.Int).Mul(&(i.i), &(i2.i))} }

func (i integer) div(i2 integer) integer { return integer{*new(big.Int).Div(&(i.i), &(i2.i))} }

func (i integer) mod(i2 integer) integer { return integer{*new(big.Int).Mod(&(i.i), &(i2.i))} }

func (i integer) neg() integer { return integer{*new(big.Int).Neg((&(i.i)))} }

func (i integer) String() string { return (&(i.i)).String() }

func (i integer) bitLen() int { return (&(i.i)).BitLen() }

// MarshalAmino for custom encoding scheme
func (i integer) MarshalAmino() (string, error) {
	bz, err := (&(i.i)).MarshalText()
	return string(bz), err
}

// UnmarshalAmino for custom decoding scheme
func (i *integer) UnmarshalAmino(text string) (err error) {
	tempInt := new(big.Int)
	err = tempInt.UnmarshalText([]byte(text))
	if err != nil {
		return
	}
	i.i = *tempInt
	return nil
}

// MarshalJSON for custom encodig scheme
func (i integer) MarshalJSON() ([]byte, error) {
	return (&(i.i)).MarshalText()
}

// UnmarshalJSON for custom decoding scheme
func (i *integer) UnmarshalJSON(bz []byte) (err error) {
	tempInt := new(big.Int)
	err = tempInt.UnmarshalText([]byte(bz))
	if err != nil {
		return
	}
	i.i = *tempInt
	return nil
}

// Int wraps integer with 256 bit range bound
// Checks overflow, underflow and division by zero
// Exists in range from -(2^255-1) to 2^255-1
type Int struct {
	i integer
}

// BigInt converts Int to big.Int
func (i Int) BigInt() *big.Int {
	return i.i.bigInt()
}

// NewInt constructs Int from int64
func NewInt(n int64) Int {
	return Int{newInteger(n)}
}

// NewIntFromBigInt constructs Int from big.Int
func NewIntFromBigInt(i *big.Int) Int {
	return Int{newIntegerFromBigInt(i)}
}

// NewIntFromString constructs Int from string
func NewIntFromString(s string) (res Int, ok bool) {
	i, ok := newIntegerFromString(s)
	if !ok {
		return
	}
	// Check overflow
	if i.bitLen() > 255 {
		return
	}
	return Int{i}, true
}

// NewIntWithDecimal constructs Int with decimal
// Result value is n*10^dec
func NewIntWithDecimal(n int64, dec int) (res Int, ok bool) {
	i := newIntegerWithDecimal(n, dec)
	// Check overflow
	if i.bitLen() > 255 {
		return
	}
	return Int{i}, true
}

// ZeroInt returns Int value with zero
func ZeroInt() Int { return Int{zeroInt()} }

// OneInt returns Int value with one
func OneInt() Int { return Int{oneInt()} }

// Int64 converts Int to int64
// Panics if the value is out of range
func (i Int) Int64() int64 {
	return i.i.toInt64()
}

// IsZero returns true if Int is zero
func (i Int) IsZero() bool {
	return i.i.sign() == 0
}

// Sign returns sign of Int
func (i Int) Sign() int {
	return i.i.sign()
}

// Equal compares two Ints
func (i Int) Equal(i2 Int) bool {
	return i.i.equal(i2.i)
}

// GT returns true if first Int is greater than second
func (i Int) GT(i2 Int) bool {
	return i.i.gt(i2.i)
}

// LT returns true if first Int is lesser than second
func (i Int) LT(i2 Int) bool {
	return i.i.lt(i2.i)
}

// Add adds Int from another
func (i Int) Add(i2 Int) (res Int) {
	res = Int{i.i.add(i2.i)}
	// Check overflow
	if res.i.bitLen() > 255 {
		panic("Integer overflow")
	}
	return
}

// AddRaw adds int64 to Int
func (i Int) AddRaw(i2 int64) Int {
	return i.Add(NewInt(i2))
}

// Sub subtracts Int from another
func (i Int) Sub(i2 Int) (res Int) {
	res = Int{i.i.sub(i2.i)}
	// Check overflow
	if res.i.bitLen() > 255 {
		panic("Integer overflow")
	}
	return
}

// SubRaw subtracts int64 from Int
func (i Int) SubRaw(i2 int64) Int {
	return i.Sub(NewInt(i2))
}

// Mul multiples two Ints
func (i Int) Mul(i2 Int) (res Int) {
	// Check overflow
	if i.i.bitLen()+i2.i.bitLen() > 255 {
		panic("integer overflow")
	}
	res = Int{i.i.mul(i2.i)}
	// Check overflow if sign of both are same
	if res.i.bitLen() > 255 {
		panic("Integer overflow")
	}
	return
}

// MulRaw multipies Int and int64
func (i Int) MulRaw(i2 int64) Int {
	return i.Mul(NewInt(i2))
}

// Div divides Int with Int
func (i Int) Div(i2 Int) (res Int) {
	// Check division-by-zero
	if i2.i.sign() == 0 {
		panic("Division by zero")
	}
	return Int{i.i.div(i2.i)}
}

// DivRaw divides Int with int64
func (i Int) DivRaw(i2 int64) Int {
	return i.Div(NewInt(i2))
}

// Neg negates Int
func (i Int) Neg() (res Int) {
	return Int{i.i.neg()}
}

// MarshalAmino defines custom encoding scheme
func (i Int) MarshalAmino() (string, error) {
	return i.i.MarshalAmino()
}

// UnmarshalAmino defines custom decoding scheme
func (i *Int) UnmarshalAmino(text string) error {
	return (&(i.i)).UnmarshalAmino(text)
}

// MarshalJSON defines custom encoding scheme
func (i Int) MarshalJSON() ([]byte, error) {
	return i.i.MarshalJSON()
}

// UnmarshalJSON defines custom decoding scheme
func (i *Int) UnmarshalJSON(bz []byte) error {
	return (&(i.i)).UnmarshalJSON(bz)
}

// Int wraps integer with 256 bit range bound
// Checks overflow, underflow and division by zero
// Exists in range from to 2^256-1
type Uint struct {
	i integer
}

// BigInt converts Uint to big.Unt
func (i Uint) BigInt() *big.Int {
	return i.i.bigInt()
}

// NewUint constructs Uint from int64
func NewUint(n int64) Uint {
	return Uint{newInteger(n)}
}

// NewUintFromBigUint constructs Uint from big.Uint
func NewUintFromBigInt(i *big.Int) Uint {
	return Uint{newIntegerFromBigInt(i)}
}

// NewUintFromString constructs Uint from string
func NewUintFromString(s string) (res Uint, ok bool) {
	i, ok := newIntegerFromString(s)
	if !ok {
		return
	}
	// Check overflow
	if i.bitLen() > 256 {
		return
	}
	return Uint{i}, true
}

// NewUintWithDecimal constructs Uint with decimal
// Result value is n*10^dec
func NewUintWithDecimal(n int64, dec int) (res Uint, ok bool) {
	i := newIntegerWithDecimal(n, dec)
	// Check overflow
	if i.bitLen() > 256 {
		return
	}
	return Uint{i}, true
}

// ZeroUint returns Uint value with zero
func ZeroUint() Uint { return Uint{zeroInt()} }

// OneUint returns Uint value with one
func OneUint() Uint { return Uint{oneInt()} }

// Uint64 converts Uint to uint64
// Panics if the value is out of range
func (i Uint) Uint64() uint64 {
	return i.i.toUint64()
}

// IsZero returns true if Uint is zero
func (i Uint) IsZero() bool {
	return i.i.isZero()
}

// Sign returns sign of Uint
func (i Uint) Sign() int {
	return i.i.sign()
}

// Equal compares two Uints
func (i Uint) Equal(i2 Uint) bool {
	return i.i.equal(i2.i)
}

// GT returns true if first Uint is greater than second
func (i Uint) GT(i2 Uint) bool {
	return i.i.gt(i2.i)
}

// LT returns true if first Uint is lesser than second
func (i Uint) LT(i2 Uint) bool {
	return i.i.lt(i2.i)
}

// Add adds Uint from another
func (i Uint) Add(i2 Uint) (res Uint) {
	res = Uint{i.i.add(i2.i)}
	// Check overflow
	if res.Sign() == -1 || res.Sign() == 1 && res.i.bitLen() > 256 {
		panic("Uinteger overflow")
	}
	return
}

// AddRaw adds int64 to Uint
func (i Uint) AddRaw(i2 int64) Uint {
	return i.Add(NewUint(i2))
}

// Sub subtracts Uint from another
func (i Uint) Sub(i2 Uint) (res Uint) {
	res = Uint{i.i.sub(i2.i)}
	// Check overflow
	if res.Sign() == -1 || res.Sign() == 1 && res.i.bitLen() > 256 {
		panic("Uinteger overflow")
	}
	return
}

// SubRaw subtracts int64 from Uint
func (i Uint) SubRaw(i2 int64) Uint {
	return i.Sub(NewUint(i2))
}

// Mul multiples two Uints
func (i Uint) Mul(i2 Uint) (res Uint) {
	// Check overflow
	if i.i.bitLen()+i2.i.bitLen() > 256 {
		panic("integer overflow")
	}
	res = Uint{i.i.mul(i2.i)}
	// Check overflow
	if res.Sign() == -1 || res.Sign() == 1 && res.i.bitLen() > 256 {
		panic("Uinteger overflow")
	}
	return
}

// MulRaw multipies Uint and int64
func (i Uint) MulRaw(i2 int64) Uint {
	return i.Mul(NewUint(i2))
}

// Div divides Uint with Uint
func (i Uint) Div(i2 Uint) (res Uint) {
	// Check division-by-zero
	if i2.Sign() == 0 {
		return
	}
	return Uint{i.i.div(i2.i)}
}

// Div divides Uint with int64
func (i Uint) DivRaw(i2 int64) Uint {
	return i.Div(NewUint(i2))
}

// MarshalAmino defines custom encoding scheme
func (i Uint) MarshalAmino() (string, error) {
	return i.i.MarshalAmino()
}

// UnmarshalAmino defines custom decoding scheme
func (i *Uint) UnmarshalAmino(text string) (err error) {
	return (&(i.i)).UnmarshalAmino(text)
}

// MarshalJSON defines custom encoding scheme
func (i Uint) MarshalJSON() ([]byte, error) {
	return i.i.MarshalJSON()
}

// UnmarshalJSON defines custom decoding scheme
func (i *Uint) UnmarshalJSON(bz []byte) error {
	return (&(i.i)).UnmarshalJSON(bz)
}
