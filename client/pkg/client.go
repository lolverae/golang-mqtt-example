package mqttclient

import (
	"fmt"
	"log"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)


func InitMqttClient(broker, topic string) {
  opts := mqtt.NewClientOptions()
  opts.AddBroker(fmt.Sprintf("tcp://%s:1883", broker))
  opts.SetClientID("go_mqtt_client")
  opts.SetUsername("emqx")
  opts.SetPassword("public")
  opts.SetDefaultPublishHandler(messagePubHandler)
  opts.OnConnect = connectHandler
  opts.OnConnectionLost = connectLostHandler

  client := mqtt.NewClient(opts)
  if token := client.Connect(); token.Wait() && token.Error() != nil {
    log.Fatal(fmt.Sprintf("Error connecting to MQTT broker: %v", token.Error()))
  }
  defer client.Disconnect(250)

  SubscribeToTopic(topic, client)

}

