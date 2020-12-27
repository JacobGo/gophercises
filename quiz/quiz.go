package main

import (
	"bufio"
	"bytes"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"github.com/denisbrodbeck/machineid"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

type MyScore struct {
	Id string
	Points int
}

func showScoreAndQuit(score int, questions int) {
	fmt.Printf("\nYou scored %d out of %d\n", score, questions)

	url := "http://localhost:8080/scoreboard"

	client := http.Client{
		Timeout: time.Second * 2,
	}

	id, err := machineid.ID()
	myScore := MyScore{id, score}

	data, err := json.Marshal(myScore)
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(data))
	if err != nil {
		log.Fatal(err)
	}
	res, getErr := client.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}
	bodyBytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	bodyString := string(bodyBytes)
	fmt.Println(bodyString)


	os.Exit(0)
}
func main() {

	open, err := os.Open("./quiz/problems.csv")
	if err != nil {
		fmt.Print(err)
	}

	problems, err := csv.NewReader(open).ReadAll()
	if err != nil {
		fmt.Print(err)
	}


	reader := bufio.NewReader(os.Stdin)
	score := 0

	seconds := 5

	fmt.Printf("Answer as many questions you can within %d seconds! Press enter to start: ", seconds)
	_, _ = reader.ReadString('\n')

	timer := time.NewTimer(time.Duration(seconds) * time.Second)
	go func() {
		<-timer.C
		showScoreAndQuit(score, len(problems))
	}()


	for i, problem := range problems {
		question, answer := problem[0], problem[1]
		fmt.Printf("Problem #%d: %s = ", i, question)
		response, _ := reader.ReadString('\n')
		if strings.TrimSpace(response) == answer {
			score++
		}
	}
	showScoreAndQuit(score, len(problems))

}
