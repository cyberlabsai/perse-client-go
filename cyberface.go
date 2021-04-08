package cyberface

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"strings"
	"time"
)

// FaceRecClient Represents the face recognition api
type FaceRecClient struct {
	apiKey     string
	url        string
	httpClient http.Client
}

const httpRequestTimeOut = 5 * time.Second
const apiAddress = "https://76f5ey2m6j.execute-api.us-east-1.amazonaws.com"

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

func (faceRecClient *FaceRecClient) FaceCompare(images []interface{}) ([]byte, error) {
	if len(images) != 2 {
		return nil, errors.New("images must have 2 elements")
	}

	buffer := new(bytes.Buffer)

	writer := multipart.NewWriter(buffer)

	for index, data := range images {
		if testedData, ok := data.(*os.File); ok {
			fileWriter, err := writer.CreateFormFile(fmt.Sprintf("image_file%d", index+1), testedData.Name())
			if err != nil {
				return nil, err
			}

			if _, err = io.Copy(fileWriter, testedData); err != nil {
				return nil, err
			}

		} else {
			if testedData, ok := data.(string); ok {
				fileWriter, err := writer.CreateFormField(fmt.Sprintf("image_token%d", index+1))
				if err != nil {
					return nil, err
				}

				readerData := strings.NewReader(testedData)

				if _, err := io.Copy(fileWriter, readerData); err != nil {
					return nil, err
				}

			} else {
				return nil, errors.New("image in the list must be either a *os.File (the image file) or a string (image token)")
			}
		}
	}

	err := writer.Close()
	if err != nil {
		return nil, err
	}

	urlWithPath := fmt.Sprintf("%s%s", faceRecClient.url, "/v0/face/compare")

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

func (faceRecClient *FaceRecClient) DetectFaces(imageData *os.File) ([]byte, error) {
	buffer := new(bytes.Buffer)

	writer := multipart.NewWriter(buffer)

	fileWriter, err := writer.CreateFormFile("image_file", imageData.Name())
	if err != nil {
		return nil, err
	}

	_, err = io.Copy(fileWriter, imageData)
	if err != nil {
		return nil, err
	}

	err = writer.Close()
	if err != nil {
		return nil, err
	}

	urlWithPath := fmt.Sprintf("%s%s", faceRecClient.url, "/v0/face/detect")

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
