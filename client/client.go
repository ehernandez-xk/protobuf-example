package main

import (
	"bytes"
	"fmt"
	"net/http"

	pb "github.com/ehernandez-xk/protobuf-example/service"
	"github.com/golang/protobuf/proto"
)

func main() {
	myClient := pb.Client{Id: 526, Name: "John Doe", Email: "johndoe@example.com", Country: "US"}
	clientInbox := make([]*pb.Client_Mail, 0, 20)
	clientInbox = append(clientInbox, &pb.Client_Mail{RemoteEmail: "jannetdoe@example.com", Body: "Hello. Greetings. Bye."}, &pb.Client_Mail{RemoteEmail: "WilburDoe@example.com", Body: "Bye, Greetings, hello."})

	myClient.Inbox = clientInbox

	data, err := proto.Marshal(&myClient)
	if err != nil {
		fmt.Println(err)
		return
	}

	_, err = http.Post("http://localhost:3000", "", bytes.NewBuffer(data))
	if err != nil {
		fmt.Println(err)
		return
	}

}
