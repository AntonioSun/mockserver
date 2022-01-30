## {{toc 5}}
- [Install Debian/Ubuntu package](#install-debianubuntu-package)
- [Download/install binaries](#downloadinstall-binaries)
  - [The binary executables](#the-binary-executables)
- [Install Source](#install-source)
- [Author](#author)
- [Contributors](#contributors-)

## {{.Name}} - super slim & blazing fast mock server

Run a blazing fast mock server in just seconds! 🚀

All you need is to make a json file that contains request and response mapping. See an example [here](#sample-mockjson-file).

## Run
With defaults - 
```bash
./mockserver
```
**Defaults: `addr=localhost:7070` , `file=mock.json`**


With custom flags - 
```bash
./mockserver -addr <YOUR_HOST_AND_PORT> -file <MOCK_JSON_FILE_LOCATION>
```


For windows - 
```powershell
mockserver.exe -addr <YOUR_HOST_AND_PORT> -file <MOCK_JSON_FILE_LOCATION>
```

## Sample mock.json file


Example - 
```json
[
  {
    "httpRequest" : {
      "method": "POST",
      "path" : "/login",
      "headers": {
        "Content-Type": ["application/x-www-form-urlencoded; charset=UTF-8"]
      },
      "body": {
        "type": "STRING",
        "string": "username=john&password=john",
        "contentType" : "application/x-www-form-urlencoded; charset=UTF-8"
      }
    },
  {
    "httpRequest": {
      "method": "GET",
      "path": "/api/books",
      "queryStringParameters": {
        "limit": [ "[0-9]+" ]
      },
      "cookies": {
        "sessionId" : "055CA455-1DF7-45BB-8535-4F83E7266092"
      }
    },
    "httpResponse": {
      "body": "[{\"id\": 234, \"title\": \"Book number 234\"},{\"id\": 432, \"title\": \"Book number 432\"}]",
      "delay": {
        "timeUnit": "MILLISECONDS",
        "value": 750
      }
    }
  }
]
```

These `httpRequest.path`s will be matched and the response will be sent. E.g., if a request lands in the server in path `/api/books` the json object inside `httpResponse.body` will be sent as response. The full sample is available [here](https://github.com/AntonioSun/mockserver/blob/main/mock.json), and here is how it works:

``` sh
curl -X POST -d "username=john&password=john" localhost:7070/login
<html><head><meta name=\"trackId\" content=\"19293921933\"></head><body></body</html>

$ curl -L localhost:7070/api/books
"[{\"id\": 234, \"title\": \"Book number 234\"},{\"id\": 432, \"title\": \"Book number 432\"}]"

$ curl -L localhost:7070/api/books/234
"{\"id\": 234, \"title\": \"Book number 234\", \"author\": {\"id\": 523, \"name\": \"Author Name\"}}"

$ curl -L localhost:7070/api/authors/523
"{\"id\": 523, \"name\": \"Author Name\", \"bio\": \"Author bio\"}"
```

Notes:

- The mock file defines the rules that determine how the server should respond to a request.
- We use a rule-based system to match requests to responds. Therefore, you have to organize them from most restrictive to least. 
- **The request type [POST or GET] doesn't matter.**

## Build
For mac/linux - 
```bash
go mod download
go build
```

For windows -
```bash
go mod download
GOOS=windows GOARCH=amd64 go build 
```

**If the build/binary doesn't work for you, you can do this -

- Check your os and arch using this command - `go env GOOS GOARCH`
- Use the output os and arch to build the binary - `GOOS=<YOUR_OS> GOARCH=<YOUR_ARCH> go build`
