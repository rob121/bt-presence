package hubitat_plugin

import "fmt"
import "../../plugin_registry"
import "net/http"
import "github.com/tidwall/gjson"
import "os"
import "time"
import "io/ioutil"

/*
To create a driver all you have to do is create a Driver sturct, init as below and then define an Arrived Method	
**/

type Driver struct {}
var config string
var host string
var access_token string

func init() {
	
	

    http.DefaultClient.Timeout = time.Second * 10

	
	fmt.Println("Hubitat Driver Loaded")
    dr := &Driver{}
    plugin_registry.RegisterNotifcation(dr)
    json, err := os.Open("./hubitat_config.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	
	// defer the closing of our jsonFile so that we can parse it later on
	defer json.Close()
	
	byteValue, _ := ioutil.ReadAll(json)
    
    config = string(byteValue)
    
    host         = gjson.Get(config, "host").String()
	access_token = gjson.Get(config,"token").String()

    
}

func (d *Driver) Arrived(mac string,room string) {
	    
	    
	    
	     
	     device_id := gjson.Get(config, "devices."+mac+"."+room)
	     
	     
	     if(!device_id.Exists()){
		     fmt.Println("Device "+mac+" - "+room+" not configured for hubitat!")
		     return
	     }
	     
	     
	     url := "http://"+host+"/apps/api/8/devices/"+device_id.String()+"/arrived?access_token="+access_token
    
      //   fmt.Println(url)
         
        resp, err := http.Get(url)
        
        if err != nil {
			fmt.Println(err)
			return
		}
		
		
		resp.Body.Close()
		
		//fmt.Printf("Api Response: %d",resp.StatusCode)
	//	fmt.Println("")
    
}



func (d *Driver) Departed(mac string, room string){	
	
	     device_id := gjson.Get(config, "devices."+mac+"."+room)
	     
	     
	     if(!device_id.Exists()){
		     fmt.Println("Device "+mac+" - "+room+" not configured for hubitat!")
		     return
	     }
	     
	     
	     url := "http://"+host+"/apps/api/8/devices/"+device_id.String()+"/departed?access_token="+access_token
    
       //  fmt.Println(url)
         
        resp, err := http.Get(url)
        
       
        if err != nil {
			fmt.Println(err)
			return
		}
		
		resp.Body.Close()
		
	//	fmt.Printf("Api Response: %d",resp.StatusCode)
	//	fmt.Println("")
	
}


