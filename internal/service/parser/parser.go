package parser

import (
	"encoding/json"
	"net/url"
	"regexp"
	"time"

	"github.com/gocolly/colly"
)

type (
	Parser struct {
		Cfg *Config
	}
	Config struct {
		Visit       []Visit
		Delay       int
		RandomDelay int
	}
	Visit struct {
		URL string
		Get []Get
	}
	Get struct {
		Elem  string
		Strip bool
		To    To
	}
	To struct {
		Name  string
		Attrs []Attr
	}
	Attr struct {
		Name  string
		Value string
	}
	Result struct {
		StartTime time.Time
		EndTime   time.Time
		Elapsed   time.Duration
		Items     []ResultItem
	}
	ResultItem struct {
		Name  string
		Value string
		Attrs []Attr
	}
)

func ConfigFromJSON(data []byte) (*Config, error) {
	var cfg Config
	err := json.Unmarshal(data, &cfg)
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}

func New(cfg *Config) *Parser {
	return &Parser{
		Cfg: cfg,
	}
}

func (p *Parser) Run() (*Result, error) {
	startTime := time.Now()

	reg, err := regexp.Compile("[^[:graph:] ]+")
	if err != nil {
		return nil, err
	}

	var resultItems []ResultItem
	for _, visit := range p.Cfg.Visit {
		baseURL, err := url.Parse(visit.URL)
		if err != nil {
			return nil, err
		}

		// TODO maybe to put it outside the loop
		// to limit requests over all "visits"
		c := colly.NewCollector()

		err = c.Limit(&colly.LimitRule{
			DomainGlob:  baseURL.Host + "*",
			Parallelism: 1,
			Delay:       time.Duration(p.Cfg.Delay) * time.Second,
			RandomDelay: time.Duration(p.Cfg.RandomDelay) * time.Second,
		})
		if err != nil {
			return nil, err
		}

		c.OnHTML("body", func(e *colly.HTMLElement) {
			for _, get := range visit.Get {
				// e.ForEach(get.Elem, func(_ int, el *colly.HTMLElement) {
				value := e.ChildText(get.Elem)
				if get.Strip {
					value = reg.ReplaceAllString(value, "")
				}

				item := ResultItem{
					Name:  get.To.Name,
					Value: value,
					Attrs: get.To.Attrs,
				}

				resultItems = append(resultItems, item)
				// })
			}
		})

		err = c.Visit(visit.URL)
		if err != nil {
			return nil, err
		}
	}

	endTime := time.Now()

	pr := Result{
		StartTime: startTime,
		EndTime:   endTime,
		Elapsed:   endTime.Sub(startTime),
		Items:     resultItems,
	}

	return &pr, nil
}
