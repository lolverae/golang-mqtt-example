package mqttclient

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)


func SubscribeToTopic(topic string, client mqtt.Client) {

	if token := client.Subscribe(topic, 0, onMessageReceived); token.Wait() && token.Error() != nil {
		log.Fatal(fmt.Sprintf("Error subscribing to topic: %v", token.Error()))
	}

  log.Printf("Subscribed to topic %s", topic)

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan

	client.Unsubscribe(topic)
}
