package main

import ht "./.."
import "git.apache.org/thrift.git/lib/go/thrift"

import "fmt"
import "os"

func main() {

	transportFactory := thrift.NewTTransportFactory()
	//protocolFactory  := thrift.NewTBinaryProtocolFactoryDefault();
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()

	transportFactory = thrift.NewTFramedTransportFactory(transportFactory)

	socket, err := thrift.NewTSocket("localhost:38080")
	if err != nil {
		fmt.Print("Error opening connection: ", err.Error(), "\n")
		os.Exit(1)
	}

	transport := transportFactory.GetTransport(socket)
	defer transport.Close()
	if err := transport.Open(); err != nil {
		fmt.Printf("Error opening transport: %v\n", err)
		return
	}

	clientServiceClient := ht.NewClientServiceClientFactory(transport, protocolFactory)

	fmt.Printf("protocolFactory:  %v\n", protocolFactory)
	fmt.Printf("transportFactory: %v\n", transportFactory)
	fmt.Printf("clientServiceClient: %v\n", clientServiceClient)

	ns, nserr, err := clientServiceClient.NamespaceOpen("test")

	if err != nil || nserr != nil {
		fmt.Printf("Error opening namespace: %v\n", err.Error())
		fmt.Printf("NSerr: %v\n", nserr)
		return
	}

	fmt.Printf("Namespace: %v\n", ns)
	row, exception, err := clientServiceClient.GetRow(ns, "test", "1")

	fmt.Printf("Exception: %v\nError: %v\n", exception, err)

	fmt.Printf("Row: %v", row)
}
