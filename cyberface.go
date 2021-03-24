package cyberface

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

// FaceRecClient Represents the face recognition api
type FaceRecClient struct {
	apiKey     string
	url        string
	httpClient http.Client
}

const httpRequestTimeOut = 5 * time.Second
const apiAddress = "https://vdvubhbr8e.execute-api.us-east-1.amazonaws.com"

// New creates a new instance of a face recog client
func New(key string, client *http.Client) *FaceRecClient {
	if client == nil {
		return &FaceRecClient{
			apiKey: key,
			url:    apiAddress,
			httpClient: http.Client{
				Timeout: httpRequestTimeOut,
			},
		}
	}

	return &FaceRecClient{
		apiKey:     key,
		url:        apiAddress,
		httpClient: *client,
	}
}

// UploadImageFromPath opens an image from "imagePath" as binary and sends to the server
// returns the image UUID if successful or an error with the error message sent by the server
func (faceRecClient *FaceRecClient) UploadImageFromPath(imagePath string) (string, error) {
	imageData, err := os.Open(imagePath)
	if err != nil {
		return "", err
	}
	defer imageData.Close()

	urlWithPath := fmt.Sprintf("%s/%s", faceRecClient.url, "/v0/upload")

	request, err := http.NewRequest("POST", urlWithPath, imageData)
	if err != nil {
		return "", err
	}

	request.Header.Add("content-type", "application/json")
	request.Header.Add("x-api-key", faceRecClient.apiKey)

	response, err := faceRecClient.httpClient.Do(request)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	responseData := struct {
		Body struct {
			UUID string `json:"UUID"`
		} `json:"body"`
	}{}

	err = json.NewDecoder(response.Body).Decode(&responseData)
	if err != nil {
		return "", err
	}

	return responseData.Body.UUID, nil
}

// FaceCompareUUID gets the json with the results of the face recognition for a given uuid
// returns the []byte with the json data and the error sent by the server or parser
func (faceRecClient *FaceRecClient) FaceCompareUUID(uuid1 string, uuid2 string) ([]byte, error) {
	urlWithPath := fmt.Sprintf("%s/%s", faceRecClient.url, "/v0/facecompare")

	request, err := http.NewRequest("GET", urlWithPath, nil)
	if err != nil {
		return make([]byte, 0), err
	}

	request.Header.Add("content-type", "application/json")
	request.Header.Add("x-api-key", faceRecClient.apiKey)

	url := request.URL.Query()

	url.Add("UUID_1", uuid1)
	url.Add("UUID_2", uuid2)

	request.URL.RawQuery = url.Encode()

	response, err := faceRecClient.httpClient.Do(request)
	if err != nil {
		return make([]byte, 0), err
	}
	defer response.Body.Close()

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return make([]byte, 0), err
	}

	return responseData, nil
}

// DetectFacesUUID gets the json with the results of the face detection for a given uuid
// returns the []byte with the json data and the error sent by the server or parser
func (faceRecClient *FaceRecClient) DetectFacesUUID(uuid string) ([]byte, error) {
	urlWithPath := fmt.Sprintf("%s/%s", faceRecClient.url, "/v0/facedetect")

	request, err := http.NewRequest("GET", urlWithPath, nil)
	if err != nil {
		return make([]byte, 0), err
	}

	request.Header.Add("content-type", "application/json")
	request.Header.Add("x-api-key", faceRecClient.apiKey)

	response, err := faceRecClient.httpClient.Do(request)
	if err != nil {
		return make([]byte, 0), err
	}
	defer response.Body.Close()

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return make([]byte, 0), err
	}

	return responseData, nil
}

// DetectFaces shortcut to open a file and detect a face
func (faceRecClient *FaceRecClient) DetectFaces(imagePath string) ([]byte, error) {
	UUID, err := faceRecClient.UploadImageFromPath(imagePath)
	if err != nil {
		return make([]byte, 0), err
	}

	imageData, err := faceRecClient.DetectFacesUUID(UUID)
	if err != nil {
		return make([]byte, 0), err
	}

	return imageData, nil
}

// FacRecognize shortcut to open the two image files and compare their faces
func (faceRecClient *FaceRecClient) FacRecognize(imagePath1 string, imagePath2 string) ([]byte, error) {
	UUID1, err := faceRecClient.UploadImageFromPath(imagePath1)
	if err != nil {
		return make([]byte, 0), err
	}

	UUID2, err := faceRecClient.UploadImageFromPath(imagePath2)
	if err != nil {
		return make([]byte, 0), err
	}

	imageData, err := faceRecClient.FaceCompareUUID(UUID1, UUID2)
	if err != nil {
		return make([]byte, 0), err
	}

	return imageData, nil
}
