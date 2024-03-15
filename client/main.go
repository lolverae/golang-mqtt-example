package main

import mqttclient "main/pkg"

func main() {
  topic  := "topic/waterevent"
  broker := "127.0.0.1"

  mqttclient.InitMqttClient(broker, topic)
}
