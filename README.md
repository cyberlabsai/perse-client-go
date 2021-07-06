[<img src="https://raw.githubusercontent.com/Yoonit-Labs/nativescript-yoonit-camera/development/logo_cyberlabs.png" width="300">](https://cyberlabs.ai/)

# perse-client-go

This is an example client implementation written in Go to interact the Perse API.

For more information, read the [API documentation](https://apidocs.cyberface.ai/).

## Run the example

```bash
export API_KEY=<provided api key>

cd example

go run main.go
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
faces, err := client.DetectFaces(image_with_some_face)
```

```bash
// function header
(faceRecClient *FaceRecClient) DetectFaces(image []byte) ([]byte, error)
```

It takes an []byte (the contents of an image file). It returns a []byte with the data received from the server. You can use the json.Unmarshal to parse it. the faces will be nil in case of error.

### Comparing faces

You can use the method "FaceCompare" to compare two faces.

```
compare, err := client.FaceCompare(image_with_some_face, another_image_with_some_face)
```

```
// function header
(faceRecClient *FaceRecClient) FaceCompare(image1 []byte, image2 []byte) ([]byte, error)
```

It takes the contents of two images as []bytes.

It will return a []byte with the server's response and a error.


## Examples

For further details, take a look on the "examples" directory. You will find some usefull pieces of code.

[![GitHub license](https://img.shields.io/github/license/Naereen/StrapDown.js.svg)](https://github.com/Naereen/StrapDown.js/blob/master/LICENSE)