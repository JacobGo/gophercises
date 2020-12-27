package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

var scores = make(map[string]int)
type Score struct {
	Id string
	Points int
}

func handleScoreboard(w http.ResponseWriter, req *http.Request)  {
	switch req.Method {
	case "GET":
		data, err := json.Marshal(scores)
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprint(w, err)
		}
		fmt.Fprint(w, string(data))
	case "POST":
		decoder := json.NewDecoder(req.Body)
		var score Score
		err := decoder.Decode(&score)
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprint(w, err)
		}
		if scores[score.Id] < score.Points {
			prev := scores[score.Id]
			scores[score.Id] = score.Points
			fmt.Fprintf(w, "New high score! Congrats! You beat your previous score by %d", score.Points - prev)
		} else {
			if scores[score.Id] - score.Points == 0 {
				fmt.Fprintf(w, "Nice work, you tied your high score of %d", score.Points)
			} else {
				fmt.Fprintf(w, "So close! You missed your high score by %d", scores[score.Id] - score.Points)
			}

		}
	default:
		fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
	}

}

func main() {
	http.HandleFunc("/scoreboard", handleScoreboard)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
	}
}
