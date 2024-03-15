package mqttclient

import (
	"log"
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

var MessagePubHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
  log.Printf("Received message: %s from topic: %s\n", msg.Payload(), msg.Topic())
}

var ConnectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
  log.Printf("Connected to broker")
}

var ConnectLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
  log.Printf("Connection to broker lost: %v", err)
}

