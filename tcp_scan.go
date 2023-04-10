package main

import (
	"errors"
	"fmt"
	"net"
	"os"
	"strconv"
	"time"
)

type scanRange struct {
	// option string
	ip        string
	portStart int
	portEnd   int
}

var tcpArgv scanRange

var helpStr string = `
-h	ip startPort endPort
`
var errNumberOfArgs string = "The number of parameters is incorrect"

func parseArgs() error {
	if len(os.Args) == 0 {
		fmt.Println(helpStr)
	}
	switch os.Args[1] {
	case "-i":
		if len(os.Args) != 5 {
			fmt.Println(errNumberOfArgs)
		}
		tcpArgv.ip = os.Args[2]

		port, err := strconv.Atoi(os.Args[3])
		if err != nil {
			return err
		}
		tcpArgv.portStart = port

		port, err = strconv.Atoi(os.Args[4])
		if err != nil {
			return err
		}
		tcpArgv.portEnd = port
		// if port, err := strconv.Atoi(os.Args[3]); err != nil {
		// 	tcpArgv.portStart = port
		// }
		// if port, err := strconv.Atoi(os.Args[4]); err != nil {
		// 	tcpArgv.portEnd = port
		// }
		return nil
	}
	return errors.New("parse input arguments error")
}

func isIPConnected(ip string) error {
	conn, err := net.DialTimeout("tcp", ip, 2*time.Second)
	if err != nil {
		return errors.New(fmt.Sprintf("ip %+v cannot to be connected", ip))
	}
	conn.Close()
	return nil
}

func isPortConnectable(port int) error {
	if port < 0 && port > 2<<16 {
		return errors.New(fmt.Sprintf("port %+v is out of connectable range\n", port))
	}
	return nil
}

func checkScanRange(sr scanRange) error {

	if err := isIPConnected(sr.ip); err != nil {
		return err
	}

	if err := isPortConnectable(sr.portStart); err != nil {
		return err
	}

	if err := isPortConnectable(sr.portStart); err != nil {
		return err
	}

	return nil
}

func tcpScan(sr scanRange) error {
	// if err := checkScanRange(sr); err != nil {
	// 	return err
	// }

	for i := sr.portStart; i < sr.portEnd; i++ {
		address := fmt.Sprintf("%s:%d", sr.ip, i)
		conn, err := net.Dial("tcp", address)
		if err != nil {
			return err
		}
		conn.Close()

	}

	return nil
}

func main() {

	// if err := parseArgs(); err != nil {
	// 	fmt.Println(err)
	// 	os.Exit(1)
	// }

	// fmt.Println(tcpArgv)

	// err := tcpScan(tcpArgv)
	// if err != nil {
	// 	fmt.Println(err)
	// 	os.Exit(1)
	// }

	tcpArgv = scanRange{
		ip: "127.0.0.1",
		// ip: "39.156.66.10",
		portStart: 21,
		portEnd: 550,
	}
	target := tcpArgv.ip
	for i := tcpArgv.portStart; i <= tcpArgv.portEnd; i++ {
	    address := fmt.Sprintf("%s:%d", target, i)
	    conn, err := net.Dial("tcp", address)
	    if err == nil {
	        fmt.Printf("Port %d is open.\n", i)
	        conn.Close()
	    }
	}
}
