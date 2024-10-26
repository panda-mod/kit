package stringx

import (
	"strings"
	"unicode"
)

// Words 将字符串转换为单词数组
func Words(s string) []string {
	s = splitWithUnicode(s)
	var builder strings.Builder
	for _, r := range s {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			builder.WriteRune(r)
		} else {
			builder.WriteRune(' ')
		}
	}
	return strings.Fields(builder.String())
}

// KebabCase 返回字符串转换为 kebab-case 格式
func KebabCase(s string) string {
	items := Words(s)
	for i := range items {
		items[i] = strings.ToLower(items[i])
	}
	return strings.Join(items, "-")
}

// SnakeCase 返回字符串在转换为 snake_case 格式
func SnakeCase(s string) string {
	items := Words(s)
	for i := range items {
		items[i] = strings.ToLower(items[i])
	}
	return strings.Join(items, "_")
}

// SplitWithUnicode 按照 Unicode 标准分割字符串
func splitWithUnicode(s string) string {
	if s == "" {
		return s
	}
	var builder strings.Builder
	builder.Grow(len(s) * 2)

	runes := []rune(s)
	lastPos := 0

	for i := 1; i < len(runes); i++ {
		curr := runes[i]
		prev := runes[i-1]

		// 检查是否需要插入分隔符
		needSplit := false

		switch {
		// 小写字母后跟大写字母或数字
		case unicode.IsLower(prev) && (unicode.IsUpper(curr) || unicode.IsNumber(curr)):
			needSplit = true
		// 字母后跟数字
		case unicode.IsLetter(prev) && unicode.IsNumber(curr):
			needSplit = true
		// 数字后跟字母
		case unicode.IsNumber(prev) && unicode.IsLetter(curr):
			needSplit = true
		// 大写字母后跟大写字母，且后面还有小写字母
		case i+1 < len(runes) &&
			unicode.IsUpper(prev) &&
			unicode.IsUpper(curr) &&
			unicode.IsLower(runes[i+1]):
			needSplit = true
		}

		if needSplit {
			builder.WriteString(string(runes[lastPos:i]))
			builder.WriteRune(' ')
			lastPos = i
		}
	}

	builder.WriteString(string(runes[lastPos:]))
	return builder.String()
}
