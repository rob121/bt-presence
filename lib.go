package main 

import (
 
    "net"
    "time"
    "fmt"
)


func uptime() int {
    return int(time.Since(startTime).Seconds())
}


func getoutboundip() string {
    conn, err := net.Dial("udp", "8.8.8.8:80")
    if err != nil {
            
    }
    defer conn.Close()

    localAddr := conn.LocalAddr().(*net.UDPAddr)

    return localAddr.IP.String()
}

func master_address() string{
	
    min := 0
    resp := ""
    

	for ip, val := range peer_list {
		
		    //fmt.Println("IP:",ip," INT:",val)
		
	        if (val.Uptime > min) {
	            min = val.Uptime
	            resp = ip
	        }
	}

    master_host = resp
    
    if(master_host==getoutboundip()){
	    
	    master = true;
	    slave = false;
	    
	    fmt.Println("I AM THE MASTER")
	    
    }else{
	    
	     fmt.Println("I AM THE SLAVE")
	    
	    master = false;
	    slave = true;
    }
	
	return resp
	
	
}