package cache

import (
	"hash/fnv"
	"math"
	"reflect"
)

// Hash is an interface implemented by cache keys to
// override default hash function.
type Hash interface {
	Sum64() uint64
}

// sum calculates hash value of the given key.
func sum(k interface{}) uint64 {
	if h, ok := k.(Hash); ok {
		return h.Sum64()
	}
	switch h := k.(type) {
	case int:
		return uint64(h)
	case int8:
		return uint64(h)
	case int16:
		return uint64(h)
	case int32:
		return uint64(h)
	case int64:
		return uint64(h)
	case uint:
		return uint64(h)
	case uint8:
		return uint64(h)
	case uint16:
		return uint64(h)
	case uint32:
		return uint64(h)
	case uint64:
		return h
	case uintptr:
		return uint64(h)
	case float32:
		return uint64(math.Float32bits(h))
	case float64:
		return math.Float64bits(h)
	case bool:
		if h {
			return 1
		}
		return 0
	case string:
		return hashBytes([]byte(h))
	}
	// TODO: complex64 and complex128
	if h, ok := hashPointer(k); ok {
		return h
	}
	// TODO: use gob to encode k to bytes then hash.
	return 0
}

func hashBytes(b []byte) uint64 {
	h := fnv.New64a()
	_, _ = h.Write(b)
	return h.Sum64()
}

func hashPointer(k interface{}) (uint64, bool) {
	v := reflect.ValueOf(k)
	switch v.Kind() {
	case reflect.Ptr, reflect.UnsafePointer, reflect.Func, reflect.Slice, reflect.Map, reflect.Chan:
		return uint64(v.Pointer()), true
	}
	return 0, false
}