package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"os/exec"
)

type VolumeData struct{
	VolumeInPercentage string
}

type API int

func (r *API) SetVolume(volume VolumeData, reply *VolumeData) error{
	log.Printf("Setting volume to: ", volume.VolumeInPercentage);
	exec.Command("/bin/sh", "/home/swopnil/PycharmProjects/handtracker/handtracker.sh", volume.VolumeInPercentage).Output()
	*reply = volume
	return nil
}

func (r *API) GetVolume(getVolume string, reply *VolumeData)error{
	log.Printf("Getting volume data")
	cmd, _ := exec.Command("/bin/sh", "/home/swopnil/PycharmProjects/handtracker/handtracker.sh", "get").Output()
	volumedata := VolumeData{VolumeInPercentage: string(cmd)[4:7]}
	*reply = volumedata
	return nil
}

func main(){
	fmt.Println("Hello world")
	api := new(API)
	err := rpc.Register(api)
	if err != nil {
		log.Fatal("error in registration")
	}

	rpc.HandleHTTP();
	listener, err := net.Listen("tcp", ":4040")

	if err != nil{
		log.Fatal("Error in listening")
	}

	err = http.Serve(listener, nil)
	if err != nil{
		log.Fatal("error in serving")
	}
}