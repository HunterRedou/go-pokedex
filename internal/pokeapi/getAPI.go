package pokeapi

import(
	"io"
	"net/http"
	"fmt"
	"encoding/json"
	"time"
	"github.com/HunterRedou/pokedex/internal/pokecache"
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

type Catch struct {
  Name   string `json:"name"`
  Height int    `json:"height"`
  Weight int    `json:"weight"`
	BaseExp int		`json:"base_experience"`
  Stats  []struct {
      BaseStat int `json:"base_stat"`
      Stat     struct {
          Name string `json:"name"`
        }`json:"stat"`
    }`json:"stats"`
  Types []struct {
      Type struct {
          Name string `json:"name"`
        }`json:"type"`
    }`json:"types"`
}

type pokemon struct{
	Name	string `json:"name"`
}

type encounter struct{
	Pokemon pokemon `json:"pokemon"`
}

type top struct{
	Encounter []encounter `json:"pokemon_encounters"`
}

type Client struct{
	cache *pokecache.Cache
	httpClient *http.Client
}

func NewClient(httpClient *http.Client, cacheInterval time.Duration) *Client{
	return &Client{
		cache: pokecache.NewCache(cacheInterval),
		httpClient: httpClient,
	}
}

func (c *Client) GetBody(pageURL *string) (locationNames, error){
	baseUrl := "https://pokeapi.co/api/v2/location-area"

	if pageURL != nil{
		baseUrl = *pageURL
	}

	if val, ok := c.cache.Get(baseUrl); ok {
		location := locationNames{}
		err := json.Unmarshal(val, &location)
		if err != nil {
			return locationNames{}, err
		}

		return location, nil
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

	c.cache.Add(baseUrl, body)
	return loc, nil

}

func GetNames(loc locationNames){
	for i := range loc.Results{
		fmt.Println(loc.Results[i].Name)
	}
}

func (c *Client)GetPokemon(city string) (top, error){
	pokeUrl := "https://pokeapi.co/api/v2/location-area" + "/" + city

	if val, ok := c.cache.Get(pokeUrl); ok {
		t := top{}
		err := json.Unmarshal(val, &t)
		if err != nil {
			return top{}, err
		}

		return t, nil
	}
	
	req, err := http.NewRequest("GET", pokeUrl, nil)
	if err != nil{
		return top{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil{
		return top{}, err
	}

	defer res.Body.Close()

	if res.StatusCode > 299{
		return top{}, fmt.Errorf("bad status: %d", res.StatusCode)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil{
		return top{}, err
	}

	var t top
	err = json.Unmarshal(body, &t)
	if err != nil{
		return top{}, err
	}
	
	c.cache.Add(pokeUrl, body)
	return t, nil
}


func (c *Client)CatchPokemon(poke string) (Catch, error){
	pokeUrl := "https://pokeapi.co/api/v2/pokemon" + "/" + poke

	if val, ok := c.cache.Get(pokeUrl); ok {
		cat := Catch{}
		err := json.Unmarshal(val, &cat)
		if err != nil {
			return Catch{}, err
		}

		return cat, nil
	}
	
	req, err := http.NewRequest("GET", pokeUrl, nil)
	if err != nil{
		return Catch{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil{
		return Catch{}, err
	}

	defer res.Body.Close()

	if res.StatusCode > 299{
		return Catch{}, fmt.Errorf("bad status: %d", res.StatusCode)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil{
		return Catch{}, err
	}

	var cat Catch
	err = json.Unmarshal(body, &cat)
	if err != nil{
		return Catch{}, err
	}
	
	c.cache.Add(pokeUrl, body)
	return cat, nil
}
