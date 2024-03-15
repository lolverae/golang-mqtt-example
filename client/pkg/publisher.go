package mqttclient

import (
	"fmt"
	"log"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func publish(client mqtt.Client) {
  num := 10
  for i := 0; i < num; i++ {
    message := fmt.Sprintf("Publishing message %d", i)
    token := client.Publish("topic/test", 0, false, message)
    token.Wait()
    time.Sleep(time.Second)

		log.Println("Published:", message)
		time.Sleep(1 * time.Second)
  }
}

