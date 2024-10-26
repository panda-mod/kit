package convert

import (
	"strconv"

	"github.com/panda-mod/kit/constraints"
	"github.com/panda-mod/kit/slice"
)

func Uint(val any, def ...uint) uint {
	return uintWithUnsigned(val, def...)
}

func Uint8(val any, def ...uint8) uint8 {
	return uintWithUnsigned(val, def...)
}

func Uint16(val any, def ...uint16) uint16 {
	return uintWithUnsigned(val, def...)
}

func Uint32(val any, def ...uint32) uint32 {
	return uintWithUnsigned(val, def...)
}

func Uint64(val any, def ...uint64) uint64 {
	return uintWithUnsigned(val, def...)
}

func UintE(val any) (uint, error) {
	return uintEWithUnsigned[uint](val)
}

func Uint8E(val any) (uint8, error) {
	return uintEWithUnsigned[uint8](val)
}

func Uint32E(val any) (uint32, error) {
	return uintEWithUnsigned[uint32](val)
}

func Uint64E(val any) (uint64, error) {
	switch ret := val.(type) {
	case uint64:
		return ret, nil
	case int:
		return uint64(ret), nil
	case int8:
		return uint64(ret), nil
	case int16:
		return uint64(ret), nil
	case int32:
		return uint64(ret), nil
	case int64:
		return uint64(ret), nil
	case uint:
		return uint64(ret), nil
	case uint8:
		return uint64(ret), nil
	case uint32:
		return uint64(ret), nil
	case float32:
		return uint64(ret), nil
	case float64:
		return uint64(ret), nil
	case []byte:
		return Uint64EWithString(string(ret))
	case string:
		return Uint64EWithString(ret)
	case bool:
		return 0, nil
	default:
		return 0, buildError(ret, "uint64")
	}
}

func uintWithUnsigned[T constraints.Unsigned](val any, def ...T) T {
	if ret, err := uintEWithUnsigned[T](val); err == nil {
		return ret
	}
	return slice.First(def)
}

func uintEWithUnsigned[T constraints.Unsigned](val any) (T, error) {
	ret, err := Uint64E(val)
	if err != nil {
		return 0, err
	}
	return T(ret), nil
}

func Uint64EWithString(val string) (uint64, error) {
	if ret, err := strconv.ParseFloat(val, 64); err == nil {
		return uint64(ret), nil
	}
	return 0, buildError(val, "uint64")
}
