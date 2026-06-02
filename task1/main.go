package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Student struct {
	Name         string     `json:"name"`
	NrOfSubjects int        `json:"nrOfSubjects"`
	EachSubject  []subjects `json:"eachSubject"`
	Average      float64    `json:"average"`
}

type subjects struct {
	Name  string  `json:"name"`
	Grade float64 `json:"grade"`
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	var studentul Student
	var total float64

	file, _ := os.Create("data/output.json")

	defer file.Close()
	fmt.Print("Enter your name: ")
	name, _ := reader.ReadString('\n')
	studentul.Name = strings.TrimSpace(name)

	fmt.Print("Hello, %s!\n", name, ", please insert the number of subjects:")
	nrSub, _ := reader.ReadString('\n')
	nrSub = strings.TrimSpace(nrSub)
	studentul.NrOfSubjects, _ = strconv.Atoi(nrSub)

	fmt.Print("You have ", nrSub, " subjects.")

	for i := 1; i <= studentul.NrOfSubjects; i++ {
		var sub subjects
		fmt.Print("Enter the ", i, " subject name: ")
		name, _ := reader.ReadString('\n')
		sub.Name = strings.TrimSpace(name)

		fmt.Print("Enter the ", name, " grade: ")
		grade, _ := reader.ReadString('\n')
		sub.Grade, _ = strconv.ParseFloat(strings.TrimSpace(grade), 64)
		total += sub.Grade

		studentul.EachSubject = append(studentul.EachSubject, sub)
	}

	studentul.Average = total / float64(studentul.NrOfSubjects)

	b, _ := json.Marshal(studentul)

	encoder := json.NewEncoder(file)
	err := encoder.Encode(studentul)
	if err != nil {
		fmt.Println("Error encoding JSON:", err)
		return
	}

	fmt.Print("Media:", total/float64(studentul.NrOfSubjects))

	fmt.Print(string(b))

}
