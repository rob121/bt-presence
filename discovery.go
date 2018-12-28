package main

import (
	"fmt"
	"log"
	"time"
	"net"
    "encoding/binary"
   	"github.com/schollz/peerdiscovery"
)


func discover_peers(){
	
	go peer_handler()
	
	timelimit := 3
	loop := 0
	
	for {
	 fmt.Println("Scanning for 10 seconds to find LAN peers")
	// discover peers
	discoveries, err := peerdiscovery.Discover(peerdiscovery.Settings{
		Limit:     -1,
		Payload:   []byte(room),
		Delay:     800 * time.Millisecond,
		TimeLimit: time.Duration(timelimit) * time.Second,
	})

	// print out results
	if err != nil {
		log.Fatal(err)
	} else {
		if len(discoveries) > 0 {
			fmt.Printf("Found %d other nodes\n", len(discoveries))
			for i, d := range discoveries {
				
				peers <- d.Address //send to channel
				fmt.Printf("%d) '%s' with payload '%s'\n", i, d.Address, d.Payload)
			}
		} else {
			fmt.Println("Found no devices. You need to run this on another computer at the same time.")
		}
	}
	loop++
	if(loop>20){
		
		timelimit = 180 
		
	}
	

	}
	
}

func peer_handler(){
	
	 //go device_poller() //uncomment when ready
 for {
    select {
    case ip := <-peers:
    
        resp,_,_ := net.ParseCIDR(ip+"/30")
        peer_list[ip] = ip2int(resp)
        //fmt.Println("received message", ip)
        //fmt.Println(peer_list)
        
        master_address()
        //fmt.Printf("Master address %s",master)
   
    }
     time.Sleep(1 * time.Second)
 }
	
	
}

func ip2int(ip net.IP) uint32 {
    
	if len(ip) == 16 {
		return binary.BigEndian.Uint32(ip[12:16])
	}
	return binary.BigEndian.Uint32(ip)
}

func int2ip(nn uint32) net.IP {
	ip := make(net.IP, 4)
	binary.BigEndian.PutUint32(ip, nn)
	return ip
}