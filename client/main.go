package main

import mqttclient "main/pkg"

func main() {
  topic  := "topic/waterevent"
  broker := "broker.emqx.io"
  mqttclient.InitMqttClient(broker, topic)
}
