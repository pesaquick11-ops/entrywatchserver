package mqtt

import (
	"encoding/json"
	"log"
	"time"

	"entrywatchserver/internal/ws"
	paho "github.com/eclipse/paho.mqtt.golang"
)

type Message struct {
	Topic     string    `json:"topic"`
	Payload   string    `json:"payload"`
	Timestamp time.Time `json:"timestamp"`
}

type Config struct {
	Broker   string
	Username string
	Password string
	ClientID string
	Topic    string
}

// Connect sets up the MQTT client and pushes every message onto hub
func Connect(cfg Config, hub *ws.Hub) (paho.Client, error) {
	opts := paho.NewClientOptions().
		AddBroker(cfg.Broker).
		SetClientID(cfg.ClientID).
		SetUsername(cfg.Username).
		SetPassword(cfg.Password).
		SetAutoReconnect(true).
		SetConnectRetry(true).
		SetConnectRetryInterval(5 * time.Second).
		SetConnectTimeout(10 * time.Second)

	opts.OnConnect = func(c paho.Client) {
		log.Println("✅ Connected to MQTT broker")
		token := c.Subscribe(cfg.Topic, 1, func(client paho.Client, msg paho.Message) {
			m := Message{
				Topic:     msg.Topic(),
				Payload:   string(msg.Payload()),
				Timestamp: time.Now(),
			}
			log.Println("MQTT →", m.Topic, m.Payload)

			// Optional: if payload is JSON, forward it decoded instead of as a string
			var decoded interface{}
			if err := json.Unmarshal(msg.Payload(), &decoded); err == nil {
				hub.Broadcast(map[string]interface{}{
					"topic":     m.Topic,
					"payload":   decoded,
					"timestamp": m.Timestamp,
				})
			} else {
				hub.Broadcast(m)
			}
		})
		token.Wait()
		if err := token.Error(); err != nil {
			log.Println("❌ Subscribe failed:", err)
		}
	}
	opts.OnConnectionLost = func(c paho.Client, err error) {
		log.Println("❌ MQTT connection lost:", err)
	}

	client := paho.NewClient(opts)
	token := client.Connect()
	token.Wait()
	return client, token.Error()
}
