package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path"
)

func main() {
	ip := "40.73.38.99"
	userId := 1000
	flowId := 1000
	lbHome := "./"

	trainLogDir := path.Join(lbHome, "logs/trainlogs")
	err := os.MkdirAll(trainLogDir, 0777)
	if err != nil {
		log.Fatal("Can't create train log dir.")
		return
	}
	trainLogName := fmt.Sprintf("%s_user_%d_flow_%d.log", ip, userId, flowId)
	trainLogPath := path.Join(trainLogDir, trainLogName)
	log.Println("trainLogPath : ", trainLogPath)

	f, err := os.Create(trainLogPath)
	if err != nil {
		log.Fatal("Failed to create train log file.")
		return
	}
	defer f.Close()

	trainerPort := 28080
	url := fmt.Sprintf("http://%s:%d/train/log/littleboy-trainer.log", ip, trainerPort)
	log.Println("Get Train log from ", url)
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal("Failed to new request:", err)
		return
	}

	client := http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		log.Fatal("Failed to do request: ", err)
		return
	}
	defer resp.Body.Close()

	_, err = io.Copy(f, resp.Body)
	if err != nil {
		log.Fatal("Failed to get train log: ", err)
		return
	}

	return
}
