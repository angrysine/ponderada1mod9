package main

import (
	"math/rand"
	"strconv"
	"time"

	MQTT "github.com/eclipse/paho.mqtt.golang"
)

func Pub() {
	opts := MQTT.NewClientOptions().AddBroker("tcp://localhost:1891")
	opts.SetClientID("go_publisher")

	client := MQTT.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	for {
		text :=  strconv.Itoa(rand.Intn(1280)) +" W/m2"
		token := client.Publish("solar/sensor", 0, false, text)
		token.Wait()
		time.Sleep(2 * time.Second)
	}
}
// func main() {
// 	Pub()
// }
