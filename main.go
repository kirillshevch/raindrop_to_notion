package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/PuerkitoBio/goquery"
	"github.com/dstotijn/go-notion"
)

const ParentID string = ""
const NotionSecret string = ""

func main() {
	data, err := os.Open("import.html")

	if err != nil {
		log.Fatal(err)
	}

	doc, err := goquery.NewDocumentFromReader(data)
	if err != nil {
		log.Fatal(err)
	}

	// Notion Client
	client := notion.NewClient(NotionSecret)

	// Find a Category
	doc.Find("dt").Each(func(i int, s *goquery.Selection) {
		category := s.Find("h3").Text()

		if category != "" {
			params := notion.CreatePageParams{

				ParentType: notion.ParentTypePage,
				ParentID:   ParentID,

				Title: []notion.RichText{
					{
						Text: &notion.Text{
							Content: category,
						},
					},
				},
			}

			// Create a Page
			page, err := client.CreatePage(context.Background(), params)

			if err != nil {
				fmt.Println("Cant create page: ", err)
			} else {
				s.Find("dt").Each(func(i int, s *goquery.Selection) {
					a := s.Find("a")
					link, _ := a.Attr("href")

					// Add Bookmark
					_, err := client.AppendBlockChildren(context.Background(), page.ID, []notion.Block{
						{
							Object: "block",
							Type:   notion.BlockTypeBookmark,
							Bookmark: &notion.Bookmark{
								URL: link,
								Caption: []notion.RichText{
									{
										Text: &notion.Text{
											Content: a.Text(),
										},
									},
								},
							},
						},
					})

					if err != nil {
						fmt.Println(err)
					} else {
						fmt.Println("Added: ", link)
					}
				})
			}
		}
	})
}
