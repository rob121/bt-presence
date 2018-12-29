package main

import (
	"fmt"
	"log"
	"time"
	"net"
	"strconv"
	"strings"
    "encoding/binary"
   	"github.com/schollz/peerdiscovery"
)



type Host struct{
	
	Address string
	Uptime int
	Room string
}

func discover_peers(){
	
	
	//send this client to peer_handler, use discover to get the others on the network
	
	
	go peer_handler() //notifcation of peers found!ß
	
	ut := 0
	timelimit := 1
	loop := 0
	
	for {
	 
	fmt.Println("Scanning for 10 seconds to find LAN peers")
	// discover peers
	
	uptm := strconv.Itoa(uptime()) //pass along each client uptime, for electing master
	
	
	
	discoveries, err := peerdiscovery.Discover(peerdiscovery.Settings{
		Limit:     -1,
		Payload:   []byte(room+"|"+uptm),
		Delay:     800 * time.Millisecond,
		TimeLimit: time.Duration(timelimit) * time.Second,
	})

    

	// print out results
	if err != nil {
		log.Fatal(err)
	} else {
		if len(discoveries) > 0 {
			
			fmt.Printf("Found %d other nodes\n", len(discoveries))
			
			peers <- Host{Address: getoutboundip(),Uptime: uptime(),Room: room}//send this machine
			
			for _, d := range discoveries {
				
				//peers <- d.Address //send to channel peers for notifcation in peer_handlers
				ut = 0
				
				payload := string(d.Payload)
				
				parts := strings.Split(payload,"|")
				
				if(len(parts)>1){
				
				ut, _ = strconv.Atoi(parts[1])
				
				}
				
				peers <- Host{Address: d.Address,Uptime: ut,Room: parts[0]}
			
				
				}
		} else {
			fmt.Println("Found no other devices. You need to run this on another computer at the same time.")
		}
	}
	loop++
	if(loop>100){
		
		timelimit = 180 
		
	}
	

	}
	
}

func peer_handler(){
	
	 //go device_poller() //uncomment when ready
 for {
    select {
    case host := <-peers:
    
        fmt.Printf("Found '%s' with payload '%s' and uptime of '%d'\n", host.Address, host.Room,host.Uptime)

        peer_list[host.Address] = host
        master_address()// sort out who the master is
 
    }
    time.Sleep(100 * time.Millisecond)
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