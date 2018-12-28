package main 

import (
    "fmt"
    "os/exec"
    "strings"
    "bufio"
    "os"
    "time"
    "strconv"
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

lines,err := devices_list()
	
if(err!=nil){}	
	
for {	

 // go func() {
  for _,id := range lines {
	  
	fmt.Printf("Scanning for %s%s",id,"\n")  
    strength := device_strength(id)
    time.Sleep(10 * time.Second)
    client_send(id,strength)
    
    
       
  }
  //}()
  //time.Sleep(5 * time.Second)
 
    
 }   
	
}

func devices_list() ([]string, error) { 
	
  path := "./devices"

  file, err := os.Open(path)
  if err != nil {
    return nil, err
  }
  defer file.Close()

  var lines []string
  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    lines = append(lines, scanner.Text())
  }
  return lines, scanner.Err()
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