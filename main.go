package main

import (
	"fmt"
	"log"

	desc "github.com/golang/protobuf/descriptor"
	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
	descpb "github.com/golang/protobuf/protoc-gen-go/descriptor"
	pb "github.com/updogliu/pboptions/proto"
)

func main() {
	req := (*pb.RequestType)(nil)
	fileDescriptor, _ /*msgDescriptor*/ := desc.ForMessage(req)

	// fmt.Println("fileDescriptor: ", pb2json(fileDescriptor))
	// fmt.Println("msgDescriptor: ", pb2json(msgDescriptor))
	// fmt.Printf("msgDescriptor: %+v", msgDescriptor)

	method := FindMethod(fileDescriptor, "/mypb.MyService/MyMethod")
	fmt.Println("method: ", pb2json(method))
}

func FindMethod(fd *descpb.FileDescriptorProto, fullMethodName string) *descpb.MethodDescriptorProto {
	for _, service := range fd.Service {
		for _, method := range service.Method {
			if "/"+*fd.Package+"."+*service.Name+"/"+*method.Name == fullMethodName {
				return method
			}
		}
	}
	log.Fatal("method not found")
  return nil
}

func pb2json(pb proto.Message) string {
	marshaler := jsonpb.Marshaler{Indent: "  "}
	str, err := marshaler.MarshalToString(pb)
	if err != nil {
		panic(err)
	}
	return str
}
