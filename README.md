# Bluetooth Presence 

This code allows you to setup bluetooth nodes on your network that detect which device you are closer to

## Building 

This is used to build for a raspberry pi
```
env GOOS=linux GOARCH=arm GOARM=5 go build -o discovery-pi
```

## Web Interface

By default there is a web interface exposed on port 15784 that lets you see nodes, etc.

## Operating Rationale

Each node looks for others on the network, through a simple conseus process they elect a master and then send all the location data to that node, that node then connects to whatever automation system is configured in drivers.

## Drivers

Drivers can be added to the drivers folder and you must create two methods 

Arrived(mac string,room string)
Departed(mac string,room string)
