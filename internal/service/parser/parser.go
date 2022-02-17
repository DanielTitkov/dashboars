package parser

import (
	"encoding/json"
	"log"
	"net/url"
	"regexp"
	"time"

	"github.com/gocolly/colly"
)

const defaultStripRegexp = "[^[:graph:] ]+"

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
		Elem        string
		Strip       bool
		StripRegexp string
		To          To
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

	reg, err := regexp.Compile(defaultStripRegexp)
	if err != nil {
		return nil, err
	}

	var resultItems []ResultItem
	for _, visit := range p.Cfg.Visit {
		log.Println("NEXT VISIT") // FIXME
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
				value := e.ChildText(get.Elem)
				if get.Strip {
					var err error
					if get.StripRegexp != "" {
						reg, err = regexp.Compile(get.StripRegexp)
						if err != nil {
							// TODO: probably send this errors to some channel
							continue
						}
					}
					value = reg.ReplaceAllString(value, "")
				}

				item := ResultItem{
					Name:  get.To.Name,
					Value: value,
					Attrs: get.To.Attrs,
				}

				resultItems = append(resultItems, item)
			}
		})

		err = c.Visit(visit.URL)
		if err != nil {
			return nil, err
		}
		log.Println("VISIT DONE") // FIXME
	}

	endTime := time.Now()
	return &Result{
		StartTime: startTime,
		EndTime:   endTime,
		Elapsed:   endTime.Sub(startTime),
		Items:     resultItems,
	}, nil
}
