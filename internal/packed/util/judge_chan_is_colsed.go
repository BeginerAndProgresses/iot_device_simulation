package util

// IsChanIntClose 判断int通道是否关闭
func IsChanIntClose(ch chan int) bool {
	select {
	case _, received := <-ch:
		return !received
	default:
	}
	return false
}

// IsChanStringClose 判断string通道是否关闭
func IsChanStringClose(ch chan string) bool {
	select {
	case _, received := <-ch:
		return !received
	default:
	}
	return false
}
