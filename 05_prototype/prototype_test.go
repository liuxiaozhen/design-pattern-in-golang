package prototype

import (
	"fmt"
	"reflect"
	"testing"
)

type Student struct {
	School string
	Name   string
	Age    int
}

func NewStudent() *Student {
	return &Student{
		School: "中心小学",
		Name:   "",
		Age:    0,
	}
}

func (s *Student) Clone() Cloneable {
	t := *s
	return &t
}

var proto *PrototypeManager

func init() {
	proto = NewPrototypeManager()
	proto.Set("student", NewStudent())
}
func Test_Prototype(t *testing.T) {
	zhang := proto.Get("student") //构造不同的实体
	zhang.(*Student).Name = "张三"
	zhang.(*Student).Age = 10

	wang := proto.Get("student") //构造不同的实体
	wang.(*Student).Name = "王五"
	wang.(*Student).Age = 16
	if zhang == wang || reflect.DeepEqual(zhang, wang) {
		t.Fatalf("equal failed")
	}
	fmt.Println(zhang)
	fmt.Println(wang)
}
