package services

import (
	"encoding/json"
	"log"
	"pulsefin/config"
	"pulsefin/models"
	"strconv"
	"time"

	"github.com/gorilla/websocket"
)

// Kullanıcı ve sembol bazlı bağlantılar
var WebSocketConnections = make(map[int]map[string]*websocket.Conn)

func StartTickerWebSocketForUser(userID int) error {
	watchlist, err := GetWatchlist(userID)
	if err != nil {
		return err
	}

	// Kullanıcı için bağlantı haritasını başlat
	if WebSocketConnections[userID] == nil {
		WebSocketConnections[userID] = make(map[string]*websocket.Conn)
	}

	for _, item := range watchlist {
		go func(symbol string) {
			err := startWebSocketForSymbol(userID, symbol)
			if err != nil {
				log.Printf("WebSocket başlatılamadı (%s): %v\n", symbol, err)
			}
		}(item.StockSymbol)
	}
	return nil
}

func startWebSocketForSymbol(userID int, symbol string) error {
	url := config.AppConfig.WEBSOCKETURL

	conn, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		log.Printf("WebSocket bağlantısı başarısız (%s): %v\n", symbol, err)
		return err
	}

	// Kullanıcı bağlantısını sembol ile ilişkilendir
	WebSocketConnections[userID][symbol] = conn
	log.Printf("WebSocket bağlantısı başarılı: User ID: %d, Symbol: %s\n", userID, symbol)

	subscription := map[string]interface{}{
		"op": "subscribe",
		"args": []map[string]string{
			{"channel": "tickers", "instId": symbol},
		},
	}

	err = conn.WriteJSON(subscription)
	if err != nil {
		log.Printf("Abonelik mesajı gönderilemedi (%s): %v\n", symbol, err)
		return err
	}
	log.Printf("tickers kanalına abone olundu: %s\n", symbol)

	go listenToWebSocket(conn, userID, symbol)
	return nil
}

func listenToWebSocket(conn *websocket.Conn, userID int, symbol string) {
	defer func() {
		log.Printf("WebSocket bağlantısı kapatılıyor: Symbol: %s, User ID: %d\n", symbol, userID)
		conn.Close()
		delete(WebSocketConnections[userID], symbol) // Sembol bazlı bağlantıyı kaldır

		// Kullanıcının tüm sembolleri kapandıysa haritayı temizle
		if len(WebSocketConnections[userID]) == 0 {
			delete(WebSocketConnections, userID)
		}
	}()

	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Printf("WebSocket mesajı alınamadı (%s): %v\n", symbol, err)
			break
		}

		var response map[string]interface{}
		err = json.Unmarshal(message, &response)
		if err != nil {
			log.Printf("Mesaj ayrıştırılamadı (%s): %v\n", symbol, err)
			continue
		}

		if data, ok := response["data"].([]interface{}); ok && len(data) > 0 {
			processTickerData(data[0], symbol)
		}
	}
}

func processTickerData(rawData interface{}, symbol string) {
	tickerData, ok := rawData.(map[string]interface{})
	if !ok {
		log.Printf("Geçersiz ticker verisi (%s): %v\n", symbol, rawData)
		return
	}

	ticker := models.Ticker{
		Symbol:    symbol,
		LastPrice: parseFloat(tickerData["last"]),
		High24h:   parseFloat(tickerData["high24h"]),
		Low24h:    parseFloat(tickerData["low24h"]),
		Volume24h: parseFloat(tickerData["vol24h"]),
		Timestamp: time.Now(),
	}

	log.Printf("Anlık Fiyat (%s): %+v\n", symbol, ticker)
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

func StopWebSocketForUser(userID int) {
	if userConnections, exists := WebSocketConnections[userID]; exists {
		for symbol, conn := range userConnections {
			if conn != nil {
				conn.Close()
				log.Printf("WebSocket bağlantısı kapatıldı: Symbol: %s, User ID: %d\n", symbol, userID)
			}
		}
		delete(WebSocketConnections, userID)
		log.Printf("Kullanıcıya ait tüm WebSocket bağlantıları kapatıldı: User ID: %d\n", userID)
	} else {
		log.Printf("Kullanıcıya ait aktif WebSocket bağlantısı bulunamadı: User ID: %d\n", userID)
	}
}
