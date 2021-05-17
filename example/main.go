package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	cyberFace "github.com/cyberlabsai/cyberface-client-go"
)

func getApiKey() string {
	value := os.Getenv("API_KEY")

	if value == "" {
		panic("API_KEY not found")
	}

	return value
}

func main() {
	fmt.Println("begin")

	img1, err := ioutil.ReadFile("./images/img1.jpeg")
	if err != nil {
		panic(fmt.Sprintf("problems opening first image:\n%s", err.Error()))
	}

	img2, err := ioutil.ReadFile("./images/img2.jpeg")
	if err != nil {
		panic(fmt.Sprintf("problems opening second image:\n%s", err.Error()))
	}

	client := cyberFace.New(getApiKey(), nil)

	fmt.Println("face detect")
	faces, err := client.DetectFaces(img1)

	fmt.Println(string(faces), err)

	fmt.Println("compare two images")

	frameData := struct {
		ImageToken string `json:"image_token"`
	}{}

	err = json.Unmarshal(faces, &frameData)
	if err != nil {
		panic(fmt.Sprintf("error while deserializing face detect data:\n%s", err.Error()))
	}

	result, err := client.FaceCompare(img1, img2)

	fmt.Println(string(result), err)

	fmt.Println("end")
}
