package convert

import (
	"slices"
	"strings"

	"github.com/panda-mod/kit/slice"
)

func BoolE(val any) (bool, error) {
	switch ret := val.(type) {
	case bool:
		return ret, nil
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, float32, float64:
		return val != 0, nil
	case []byte:
		return BoolWithString(string(ret)), nil
	case string:
		return BoolWithString(ret), nil
	case nil:
		return false, nil
	default:
		return false, buildError(ret, "bool")
	}
}

func Bool(val any, def ...bool) bool {
	if ret, err := BoolE(val); err == nil {
		return ret
	}
	return slice.First(def)
}

func BoolWithString(val string) bool {
	return !slices.Contains([]string{"", "0", "false"}, strings.ToLower(val))
}
