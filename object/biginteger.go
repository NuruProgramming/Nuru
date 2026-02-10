package object

import (
	"hash/fnv"
	"math/big"
)

type BigInteger struct {
	Value *big.Int
	BaseRefCountable
}

func (b *BigInteger) Type() ObjectType { return BIG_INTEGER_OBJ }

func (b *BigInteger) Inspect() string {
	if b.Value == nil {
		return "0"
	}
	return b.Value.String()
}

func (b *BigInteger) HashKey() HashKey {
	if b.Value == nil {
		return HashKey{Type: b.Type(), Value: 0}
	}
	h := fnv.New64a()
	h.Write([]byte(b.Value.String()))
	return HashKey{Type: b.Type(), Value: h.Sum64()}
}

func (b *BigInteger) GetOutgoingReferences() []Object { return nil }

// NewBigInteger creates a BigInteger from a string (decimal) or from an int64.
func NewBigIntegerFromString(s string) (*BigInteger, bool) {
	z := new(big.Int)
	_, ok := z.SetString(s, 10)
	if !ok {
		return nil, false
	}
	return &BigInteger{Value: z}, true
}

func NewBigIntegerFromInt64(n int64) *BigInteger {
	return &BigInteger{Value: big.NewInt(n)}
}

// Int64 returns the int64 value if it fits; otherwise ok is false.
func (b *BigInteger) Int64() (v int64, ok bool) {
	if b.Value == nil {
		return 0, true
	}
	return b.Value.Int64(), b.Value.IsInt64()
}
