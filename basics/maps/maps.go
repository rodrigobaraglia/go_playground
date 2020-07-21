package main

import "fmt"

var tasks = map[string]string{
	"Monday":    "Study Golang",
	"Tuesday":   "Buy Cat Food",
	"Wednesday": "Finish S/Z"}

func main() {
	fmt.Println(tasks)
	delete(tasks, "Tuesday")
	for key, value := range tasks {
		fmt.Println(key, ": ", value)
	}

}
