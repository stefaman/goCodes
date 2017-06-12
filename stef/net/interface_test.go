package mynet

import (
	"fmt"
	"net"
)

// func TestAddr(t *testing.T) {
// 	addrs, err := net.InterfaceAddrs()
// 	if err != nil {
// 		fmt.Fprint(os.Stderr, err)
// 	}
// 	printAddrs(addrs)
//
// 	inters, err := net.Interfaces()
// 	if err != nil {
// 		fmt.Fprintln(os.Stderr, err)
// 	}
// 	for _, inter := range inters {
// 		// j, _ := json.MarshalIndent(inter, "", "  ")
// 		// fmt.Println(string(j))
// 		display.Display("inter", inter)
// 		addrs, err := inter.Addrs()
// 		if err != nil {
// 			fmt.Fprintln(os.Stderr, err)
// 		} else {
// 			fmt.Println("Addrs")
// 			printAddrs(addrs)
// 		}
// 		multicastAddrs, err := inter.MulticastAddrs()
// 		if err != nil {
// 			fmt.Fprintln(os.Stderr, err)
// 		} else {
// 			fmt.Println("MuticastAddrs")
// 			printAddrs(multicastAddrs)
//
// 		}
// 	}
// }

func printAddrs(addrs []net.Addr) {
	// display.Display("addrs", addrs)
	for _, addr := range addrs {
		fmt.Println("Network:", addr.Network(), "String:", addr.String())
	}
}
