package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
)

type Content struct {
	Price    float64
	Currency string
}

type Response struct {
	Id      string
	Content Content
	Partial bool
}

type ApiResponse struct {
	Id         string        `json:"id"`
	MarketData *CurrentPrice `json:"market_data,omitempty"`
}

type CurrentPrice struct {
	CurrentPrice map[string]float64 `json:"current_price"`
}

func main() {
	r := gin.Default()
	r.GET("/myapi", func(c *gin.Context) {
		data := c.Query("data")
		resp:= getResponseApiExternal(data)

		if resp.Content.Price != 0 {
			c.JSON(200, gin.H{
				"Id": resp.Id,
				"Content": resp.Content,
				"Partial": resp.Partial,
			})
		} else {
			c.JSON(206, gin.H{
				"Id": data,
				"Partial": true,
			})
		}
	})
	r.Run(":8081")
}

func getResponseApiExternal(coin string) Response {
	url := "https://api.coingecko.com/api/v3/coins/" + coin
	resp, err := http.Get(url)

	if err != nil {
		log.Fatalln(err)
	}

	body, err2 := ioutil.ReadAll(resp.Body)
	if err2 != nil {
		log.Fatalln(err)
	}

	var apiResponse = ApiResponse{
		Id:         "",
		MarketData: &CurrentPrice{},
	}
	err3 := json.Unmarshal(body, &apiResponse)
	if err3 != nil {
		log.Fatalln(err)
	}

	respInStruct := Response{
		Id: apiResponse.Id,
		Content: Content{
			Price: apiResponse.MarketData.CurrentPrice["usd"],
			Currency: "USD",
		},
		Partial: false,
	}

	return respInStruct
}
