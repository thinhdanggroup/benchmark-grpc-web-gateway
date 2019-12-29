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

	target := "bench-server:6790"
	conn, err := grpc.Dial(target, grpc.WithInsecure())

	if err != nil {
		log.Fatal("error request", err)
	}
	defer conn.Close()

	start := makeTimestamp()
	server := pkg.NewCoreServiceClient(conn)
	lock := &sync.WaitGroup{}

	total := int64(50000)

	for i := int64(0); i < total; i++ {
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

	execTime := (makeTimestamp() - start)
	fmt.Println("Benchmark result for ", total, " request in ", execTime," microsecond")
	fmt.Println("Throughtput :", total * 1000000/ execTime, "req/s" )
	fmt.Println("Latency :", execTime / total, "microsecond")  
}

func makeTimestamp() int64 {
	return time.Now().UnixNano() / 1000
}
