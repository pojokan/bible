package bibleis

import (
	"net/http"
	"os"
)

type Client struct {
	endpoint   string
	key        string
	httpClient *http.Client
}

func NewClient() *Client {
	endpoint := os.Getenv("BIBLEIS_ENDPOJNT")
	key := os.Getenv("BIBLEIS_KEY")
	httpClient := &http.Client{}
	c := &Client{
		endpoint:   endpoint,
		key:        key,
		httpClient: httpClient,
	}
	return c
}

type DramaType string

const (
	Drama    DramaType = "2"
	NonDrama DramaType = "1"
)

type MediaType string

const (
	ElectronicText MediaType = "ET"
	DigitalAudio   MediaType = "DA"
	DigitalVideo   MediaType = "DV"
)

type LanguageFamily string
type DamID string
