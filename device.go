package main 

import (
    "fmt"
    "os/exec"
    "strings"
   // "bufio"
   // "os"
    "io/ioutil"
    "time"
    "strconv"
    "net/http"
    "github.com/tidwall/gjson"
)


type Device struct {
	Host string
	Mac string
	Label string
	Rssi string
	Addr string
	Power int
	Location [2]int
}

func convert_power(rssi string) int{ 
	
	pwr,_ := strconv.Atoi(rssi)
	
	
	pw := 0
	
	
    apwr := (pwr * -1)
    
    
	if apwr < 10 { 
		
	  pw = 100 - apwr 
		
	}else if apwr < 20 {
		
	
	  pw = 100 - (apwr * 2)	
		
	} else if apwr < 40 {
		
	   pw = 100 - (apwr * 3)		
		
	}else {
		
	  	
	   pw = 100 - (apwr * 4)		 
		
	}  
	
	
  
	
    if pw < 0 {
	    
	    return 0
	    
    }
    
    return pw
		
}

func device_poller(){


	

	
for {	
	
	
  if(master_host!=""){	
   
   resp := httpGet("http://"+master_host+":15784/devices")
   
   result := gjson.Get(resp,"devices")
   
   result.ForEach(func(key, value gjson.Result) bool {
	v := value.String()
	devices[strings.TrimSpace(v)]=strings.TrimSpace(v)
	return true // keep iterating
   })
   
  } 
  
  	
  //devices is a map that must be populated
 // go func() {
  for _,id := range devices {
	  
	fmt.Printf("Scanning for %s%s",id,"\n")  
    strength := device_strength(id)
    time.Sleep(10 * time.Second)
    client_send(id,strength)
    
       
  }
  //}()
  //time.Sleep(5 * time.Second)
 
    
 }   
	
}



func httpGet(url string) string{
	
	client := http.Client{
		Timeout: time.Second * 2, // Maximum of 2 secs
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		
	}

	req.Header.Set("User-Agent", "device-client")

	res, getErr := client.Do(req)
	if getErr != nil {
		
	}

	body, _ := ioutil.ReadAll(res.Body)
	
	return string(body)
	
	
}	
	


func device_strength(id string) string{
	
	
	
    cmd :="hcitool"
	args := []string{"-i","hci0","cc",id}
	_, err := exec.Command(cmd,args...).Output()
    if err!=nil {
        fmt.Println(err.Error())
        return "-1000"
    }
    
	args2 := []string{"-i","hci0","rssi",id}
	output2, err2 := exec.Command(cmd,args2...).Output()
    if err2!=nil {
        fmt.Println(err2.Error())
        return "-1000"
    }
    
    args3 := []string{"-i","hci0","dc",id,"19"} //19 is the reason - user disconnect
    	
	_, err3 := exec.Command(cmd,args3...).Output()
    if err3!=nil {
        fmt.Println(err3.Error())
        
    }
    
    //fmt.Println(string(output2))
		
	op := string(output2)
	


	stat := strings.Contains(op,"RSSI return value")
	
	
	if(stat==true){
		
		r1 := strings.Replace(op,"RSSI return value: ","",-1)
		r2 := strings.Trim(r1,"%0A")
		r3 := strings.Replace(r2,"\n","",-1)
		
		return r3
		
	}
	
    return "-1000"
	
}