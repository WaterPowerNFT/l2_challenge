package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"
	"time"
)

type ipPort struct {
	ip      string
	port    string
	timeout int64
}

func getIPPortTimeout() ipPort {

	if len(os.Args) > 3 || len(os.Args) == 1 {
		fmt.Println(os.Args)
		panic("wrong number of elements")
	}
	massiveIP := strings.Split(os.Args[1], ":")
	var timeout int64 = 10
	if len(os.Args) == 3 {
		var err error
		trimmedStr := strings.Trim(os.Args[2], "--timeout=")
		timeout, err = strconv.ParseInt(trimmedStr, 10, 32)
		if err != nil {
			panic(err)
		}
	}
	newIPPort := &ipPort{ip: massiveIP[0], port: massiveIP[1], timeout: timeout}
	fmt.Println(newIPPort)
	return *newIPPort
}

func (ipp *ipPort) connect() error {
	address := fmt.Sprintf("%s:%s", ipp.ip, ipp.port)

	fmt.Println("Connecting to", address, "...")
	conn, err := net.DialTimeout("tcp", address, time.Duration(ipp.timeout))
	if err != nil {
		return err
	}
	defer conn.Close()
	fmt.Println("Connected")
	sigChannel := make(chan os.Signal, 1)
	errChannel := make(chan error, 1)
	signal.Notify(sigChannel, syscall.SIGINT)

	go func(connection net.Conn, error_channel chan error) {
		input := make([]byte, 1024)
		for {
			n, err := conn.Read(input)
			if err != nil {
				error_channel <- fmt.Errorf("remoute server stopped: %v", err)
				return
			}
			fmt.Println(string(input[:n]))
		}
	}(conn, errChannel)

	go func(connection net.Conn, error_channel chan error) {
		reader := bufio.NewReader(os.Stdin)
		for {
			text, err := reader.ReadBytes('\n')
			if err != nil {
				error_channel <- err
				return
			}
			text = text[:len(text)-1]

			_, err = conn.Write(text)
			if err != nil {
				error_channel <- err
				return
			}
		}
	}(conn, errChannel)

	select {
	case err := <-errChannel:
		{
			return err
		}
	case <-sigChannel:
		{
			fmt.Println("closed by system call: ")
		}
	}
	return nil
}

func main() {
	ipPortUnit := getIPPortTimeout()
	err := ipPortUnit.connect()
	if err != nil {
		fmt.Println(err)
	}
}
