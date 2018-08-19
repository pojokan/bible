package bibleis

import (
    "net/http"    
)

type Language struct {
	Family      LanguageFamily `json:"language_family_code"`
	Name        string         `json:"language_family_name"`
	English     string         `json:"language_family_english"`
	ISO         string         `json:"language_family_iso"`
	Languages   []string       `json:"language"`
	Medias      []string       `json:"media"`
	Deliveries  []string       `json:"delivery"`
	Resolutions []string       `json:"resolution"`
}

func (c *Client) GetLanguages() []*Language {
	request, errNewRequest := http.NewRequest("GET", c.endpoint, nil)
	if errNewRequest != nil {
	    return nil
	}
	request.URL.Path = "volumelanguagefamily"
	query := request.URL.Query()
	query.Add("key", c.key)
	return nil
}
