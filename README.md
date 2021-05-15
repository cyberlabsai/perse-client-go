[<img src="https://raw.githubusercontent.com/Yoonit-Labs/nativescript-yoonit-camera/development/logo_cyberlabs.png" width="300">](https://cyberlabs.ai/)

# perse-client-go

This is an example client implementation written in Go to interact the Perse API.

For more information, read the [API documentation](https://apidocs.cyberface.ai/).

## Run the example

```bash
export API_KEY=<provided api key>

go build -o example/main example/main.go

cd example

./main
```

## How to use

### Getting the package

First, download the package with:

```bash
go get -v -u github.com/cyberlabsai/perse-client-go
```

Then, make sure to import the package using:

```bash
perse "github.com/cyberlabsai/perse-client-go"
```

### Creating the client instance

Then, you need to create a client with:

```
client := perse.New(apiKey, nil)
```

To get an apiKey ( string ), contact our sales team C:

the second parameter ( *http.Client ) is an istance of go's http client. Create one if you want to customize it's http options. If you send nil, a default one will be created.

### Detecting faces

To detect faces in a frame, use the "DetectFaces" method.

```bash
faces, err := client.DetectFaces(img1, "image/jpg")
```

```bash
// function header
(faceRecClient *FaceRecClient) DetectFaces(imgData []byte, contentType string) ([]byte, error)
```

It takes an *os.File (aka a file that you can open with os.Open). It returns a []byte with the data received from the server. You can use the json.Unmarshal to parse it. the faces will be nil in case of error.

### Comparing faces

You can use the method "FaceCompare" to compare two faces.

```
compare, err := client.FaceCompare(img1, img2, "image/jpg")
```

```
// function header
(faceRecClient *FaceRecClient) FaceCompare(img1 []byte, img2 []byte, contentType string) ([]byte, error)
```

It takes a list with either the image tokens returned by the detectFace or the *os.File (or even both).

It will return a []byte with the server's response and a error.


## Examples

For further details, take a look on the "examples" directory. You will find some usefull pieces of code.

[![GitHub license](https://img.shields.io/github/license/Naereen/StrapDown.js.svg)](https://github.com/Naereen/StrapDown.js/blob/master/LICENSE)