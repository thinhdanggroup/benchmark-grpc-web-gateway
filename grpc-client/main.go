package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"grpc-client/pkg"
	"log"
	"sync"
	"time"
)

func main() {
	request := &pkg.PingRequest{}

	conn, err := grpc.Dial("localhost:6790", grpc.WithInsecure())

	if err != nil {
		log.Fatal("error request", err)
	}
	defer conn.Close()

	start := makeTimestamp()
	server := pkg.NewCoreServiceClient(conn)
	lock := &sync.WaitGroup{}

	for i := 0; i < 10000; i++ {
		lock.Add(1)

		go func() {
			_, error := server.Ping(context.Background(), request)

			if error != nil {
				log.Print(error)
			}

			lock.Done()
		}()

		//log.Print(response)
	}

	lock.Wait()

	fmt.Print(makeTimestamp() - start)
}

func makeTimestamp() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}
