package pokeapi

import(
	"io"
	"net/http"
	"fmt"
	"encoding/json"
)

type locationNames struct{
	Count 		int `json:"count"`
	Next 			*string `json:"next"`
	Previous 	*string `json:"previous"`
	Results []struct{
		Name 		string `json:"name"`
		Url 		string `json:"url"`
	}`json:"results"`
}

type Client struct{
	httpClient *http.Client
}

func NewClient(httpClient *http.Client) *Client{
	return &Client{
		httpClient: httpClient,
	}
}

func (c *Client) GetBody(pageURL *string) (locationNames, error){
	baseUrl := "https://pokeapi.co/api/v2/location-area"

	if pageURL != nil{
		baseUrl = *pageURL
	}
	
	req, err := http.NewRequest("GET", baseUrl, nil)
	if err != nil{
		return locationNames{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil{
		return locationNames{}, err
	}

	defer res.Body.Close()

	if res.StatusCode > 299{
		return locationNames{}, fmt.Errorf("bad status: %d", res.StatusCode)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil{
		return locationNames{}, err
	}

	var loc locationNames
	err = json.Unmarshal(body, &loc)
	if err != nil{
		return locationNames{}, err
	}

	return loc, nil

}

func GetNames(loc locationNames){
	for i := range loc.Results{
		fmt.Println(loc.Results[i].Name)
	}
}
