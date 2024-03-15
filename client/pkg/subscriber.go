package mqttclient

import (
  "log"
  "fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func onMessageReceived(client mqtt.Client, message mqtt.Message) {
	fmt.Printf("Received message: %s from topic: %s\n", message.Payload(), message.Topic())
}

func sub(client mqtt.Client) {
    topic := "topic/waterevent"
    token := client.Subscribe(topic, 1, nil)
    token.Wait()
    log.Printf("Subscribed to topic %s", topic)
}
