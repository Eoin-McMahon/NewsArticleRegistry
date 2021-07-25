package data

import (
    "time"
    "io"
    "encoding/json"
)

type Article struct {
    ID          int     `json:"id"`
    Title       string  `json:"title"`
    Abstract    string  `json:"abstract"`
    Body        string  `json:"article_body"`
    Publisher   string  `json:"publisher"`
    Author      string  `json:"author"`
    WrittenOn   string  `json:"-"`
}

type Articles []*Article

func (art*Articles) ToJSON(writer io.Writer) error {
    enc := json.NewEncoder(writer)
    return enc.Encode(art)
}

func GetArticles() Articles {
    return articleList
}

var articleList = Articles {
    &Article{
        ID:             001,
        Title:          "Serial killer on death row Rodney Alcala dies of natural causes",
        Abstract:       "A man sentenced to death in the US state of California for murdering a 12-year-old girl and four other women has died of natural causes, officials say.",
        Body:           "Rodney Alcala, 77, died at a hospital near California's Corcoran state prison in the early hours of Saturday.Alcala, who was known as the 'Dating Game Killer' after taking part in a US TV show, was convicted in 2010.",
        Publisher:      "BBC",
        Author:         "BBC",
        WrittenOn:      time.Now().UTC().String(),
    },
    &Article{
        ID:             002,
        Title:          "Tokyo Olympics: What is 3x3 basketball all about, and who stood out on day one?",
        Abstract:       "It's been a long time coming, but on Saturday we finally saw what Olympic 3x3 basketball is all about. ",
        Body:           "One of four new sports for this year's Games - along with skateboarding, surfing and climbing - 3x3 is played on a half court, with both teams shooting at the same basket.",
        Publisher:      "BBC",
        Author:         "BBC",
        WrittenOn:      time.Now().UTC().String(),
    },
}
