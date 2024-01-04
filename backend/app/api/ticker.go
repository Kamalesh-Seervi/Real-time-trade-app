package api

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/segmentio/kafka-go"

	"github.com/kamalesh-seervi/consumer/core"
)

func GetAllTickers(c *gin.Context) {
	c.JSON(200, core.GetAllTickers())
}

func ListenTicker(c *gin.Context) {
	conn, err := websocket.Upgrade(c.Writer, c.Request, nil, 1024, 1024)
	if err != nil {
		log.Println("WebSocket Upgrade Error: ", err)
		return
	}
	defer conn.Close()

	currTicker := c.Param("ticker")
	log.Println("Current ticker: ", currTicker)

	if !core.IsTickerAllowed(currTicker) {
		conn.WriteMessage(websocket.CloseUnsupportedData, []byte("Ticker is not allowed"))
		log.Println("Ticker not allowed ticker: ", currTicker)
		return
	}

	topic := "trades-" + strings.ToLower(currTicker)
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{core.KAFKA_HOST + ":" + core.KAFKA_PORT},
		Topic:   topic,
	})
	reader.SetOffset(-1)
	defer reader.Close()

	conn.SetCloseHandler(func(code int, text string) error {
		reader.Close()
		log.Printf("Received connection close request. Closing connection .....")
		return nil
	})

	go func() {
		code, wsMessage, err := conn.NextReader()
		if err != nil {
			log.Println("Error reading last message from WS connection. Exiting ...")
			return
		}
		fmt.Printf("CODE : %d MESSAGE : %s\n", code, wsMessage)
	}()

	for {
		message, err := reader.ReadMessage(context.Background())
		if err != nil {
			log.Println("Error: ", err)
			return
		}
		fmt.Println("Reading..... ", string(message.Value))

		err = conn.WriteMessage(websocket.TextMessage, message.Value)
		if err != nil {
			log.Println("Error writing message to WS connection: ", err)
			return
		}
	}
}
