package convert

import (
	"strconv"

	"github.com/panda-mod/kit/constraints"
	"github.com/panda-mod/kit/slice"
)

func Int(val any, def ...int) int {
	return intWithSigned(val, def...)
}

func Int8(val any, def ...int8) int8 {
	return intWithSigned(val, def...)
}

func Int16(val any, def ...int16) int16 {
	return intWithSigned(val, def...)
}

func Int32(val any, def ...int32) int32 {
	return intWithSigned(val, def...)
}

func Int64(val any, def ...int64) int64 {
	return intWithSigned(val, def...)
}

func IntE(val any) (int, error) {
	return intEWithSigned[int](val)
}

func Int8E(val any) (int8, error) {
	return intEWithSigned[int8](val)
}

func Int16E(val any) (int16, error) {
	return intEWithSigned[int16](val)
}

func Int32E(val any) (int32, error) {
	return intEWithSigned[int32](val)
}

func Int64E(val any) (int64, error) {
	switch ret := val.(type) {
	case int64:
		return ret, nil
	case int:
		return int64(ret), nil
	case int8:
		return int64(ret), nil
	case int16:
		return int64(ret), nil
	case int32:
		return int64(ret), nil
	case float32:
		return int64(ret), nil
	case float64:
		return int64(ret), nil
	case []byte:
		return Int64WithString(string(ret))
	case string:
		return Int64WithString(ret)
	default:
		return 0, buildError(ret, "int64")
	}
}

func Int64WithString(val string) (int64, error) {
	ret, err := strconv.ParseInt(val, 10, 0)
	if err == nil {
		return ret, nil
	}
	return 0, buildError(ret, "int64")
}

func intWithSigned[T constraints.Signed](val any, def ...T) T {
	if ret, err := intEWithSigned[T](val); err == nil {
		return ret
	}
	return slice.First(def)
}

func intEWithSigned[T constraints.Signed](val any) (T, error) {
	ret, err := Int64E(val)
	if err != nil {
		return 0, err
	}
	return T(ret), nil
}
