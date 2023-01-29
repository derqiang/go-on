package _json

import (
	"encoding/json"
	"fmt"
	"go_try/ch4/json/github"
	"log"
	"strings"
)

type JsonRunner int64

func (jr JsonRunner) Run() {
	var movies = []Movie{
		{Title: "Casablanca", Year: 1942, Color: false,
			Actors: []string{"Humphrey Bogart", "Ingrid Bergman"}},
		{Title: "Cool Hand Luke", Year: 1967, Color: true,
			Actors: []string{"Paul Newman"}},
		{Title: "Bullitt", Year: 1968, Color: true,
			Actors: []string{"Steve McQueen", "Jacqueline Bisset"}},
	}

	//data, err := json.Marshal(movies)
	data, err := json.MarshalIndent(movies, "", "    ")
	if err != nil {
		log.Fatalf("JSON marchaling failed : %s", err)
	}
	fmt.Printf("%s\n", data)

	result, err := github.SearchIssues(strings.Split("repo:golang/go is:open json decoder", " "))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues: \n", result.TotalCount)
	for _, item := range result.Items {
		fmt.Printf("#%-50d %9.9s %.55s\n", item.Number, item.User.Login, item.Title)
	}
}

type Movie struct {
	Title  string
	Year   int  `json:"released"`
	Color  bool `json:"color,omitempty"`
	Actors []string
}
