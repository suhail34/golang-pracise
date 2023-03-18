package main

import "fmt"

func main() {
	fmt.Println("Maps in golang")

	languages := make(map[string]string)

	languages["JS"] = "Javascript"
	languages["PY"] = "Python"
	languages["TS"] = "Typescript"
	languages["GO"] = "Golang"

	fmt.Println("HashMap of languages : ", languages)
	fmt.Println("JS Short For : ", languages["JS"])

	delete(languages, "TS")
	fmt.Println("HashMap of languages : ", languages)

	for _, value := range languages {
		fmt.Printf("For key v value is %v\n", value)
	}

}
