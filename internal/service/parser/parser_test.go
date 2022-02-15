package parser

import (
	"io/ioutil"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var testConfig Config = Config{
	Visit: []Visit{
		{
			URL: "https://www.wikipedia.org/",
			Get: []Get{
				{
					Elem:  "div.footer-sidebar-text:nth-child(2)",
					Strip: true,
					To: To{
						Name: "footer_text",
						Attrs: []Attr{
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
				{
					Elem:  "p.jsl10n",
					Strip: false,
					To: To{
						Name: "save_text",
						Attrs: []Attr{
							{
								Name:  "source",
								Value: "wiki",
							},
							{
								Name:  "location",
								Value: "very-footer",
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

func TestConfigFromJSON(t *testing.T) {
	data, err := ioutil.ReadFile("./test/test_config.json")
	if err != nil {
		t.Error("failed to load test config", err)
	}
	cfg, err := ConfigFromJSON(data)
	if err != nil {
		t.Error("failed to read test config", err)
	}

	assert.Equal(t, &testConfig, cfg)
}

func TestParserRun(t *testing.T) {
	exp := Result{
		Items: []ResultItem{
			{
				Name:  "footer_text",
				Value: "Wikipedia is hosted by the Wikimedia Foundation, a non-profit organization that also hosts a range of other projects.",
				Attrs: []Attr{
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
			{
				Name:  "save_text",
				Value: "Save your favorite articles to read offline, sync your reading lists across devices and customize your reading experience with the official Wikipedia app.",
				Attrs: []Attr{
					{
						Name:  "source",
						Value: "wiki",
					},
					{
						Name:  "location",
						Value: "very-footer",
					},
				},
			},
		},
	}

	p := New(&testConfig)
	res, err := p.Run()
	assert.Nil(t, err)
	assert.Equal(t, exp.Items, res.Items)
	assert.True(t, res.StartTime.Before(res.EndTime), "end time not before start time")
	assert.NotEqual(t, time.Duration(0), res.Elapsed)
}
