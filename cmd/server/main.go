package main

import (
	"bytes"
	"coap-cbor-demo/internal"
	"io"
	"log"

	"github.com/fxamacker/cbor/v2"
	"github.com/plgd-dev/go-coap/v3"
	"github.com/plgd-dev/go-coap/v3/message"
	"github.com/plgd-dev/go-coap/v3/message/codes"
	"github.com/plgd-dev/go-coap/v3/mux"
)

func main() {
	r := mux.NewRouter()
	_ = r.Handle("/data", mux.HandlerFunc(func(w mux.ResponseWriter, r *mux.Message) {
		payloadReader := r.Message.Body()
		payload, _ := io.ReadAll(payloadReader)

		var data internal.SensorData
		if err := cbor.Unmarshal(payload, &data); err != nil {
			log.Printf("CBOR Decode Error: %v", err)
			_ = w.SetResponse(codes.BadRequest, message.TextPlain, nil)
			return
		}
		log.Printf("Received: Temp=%.1f, Humi=%d", data.Temp, data.Humi)

		resp := map[string]string{"s": "ok"}
		respPayload, _ := cbor.Marshal(resp)

		_ = w.SetResponse(codes.Content, message.AppCBOR, bytes.NewReader(respPayload))
	}))

	log.Println("CoAP+CBOR Server started on :5683")
	log.Fatal(coap.ListenAndServe("udp", ":5683", r))
}
