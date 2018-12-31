package main 

import (
 
    "net"
    "time"
    "fmt"
    "os"
    "encoding/gob"
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

func get_devices() map[string]string{
	
	
	dataFile, err := os.OpenFile("devices.gob", os.O_RDONLY|os.O_CREATE, 0666)
	
	defer dataFile.Close()

 	if err != nil {
 		fmt.Println(err)
 		return nil
 	}

 	dataDecoder := gob.NewDecoder(dataFile)
 	err = dataDecoder.Decode(&devices)

 	if err != nil {
 		fmt.Println(err)
 		return nil
 	}
 	
 	
 	if(len(devices)<1){
		return make(map[string]string)
	}
 	
 	
 	return devices
	
	
	
}

func set_devices(devices map[string]string){
	
	if(len(devices)<1){
		return
	}
	
	dataFile, err := os.Create("devices.gob")
    defer dataFile.Close()
 	if err != nil {
 		fmt.Println(err)
 		return 
 	}

      // serialize the data
 	dataEncoder := gob.NewEncoder(dataFile)
 	dataEncoder.Encode(devices)

 	
 	
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
	    //retrieve the devices?
	    master = false;
	    slave = true;
    }
	
	return resp
	
	
}