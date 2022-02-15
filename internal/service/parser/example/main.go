package main

import (
	"fmt"
	"log"

	"github.com/DanielTitkov/dashboars/internal/service/parser"
)

func main() {
	cfg := parser.Config{
		Visit: []parser.Visit{
			{
				URL: "https://www.wikipedia.org/",
				Get: []parser.Get{
					{
						Elem: "div.footer-sidebar-text:nth-child(2)",
						To: parser.To{
							Name: "footer_text",
							Attrs: []parser.Attr{
								{
									Name:  "source",
									Value: "wiki",
								},
								{
									Name:  "location",
									Value: "footer",
								},
							},
						},
					},
				},
			},
		},
		Delay:       2,
		RandomDelay: 1,
	}

	fmt.Printf("parser config:\n%+v\n", cfg)

	p := parser.New(&cfg)
	result, err := p.Run()
	if err != nil {
		log.Fatal("parser error", err)
	}

	fmt.Printf("%+v", result)

}
