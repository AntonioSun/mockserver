# mockserver
<!-- ALL-CONTRIBUTORS-BADGE:START - Do not remove or modify this section -->
[![All Contributors](https://img.shields.io/badge/all_contributors-3-orange.svg?style=flat-square)](#contributors-)
<!-- ALL-CONTRIBUTORS-BADGE:END -->

[![MIT License](http://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)
[![GoDoc](https://godoc.org/github.com/AntonioSun/mockserver?status.svg)](http://godoc.org/github.com/AntonioSun/mockserver)
[![Go Report Card](https://goreportcard.com/badge/github.com/AntonioSun/mockserver)](https://goreportcard.com/report/github.com/AntonioSun/mockserver)
[![Build Status](https://github.com/AntonioSun/mockserver/actions/workflows/go-release-build.yml/badge.svg?branch=master)](https://github.com/AntonioSun/mockserver/actions/workflows/go-release-build.yml)
[![PoweredBy WireFrame](https://github.com/go-easygen/wireframe/blob/master/PoweredBy-WireFrame-B.svg)](http://godoc.org/github.com/go-easygen/wireframe)

mock server to replace the Java/NPM counterpart mockserver

## TOC
- [mockserver - super slim & blazing fast mock server](#mockserver---super-slim-&-blazing-fast-mock-server)
- [Run](#run)
- [Sample mock.json file](#sample-mockjson-file)
- [Build](#build)
- [Install Debian/Ubuntu package](#install-debianubuntu-package)
- [Download/install binaries](#downloadinstall-binaries)
  - [The binary executables](#the-binary-executables)
- [Install Source](#install-source)
- [Author](#author)
- [Contributors](#contributors-)

## mockserver - super slim & blazing fast mock server

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

## Download/install binaries

- The latest binary executables are available 
as the result of the Continuous-Integration (CI) process.
- I.e., they are built automatically right from the source code at every git release by [GitHub Actions](https://docs.github.com/en/actions).
- There are two ways to get/install such binary executables
  * Using the **binary executables** directly, or
  * Using **packages** for your distro

### The binary executables

- The latest binary executables are directly available under  
https://github.com/AntonioSun/mockserver/releases/latest 
- Pick & choose the one that suits your OS and its architecture. E.g., for Linux, it would be the `mockserver_verxx_linux_amd64.tar.gz` file. 
- Available OS for binary executables are
  * Linux
  * Mac OS (darwin)
  * Windows
- If your OS and its architecture is not available in the download list, please let me know and I'll add it.
- The manual installation is just to unpack it and move/copy the binary executable to somewhere in `PATH`. For example,

``` sh
tar -xvf mockserver_*_linux_amd64.tar.gz
sudo mv -v mockserver_*_linux_amd64/mockserver /usr/local/bin/
rmdir -v mockserver_*_linux_amd64
```

## Install Source

To install the source code instead:

```
go get -v -u github.com/AntonioSun/mockserver
```

## Author

Antonio SUN

_Powered by_ [**WireFrame**](https://github.com/go-easygen/wireframe)  
[![PoweredBy WireFrame](https://github.com/go-easygen/wireframe/blob/master/PoweredBy-WireFrame-Y.svg)](http://godoc.org/github.com/go-easygen/wireframe)  
the _one-stop wire-framing solution_ for Go cli based projects, from _init_ to _deploy_.

## Contributors ✨

Thanks goes to these wonderful people ([emoji key](https://allcontributors.org/docs/en/emoji-key)):

<!-- ALL-CONTRIBUTORS-LIST:START - Do not remove or modify this section -->
<!-- prettier-ignore-start -->
<!-- markdownlint-disable -->
<table>
  <tr>
    <td align="center"><a href="https://github.com/AntonioSun"><img src="https://avatars.githubusercontent.com/u/2840074?v=4?s=100" width="100px;" alt=""/><br /><sub><b>AntonioSun</b></sub></a><br /><a href="https://github.com/AntonioSun/mockserver/commits?author=AntonioSun" title="Code">💻</a> <a href="#ideas-AntonioSun" title="Ideas, Planning, & Feedback">🤔</a> <a href="#design-AntonioSun" title="Design">🎨</a> <a href="#data-AntonioSun" title="Data">🔣</a> <a href="https://github.com/AntonioSun/mockserver/commits?author=AntonioSun" title="Tests">⚠️</a> <a href="https://github.com/AntonioSun/mockserver/issues?q=author%3AAntonioSun" title="Bug reports">🐛</a> <a href="https://github.com/AntonioSun/mockserver/commits?author=AntonioSun" title="Documentation">📖</a> <a href="#blog-AntonioSun" title="Blogposts">📝</a> <a href="#example-AntonioSun" title="Examples">💡</a> <a href="#tutorial-AntonioSun" title="Tutorials">✅</a> <a href="#tool-AntonioSun" title="Tools">🔧</a> <a href="#platform-AntonioSun" title="Packaging/porting to new platform">📦</a> <a href="https://github.com/AntonioSun/mockserver/pulls?q=is%3Apr+reviewed-by%3AAntonioSun" title="Reviewed Pull Requests">👀</a> <a href="#question-AntonioSun" title="Answering Questions">💬</a> <a href="#maintenance-AntonioSun" title="Maintenance">🚧</a> <a href="#infra-AntonioSun" title="Infrastructure (Hosting, Build-Tools, etc)">🚇</a></td>
    <td align="center"><a href="https://ananto.netlify.app/"><img src="https://avatars.githubusercontent.com/u/15931537?v=4?s=100" width="100px;" alt=""/><br /><sub><b>Ananto</b></sub></a><br /><a href="https://github.com/AntonioSun/mockserver/commits?author=Ananto30" title="Code">💻</a> <a href="#ideas-Ananto30" title="Ideas, Planning, & Feedback">🤔</a> <a href="#design-Ananto30" title="Design">🎨</a> <a href="#data-Ananto30" title="Data">🔣</a> <a href="https://github.com/AntonioSun/mockserver/commits?author=Ananto30" title="Tests">⚠️</a> <a href="https://github.com/AntonioSun/mockserver/commits?author=Ananto30" title="Documentation">📖</a> <a href="#example-Ananto30" title="Examples">💡</a></td>
    <td align="center"><a href="https://github.com/smicyk"><img src="https://avatars.githubusercontent.com/u/14974939?v=4?s=100" width="100px;" alt=""/><br /><sub><b>smicyk</b></sub></a><br /><a href="#data-smicyk" title="Data">🔣</a> <a href="https://github.com/AntonioSun/mockserver/commits?author=smicyk" title="Tests">⚠️</a></td>
  </tr>
</table>

<!-- markdownlint-restore -->
<!-- prettier-ignore-end -->

<!-- ALL-CONTRIBUTORS-LIST:END -->

This project follows the [all-contributors](https://github.com/all-contributors/all-contributors) specification. Contributions of any kind welcome!
