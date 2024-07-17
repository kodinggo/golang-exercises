package main

import (
	"fmt"
	"time"

	pb "go-protobuf/pb/greeting"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func main() {
	pbAddress := &pb.Address{
		Street: "jl. A",
		City:   "Jakarta",
	}
	fmt.Println(pbAddress)

	pbPerson := &pb.Person{
		Name:    "John",
		Age:     12,
		Gender:  pb.Gender_FEMALE,
		Address: &pb.Address{},
		// Others: &pb.Person_Hobby{
		// 	Hobby: "swimming",
		// },
		Others: &pb.Person_Job{
			Job: "IT",
		},
	}
	fmt.Println(pbPerson)

	// Konversi time.Time (golang) ke timestamp protobuf
	now := time.Now()
	pbNow := timestamppb.New(now)
	now = pbNow.AsTime()
	fmt.Println(now)
}
