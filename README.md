# Golang-zerotier-api
A Golang wrapper for the Zerotier Central API.

This package was created for usage in the [GameKube](https://github.com/BumbleB-NL/gamekube) project. 

# Calls

The following API calls are currently implemented:

- Get network list
- Get network
- Create network
- Update network
- Delete network
- Get network member list
- Get network member
- Update network member (WIP)
- Delete network member

# Installation
Install the package with:

`go get github.com/PimSanders/golang-zerotier-api` 

# Usage
Create a new client:

```golang
zt := golangzerotierapi.NewClient("https://api.zerotier.com/api/v1", "AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA", false)
```

Get a list of networks:

```golang
netlist, err := zt.GetNetworkList()

if err != nil {
	log.Panic(err)
}
for l := range netlist {
	fmt.Println(netlist[l].Config.Name)
}
```
Update a network:

```golang
net, _ := zt.GetNetwork("BBBBBBBBBBBBBBBB")

var update golangzerotierapi.UpdateNetwork
update.Config = net.Config
update.Config.Name = "Example Name"

updateNet, err := updateNetwork("BBBBBBBBBBBBBBBB", update)
if err != nil {
	log.Panic(err)
}

fmt.Println(updateNet.Config.ID + " = " + updateNet.Config.Name)
``` 
