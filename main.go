package main

import (
    "fmt"
    "plugin"
)

func loadPluginAndConnect(pluginPath, broker, topic string) (func(), error) {
    p, err := plugin.Open(pluginPath)
    if err != nil {
        return nil, fmt.Errorf("error loading plugin: %v", err)
    }

    connectSymbol, err := p.Lookup("Connect")
    connectFunc, _ := connectSymbol.(func(string, string) error)

    disconnectSymbol, err := p.Lookup("Disconnect")
    disconnectFunc, _ := disconnectSymbol.(func())
    err = connectFunc(broker, topic)

    return disconnectFunc, nil
}

func main() {
    broker := "tcp://broker.hivemq.com:1883"
    topic1 := "test/topic1"
    topic2 := "test/topic2"

    disconnect1, err := loadPluginAndConnect("plugin1/plugin1.so", broker, topic1)
    if err != nil {
        fmt.Println(err)
        return
    }

    disconnect2, err := loadPluginAndConnect("plugin2/plugin2.so", broker, topic2)
    if err != nil {
        fmt.Println(err)
        return
    }

    fmt.Println("Connected to MQTT broker and subscribed to topics")

    // Wait for messages
    fmt.Println("Press Enter to exit")
    fmt.Scanln()

    disconnect1()
    disconnect2()
    fmt.Println("Disconnected from MQTT broker")
}
