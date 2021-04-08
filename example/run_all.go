package main

import (
	"fmt"

	cyberFace "github.com/cyberlabsai/cyberface-client-go"
)

func main() {
	fmt.Println("inicio")

	client := cyberFace.New("key here", nil)

	fmt.Println("upando 1")
	id1, err1 := client.UploadImageFromPath("lula1.jpeg")
	fmt.Println("upando 2")
	id2, err2 := client.UploadImageFromPath("lula2.jpeg")

	fmt.Println(id1, err1)
	fmt.Println(id2, err2)

	fmt.Println("detectando 1")
	faces1, err1 := client.DetectFacesUUID(id1)
	fmt.Println("detectando 2")
	faces2, err2 := client.DetectFacesUUID(id2)

	fmt.Println(string(faces1), err1)
	fmt.Println(string(faces2), err2)

	fmt.Println("comparando")
	compare, err := client.FaceCompareUUID(id1, id2)

	fmt.Println(string(compare), err)

	fmt.Println("fim")
}
