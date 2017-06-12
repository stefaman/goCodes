// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 333.

// Package display provides a means to displayC structured data.
package display

import (
	"fmt"
	"reflect"
	"unsafe"
)

//record seen {type, address}
type see struct {
	typ  reflect.Type
	addr unsafe.Pointer
}

//DisplayC display any variable, including circle infinite type(except circle map)
//显示全部，比json.Marshal 看得更多
func DisplayC(name string, x interface{}) {
	seen := make(map[see]bool)
	fmt.Printf("Display %s (%T):\n", name, x)
	displayC(name, reflect.ValueOf(x), seen)
}

//displayC 递归的显示内部，因为参数的名字会丢失，所以需要单独传入变量的identifier
//另一方面，单独传入name也是递归的path
func displayC(path string, v reflect.Value, seen map[see]bool) {
	if v.Kind() == reflect.Invalid {
		fmt.Printf("%s = invalid\n", path)
		return
	}
	if v.CanAddr() {
		c := see{typ: v.Type(),
			addr: unsafe.Pointer(v.UnsafeAddr())}
		if seen[c] {
			fmt.Printf("%s = (circle pointer %#x)\n", path, c.addr)
			return
		}
		seen[c] = true
	}
	//call v.String()
	if f := v.MethodByName("String"); f.Kind() != reflect.Invalid {
		fmt.Printf("%s = %s\n", path, f.Call(nil))
	}
	switch v.Kind() {
	case reflect.Slice:
		if v.IsNil() {
			fmt.Printf("%s = nil\n", path)
			return
		}
		//[]byte printed as string
		if v.Type().Elem() == (reflect.TypeOf(uint8(0))) {
			fmt.Printf("%s = %q\n", path, v.Convert(reflect.TypeOf("")))
			return
		}
		fallthrough
	case reflect.Array: //if nil or empty
		if v.Len() == 0 {
			fmt.Printf("%s = []\n", path)
		}
		for i := 0; i < v.Len(); i++ {
			displayC(fmt.Sprintf("%s[%d]", path, i), v.Index(i), seen)
		}
	case reflect.Struct:
		if v.NumField() == 0 { //显示empty struct
			fmt.Printf("%s = {}\n", path)
			return
		}
		for i := 0; i < v.NumField(); i++ { //内嵌成员没有区别对待
			fieldPath := fmt.Sprintf("%s.%s", path, v.Type().Field(i).Name)
			displayC(fieldPath, v.Field(i), seen)
		}
	case reflect.Map:
		if v.IsNil() { // if nil
			fmt.Printf("%s = nil\n", path)
			return
		}
		if v.Len() == 0 { //if empty
			fmt.Printf("%s = [][]\n", path)
			return
		}
		for _, key := range v.MapKeys() {
			displayC(fmt.Sprintf("%s[%s]", path,
				fmtKey(key)), v.MapIndex(key), seen)
		}

	case reflect.Ptr:
		if v.IsNil() {
			fmt.Printf("%s = nil\n", path)
		} else {
			displayC(fmt.Sprintf("(*%s)", path), v.Elem(), seen)
		}
	case reflect.Interface:
		if v.IsNil() {
			fmt.Printf("%s = nil\n", path)
		} else {
			fmt.Printf("%s.type = %s\n", path, v.Elem().Type())
			displayC(path+".value", v.Elem(), seen)
		}
	default: // basic types, channels, funcs
		fmt.Printf("%s = %s\n", path, formatAtom(v))
	}
}

//!-displayC
