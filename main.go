package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
)

// HostState keep host
type HostState struct {
	addr string
	isUp int
}

func main() {
	// Check for hosts file
	if len(os.Args) == 1 {
		if _, err := os.Stat("hosts"); os.IsNotExist(err) {
			fmt.Println("hosts file not found, run 'camdif -c' to create it")
			return
		}
	}

	myIP := GetOwnIP()
	fmt.Println("Own IP:", myIP)
	fmt.Println("Scanning network...")
	cut := strings.Split(myIP, ".")
	addr := cut[0] + "." + cut[1] + "." + cut[2] + "."

	// Start scan
	scansz := 255
	var ports = []int{21, 22, 23, 80, 90, 135, 139, 443, 445, 554, 8000, 8022, 8080, 8088}
	ch := make(chan HostState, scansz)
	for i := 0; i < scansz; i++ {
		go func(num int) {
			ip := addr + strconv.Itoa(num)
			out := make(chan int, len(ports))
			for _, p := range ports {
				go ScanPort(ip, p, out)
			}
			host := HostState{addr: ip}
			for j := 0; j < len(ports); j++ {
				res := <-out
				host.isUp |= res
			}
			ch <- host
		}(i + 1)
	}

	// Gathering ip's back
	var pool []string
	for i := 0; i < scansz; i++ {
		host := <-ch
		if host.isUp > 0 {
			pool = append(pool, host.addr)
		}
	}

	// If -c parameter is specified, than rewrite hosts file
	// Else compare new pool with existing (from file)
	args := os.Args
	if len(args) == 2 && args[1] == "-c" {
		if len(pool) > 0 {
			// Rebuild hosts file
			str := strings.Join(pool, "\n")
			ioutil.WriteFile("hosts", []byte(str), os.ModePerm)
			fmt.Println("Found", len(pool), "hosts, saved to hosts file")
		} else {
			fmt.Println("No hosts found")
		}
	} else {
		if hosts, err := ioutil.ReadFile("hosts"); err == nil {
			hs := strings.Split(string(hosts), "\n")
			p1 := PrintDiff(hs, pool, "-")
			p2 := PrintDiff(pool, hs, "+")
			if p1+p2 == 0 {
				fmt.Println("Hosts map are not changed")
			}
		} else {
			fmt.Println("hosts file not found, run 'camdif -c' to create it")
		}
	}
}

// ScanPort return true if port is available, false otherwize
func ScanPort(ip string, port int, out chan int) {
	target := fmt.Sprintf("%s:%d", ip, port)
	for {
		conn, err := net.DialTimeout("tcp", target, 1*time.Second)
		if err != nil {
			if strings.Contains(err.Error(), "too many open files") {
				time.Sleep(100 * time.Millisecond)
				continue
			}
			out <- 0
			return
		}
		conn.Close()
		break
	}
	out <- 1
}

// PrintDiff is to print array differences, prefixed with d param
// Return number of printed (changed) items
func PrintDiff(p1, p2 []string, d string) int {
	changes := 0
	for _, d1 := range p1 {
		exists := false
		for _, d2 := range p2 {
			if d1 == d2 {
				exists = true
			}
		}
		if exists == false {
			fmt.Println(d, d1)
			changes++
		}
	}
	return changes
}

// GetOwnIP return preferred outbound ip of this machine
func GetOwnIP() string {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)
	return localAddr.IP.String()
}
