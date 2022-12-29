
# Gogix - Golang HTTP client

* [Description](#description)
* [Installation](#installation)
* [Usage](#usage)
  + [Making a simple `GET` request](#making-a-simple-get-request)
  + [Making a `POST` request with payload and header](#making-a-post-request-with-payload-and-header)
  + [Custom HTTP clients](#custom-http-clients)
* [Plugins](#plugins)
* [Help and docs](#help-and-docs)
* [FAQ](#faq)
* [License](#license)

## Description

Gogix is a Golang HTTP client that makes it easy to send HTTP requests and trivial to integrate with web services.
- Simple code for building HTTP request(GET,POST,PUT,PATCH,DELETE)
- Create clients with different timeouts for every request
- Create clients with Custom 

## Installation
```
go get -u github.com/doniantoro/gogix
```
## Usage 
### Making a simple `GET` request
The below example will print the contents of the google home page:
```go
// Create a new HTTP client with a default timeout
client := gogix.NewClient(10)
// Use the clients GET method to create and execute the request
res,code, err := client.Get("https://api.publicapis.org/entries", nil)
if err != nil{
	panic(err)
}
// Gogix returns the []byte
fmt.Println(string(response))
fmt.Println(code)
```

### Making a `POST` request with payload and header
The below example will print the contents of the google home page:
```go
// Create a new HTTP client with a default timeout
client := gogix.NewClient(10)


header := gogix.Header()
header.Set("Content-Type", "application/json")
header.Set("AccessToken", "test")

//payload should be struct
payload := Order{}
payload.Msisdn = "089526265660"
  
// Use the clients POST method to create and execute the request
// *note the example endpoint is only for get , you can change with your api
res,code, err := client.Post("https://api.publicapis.org/entries", header,payload)
if err != nil{
	panic(err)
}
// Gogix returns the []byte
fmt.Println(string(response))
fmt.Println(code)
```


### Custom HTTP clients
The below example will print the contents of the google home page:
```go
// Create a new HTTP client with a default timeout
tr := &http.Transport{
		TLSClientConfig:     &tls.Config{InsecureSkipVerify: true},
		MaxIdleConns:        100,
		MaxIdleConnsPerHost: 100,
	}
	client := &http.Client{
		Transport: tr,
		Timeout:   time.Duration(30) * time.Second,
	}
client := gogix.CustomClient(tr)
// Use the clients GET method to create and execute the request
res,code, err := client.Get("https://api.publicapis.org/entries", nil)
if err != nil{
	panic(err)
}
// Gogix returns the []byte
fmt.Println(string(response))
fmt.Println(code)
```


## Help and docs
We use GitHub issues only to discuss bugs and new features. For support please refer to:

- [Documentation](https://pkg.go.dev/github.com/doniantoro/gogix)
- [Medium ( soon) ](https://)

## License

```
Copyright 2022, Doni Antoro

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
```
