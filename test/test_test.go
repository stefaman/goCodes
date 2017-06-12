package test

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"reflect"
	"testing"
)

func Test(t *testing.T) {
	x := 2
	xp := reflect.ValueOf(&x).Elem()
	fmt.Println(xp.CanAddr(), xp.CanSet())
	xd := xp.Addr().Interface().(*int)
	*xd = 3
	fmt.Println(x)

	xp.Set(reflect.ValueOf(4))
	fmt.Println(x)

	xp.SetInt(5)
	fmt.Println(x)

	e := []int{1, 2, 3}
	ev := reflect.ValueOf(e).Index(2)
	fmt.Println(ev.CanAddr(), ev.CanSet())
	ev.SetInt(55)
	fmt.Println(e)

	// xp.Set(reflect.ValueOf(int64(4)))//panic: value of type int64 is not assignable to type int

	// x = 2
	// b := reflect.ValueOf(x)
	// b.Set(reflect.ValueOf(3)) // panic: Set using unaddressable value

	var y interface{}
	ry := reflect.ValueOf(&y).Elem()
	// ry.SetInt(2)                     // panic: SetInt called on interface Value
	ry.Set(reflect.ValueOf(3)) // OK, y = int(3)
	// ry.SetString("hello")            // panic: SetString called on interface Value
	ry.Set(reflect.ValueOf("hello")) // OK, y = "hello"

	stdout := reflect.ValueOf(os.Stdout).Elem() // *os.Stdout, an os.File var
	fmt.Println(stdout.Type())                  // "os.File"
	fd := stdout.FieldByName("fd")
	fmt.Println(fd.CanAddr(), fd.CanSet())               //true false
	fmt.Println(fd.Kind() == reflect.Uintptr, fd.Uint()) //windows:ture, XXX
	// fmt.Println(fd.Kind() == reflect.Int, fd.Int()) //linux:true, 1
	// fd.SetInt(2)  // panic: using value obtained using unexported field
	// fmt.Println(fd.Interface()) //cannot return value obtained from unexported field or method
	fmt.Println(fd.CanInterface()) //false

	var _, ok = (interface{})((*bytes.Buffer)(nil)).(io.Writer)
	t.Log(ok)

	type Com struct {
		io.Writer
	}
	type C Com
	c := C{Writer: os.Stdout}
	c.Write([]byte("蔡"))

}

// fmt.Printf("%x\n", []byte("蔡国准"))
// fmt.Printf("%x\n", []rune("蔡国准"))
// // var f = 1e17
// var i = 4300000000000000000000000000
// fmt.Println(i, reflect.TypeOf(i))
//test reflect

// var w io.Writer
// var e interface{Close() error}
// e = w.(io.Closer)
// // w = e.(io.Writer)
// w.Write([]byte("kkk"))
/*
s := "kkk"
v := reflect.ValueOf(s)
fmt.Printf("%T: %v, %s\n", v, v, v.String())
fmt.Printf("%s\n", reflect.TypeOf(reflect.TypeOf(1)))
fmt.Printf("%s\n", reflect.TypeOf(reflect.ValueOf(1)))
fmt.Printf("%s\n", reflect.TypeOf(time.Second))
fmt.Printf("%s\n", reflect.TypeOf(byte(1)))
fmt.Printf("%s\n", reflect.TypeOf(rune(1)))
//Output
// *reflect.rtype
// reflect.Value
// time.Duration
// uint8
// int32
// *rand.rngSource
fmt.Printf("%s\n", reflect.TypeOf(rand.NewSource(1)))
i := 333
fmt.Printf("%s, %[1]v\n", reflect.TypeOf(i))
val := reflect.ValueOf(i)
fmt.Printf("%T: %v, %s\n", val, val, val)

*/
//test goroutine
/*
go func(){
	fmt.Println("routine begin")
	c := make(chan int)
	go func(){
		// for {
		// }
		fmt.Println("begin, may be stuck in waiting")
		time.Sleep(3*time.Second)
		c <- 1
		fmt.Println("over")
	}()
	time.Sleep(time.Second)
	if true {
		fmt.Println("routine over ahead")
		return
	}
	<- c
	fmt.Println("routine over normal")
}()

time.Sleep(5*time.Second)
*/

// scanner := new(scanner.Scanner)
// scanner.Line = 1;
// fmt.Println(scanner.IsValid())
/*
fmt.Printf("%T: %[1]v\n", "abcd"[1:])
expr := "1 + +x * 3 - -4 + f(1,2,3)"
expr = "x"
exp, err := eval.Parse(expr)
fmt.Printf("%#v %v\n", exp, err)
fmt.Println(exp)


switch exp.(type) {
case eval.Var:
	fmt.Printf("%T\n", exp)
case eval.Expr:
	fmt.Printf("%T\n", exp)
}
*/

// fmt.Println(KiB, MiB, GiB, TiB, PiB, EiB)
//
// a := [32]int{3:99}
// fmt.Printf("%v\n%#[1]x\n",&a)
// fmt.Println(*zeroA32(&a))
//
// s := []int{1,2,3,4,5,6}
// copy(s[2:], s)
// fmt.Println(s)

//
/*
type father struct{
	blood int
	weight int
}
type mother struct{
	weight int
}
type child struct{
	// blood int
	father
	mother
}

// c := child{
// 	blood: 2
// 	weight
// }

var c child
c = child{
	// blood:2
	mother: mother{
		weight: 50,
	},
	father: father{
		weight:60,
		blood: 3,
	},

}
c.blood = 2
c.mother.weight = 52
fmt.Printf("child type is %T\v\n%#[1]v\n%[1]v\n", c)
*/

// fmt.Printf("\ntest path")
// fmt.Printf("path.Base()\n")

/*
for i := 0; i < 5; i++ {
	defer func(i int){ test(i)}(i)
}

test :=func(i int) (ret int){
	defer func() {
		if v := recover(); v == i {
			ret = i
		}else{
			panic(v)
		}
	}()

	panic(i)
}

fmt.Println(test(33))
*/
// fmt.Println(-5/3, -5%3)

//test for struct and method

// var s *intset.IntSet //bad coding, s is not initial pointer
/*
type T struct{
	 IntSet *intset.IntSet
	padding int
}
s := new(T) //ok
s.IntSet = new(intset.IntSet)

fmt.Printf("%s", "\u00b0C")

// fmt.Println((*intset.IntSet)(nil).Has(1)) //error, nil struct poiner is dangerous
fmt.Println((&intset.IntSet{}).Has(1)) //ok
*/

/*
s.IntSet.AddAll(1,2,3,4,5,6,7,8)
s.IntSet = s.IntSet.Copy()
// s.Clear()
s.IntSet.Add(1)
fmt.Println(s)

type Func func(x, y int) bool
funcs := map[string]Func {
	"byArtist" : func(x, y int) bool { return x>y },

}

fmt.Println(funcs)
*/
/*
func zeroA32(a *[32]int) *[32]int {
	for i := range a {
		a[i] = 0
	}
	return a
}

func test(i int) {
	fmt.Println(i)
}
*/
