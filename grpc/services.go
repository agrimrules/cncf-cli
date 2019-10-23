/*
Copyright © 2019 Agrim Asthana <dev@agrimasthana.com>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package main

import (
	"log"
	"golang.org/x/net/context"
	messages "github.com/agrimrules/cncf/services/grpc/proto"
	"google.golang.org/grpc"
)

func main() {
	var conn *grpc.ClientConn

	conn, err := grpc.Dial(":50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %s", err)
	}
	log.Printf("successfully connected to GRPC SERVER")
	defer conn.Close()

	c := messages.NewMessageServiceClient(conn)

	res, err := c.ListMessage(context.Background(), nil)
	if err != nil {
		log.Fatalf("couldn't find any messages : %s", err)
	}
	log.Printf("Response got: %s", res)
}