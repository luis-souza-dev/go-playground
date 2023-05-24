package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"

	"github.com/gocolly/colly"
)

// Mail is the container of a single e-mail
type Table struct {
	Team    string `json:"team"`
	Position int `json:"position"`
	Points int `json:"points"`
	GoalDifference int `json:"goalDifference"`
	MatchesPlayed int `json:"matchesPlayed"`
}

func main() {
	
	premierTable := make([]Table, 0, 20);
	
	c := colly.NewCollector()

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})
	
	c.OnScraped(func(r *colly.Response) {
		fmt.Println(r.Request.URL, " scraped!");
		js, _ := json.Marshal(premierTable);
		os.WriteFile("premiertable.json", js, 0644)
	})

	c.OnHTML("div.box.tab-print div.box div[id=yw5] table.items tbody tr", func(h *colly.HTMLElement) {

		pos, _ := strconv.Atoi(h.ChildText("td:nth-child(1)"))
		matchesPlayed, _ := strconv.Atoi(h.ChildText("td:nth-child(4)"))
		gd, _ := strconv.Atoi(h.ChildText("td:nth-child(5)"))
		pts, _ := strconv.Atoi(h.ChildText("td:nth-child(6)"))
		obj := &Table{
			Position: pos,
			Team: h.ChildText("td:nth-child(3)"),
			MatchesPlayed: matchesPlayed,
			GoalDifference: gd,
			Points: pts,
		}

		premierTable = append(premierTable, *obj)
	})

	// Find and visit all links

	c.Visit("https://www.transfermarkt.com/premier-league/startseite/wettbewerb/GB1")
}
