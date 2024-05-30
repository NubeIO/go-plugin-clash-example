package main

import (
    "fmt"
    mqtt "github.com/eclipse/paho.mqtt.golang"
)

var client mqtt.Client

func Connect(broker string, topic string) error {
    opts := mqtt.NewClientOptions().AddBroker(broker)
    client = mqtt.NewClient(opts)
    if token := client.Connect(); token.Wait() && token.Error() != nil {
        return token.Error()
    }

    if token := client.Subscribe(topic, 0, messageHandler); token.Wait() && token.Error() != nil {
        return token.Error()
    }
    return nil
}

func messageHandler(client mqtt.Client, msg mqtt.Message) {
    fmt.Printf("Received message: %s from topic: %s\n", msg.Payload(), msg.Topic())
}

func Disconnect() {
    client.Disconnect(250)
}
