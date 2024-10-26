package convert

import (
	"encoding"
	"fmt"
	"strconv"

	"github.com/panda-mod/kit/slice"
	"github.com/panda-mod/kit/zerocopy"
)

func String(val any, def ...string) string {
	if ret, err := StringE(val); err == nil {
		return ret
	}
	return slice.First(def)
}

func StringE(val any) (string, error) {
	switch ret := val.(type) {
	case string:
		return ret, nil
	case []byte:
		return zerocopy.String(ret), nil
	case []rune:
		return string(ret), nil
	case int64, int, int8, int32, uint, uint8, uint32, uint64, float32, float64:
		return fmt.Sprintf("%v", val), nil
	case bool:
		return strconv.FormatBool(ret), nil
	case fmt.Stringer:
		return ret.String(), nil
	case error:
		return ret.Error(), nil
	case encoding.TextMarshaler:
		v, err := ret.MarshalText()
		if err != nil {
			return "", err
		}
		return string(v), nil
	default:
		return "", buildError(ret, "string")
	}
}
