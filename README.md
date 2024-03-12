# golang-zerotier-api
A Golang wrapper for the Zerotier Central API

# Examples
```golangnetlist, err := getNetworkList()

if err != nil {
	log.Panic(err)
}
for l := range netlist {
	fmt.Println(netlist[l].Config.Name)
}
``` 

```golang
member, err := getNetworkMemberList("48d6023c4689edac")

if err != nil {
	log.Panic(err)
}

for m := range member {
	if len(member[m].Config.IPAssignments) > 0 {
		fmt.Println(member[m].Name + " = " + member[m].Config.IPAssignments[0])
	}
}
``` 

```golang
newNet, _ := createNetwork()

fmt.Println(newNet.ID)
``` 

```golang
net, _ := getNetwork("9e1948db63a2684c")

var update UpdateNetwork
update.Config = net.Config
update.Config.Name = "b-van-bok"

updateNet, err := updateNetwork("9e1948db63a2684c", update)
if err != nil {
	log.Panic(err)
}

fmt.Println(updateNet.Config.ID + " = " + updateNet.Config.Name)
``` 

```golang
deleteNetwork("e3918db48360780a")
``` 

```golang
member, err := getNetworkMember("48d6023c4689edac", "0db3a4fa49")

if err != nil {
	log.Panic(err)
}

if len(member.Config.IPAssignments) > 0 {
	fmt.Println(member.Name + " = " + member.ID)
}
``` 

```golang
var update UpdateNetworkMember
update.Name = "ruimtekruiser"
update.Config.Authorized = true

updateMem, err := updateNetworkMember("48d6023c4689edac", "48d6023c4689edac-0db3a4fa49", update)
if err != nil {
	log.Panic(err)
}

fmt.Println(updateMem.Config.ID + " = " + updateMem.Name)
``` 

```golang
deleteNetworkMember("48d6023c4689edac", "1137e233db")
``` 


