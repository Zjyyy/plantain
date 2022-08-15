package core

//测试NewRtdbMethod是否正确开辟了多个互相隔离的内存片
// func TestNewRtdbMethod(t *testing.T) {
// 	rtdbHandler1 := NewRtdbMethod(cache.New(0, 0))
// 	rtdbHandler2 := NewRtdbMethod(cache.New(0, 0))

// 	rtdbHandler1.Write("Tag01", 1)
// 	value1 := rtdbHandler1.Read("Tag01")

// 	rtdbHandler2.Write("Tag01", "2")
// 	value2 := rtdbHandler2.Read("Tag01")

// 	if value1 != 1 || value2 != "2" || value1 == value2 {
// 		t.Fatal("NewRtdbMethod异常")
// 	}
// }

// //测试读写方法能否支持多类型数据
// func TestWriteReadType(t *testing.T) {
// 	rtdbHandler := NewRtdbMethod(cache.New(0, 0))
// 	rtdbHandler.Write("Tag01", 1)
// 	rtdbHandler.Write("Tag02", "1")
// 	rtdbHandler.Write("Tag03", 1.1)

// 	value1 := rtdbHandler.Read("Tag01")
// 	value2 := rtdbHandler.Read("Tag02")
// 	value3 := rtdbHandler.Read("Tag03")
// 	if reflect.TypeOf(value1).Name() != "int" ||
// 		reflect.TypeOf(value2).Name() != "string" ||
// 		reflect.TypeOf(value3).Name() != "float64" {
// 		t.Fatal("RTDBMethod Write Read方法对类型处理异常")
// 	}
// }
