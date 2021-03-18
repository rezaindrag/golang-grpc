package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"

	"github.com/golang/protobuf/jsonpb"
	"github.com/rezaindrag/golang-grpc/example_1/model"
)

func main() {
	var user1 = &model.User{
		Id:       "u001",
		Name:     "Sylvana Windrunner",
		Password: "f0r Th3 H0rD3",
		Gender:   model.UserGender_FEMALE,
	}

	// Convert proto object to json
	var buf bytes.Buffer
	err1 := (&jsonpb.Marshaler{}).Marshal(&buf, user1)
	if err1 != nil {
		fmt.Println(err1.Error())
		os.Exit(0)
	}
	jsonString := buf.String()
	fmt.Println(jsonString)

	// Convert json to proto object
	buf2 := strings.NewReader(jsonString)
	protoObject := new(model.User)
	err2 := (&jsonpb.Unmarshaler{}).Unmarshal(buf2, protoObject)
	if err2 != nil {
		fmt.Println(err2.Error())
		os.Exit(0)
	}
	fmt.Println(protoObject.GetName())
}
