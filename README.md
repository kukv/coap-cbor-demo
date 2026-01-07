# CoAP + CBOR Demo

このプロジェクトは、Go言語を使用してCoAP（Constrained Application Protocol）プロトコルとCBOR（Concise Binary Object Representation）データ形式を組み合わせた、軽量な通信のデモンストレーションです。

IoTデバイスなどのリソースが制限された環境において、JSONよりも効率的にデータを送受信する方法を示しています。

## 特徴

- **CoAPプロトコル**: UDPベースの軽量なアプリケーション層プロトコル。
- **CBORデータ形式**: JSONと互換性のあるバイナリ形式で、シリアライズ後のサイズが小さくなります。
- **サイズ比較**: クライアント実行時に、同一データにおけるJSONとCBORのペイロードサイズを比較表示します。

## プロジェクト構造

- `cmd/server/main.go`: CoAPサーバー。`/data` エンドポイントでCBORデータを受信します。
- `cmd/client/main.go`: CoAPクライアント。データをCBORに変換して送信し、JSONとのサイズ比較を行います。
- `internal/sensor_data.go`: 通信に使用する共通のデータ構造体（SensorData）。

## 使い方

### 1. サーバーの起動

別のターミナルでサーバーを起動します。

```bash
go run cmd/server/main.go
```

サーバーは `localhost:5683` (UDP) で待機を開始します。

### 2. クライアントの実行

サーバーが起動している状態で、クライアントを実行します。

```bash
go run cmd/client/main.go
```

### 実行結果の例

クライアントを実行すると、以下のような出力が表示されます：

```text
--- Communication Size Comparison ---
JSON payload: {"temperature":26.5,"humidity":45} (34 bytes)
CBOR payload: [162 100 116 251 64 58 128 0 0 0 0 0 100 104 24 45] (16 bytes)
Payload Reduction: 52.9%
Estimated Total CoAP Packet: ~20 bytes
-------------------------------------
Server Response: 2.05 (Content)
```

サーバー側のログには、受信したデータが表示されます：

```text
2026/01/07 20:56:00 Received: Temp=26.5, Humi=45
```

## 使用ライブラリ

- [go-coap/v3](https://github.com/plgd-dev/go-coap): CoAPの実装。
- [fxamacker/cbor](https://github.com/fxamacker/cbor): 高速で安全なCBORの実装。
