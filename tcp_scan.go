package main

import (
	"errors"
	"fmt"
	"net"
	"os"
	"strconv"
	"sync"
	"time"
)

func scanPort(ip string, port int, wg *sync.WaitGroup) {
	defer wg.Done()

	address := fmt.Sprintf("%s:%d", ip, port)
	conn, err := net.DialTimeout("tcp", address, 1*time.Second)
	if err != nil {
		return
	}
	conn.Close()

	fmt.Printf("Port %d is open\n", port)
}

// 耗时操作, 会开子线程, 不要提前结束主线程
func tcpScan(ip string, startPort int, endPort int) {
	wg := &sync.WaitGroup{}
	for port := startPort; port <= endPort; port++ {
		wg.Add(1)
		go scanPort(ip, port, wg)
	}
	wg.Wait()
}

type tcpScanArgs struct {
	ip                 string
	startPort, endPort int
}

// 全局变量
var inputTcpScanArgs tcpScanArgs

var helpStr string = `
-h	ip startPort endPort
`

func parseArgs() error {
	if len(os.Args) == 0 {
		fmt.Println(helpStr)
	}
	switch os.Args[1] {
	case "-i":
		if len(os.Args) != 5 {
			return errors.New("参数数量错误")
		}

		inputTcpScanArgs.ip = os.Args[2]

		port, err := strconv.Atoi(os.Args[3])
		if err != nil {
			fmt.Println(err)
		}
		inputTcpScanArgs.startPort = port

		port, err = strconv.Atoi(os.Args[4])
		if err != nil {
			fmt.Println(err)
		}
		inputTcpScanArgs.endPort = port

		// 如果 `startPort` 大于 `endport`, 则返回错误
		if inputTcpScanArgs.startPort > inputTcpScanArgs.endPort {
			return errors.New("startPort 不可大于 endport")
		}
		return nil
	default:
		return errors.New("不可识别")
	}
}

func main() {
	// ip := "127.0.0.1"
	// startPort := 1
	// endPort := (2 << 16) - 1

	err := parseArgs()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	tcpScan(inputTcpScanArgs.ip, inputTcpScanArgs.startPort, inputTcpScanArgs.endPort)
	// wg := &sync.WaitGroup{}

	// for port := startPort; port <= endPort; port++ {
	// 	wg.Add(1)
	// 	go scanPort(ip, port, wg)
	// }

	// wg.Wait()
}
