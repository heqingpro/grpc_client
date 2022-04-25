/*
 *
 * Copyright 2015 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

// Package main implements a client for Greeter service.
package main

import (
	"context"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	// "time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	defaultName = "world"
)

var (
	addr = flag.String("addr", "localhost:9000", "the address to connect to")
	name = flag.String("name", defaultName, "Name to greet")
)

func main() {
	flag.Parse()
	// Set up a connection to the server.
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := NewGreeterClient(conn)

	// Contact the server and print out its response.
	// ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	// defer cancel()
	// r, err := c.SayHello(ctx, &HelloRequest{Name: *name})
	// if err != nil {
	// 	log.Fatalf("could not greet: %v", err)
	// }
	// log.Printf("Greeting: %s", r.GetMessage())

	UploadFile(c)
	fmt.Println("finish")
}

func UploadFile(cc GreeterClient) {
	fmt.Println("start UploadFile")
	file, err := os.Open("./go.mod")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	content, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}
	fmt.Printf("UploadFile: len%d", len(content))
	client, err := cc.UploadFile(context.TODO())
	if err != nil {
		fmt.Println(err)
		return
	}

	// if err := client.Send(&UploadRequest{
	// 	Content: content,
	// 	}); err != nil{
	// 	fmt.Println(err)
	// 	return
	// }
	// status, err := client.CloseAndRecv()
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	for n := 0; n < len(content); n += 100 {
		var buf []byte
		if 100+n > len(content) {
			buf = content[n:]
		} else {
			buf = content[n : n+100]
		}
		err := client.Send(&UploadRequest{
			Content: buf,
		})
		if err != nil {
			fmt.Println(err)
			return
		}
	}
	status, err := client.CloseAndRecv()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("uploadResp%v: ", status)
}
