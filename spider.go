// spider
package main

import (
	"fmt"
	"log"

	"github.com/PuerkitoBio/goquery"
)

type Post struct {
	Title string
	Url   string
	Img   string
	Text  string
}

func getPosts() []Post {
	doc, err := goquery.NewDocument(URL)
	var posts []Post

	db := initializationDatabase()
	defer db.Close()

	if err != nil {
		log.Fatal(err)
	}

	doc.Find(RULES.Item).Each(func(index int, item *goquery.Selection) {
		var news Post

		//		log.Printf("Item: %s; Title: %s; Text: %s", RULES.Item, RULES.Title, RULES.Text)
		log.Printf("ITEM-TEXT: %s", item.Text())
		linkTag := item.Find("a")
		link, _ := linkTag.Attr("href")

		imgTag := item.Find("img")
		img, _ := imgTag.Attr("src")

		title := item.Find(RULES.Title).Text()
		text := item.Find(RULES.Text).Text()

		//		if RULES.Text == "" && RULES.Title == "" {
		//			title = item.Text()
		//		}

		news.Title = title
		news.Url = link
		news.Img = img
		news.Text = text

		insertNews(&db, news)

		posts = append(posts, news)

		//fmt.Printf("Post #%d: %s - %s\n", index, posts[index].Title, posts[index].Url)
		//GetPost(link)

	})
	return posts
}

func getPost(url string) {
	fmt.Printf("first: %s", url)
	news, err := goquery.NewDocument(url)
	if err != nil {
		log.Fatal(err)
	}
	news.Find(".e1-news-item .js-news-item").Each(func(index int, item *goquery.Selection) {
		content := item.Text()
		fmt.Printf("Content #%d: %s", index, content)
	})
}
