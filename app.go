// Projekt: Loxone ↔ Homematic Bridge (Go)
// Autor: Niclas Schnell
// Lizenz: MIT

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/websocket"
)

// Konfiguration
const (
	wsURL     = "ws://<homematic-ip>:<port>/api/ws"
	token     = "<dein_token>"
	httpPort  = ":8080"
)

var conn *websocket.Conn

func sendPluginState() error {
	msg := map[string]interface{}{
		"type": "PluginStateResponse",
		"pluginState": map[string]interface{}{
			"pluginReadinessStatus": "READY",
			"message": "Plugin bereit",
		},
	}
	return conn.WriteJSON(msg)
}

func connectWebSocket() {
	header := http.Header{}
	header.Add("Authorization", "Bearer "+token)

	for {
		c, _, err := websocket.DefaultDialer.Dial(wsURL, header)
		if err != nil {
			log.Println("WebSocket Verbindungsfehler:", err)
			time.Sleep(5 * time.Second)
			continue
		}
		conn = c
		log.Println("WebSocket verbunden")

		err = sendPluginState()
		if err != nil {
			log.Println("Fehler beim Senden von PluginState:", err)
			conn.Close()
			continue
		}

		for {
			_, msg, err := conn.ReadMessage()
			if err != nil {
				log.Println("WebSocket getrennt:", err)
				conn.Close()
				break
			}
			log.Println("Empfangen:", string(msg))
		}
	}
}

func setSwitchHandler(w http.ResponseWriter, r *http.Request) {
	deviceID := r.URL.Query().Get("device")
	state := r.URL.Query().Get("state")
	if deviceID == "" || (state != "on" && state != "off") {
		http.Error(w, "Fehlende oder ungültige Parameter", http.StatusBadRequest)
		return
	}
	if conn == nil {
		http.Error(w, "WebSocket nicht verbunden", http.StatusServiceUnavailable)
		return
	}

	msg := map[string]interface{}{
		"type": "ControlResponse",
		"control": map[string]interface{}{
			"deviceId": deviceID,
			"property": map[string]interface{}{
				"type":  "SwitchState",
				"value": stringToUpper(state),
			},
		},
	}

	err := conn.WriteJSON(msg)
	if err != nil {
		http.Error(w, "Fehler beim Senden an Homematic", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Befehl an %s gesendet: %s\n", deviceID, state)
}

func stringToUpper(s string) string {
	if len(s) == 0 {
		return s
	}
	return string(s[0]-32) + s[1:]
}

func main() {
	go connectWebSocket()

	http.HandleFunc("/setSwitch", setSwitchHandler)
	log.Println("HTTP-Server läuft auf Port", httpPort)
	err := http.ListenAndServe(httpPort, nil)
	if err != nil {
		log.Println("HTTP-Server Fehler:", err)
		os.Exit(1)
	}
}
