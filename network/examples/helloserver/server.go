package main

import (
	"context"
	"encoding/binary"
	"fmt"
	"github.com/zllangct/RockGO/network"
	"time"
	
)

//MyServer struct for testing tars tcp server
type MyServer struct{}

//ParseMessage recv request and make response.
func (s *MyServer) ParseMessage(ctx context.Context,req []byte)([]uint32,[]byte){

	fmt.Println("recv", string(req))
	return []uint32{0}, nil
}

//ParsePackage parse package from buff,check if tars package finished.
func (s *MyServer) ParsePackage(buff []byte) (pkgLen, status int) {
	if len(buff) < 4 {
		return 0, network.PACKAGE_LESS
	}
	length := binary.BigEndian.Uint32(buff[:4])

	if length > 1048576000 || len(buff) > 1048576000 { // 1000MB
		return 0, network.PACKAGE_ERROR
	}
	if len(buff) < int(length) {
		return 0, network.PACKAGE_LESS
	}
	return int(length), network.PACKAGE_FULL
}

func main() {
	s := MyServer{}
	conf := &network.ServerConf{
		Proto:           "tcp",
		PackageProtocol: &s,
		Address:         "localhost:3333",
		//MaxAccept:     500,
		MaxInvoke:     20,
		AcceptTimeout: time.Millisecond * 500,
		ReadTimeout:   time.Millisecond * 100,
		WriteTimeout:  time.Millisecond * 100,
		IdleTimeout:   time.Millisecond * 600000,
	}

	svr := network.NewServer(conf)
	err := svr.Serve()
	if err != nil {
		panic(err)
	}
}
