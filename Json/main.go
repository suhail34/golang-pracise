package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string   `json:"coursename"`
	Price    int      `json:"price"`
	Platform string   `json:"platform"`
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("Json Data in golang")
	// EncodeJson()
	DecodeJson()
}

func EncodeJson() {
	Courses := []course{
		{"ReactJS", 299, "skna.com", "abc123", []string{"web-dev", "JS"}},
		{"Angular13", 199, "skna.com", "abc124", []string{"web-dev", "TS"}},
		{"Python", 200, "skna.com", "abd123", []string{"languages", "DSA"}},
		{"Devops", 399, "skna.com", "avc123", nil},
	}

	finalJSON, err := json.MarshalIndent(Courses, "", "\t")

	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", finalJSON)
}

func DecodeJson() {
	jsonData := []byte(`
		{
			"coursename": "Python",
			"price": 200,
			"platform": "skna.com",
			"tags": ["languages","DSA"]
		}
	`)

	var course course

	valid := json.Valid(jsonData)

	if valid {
		fmt.Println("JSON is Valid")
		json.Unmarshal(jsonData, &course)
		fmt.Printf("%#v\n", course)
	} else {
		fmt.Println("JSON not valid")
	}

	var myData map[string]interface{}
	json.Unmarshal(jsonData, &myData)
	fmt.Printf("%#v\n", myData)

	for k, v := range myData {
		fmt.Printf("Key is %v and value is %v and type is %T\n", k, v, v)
	}
}
