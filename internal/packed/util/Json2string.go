package util

// RemoveBackslashesBytes 修改JSON字符串中的反斜杠
func RemoveBackslashesBytes(b []byte) []byte {
	result := make([]byte, 0, len(b))
	for _, c := range b {
		if c != '\\' {
			result = append(result, c)
		}
	}
	return result
}
