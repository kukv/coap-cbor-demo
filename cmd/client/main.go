package main

import (
	"bytes"
	"coap-cbor-demo/internal"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/fxamacker/cbor/v2"
	"github.com/plgd-dev/go-coap/v3/udp"
)

func main() {
	data := internal.SensorData{Temp: 26.5, Humi: 45}

	jsonData, _ := json.Marshal(data)
	cborData, _ := cbor.Marshal(data)

	fmt.Println("--- Communication Size Comparison ---")
	fmt.Printf("JSON payload: %s (%d bytes)\n", string(jsonData), len(jsonData))
	fmt.Printf("CBOR payload: %v (%d bytes)\n", cborData, len(cborData))

	reduction := float64(len(jsonData)-len(cborData)) / float64(len(jsonData)) * 100
	fmt.Printf("Payload Reduction: %.1f%%\n", reduction)

	fmt.Printf("Estimated Total CoAP Packet: ~%d bytes\n", len(cborData)+4)
	fmt.Println("-------------------------------------")

	co, err := udp.Dial("localhost:5683")
	if err != nil {
		log.Fatal(err)
	}
	defer co.Close()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	resp, err := co.Post(ctx, "/data", 60, bytes.NewReader(cborData))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Server Response: %v\n", resp.Code())
}
