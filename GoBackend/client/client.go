package main

import (
	"fmt"
	"log"
	"net/rpc"
)

type VolumeData struct {
	VolumeInPercentage string
}

func main() {
	var reply VolumeData
	var getReply VolumeData

	client, err := rpc.DialHTTP("tcp", "localhost:4040")

	if err != nil {
		log.Fatal("error in connection")
	}

	volume := VolumeData{VolumeInPercentage: "50"}
	client.Call("API.SetVolume", volume, &reply)

	client.Call("API.GetVolume", "get", &getReply)

	fmt.Println(getReply)

}
