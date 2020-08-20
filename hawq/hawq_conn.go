package main

import (
	// "encoding/json"
	"fmt"
	commondb "github.com/oushu/gocommon/database"
	"github.com/oushu/gocommon/other"
	"github.com/oushu/lava-cloud/cloud/database"
	"log"
)

func main() {
	commondb.ConnectDB("oushu", "", "localhost", "4432", "postgres", "disable")
	requestHandler, err := database.CreateRequestHandler()
	if err != nil {
		log.Fatal(err)
	}
	masterNode, err := requestHandler.GetMasterNodeByClusterID(1)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println("masterNode: ")
	fmt.Println(other.EncodeJson(masterNode))
	standbyNode, err := requestHandler.GetStandbyNodeByClusterID(1)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println("standbyNode: ", standbyNode)
	fmt.Println(other.EncodeJson(standbyNode))
	requestUser, err := requestHandler.GetUserByUserId(-1000)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println("requestUser: ", requestUser)
	fmt.Println(other.EncodeJson(requestUser))
	hawqport, err := requestHandler.GetHawqPort(1)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println("hawqport: ", hawqport)
	fmt.Println(other.EncodeJson(hawqport))
}
