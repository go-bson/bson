package bson_test

import (
	"fmt"
	"reflect"
	"testing"

	"gopkg.in/bson.v2"
)

type innerObj struct {
	Field1 string `bson:"f1"`
	Field2 int    `bson:"f2"`
}

type outObj struct {
	Field1 string      `bson:"f1"`
	Inner  interface{} `bson:"inner"`
}

func TestPredefined(t *testing.T) {
	out := &outObj{
		Field1: "out1",
		Inner: &innerObj{
			Field1: "inner1",
			Field2: 1,
		},
	}
	// var val interface{} = out
	fmt.Printf("->%v\n", reflect.ValueOf(out.Inner).Type())
	bys, _ := bson.Marshal(out)
	out2 := &outObj{
		Inner: &innerObj{},
	}
	err := bson.Unmarshal(bys, out2)
	if err != nil {
		t.Error("error")
		return
	}
	if inner, ok := out2.Inner.(*innerObj); ok && out2.Field1 == out.Field1 && inner.Field1 == "inner1" && inner.Field2 == 1 {
		fmt.Println(inner)
	} else {
		t.Errorf("inner:%v", reflect.TypeOf(out2.Inner))
	}
}
