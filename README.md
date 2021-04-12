# cyberface-client-go

A simple SDK to help using the CyberFace API.


## How to use

### Getting the package

First, get the package with:

```
go get -v -u github.com/cyberlabsai/cyberface-client-go
```

Then, make sure to import the package using:

```
cyberFace "github.com/cyberlabsai/cyberface-client-go"
```

### Creating the client instance

Then, you need to create a client with:

```
client := cyberFace.New(apiKey, nil)
```

To get an apiKey ( string ), contact our sales team C:

the second parameter ( *http.Client ) is an istance of go's http client. Create one if you want to customize it's http options. If you send nil, a default one will be created.

### Detecting faces

To detect faces in a frame, use the "DetectFaces" method.

```
faces, err := client.DetectFaces(img1)
```

```
// function header
(faceRecClient *FaceRecClient) DetectFaces(imageData *os.File) ([]byte, error)
```

It takes an *os.File (aka a file that you can open with os.Open). It returns a []byte with the data received from the server. You can use the json.Unmarshal to parse it. the faces will be nil in case of error.

### Comparing faces

You can use the method "FaceCompare" to compare two faces.

```
compare, err := client.FaceCompare(compareData)
```

```
// function header
(faceRecClient *FaceRecClient) FaceCompare(images []interface{}) ([]byte, error)
```

It takes a list with either the image tokens returned by the detectFace or the *os.File (or even both).

It will return a []byte with the server's response and a error.


## Examples

For further details, take a look on the "examples" directory. You will find some usefull pieces of code.