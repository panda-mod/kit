package mapx

import (
	"strconv"
	"strings"
)

// Search 根据路径查找字典中的值
// @param pattern 路径
// @param dict 字典
func Search[K string, V any](pattern string, dict map[K]V) any {
	var (
		node any = dict
		keys     = strings.Split(pattern, ".")
	)
	for i := 0; i < len(keys); i++ {
		key := keys[i]
		switch ret := node.(type) {
		case map[string]V:
			if v, ok := ret[key]; ok {
				node = v
			} else {
				node = nil
			}
		case []any: // 支持索引查找
			if idx, err := strconv.Atoi(key); err == nil && len(ret) > idx {
				node = ret[idx]
			} else {
				node = nil
			}
		}
	}
	return node
}
