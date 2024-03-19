package main

import mqttclient "main/pkg"

func main() {
  topic  := "water/levelevent"
  broker := "127.0.0.1"

  mqttclient.InitMqttClient(broker, topic)
}
