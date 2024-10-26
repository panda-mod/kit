package convert

import (
	"fmt"
	"strconv"

	"github.com/panda-mod/kit/slice"
	"github.com/panda-mod/kit/zerocopy"
)

func BytesE(val any) ([]byte, error) {
	switch ret := val.(type) {
	case []byte:
		return ret, nil
	case *[]byte:
		return *ret, nil
	case string:
		return zerocopy.Bytes(ret), nil
	case *string:
		return zerocopy.Bytes(*ret), nil
	case int64, int, int8, int32, uint, uint8, uint32, uint64, float32, float64:
		return []byte(fmt.Sprintf("%v", val)), nil
	case bool:
		return []byte(strconv.FormatBool(ret)), nil
	default:
		return nil, buildError(ret, "[]byte")
	}
}

func Bytes(val any, def ...[]byte) []byte {
	if ret, err := BytesE(val); err == nil {
		return ret
	}
	return slice.First(def)
}
