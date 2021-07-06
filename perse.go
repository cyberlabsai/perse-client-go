package perse

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"time"
)

// FaceRecClient Represents the face recognition api
type FaceRecClient struct {
	apiKey     string
	url        string
	httpClient http.Client
}

const httpRequestTimeOut = 5 * time.Second
const apiAddress = "https://api.getperse.com"
const apiVersion = "v0"

// New creates a new instance of a face recognition client
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

func (faceRecClient *FaceRecClient) FaceCompare(image1 []byte, image2 []byte) ([]byte, error) {
	buffer := new(bytes.Buffer)

	writer := multipart.NewWriter(buffer)

	// image 1 ---------------------------
	fileWriter1, err := writer.CreateFormFile("image_file1", "img1.jpeg")
	if err != nil {
		return nil, err
	}

	if _, err = io.Copy(fileWriter1, bytes.NewReader(image1)); err != nil {
		return nil, err
	}

	// image 2 ---------------------------
	fileWriter2, err := writer.CreateFormFile("image_file2", "img2.jpeg")
	if err != nil {
		return nil, err
	}

	if _, err = io.Copy(fileWriter2, bytes.NewReader(image2)); err != nil {
		return nil, err
	}

	err = writer.Close()
	if err != nil {
		return nil, err
	}

	urlWithPath := fmt.Sprintf("%s/%s/%s", faceRecClient.url, apiVersion, "face/compare")

	request, err := http.NewRequest("POST", urlWithPath, buffer)
	if err != nil {
		return nil, err
	}

	request.Header.Add("content-type", writer.FormDataContentType())
	request.Header.Add("x-api-key", faceRecClient.apiKey)

	response, err := faceRecClient.httpClient.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return responseData, nil
}

func (faceRecClient *FaceRecClient) DetectFaces(image []byte) ([]byte, error) {
	buffer := new(bytes.Buffer)

	writer := multipart.NewWriter(buffer)

	fileWriter, err := writer.CreateFormFile("image_file", "img1.jpeg")
	if err != nil {
		return nil, err
	}

	_, err = io.Copy(fileWriter, bytes.NewReader(image))
	if err != nil {
		return nil, err
	}

	err = writer.Close()
	if err != nil {
		return nil, err
	}

	urlWithPath := fmt.Sprintf("%s/%s/%s", faceRecClient.url, apiVersion, "face/detect")

	request, err := http.NewRequest("POST", urlWithPath, buffer)
	if err != nil {
		return nil, err
	}

	request.Header.Add("content-type", writer.FormDataContentType())
	request.Header.Add("x-api-key", faceRecClient.apiKey)

	response, err := faceRecClient.httpClient.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return responseData, nil
}
