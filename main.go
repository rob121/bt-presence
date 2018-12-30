package main

import(
"flag"
"fmt"
"time"
)

var master bool
var slave bool
var peers chan Host
var devices map[string]string
var rssi chan string
var peer_list map[string]Host
var master_host string
var power_rating map[string][]Device
var room string
var pos string
var startTime time.Time

type NotificationPlugin interface {
    Notify() 
}


func main() {
	
 setup()
 
 fmt.Println("Starting up listener for "+room)
 
 go start_server()
 
 go discover_peers()

 go device_poller()

 select {}//here to keep it up and running
	
}


func setup(){
	
 flag.StringVar(&room, "room","default", "Room Name")		

 pos = "1:1" //this is set for future position capability

 
 flag.Parse()
 
 startTime = time.Now()
 devices = make(map[string]string)	
 power_rating = make(map[string][]Device)	
 peers = make(chan Host,1)
 rssi  = make(chan string,1)
 peer_list = make(map[string]Host)
 
 
  

	
}