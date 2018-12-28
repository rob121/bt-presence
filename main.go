package main

import(
"net"
"flag"
"fmt"
)

var master bool
var slave bool
var peers chan string
var rssi chan string
var peer_list map[string]uint32
var master_host string
var power_rating map[string][]Device
var room string
var pos string

type NotificationPlugin interface {
    Notify() 
}




func main() {
	
 setup()
 
 fmt.Println("Starting up listener for "+room)
 
 go device_poller()
 
 go discover_peers()
 
 go start_server()
 
 select {}//here to keep it up
	
}


func setup(){
	
 flag.StringVar(&room, "room","room", "Room Name")	
 flag.StringVar(&pos, "pos","0:0", "Room Position (4x4 Grid)")	
 
 flag.Parse()
 
 power_rating = make(map[string][]Device)	
 peers = make(chan string,1)
 rssi  = make(chan string,1)
 peer_list = make(map[string]uint32)
 my_ip := GetOutboundIP()
 res,_,_ := net.ParseCIDR(my_ip+"/30")
 peer_list[my_ip] = ip2int(res)
	
}