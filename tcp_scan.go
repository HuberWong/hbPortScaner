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
	ip string
	portStart int
	portEnd int
}

var argv scanRange

var helpStr string = `
-h	ip startPort endPort
`
var errNumberOfArgs string = "The number of parameters is incorrect"

func parseArgs() error {
	if  len(os.Args) == 0 {
		fmt.Println(helpStr)
	}
	switch os.Args[1] {
	case "-i":
		if len(os.Args) != 5 {
			fmt.Println(errNumberOfArgs)
		}
		argv.ip = os.Args[2]
		if port, err := strconv.Atoi(os.Args[3]); err != nil {
			argv.portStart = port
		}
		if port, err := strconv.Atoi(os.Args[4]); err != nil {
			argv.portEnd = port
		}
		return nil
	}
	return errors.New("parse input arguments error")
}

func isIPConnected(ip string) error {
	conn, err := net.DialTimeout("tcp", ip, 2 * time.Second)
	if err != nil {
		return errors.New(fmt.Sprintf("ip %+v cannot to be connected", ip))
	}
	conn.Close()
	return nil
}

func isPortConnectable(port int) error {
	if port >= 0 && port <= (2 << 16 - 1) {
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
	if err := checkScanRange(sr); err != nil {
		return err
	}

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

	if err := parseArgs(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(argv)

    target := argv.ip
    for i := argv.portStart; i <= argv.portEnd; i++ {
        address := fmt.Sprintf("%s:%d", target, i)
        conn, err := net.Dial("tcp", address)
        if err == nil {
            fmt.Printf("Port %d is open.\n", i)
            conn.Close()
        }
    }
}
