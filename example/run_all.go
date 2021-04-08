package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	cyberFace "github.com/cyberlabsai/cyberface-client-go"
)

func getKey() string {
	file, err := os.Open(".env")
	if err != nil {
		panic(fmt.Sprintf("cannot open the .env file:\n%s", err.Error()))
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		panic(fmt.Sprintf("cannot read file contents:\n%s", err.Error()))
	}

	splited := strings.Split(string(data), "=")

	return splited[1]
}

func main() {
	fmt.Println("begin")

	img1, err := os.Open("./images/musk1.jpeg")
	if err != nil {
		panic(fmt.Sprintf("problems opening first image:\n%s", err.Error()))
	}
	defer img1.Close()

	img2, err := os.Open("./images/musk1.jpeg")
	if err != nil {
		panic(fmt.Sprintf("problems opening second image:\n%s", err.Error()))
	}
	defer img2.Close()

	apiKey := getKey()

	client := cyberFace.New(apiKey, nil)

	fmt.Println("face detect")
	faces, err := client.DetectFaces(img1)

	fmt.Println(string(faces), err)

	fmt.Println("compare")

	frameData := struct {
		ImageToken string `json:"image_token"`
	}{}

	err = json.Unmarshal(faces, &frameData)
	if err != nil {
		panic(fmt.Sprintf("problems desserializing face detect data:\n%s", err.Error()))
	}

	compareData := make([]interface{}, 0)

	compareData = append(compareData, frameData.ImageToken)
	compareData = append(compareData, img2)

	compare, err := client.FaceCompare(compareData)

	fmt.Println(string(compare), err)

	fmt.Println("end")
}
