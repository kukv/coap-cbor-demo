package internal

type SensorData struct {
	Temp float64 `cbor:"t" json:"temperature"`
	Humi int     `cbor:"h" json:"humidity"`
}
