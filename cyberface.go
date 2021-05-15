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
const APIURL = "https://api.getperse.com"

// New creates a new instance of a face recognition client
func New(key string, client *http.Client) *FaceRecClient {
	if client == nil {
		return &FaceRecClient{
			apiKey: key,
			url:    APIURL,
			httpClient: http.Client{
				Timeout: httpRequestTimeOut,
			},
		}
	}

	return &FaceRecClient{
		apiKey:     key,
		url:        APIURL,
		httpClient: *client,
	}
}

func (faceRecClient *FaceRecClient) FaceCompare(img1 []byte, img2 []byte, contentType string) ([]byte, error) {

	r := bytes.NewReader(img1)

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	fw, err := writer.CreateFormFile("image_file1", "img1.jpg")
	if err != nil {
		return nil, err
	}
	_, err = io.Copy(fw, r)
	if err != nil {
		return nil, err
	}

	r = bytes.NewReader(img2)

	fw, err = writer.CreateFormFile("image_file2", "img2.jpg")
	if err != nil {
		return nil, err
	}
	_, err = io.Copy(fw, r)
	if err != nil {
		return nil, err
	}

	urlWithPath := fmt.Sprintf("%s%s", faceRecClient.url, "/v0/face/compare")

	request, err := http.NewRequest("POST", urlWithPath, buffer)
	if err != nil {
		return nil, err
	}

	request.Header.Add("content-type", contentType)
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

func (faceRecClient *FaceRecClient) DetectFaces(imageData []byte, contentType string) ([]byte, error) {

	r := bytes.NewReader(imageData)

	urlWithPath := fmt.Sprintf("%s%s", faceRecClient.url, "/v0/face/detect")

	request, err := http.NewRequest("POST", urlWithPath, r)
	if err != nil {
		return nil, err
	}

	request.Header.Add("content-type", contentType)
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
