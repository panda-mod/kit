package convert

import (
	"strconv"

	"github.com/panda-mod/kit/slice"
)

func Float32(val any, def ...float32) float32 {
	if ret, err := Float64E(val); err == nil {
		return float32(ret)
	}
	return slice.First(def)
}

func Float64(val any, def ...float64) float64 {
	if ret, err := Float64E(val); err == nil {
		return ret
	}
	return slice.First(def)
}

func Float64E(val any) (float64, error) {
	switch ret := val.(type) {
	case float64:
		return ret, nil
	case int:
		return float64(ret), nil
	case int8:
		return float64(ret), nil
	case int32:
		return float64(ret), nil
	case int64:
		return float64(ret), nil
	case uint:
		return float64(ret), nil
	case uint8:
		return float64(ret), nil
	case uint32:
		return float64(ret), nil
	case uint64:
		return float64(ret), nil
	case float32:
		return float64(ret), nil
	case bool:
		return 0, nil
	case []byte:
		return Float64WithString(string(ret))
	case string:
		return Float64WithString(ret)
	default:
		return 0, buildError(ret, "float64")
	}
}

func Float64WithString(val string) (float64, error) {
	if ret, err := strconv.ParseFloat(val, 64); err == nil {
		return ret, nil
	}
	return 0, buildError(val, "float64")
}
