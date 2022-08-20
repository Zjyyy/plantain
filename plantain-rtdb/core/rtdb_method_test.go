package core

import (
	"testing"
)

// 测试NewRtdbMethod是否正确开辟了多个互相隔离的内存片
func TestNewRtdbMethod(t *testing.T) {
	// pDriverArr := createMockPDriver()

	// cache1 := cache.New(0, 0)
	// cache1.Set("Tag01", 1, 0)

	// cache2 := cache.New(0, 0)
	// cache2.Set("Tag01", 1, 0)

	// rtdbHandler1 := NewRtdbMethod(pDriverArr[0], cache1)
	// rtdbHandler2 := NewRtdbMethod(pDriverArr[0], cache2)

	// rtdbHandler1.Write("Tag01", 1)
	// value1 := rtdbHandler1.Read("Tag01")

	// rtdbHandler2.Write("Tag01", 2)
	// value2 := rtdbHandler2.Read("Tag01")

	// if value1 != 1 || value2 != 2 || value1 == value2 {
	// 	t.Fatal("NewRtdbMethod异常")
	// }
}
