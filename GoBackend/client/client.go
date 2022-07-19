// package main

// import (
// 	"fmt"
// 	"log"
// 	"net/rpc"
// 	"os"
// )

// type VolumeData struct {
// 	VolumeInPercentage string
// }

// // func main() {
// //     argsWithoutProg := os.Args[1:]

// // 	var reply VolumeData
// // 	var getReply VolumeData

// // 	client, err := rpc.DialHTTP("tcp", "localhost:4040")

// // 	if err != nil {
// // 		log.Fatal("error in connection")
// // 	}

// // 	volume := VolumeData{VolumeInPercentage: argsWithoutProg[0]}
// // 	client.Call("API.SetVolume", volume, &reply)

// // 	client.Call("API.GetVolume", "get", &getReply)

// // 	fmt.Println(getReply)

// // }
