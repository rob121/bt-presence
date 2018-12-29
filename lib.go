package main 

import (
 
    "net"
    "time"
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
	    
    }else{
	    
	    master = false;
	    slave = true;
    }
	
	return resp
	
	
}