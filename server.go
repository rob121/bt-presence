package main


import (
 "fmt"
 "net/http"
 "html/template"
 "strings"
 "strconv"
 //"encoding/json"
 "github.com/tidwall/sjson"
 //"github.com/kr/pretty"
 "./plugin_registry"
 _ "./drivers/hubitat_plugin" // <-- import the plugin to enable, thats it
 "gopkg.in/yaml.v2"
)

func start_server(){
	
	fmt.Println("Starting app server")

	http.HandleFunc("/deviceform", gui_device_handler)
	http.HandleFunc("/manage", gui_handler)
	http.HandleFunc("/devices",device_handler)
	http.HandleFunc("/stats", stats_handler)
	http.Handle("/resources/", http.StripPrefix("/resources/", http.FileServer(http.Dir("./www/resources"))))
	http.HandleFunc("/", server_handler)
	http.ListenAndServe(":15784", nil)
		
}

func stats_handler(w http.ResponseWriter, r *http.Request) { 
	
	data := `{"master": "`+master_host+`", "peers": "", "devices": "","hosts":""}`    
    value, _ := sjson.Set(data, "peers", peer_list)
    value2, _:= sjson.Set(value, "devices",devices)
    value3, _:= sjson.Set(value2, "hosts",power_rating)
    fmt.Fprintf(w,value3)
	
}


func device_handler(w http.ResponseWriter, r *http.Request) { 
	
	data := `{"devices": ""}`    
    value, _ := sjson.Set(data, "devices", devices)
    fmt.Fprintf(w,value)
	
}

type data struct {
  Devices map[string]string
}

func gui_device_handler(w http.ResponseWriter, r *http.Request) { 
	
	
	devices = get_devices()
	
	if(r.Method=="POST"){
		
		//save here
		
	    if err := r.ParseForm(); err != nil {
            fmt.Fprintf(w, "ParseForm() err: %v", err)
            return
        }
        
        dev := r.Form["devices"][0]
       
		devi := strings.Split(dev,"\n")
		
		if(len(devi)>0){
		
		devices = make(map[string]string)//clear it out
		
		for _,v := range devi {
		
			devices[strings.TrimSpace(v)]=strings.TrimSpace(v)		
		}
		
		set_devices(devices)
		
		}

		
	}
	
	
	
    t, err := template.ParseFiles("www/form.html") //parse the html file homepage.html
    
    if err != nil { // if there is an error
  	  fmt.Print("template parsing error: ", err) // log it
  	}
  	
    err = t.Execute(w,data{Devices:devices})
    
    if err != nil { // if there is an error
  	  fmt.Print("template executing error: ", err) //log it
  	}
  	
  
	
	
}

func gui_handler(w http.ResponseWriter, r *http.Request) { 
	
	//if this is a slave, send to the master
	if(master_host!="" && master_host!=getoutboundip()){
	    fmt.Println("I am not the master, redirecting to master")
		http.Redirect(w, r, "http://"+master_host+":15784/manage", 302)
		return
	}
	
	
    if(len(devices)<1){
	   
		http.Redirect(w, r, "/deviceform", 302)
		return
	}
	
		
    t, err := template.ParseFiles("www/index.html") //parse the html file homepage.html
    
    if err != nil { // if there is an error
  	  fmt.Print("template parsing error: ", err) // log it
  	}
  	
    err = t.Execute(w,nil)
    
    if err != nil { // if there is an error
  	  fmt.Print("template executing error: ", err) //log it
  	}
	
	
}

// "handler" is our handler function. It has to follow the function signature of a ResponseWriter and Request type
// as the arguments.
func server_handler(w http.ResponseWriter, r *http.Request) {
	// For this case, we will always pipe "Hello World" into the response writer
		
	fmt.Fprintf(w, "PONG")
	
	fmt.Println("RECEIVED HIT")
	
	   //fmt.Println("Sending: http://"+master_host+":15784/"+host+"/"+room+"/"+pos+"/"+device+"/"+strength)
		
	    parts := strings.Split(r.URL.RequestURI(),"/")
	
	    ppos := strings.Split(parts[3],":")
	    
	    
	    pos1,_ := strconv.Atoi(ppos[0])
	    pos2,_ := strconv.Atoi(ppos[1])
	
		host   := parts[1]
	    label  := parts[2]
	    pos    := [2]int{pos1,pos2} 
		mac    := parts[4]
		rssi   := parts[5]
		addr   := r.RemoteAddr

       updated := 0

       //loop over what we have, update if present

       for _,devices := range power_rating {
	       
	       for idx,device := range devices {
	   
		       if (device.Mac==mac && device.Host==host){
			       
			       //update 
			       updated = 1
			       power_rating[mac][idx].Rssi = rssi
			       power_rating[mac][idx].Power = convert_power(rssi) 
			       break
			            
		       }
	       
	       }
	       
	       
       }
       
       if( updated==0 ) {

	       
	       power_rating[mac] = append( power_rating[mac],Device{Host: host,Mac: mac,Label: label,Rssi: rssi,Addr: addr, Power: convert_power(rssi), Location: pos})
	       
       }

	  // points := make([]trilateration.Point3,3)
      // dis := make([]float64,3)  
           //points[i] = [3]float64{float64(device.Location[0]),float64(device.Location[1]), 0} 
	         //dis[i] = float64(100 - device.Power)
			 //i++ 

	   //pretty.Print(power_rating)
	   
	   
	    ym, err := yaml.Marshal(&power_rating)
        if err != nil {
                fmt.Printf("error: %v", err)
        }
        fmt.Printf("\n%s\n\n", string(ym))
	   
	    fmt.Println("")
	 
	   //i := 0
	   
	   max_idx := 0
	   max := 0
	   
	   for _,devices2 := range power_rating {
		
	    //i = 0 
	    
	    max_idx = -1
	    
	    max = 1 //devices2[0].Power
	    
		for idx2, device2 := range devices2 {

	        			 
			 if(device2.Power > max && device2.Power>0){
				 
				 max = device2.Power
				 max_idx = idx2
				 
			 }
			 
		
	   } 
	   
	    if(max_idx>-1){
	    
	    
	    //send maker request to hubitat
	    for idx2, _ := range devices2 {
	    
	    // this is the arrived
	    if(max_idx==idx2){
	        fmt.Println("Arrived in room: "+devices2[max_idx].Label+ " for device "+devices2[max_idx].Mac)	
			for _, d := range plugin_registry.Notifiers {
		        d.Arrived(devices2[max_idx].Mac,devices2[max_idx].Label)
		    }
	    
	    }else{ //this is departed!
	        fmt.Println("Departed room: "+devices2[idx2].Label+ " for device "+devices2[idx2].Mac)	
	    	for _, d := range plugin_registry.Notifiers {
		        d.Departed(devices2[idx2].Mac,devices2[idx2].Label)
		    }	    
		    
	    }
	 

	    
	    }
	    
	 	}else{
		 //we didnt get a closes room == not here!
		 
	 	}
	 	
	 	
		 
			
	 }
	 

	 
	// pretty.Print(points)
	// pretty.Print(dis)
	// fmt.Println("")
	 /*
		 Trilateration is still unsatisfactory
	  params := trilateration.Parameters3{
		Loc: points,
		Dis: dis,
	  }
	  
	  result, _ := params.SolveTrilat3()
	  
	  pretty.Print(result)
	  fmt.Println("")
	  */
	
	
}