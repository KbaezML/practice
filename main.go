package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
)

type Content struct {
	Price    string
	Currency string
}

type Response struct {
	Id      string
	Content Content
	Partial bool
}

type ApiResponse struct {
	Lprice string `json:"lprice"`
	Curr1  string `json:"curr1"`
	Curr2  string `json:"curr2"`
}

func main() {
	r := gin.Default()
	r.GET("/myapi", func(c *gin.Context) {
		resp:= getResponseApiExternal()
		c.JSON(200, gin.H{
			"Id": resp.Id,
			"Content": resp.Content,
			"Partial": resp.Partial,
		})
	})
	r.Run(":8081")
}

func getResponseApiExternal() Response {

	resp, err := http.Get("https://cex.io/api/last_price/BTC/USD")
	if err != nil {
	log.Fatalln(err)
	}
	body, err2 := ioutil.ReadAll(resp.Body)
	if err2 != nil {
	log.Fatalln(err)
	}

	var apiResponse ApiResponse
	err3 := json.Unmarshal(body, &apiResponse)
	if err3 != nil {
	log.Fatalln(err)
	}

	respInStruct := Response{
		Id: apiResponse.Curr1,
		Content: Content{
			Price: apiResponse.Lprice,
			Currency: apiResponse.Curr2,
		},
		Partial: false,
	}
	return respInStruct
}
