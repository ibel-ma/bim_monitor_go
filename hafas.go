package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

var hafasURL = "https://verkehrsauskunft.verbundlinie.at/hamm/gate"
var hafasAuth = map[string]any{"type": "AID", "aid": "wf7mcf9bv3nv8g5f"}
var hafasClient = map[string]any{"id": "VAO", "type": "WEB", "name": "webapp", "l": "vs_stv", "v": 10010}

func hafasRequest(req any) (map[string]any, error) {
	//logger := slog.Default()
	//logger.Info("Sending HAFAS request", "method", req.(map[string]any)["meth"])
	payload := map[string]any{"id": "request01", "ver": "1.59", "lang": "deu", "auth": hafasAuth, "client": hafasClient, "ext": "VAO.22", "formatted": false, "svcReqL": []any{req}}
	b, _ := json.Marshal(payload)
	url := fmt.Sprintf("%s?rnd=%d", hafasURL, time.Now().UnixMilli())
	httpReq, _ := http.NewRequest("POST", url, bytes.NewBuffer(b))
	httpReq.Header.Set("Content-Type", "application/json")
	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var out map[string]any
	if err := json.NewDecoder(resp.Body).Decode(&out); err != nil {
		return nil, err
	}
	return out, nil
}
