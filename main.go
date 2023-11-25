package main

import (
	"github.com/gin-gonic/gin"
	"github.com/shakezidin/model"
)

func main() {
	r := gin.Default()

	requestChannel := make(chan model.Request)

	resultChannel := make(chan model.ConvertedMessage)

	go worker(requestChannel, resultChannel)

	r.POST("/submit", func(c *gin.Context) {
		var req model.Request

		if err := c.BindJSON(&req); err != nil {
			c.JSON(400, gin.H{"error": "Invalid JSON format"})
			return
		}

		requestChannel <- req

		converted := <-resultChannel

		c.JSON(200, gin.H{
			"Status":    "Success",
			"message":   "Request received and processed by the worker",
			"converted": converted})
	})

	r.Run(":8080")
}

func worker(requests <-chan model.Request, results chan<- model.ConvertedMessage) {
	req := <-requests
	converted := convertRequest(req)
	results <- converted
}

func convertRequest(req model.Request) model.ConvertedMessage {
	return model.ConvertedMessage{
		Event:           req.Ev,
		EventType:       req.Et,
		AppID:           req.Id,
		UserID:          req.Uid,
		MessageID:       req.Mid,
		PageTitle:       req.T,
		PageURL:         req.P,
		BrowserLanguage: req.L,
		ScreenSize:      req.Sc,
		Attributes: map[string]model.Attribute{
			"form_varient": {Value: req.Atrv1, Type: req.Atrt1},
			"ref":          {Value: req.Atrv2, Type: req.Atrt2},
		},
		UserTraits: map[string]model.Trait{
			"name":  {Value: req.Uatrv1, Type: req.Uatrt1},
			"email": {Value: req.Uatrv2, Type: req.Uatrt2},
			"age":   {Value: req.Uatrv3, Type: req.Uatrt3},
		},
	}
}
