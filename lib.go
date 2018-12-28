package main 

import (
 
    "net"
)


func GetOutboundIP() string {
    conn, err := net.Dial("udp", "8.8.8.8:80")
    if err != nil {
            
    }
    defer conn.Close()

    localAddr := conn.LocalAddr().(*net.UDPAddr)

    return localAddr.IP.String()
}

func master_address() string{
	
    resp := ""
    
    min := uint32(4294967294)
    
    min = uint32(0)

	for ip, val := range peer_list {
		
		    //fmt.Println("IP:",ip," INT:",val)
		
	        if (val > min) {
	            min = val
	            resp = ip
	        }
	}

    master_host = resp
	
	return resp
	
	
}