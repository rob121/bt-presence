package main

import(
"fmt"
"os"
"log"
"net/http"
"io/ioutil"
)

//send a http request after receing channel

func client_send(device string,strength string){
	
	    host, err := os.Hostname()
		 
		 if err != nil {
			 
	   }	
       
       
       if(master_host==""){
	       fmt.Println("No Master Elected Yet!")
	       return
       }
       
        fmt.Println("Sending: http://"+master_host+":15784/"+host+"/"+room+"/"+pos+"/"+device+"/"+strength)
       
        resp, err := http.Get("http://"+master_host+":15784/"+host+"/"+room+"/"+pos+"/"+device+"/"+strength)
        
       
        if err != nil {
		log.Println(err)
		return
		}
		
		defer resp.Body.Close()
	
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println(err)
			return
		}
	
		fmt.Println("Response: "+string(body))
       
}