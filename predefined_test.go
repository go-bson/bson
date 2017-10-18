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

type outObjSlice struct {
	Field1 string `bson:"f1"`
	Inner  []int  `bson:"inner"`
}

type outObjSlice2 struct {
	Field1 string `bson:"f1"`
	Inner  *[]int `bson:"inner"`
}

type outObjArray struct {
	Field1 string `bson:"f1"`
	Inner  [3]int `bson:"inner"`
}

type outObjArray2 struct {
	Field1 string  `bson:"f1"`
	Inner  *[3]int `bson:"inner"`
}

func TestPredefinedObj(t *testing.T) {
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

func TestPredefinedSlice(t *testing.T) {
	out := &outObj{
		Field1: "out1",
		Inner:  []int{1, 2, 3},
	}
	// var val interface{} = out
	fmt.Printf("->%v\n", reflect.ValueOf(out.Inner).Type())
	bys, _ := bson.Marshal(out)
	{
		//
		//test predefined slice
		//
		out3 := &outObjSlice{}
		err := bson.Unmarshal(bys, out3)
		if err != nil {
			t.Error(err)
			return
		}
		if len(out3.Inner) == 3 && out3.Inner[0] == 1 {
			fmt.Println("->test slice passed")
		} else {
			t.Errorf("inner:%v", reflect.TypeOf(out3.Inner))
		}
		//
		out2 := &outObj{
			Inner: []int{},
		}
		err = bson.Unmarshal(bys, out2)
		if err != nil {
			t.Error(err)
			return
		}
		if inner, ok := out2.Inner.([]int); ok && len(inner) == 3 && (inner)[0] == 1 {
			fmt.Println("->test predefined slice passed")
		} else {
			t.Errorf("inner:%v", reflect.TypeOf(out2.Inner))
		}
		//
		out4 := &outObjSlice2{}
		err = bson.Unmarshal(bys, out4)
		if err != nil {
			t.Error(err)
			return
		}
		if len(*out4.Inner) == 3 && (*out4.Inner)[0] == 1 {
			fmt.Println("->test slice pointer passed")
		} else {
			t.Errorf("inner:%v", reflect.TypeOf(out4.Inner))
		}
		//
		out2 = &outObj{
			Inner: &[]int{},
		}
		err = bson.Unmarshal(bys, out2)
		if err != nil {
			t.Error(err)
			return
		}
		if inner, ok := out2.Inner.(*[]int); ok && len(*inner) == 3 && (*inner)[0] == 1 {
			fmt.Println("->test predefined slice pointer passed")
		} else {
			t.Errorf("inner:%v", reflect.TypeOf(out2.Inner))
		}
	}
	{
		//
		//test predefined array
		//
		outArray := &outObjArray{}
		err := bson.Unmarshal(bys, outArray)
		if err != nil {
			t.Error(err)
			return
		}
		if inner := outArray.Inner; len(inner) == 3 && (inner)[0] == 1 {
			fmt.Println("->test array passed")
		} else {
			t.Errorf("inner:%v", reflect.TypeOf(outArray.Inner))
		}
		//
		out2 := &outObj{
			Inner: [3]int{},
		}
		err = bson.Unmarshal(bys, out2)
		if err != nil {
			t.Error(err)
			return
		}
		if inner, ok := out2.Inner.([3]int); ok && len(inner) == 3 && inner[0] == 1 {
			fmt.Println("->test predefined array pointer passed")
		} else {
			t.Errorf("inner:%v", reflect.TypeOf(out2.Inner))
		}
		//
		outArray2 := &outObjArray2{
			Inner: &[3]int{},
		}
		err = bson.Unmarshal(bys, outArray2)
		if err != nil {
			t.Error(err)
			return
		}
		if inner := outArray2.Inner; len(inner) == 3 && (inner)[0] == 1 {
			fmt.Println("->test array pointer passed")
		} else {
			t.Errorf("inner:%v", reflect.TypeOf(out2.Inner))
		}
		//
		out4 := &outObj{
			Inner: &[3]int{},
		}
		err = bson.Unmarshal(bys, out4)
		if err != nil {
			t.Error(err)
			return
		}
		fmt.Println(out4.Inner)
		if inner, ok := out4.Inner.(*[3]int); ok && len(inner) == 3 && inner[0] == 1 {
			fmt.Println("->test predefined array pointer passed")
		} else {
			t.Errorf("inner:%v", reflect.TypeOf(out4.Inner))
		}
	}
}
