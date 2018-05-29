package main

import "fmt"
import "net/http"
import "strings"
import "github.com/PuerkitoBio/goquery"

const DOMAIN = "https://www.ptt.cc"

type Doc struct {
  Url string
  Title string
  Author string
  TimeStr string
  Content string
}

func parseList(board string, results chan string) {
  boardUrl := fmt.Sprintf("%s/bbs/%s/index.html", DOMAIN, board)
  resp, err := http.Get(boardUrl)
  if err == nil {
    defer resp.Body.Close()

    gq, _ := goquery.NewDocumentFromReader(resp.Body)

    gq.Find("div.title").Each(func(i int, s *goquery.Selection) {
      url, _ := s.Find("a").Attr("href")

      results <- url
    })
  }
}

func parseArticle(url string, results chan Doc) {
  articleUrl := fmt.Sprintf("%s/%s", DOMAIN, url)
  resp, err := http.Get(articleUrl)
  if err == nil {
    defer resp.Body.Close()

    gq, _ := goquery.NewDocumentFromReader(resp.Body)

    var author, title, timestr, content string
    gq.Find("div#main-content .article-metaline").Each(func(i int, s *goquery.Selection) {
      switch i {
      case 0:
	author = strings.Split(strings.Split(s.Text(), "(")[0], "作者")[1]
      case 1:
	title = s.Text()
      case 2:
	timestr = strings.Split(s.Text(), "時間")[1]
      }
    })
    gq.Find("div#main-content").Contents().Not("div").Not("span").Each(func(i int, s *goquery.Selection) {
      content += strings.TrimSpace(s.Text())
    })

    results <- Doc{url, title, author, timestr, content}
  }
}

func main() {
  boards := []string{"NBA", "Stock", "Baseball", "LoL", "MobileComm", "movie",
    "Lifeismoney", "WomenTalk", "BabyMother", "marvel", "car",
    "ONE_PIECE", "Boy-Girl", "ToS", "Hearthstone", "Tech_Job", "Japan_Travel",
    "Beauty", "e-shopping", "joke", "PlayStation", "KoreaStar",
    "MakeUp", "marriage", "creditcard", "PC_Shopping", "BTS",
    "home-sale", "AllTogether", "Tainan", "Badminton", "KR_Entertain",
    "NBA_Film", "Steam", "BuyTogether", "StupidClown", "PuzzleDragon",
    "Kaohsiung", "HatePolitics", "FATE_GO", "iOS", "Japandrama",
    "Salary", "BeautySalon", "TaichungBun", "KoreaDrama", "Elephants",
    "PokemonGO", "NSwitch", "MuscleBeach", "CFantasy", "Palmar_Drama"}

  urlsQ := make(chan string, 1000)
  docsQ := make(chan Doc, 1000)

  for _, board := range boards {
    fmt.Println(board)
    go parseList(board, urlsQ)
  }

  go func() {
    for {
      select {
      case url := <-urlsQ:
        fmt.Println(url)
        go parseArticle(url, docsQ)
      default:
        continue
      }
    }
  }()

  count := 1
  for {
    select {
    case doc := <-docsQ:
      fmt.Println("[Finish] ", count, doc.Url)
      count += 1
    default:
      continue
    }
  }

}
