package connection

import (
	"log"
	"net"
)

type ConnPool struct {
	MaxConnections int
	NumConnections int
	FreeConn       chan net.Conn
}

const (
	MaxConnections = 1000
)

func (cp *ConnPool) CreatePool(maxConnections int) error {
	cp.FreeConn = make(chan net.Conn, maxConnections)
	for i := 0; i < maxConnections; i++ {
		conn, err := net.Dial("tcp", "localhost:8081")
		if err != nil {
			log.Println("error in dialing tcp", err)
			return err
		}
		cp.FreeConn <- conn
	}
	cp.MaxConnections = maxConnections
	cp.NumConnections = maxConnections
	return nil
}

func (cp *ConnPool) Get() net.Conn {
	return <-cp.FreeConn
}

func (cp *ConnPool) Put(conn net.Conn) {
	cp.FreeConn <- conn
}

func (cp *ConnPool) CreateNewConnection() error {
	cp.NumConnections--
	if cp.NumConnections < cp.MaxConnections {
		conn, err := net.Dial("tcp", "localhost:8081")
		if err != nil {
			log.Println("error in dialing tcp", err)
			return err
		}
		cp.FreeConn <- conn
		cp.NumConnections++
	}
	return nil
}
