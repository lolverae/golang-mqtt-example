package mqttserver

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	mqtt "github.com/mochi-mqtt/server/v2"
	"github.com/mochi-mqtt/server/v2/hooks/auth"
	"github.com/mochi-mqtt/server/v2/listeners"
)

func StartMqttServer() {
  sigs := make(chan os.Signal, 1)
  signal.Notify(sigs, os.Interrupt, syscall.SIGTERM)

	authRules := &auth.Ledger{
      Auth: auth.AuthRules{
        {Username: "peach", Password: "password1", Allow: true},
      },
  }

  server := mqtt.New(nil)
	err := server.AddHook(new(auth.Hook), &auth.Options{
		Ledger: authRules,
	})

  tcpListener := listeners.NewTCP("t1", ":1883", nil)
  err = server.AddListener(tcpListener)
  if err != nil {
      log.Printf("failed to add TCP listener: %v", err)
  }

  err = server.Serve()
  if err != nil {
    log.Fatal(err)
  }
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan

}
