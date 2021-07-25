package data

import (
    "fmt"
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

// Decodes body of request into a json object
func (art *Article) FromJSON(reader io.Reader) error {
    dec := json.NewDecoder(reader)
    return dec.Decode(art)
}

// Encodes stored article into a json object
func (art*Articles) ToJSON(writer io.Writer) error {
    enc := json.NewEncoder(writer)
    return enc.Encode(art)
}

// Return list of stored articles
func GetArticles() Articles {
    return articleList
}

// Add an article to our article list
func AddArticle(art *Article) {
    art.ID = getNextID()
    articleList = append(articleList, art)
}

// Update article with specific id 
func UpdateArticle(id int, art *Article) error {
    _, pos, err := findArticleByID(id)
    if err != nil {
        return err
    }

    art.ID = id
    articleList[pos] = art

    return err
}

func getNextID() int {
    lastArticle := articleList[len(articleList)-1]
    return lastArticle.ID + 1
}

var ErrArticleNotFound = fmt.Errorf("Article Not Found")

func findArticleByID(id int) (*Article, int, error) {
    for i, art := range articleList {
        if art.ID == id {
            return art, i, nil
        }
    }
        return nil, -1, ErrArticleNotFound
}

var articleList = Articles {
    &Article{
        ID:             1,
        Title:          "Serial killer on death row Rodney Alcala dies of natural causes",
        Abstract:       "A man sentenced to death in the US state of California for murdering a 12-year-old girl and four other women has died of natural causes, officials say.",
        Body:           "Rodney Alcala, 77, died at a hospital near California's Corcoran state prison in the early hours of Saturday.Alcala, who was known as the 'Dating Game Killer' after taking part in a US TV show, was convicted in 2010.",
        Publisher:      "BBC",
        Author:         "BBC",
        WrittenOn:      time.Now().UTC().String(),
    },
    &Article{
        ID:             2,
        Title:          "Tokyo Olympics: What is 3x3 basketball all about, and who stood out on day one?",
        Abstract:       "It's been a long time coming, but on Saturday we finally saw what Olympic 3x3 basketball is all about. ",
        Body:           "One of four new sports for this year's Games - along with skateboarding, surfing and climbing - 3x3 is played on a half court, with both teams shooting at the same basket.",
        Publisher:      "BBC",
        Author:         "BBC",
        WrittenOn:      time.Now().UTC().String(),
    },
}
