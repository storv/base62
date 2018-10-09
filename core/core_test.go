package core

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)


func TestGetBase62Instance(t *testing.T) {
	base62 := GetBase62Instance("")
	Convey("Base字符串",t,func(){
		So(base62.Decode(base62.Encode("0a7812af-a426-4403-95b9-7341c454a1a5")),ShouldEqual,"0a7812af-a426-4403-95b9-7341c454a1a5")
	})
}

func TestBase62_Encode(t *testing.T) {
	base62 := GetBase62Instance("")
	Convey("Base字符串",t,func(){
		So(base62.Encode("abcd"),ShouldEqual,"6hgMUQ")
	})

}

func TestBase62_Decode(t *testing.T) {
	base62 := GetBase62Instance("")
	Convey("Base字符串",t,func(){
		So(base62.Decode("6hgMUQ"),ShouldEqual,"abcd")
	})
}

func TestTranslateStringToByteArray(t *testing.T){

	Convey("测试转换函数",t,func(){
		message := "abcd"
		result := []byte{97,98,99,100}
		//FIXME
		So(TranslateStringToByteArray(message),ShouldEqual,result)
	})

}

func TestTranslateByteArrayToString(t *testing.T) {
	Convey("测试转换函数",t,func(){
		result:= "abcd"
		message := []byte{97,98,99,100}
		//FIXME
		So(TranslateByteArrayToString(message),ShouldEqual,result)
	})
}

func TestConvert(t *testing.T) {
	Convey("测试转换函数",t,func(){
		message := []byte{2,0,0}
		result := []byte{12,8}
		//FIXME
		So(Convert(message,10,16)[:],ShouldEqual,result[:])
	})
}


func SliceEqualBCE(a, b []interface{}) bool {
	if len(a) != len(b) {
		return false
	}

	if (a == nil) != (b == nil) {
		return false
	}

	b = b[:len(a)]
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}

	return true
}
