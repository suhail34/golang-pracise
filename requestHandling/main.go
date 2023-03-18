package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("Webserver request handling")
	// PerformGetRequest("http://localhost:3000/get")
	// PerformPostJsonRequest("http://localhost:3000/post")
	PerformPostFormRequest("http://localhost:3000/postform")
}

func PerformGetRequest(myurl string) {
	response, err := http.Get(myurl)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	fmt.Println("Status Code : ", response.StatusCode)
	fmt.Println("Content Length is : ", response.ContentLength)

	var responseString strings.Builder

	content, _ := ioutil.ReadAll(response.Body)
	byteCount, _ := responseString.Write(content)

	fmt.Println("Bytecount is: ", byteCount)
	fmt.Println(responseString.String())
	// fmt.Println(content)
	// fmt.Println(string(content))
}

func PerformPostJsonRequest(myurl string) {

	requestBody := strings.NewReader(`
		{
			"course":"Golang",
			"price":0,
			"platform":"suhail12.hashnode.dev"
		}
	`)
	response, err := http.Post(myurl, "application/json", requestBody)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	content, err := ioutil.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}

	fmt.Println(string(content))
}

func PerformPostFormRequest(myurl string) {
	data := url.Values{}

	data.Add("firstname", "suhail")
	data.Add("lastname", "khan")
	data.Add("email", "skna@gmail.com")

	response, err := http.PostForm(myurl, data)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()
	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(content))

}
