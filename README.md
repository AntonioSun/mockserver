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
  - [Verbosity](#verbosity)
  - [Restriction level](#restriction-level)
  - [Prefork mode](#prefork-mode)
- [Sample mock.json file](#sample-mockjson-file)
- [Build](#build)
- [Install Debian/Ubuntu package](#install-debianubuntu-package)
- [Download/install binaries](#downloadinstall-binaries)
  - [The binary executables](#the-binary-executables)
- [Install Source](#install-source)
- [Author](#author)
- [Contributors](#contributors-)

## mockserver - super slim & blazing fast mock server

Run a super slim and blazing fast mock server in just seconds! 🚀

- How small is it? The executable is only ~6M in size.
- How fast is it (the Go fasthttp framework)? It's up to [10x faster](https://golangrepo.com/repo/valyala-fasthttp-go-network) than the default Go net/http package, [6~7 times](https://www.techempower.com/benchmarks/#section=data-r19&hw=ph&test=plaintext) faster than `nodejs`, and [38x faster](https://www.techempower.com/benchmarks/#section=data-r19&hw=ph&test=plaintext) than Java `httpserver`.

It can be used to stress test the stress-testing tools, even when they run in distributed multi-node mode. It helps to answer questions like what would the bottlenecks be when JMeter testing is running from a single node, or doubled, or tripled, etc? Therefore, the delay-responding will not likely be implemented in this tool, as there are so many tools out there already doing the throttling.

Now you can have a mock server that itself is not the bottleneck. All you need, is to make a json file that contains request and response mapping. See an example [here](#sample-mockjson-file).

## Run
With defaults - 
```bash
./mockserver
```
**Defaults: `MS_ADDR=:7070 MS_FILE=mock.json MS_VERBOSE=1 MS_RESTRICT=0`**, and `MS_COMPRESS` and `MS_PREFORK` will be `false` by default.


With custom settings - 
```bash
export MS_ADDR=<YOUR_HOST_AND_PORT> MS_FILE=<MOCK_JSON_FILE_LOCATION>
MS_VERBOSE=2 ./mockserver
```

Full list of custom environment settings:

- **MS_ADDR**: Server address (string=":7070")
- **MS_COMPRESS**: Enable transparent response compression (bool)
- **MS_FILE**: Mock json file location (string="mock.json")
- **MS_PREFORK**: Boost performance by prefork (bool)
- **MS_RESTRICT**: Restriction level (default: relaxed, only request's path will be matched) (int)
- **MS_VERBOSE**: Verbose mode (higher numbers increase the verbosity) (int="1")

### Verbosity

The default verbosity setting is `MS_VERBOSE=1`. It'll print all available paths on program start:

``` bash
$ mockserver
✔ Successfully opened: mock.json
✔ Successfully parsed: mock.json
Available paths: 
=> /login
=> /api/books/234/comments
=> /api/books/234
=> /api/books
=> /api/authors/523
Starting server on :7070
```

To suspense the printing of available paths on program start, set `MS_VERBOSE=0`:

``` bash
$ MS_VERBOSE=0 mockserver
✔ Successfully opened: mock.json
✔ Successfully parsed: mock.json
Starting server on :7070
^C
```

When verbosity is `2`, it'll print a log of all requesting paths:

``` bash
$ MS_VERBOSE=2 mockserver
✔ Successfully opened: mock.json
✔ Successfully parsed: mock.json
Available paths: 
=> /login
=> /api/books/234/comments
=> /api/books/234
=> /api/books
=> /api/authors/523
Starting server on :7070
2022/01/30 16:27:28 /login
```

With verbosity being `3`, the `mock.json` input is dumped in Go internal format, and the IP of the requesting is logged on console as well:

```shell
$ MS_VERBOSE=3 mockserver
✔ Successfully opened: mock.json
✔ Successfully parsed: mock.json
] [{HTTPRequest:{Headers:{ContentType:[application/x-www-form-urlencoded; charset=UTF-8]} Method:POST Path:/login QueryStringParameters:map[] Cookies:map[] Body:{Type:STRING String:username=john&password=john ContentType:application/x-www-form-urlencoded; charset=UTF-8}} HTTPResponse:{StatusCode:0 Body:<html><head><meta name="trackId" content="19293921933"></head><body></body</html> Cookies:map[sessionId:055CA455-1DF7-45BB-8535-4F83E7266092]}} {HTTPRequest:{Headers:{ContentType:[application/json]} Method:POST Path:/api/books/234/comments QueryStringParameters:map[] Cookies:map[sessionId:055CA455-1DF7-45BB-8535-4F83E7266092] Body:{Type: String: ContentType:}} HTTPResponse:{StatusCode:200 Body:{"Success": True} Cookies:map[]}} {HTTPRequest:{Headers:{ContentType:[]} Method:GET Path:/api/books/234 QueryStringParameters:map[] Cookies:map[sessionId:055CA455-1DF7-45BB-8535-4F83E7266092] Body:{Type: String: ContentType:}} HTTPResponse:{StatusCode:0 Body:{"id": 234, "title": "Book number 234", "author": {"id": 523, "name": "Author Name"}} Cookies:map[]}} {HTTPRequest:{Headers:{ContentType:[]} Method:GET Path:/api/books QueryStringParameters:map[limit:[[0-9]+]] Cookies:map[sessionId:055CA455-1DF7-45BB-8535-4F83E7266092] Body:{Type: String: ContentType:}} HTTPResponse:{StatusCode:0 Body:[{"id": 234, "title": "Book number 234"},{"id": 432, "title": "Book number 432"}] Cookies:map[]}} {HTTPRequest:{Headers:{ContentType:[]} Method:GET Path:/api/authors/523 QueryStringParameters:map[] Cookies:map[sessionId:055CA455-1DF7-45BB-8535-4F83E7266092] Body:{Type: String: ContentType:}} HTTPResponse:{StatusCode:0 Body:{"id": 523, "name": "Author Name", "bio": "Author bio"} Cookies:map[]}}]
Available paths: 
=> /login
=> /api/books/234/comments
=> /api/books/234
=> /api/books
=> /api/authors/523
Starting server on :7070
2022/01/31 13:37:08 /login (127.0.0.1)
^C
```

### Restriction level

Restriction level: `int`. The default is relaxed, i.e., only request's path will be matched, which is the current implementation.

### Prefork mode

If somehow you need that last mile for `mockserver` to be more performant, to be [10% more faster](https://www.techempower.com/benchmarks/#section=data-r19&hw=ph&test=plaintext), you can enable the [prefork mode](https://pkg.go.dev/github.com/valyala/fasthttp/prefork) by setting the `MS_PREFORK` environment variable to `true`:

```shell
$ MS_PREFORK=true mockserver
✔ Successfully opened: mock.json
✔ Successfully parsed: mock.json
Available paths: 
=> /login
=> /api/books/234/comments
=> /api/books/234
=> /api/books
=> /api/authors/523
Starting preforked server on :7070
✔ Successfully opened: mock.json
✔ Successfully opened: mock.json
✔ Successfully opened: mock.json
✔ Successfully parsed: mock.json
. . .
```

It'll start multiple independent sub-processes, each process processes http requests independently. The number of child processes will be [the same as the  number of CPU cores](https://chowdera.com/2021/10/20211009000611119p.html) your machine has.


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
$ curl -X POST -d "username=john&password=john" localhost:7070/login
<html><head><meta name="trackId" content="19293921933"></head><body></body</html>

$ curl -L localhost:7070/api/books
[{"id": 234, "title": "Book number 234"},{"id": 432, "title": "Book number 432"}]

$ curl -L 'localhost:7070/api/books?limits=9'
[{"id": 234, "title": "Book number 234"},{"id": 432, "title": "Book number 432"}]

$ curl -L localhost:7070/api/books/234
{"id": 234, "title": "Book number 234", "author": {"id": 523, "name": "Author Name"}}

$ curl -L localhost:7070/api/authors/523
{"id": 523, "name": "Author Name", "bio": "Author bio"}
```

Notes:

- The mock file defines the rules that determine how the server should respond to a request.
- We use a rule-based system to match requests to responds. Therefore, you have to organize them from most restrictive to least. 
- *The request type [POST or GET] doesn't matter.*

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
