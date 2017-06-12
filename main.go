package main

import (
	"fmt"
	"os"
	"reflect"
)

func main() {
	// var t  testing.T
	stdout := reflect.ValueOf(os.Stdout).Elem() // *os.Stdout, an os.File var
	fmt.Println(stdout.Type())                  // "os.File"
	fmt.Println(stdout, stdout.CanAddr(), stdout.CanSet(), stdout.CanInterface())
	stdout.Set(reflect.ValueOf(*os.Stderr))
	stdout.Set(reflect.ValueOf(*os.Stdin)) //设置为*os.Stdin test无法输出???
	fd := stdout.FieldByName("fd")
	fmt.Println(fd.CanAddr(), fd.CanSet(), fd.CanInterface()) //true false false
	fmt.Println(fd.Kind() == reflect.Uintptr, fd.Uint())      //windows:ture, XXX
	// fmt.Println(fd.Kind() == reflect.Int, fd.Int()) //linux:true, 1
	// fd.SetInt(2) // panic: using value obtained using unexported field
	// fmt.Println(fd.Interface()) //cannot return value obtained from unexported field or method
}
