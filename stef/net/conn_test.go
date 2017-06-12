package mynet

import (
	"fmt"
	"net"
	"os"
	"testing"
)

// func TestConnFile(t *testing.T) {
// 	f, err := os.Create(os.TempDir() + "tmp.txt")
// 	if err != nil {
// 		log.Fatalln(err)
// 	}
// 	defer f.Close()
// 	conn, err := net.FileConn(f)
// 	if err != nil {
// 		log.Fatalln(err)
// 	}
// 	defer conn.Close()
// 	f.WriteString("cai")
// 	buf := make([]byte, 0, 200)
// 	n, err := f.Read(buf)
// 	if err != nil {
// 		log.Fatalln(err)
// 	}
// 	fmt.Println(string(buf[:n]))
// }

func TestNames(t *testing.T) {
	host := "baidu.com"
	cname, err := net.LookupCNAME(host)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	fmt.Println(cname)
	hosts, err := net.LookupHost(host)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	fmt.Println(hosts)

	ips, err := net.LookupHost(host)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	fmt.Println(ips)

	for _, addr := range hosts {
		names, err := net.LookupAddr(addr)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
		fmt.Println(addr, names)
	}

}
