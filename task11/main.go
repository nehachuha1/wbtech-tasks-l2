package task11

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"os/signal"
	"time"
)

type Config struct {
	Timeout time.Duration
	Host    string
	Port    string
}

func NewConfig() *Config {
	c := Config{}

	flag.DurationVar(&c.Timeout, "timeout", 10*time.Second, "timeout to socket connection")
	flag.Parse()
	args := flag.Args()
	c.Host = args[0]
	c.Port = args[1]

	return &c
}

func read(reader *bufio.Reader, error chan<- error) {
	for {
		buff, err := reader.ReadBytes('\n')
		if err == io.EOF {
			error <- err
			return
		}

		if err != nil {
			error <- err
		}

		fmt.Printf("Server: %s", buff)
	}
}

func write(conn net.Conn, reader *bufio.Reader, error chan<- error) {
	for {
		buff, err := reader.ReadBytes('\n')
		if err != nil {
			error <- err
		}
		if _, err := conn.Write(buff); err != nil {
			error <- err
		}
	}
}

func main() {
	errorChan := make(chan error)
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, os.Kill)
	cfg := NewConfig()

	address := net.JoinHostPort(cfg.Host, cfg.Port)
	conn, err := net.DialTimeout("tcp", address, cfg.Timeout)
	if err != nil {
		time.Sleep(cfg.Timeout)
		log.Fatalln("error connection to server: ", err.Error())
	}
	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {
			log.Fatalln("error closing connection: ", err.Error())
		}
	}(conn)

	fmt.Println("connected to server")

	serverReader := bufio.NewReader(conn)
	inputReader := bufio.NewReader(os.Stdin)

	go read(serverReader, errorChan)
	go write(conn, inputReader, errorChan)

	select {
	case err := <-errorChan:
		fmt.Println(err.Error())
	case <-done:
	}
}
