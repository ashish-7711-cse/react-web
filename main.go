// package main

// import "fmt"

// func main() {
// 	adidasFactory, _ := GetSportsFactory("adidas")
// 	nikeFactory, _ := GetSportsFactory("nike")

// 	nikeShoe := nikeFactory.makeShoe()
// 	nikeShirt := nikeFactory.makeShirt()

// 	adidasShoe := adidasFactory.makeShoe()
// 	adidasShirt := adidasFactory.makeShirt()

// 	printShoeDetails(nikeShoe)
// 	printShirtDetails(nikeShirt)

// 	printShoeDetails(adidasShoe)
// 	printShirtDetails(adidasShirt)
// }

// func printShoeDetails(s IShoe) {
// 	fmt.Printf("Logo: %s", s.getLogo())
// 	fmt.Println()
// 	fmt.Printf("Size: %d", s.getSize())
// 	fmt.Println()
// }

// func printShirtDetails(s IShirt) {
// 	fmt.Printf("Logo: %s", s.getLogo())
// 	fmt.Println()
// 	fmt.Printf("Size: %d", s.getSize())
// 	fmt.Println()
// }

package main

import (
	"fmt"
)

// func main() {
// 	ak47, _ := getGun("ak47")
// 	musket, _ := getGun("musket")

// 	printDetails(ak47)
// 	printDetails(musket)
// }

// func printDetails(g IGun) {
// 	fmt.Printf("Gun: %s", g.getName())
// 	fmt.Println()
// 	fmt.Printf("Power: %d", g.getPower())
// 	fmt.Println()
// }

type InputValues struct {
	id    string
	value string
}

type ProvisionInputCommon struct {
	Network []InputValues
	Volumes []InputValues
	Name    string
	Version string
}

type InputStruct struct {
	peer       string
	InstanceId string
	network    []InputValues
	volume     []InputValues
	envvars    []InputValues
	name       string
	version    string
}

type Response struct {
	RespStatusCode int
	ResponseBytes  []byte
}

type ProvisionContainerInterface interface {
	Provision() Response
	start() Response
	stop() Response
	delete() Response
}

type Instance struct {
	ProvisionInput ProvisionInputCommon
	Instance_ID    string
	Peer_ID        string
	ConfigVersion  string
}

type activeInstance struct {
	Instance
}

type standbyInstance struct {
	Instance
}

func newActive(input InputStruct) ProvisionContainerInterface {
	return &activeInstance{
		Instance: Instance{
			Instance_ID: input.name,
		},
	}
}

func newStandby(input InputStruct) ProvisionContainerInterface {
	return &standbyInstance{
		Instance: Instance{
			Instance_ID: input.name,
		},
	}
}

func (i *Instance) start() Response {
	fmt.Println("Started")
	createResp := Response{}
	createResp.RespStatusCode = 200
	return createResp
}

func (i *Instance) stop() Response {
	fmt.Println("Stopped")
	createResp := Response{}
	createResp.RespStatusCode = 200
	return createResp
}

func (i *Instance) delete() Response {
	fmt.Println("Deleted")
	createResp := Response{}
	createResp.RespStatusCode = 200
	return createResp
}

func getInstanceFactory(instanceType string, input InputStruct) (ProvisionContainerInterface, error) {
	if instanceType == "active" {
		return newActive(input), nil
	} else if instanceType == "standby" {
		// In case of standby instance, userInput will have additional attributes like peer-id and replica-id
		return newStandby(input), nil
	}
	return nil, fmt.Errorf("Wrong gun type passed")
}

func (a *activeInstance) Provision() Response {
	createResp := Response{}
	createResp.RespStatusCode = 200
	fmt.Println("Active Created ______")
	return createResp
}

func (s *standbyInstance) Provision() Response {
	createResp := Response{}
	createResp.RespStatusCode = 400
	fmt.Println("Standby Created_____")
	return createResp
}

func main() {
	active, _ := getInstanceFactory("active", InputStruct{})
	active.Provision()
	standby, _ := getInstanceFactory("standby", InputStruct{})
	standby.Provision()
}
