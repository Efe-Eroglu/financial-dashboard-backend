package services

import (
	"encoding/json"
	"log"
	"pulsefin/config"
	"pulsefin/models"
	"pulsefin/utils"
	"strconv"
	"time"

	"github.com/gorilla/websocket"
)

var WebSocketConnection *websocket.Conn

func StartTickerWebSocket(symbol string) error {
	url := config.AppConfig.WEBSOCKETURL

	conn, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		log.Println("WebSocket bağlantısı başarısız:", err)
		return err
	}
	WebSocketConnection = conn
	log.Println("WebSocket bağlantısı başarılı.")

	subscription := map[string]interface{}{
		"op": "subscribe",
		"args": []map[string]string{
			{"channel": "tickers", "instId": symbol},
		},
	}

	err = conn.WriteJSON(subscription)
	if err != nil {
		log.Println("Abonelik mesajı gönderilemedi:", err)
		return err
	}
	log.Printf("tickers kanalına abone olundu: %s\n", symbol)

	go listenToWebSocket(conn)
	return nil
}

func listenToWebSocket(conn *websocket.Conn) {
	defer func() {
		log.Println("WebSocket bağlantısı kapatılıyor...")
		conn.Close()
	}()

	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("WebSocket mesajı alınamadı veya bağlantı kapandı:", err)
			break
		}

		var response map[string]interface{}
		err = json.Unmarshal(message, &response)
		if err != nil {
			log.Println("Mesaj ayrıştırılamadı:", err)
			continue
		}

		if data, ok := response["data"].([]interface{}); ok && len(data) > 0 {
			processTickerData(data[0])
		} else if event, ok := response["event"].(string); ok && event == "error" {
			log.Println("WebSocket Hata Mesajı:", response)
		}
	}
}

func processTickerData(rawData interface{}) {
	tickerData, ok := rawData.(map[string]interface{})
	if !ok {
		log.Println("Geçersiz ticker verisi:", rawData)
		return
	}

	ticker := models.Ticker{
		Symbol:    utils.SafeString(tickerData["instId"]),
		LastPrice: parseFloat(tickerData["last"]),
		High24h:   parseFloat(tickerData["high24h"]),
		Low24h:    parseFloat(tickerData["low24h"]),
		Volume24h: parseFloat(tickerData["vol24h"]),
		Timestamp: time.Now(),
	}

	log.Printf("Anlık Fiyat: %+v\n", ticker)
}

func parseFloat(value interface{}) float64 {
	if value == nil {
		return 0
	}
	strValue, ok := value.(string)
	if !ok {
		return 0
	}
	parsed, err := strconv.ParseFloat(strValue, 64)
	if err != nil {
		return 0
	}
	return parsed
}
