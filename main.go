package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println(time.Now().Add(120*time.Second).UTC().Format("Mon, 02 Jan 2006 15:04:05 GMT"))
fmt.Println(time.Now().Add(120*time.Second).UTC().Format(time.RFC1123))
}
