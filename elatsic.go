package logger

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
)

type ElasticConfig struct {
	Addresses     []string
	Username      string
	Password      string
	CAFingerprint string // CA fingerprint
	Index         string
}

var _elastic_client *elasticsearch.Client
var _index string

func InitElastic(cfg *ElasticConfig) error {
	var err error
	_elastic_client, err = elasticsearch.NewClient(elasticsearch.Config{
		Addresses:              cfg.Addresses,
		Username:               cfg.Username,
		Password:               cfg.Password,
		CertificateFingerprint: cfg.CAFingerprint,
	})

	if err != nil {
		return err
	}

	_index = cfg.Index

	res, err := _elastic_client.Indices.Create(_index, _elastic_client.Indices.Create.WithBody(strings.NewReader(`{
			"settings": {"number_of_shards": 1},
			"mappings": {
				"properties": {
					"@timestamp": {
						"type": "date"
					},
					"level": {
						"type": "keyword"
					},
					"identifier": {
						"type": "keyword"
					},
					"message": {
						"type": "text"
					}
				}
			}
		}`)))
	if err != nil {
		return err
	}

	defer res.Body.Close()

	if res.IsError() {
		r := make(map[string]interface{})
		if err = json.NewDecoder(res.Body).Decode(&r); err != nil {
			return err
		} else if r["error"].(map[string]interface{})["type"] == "resource_already_exists_exception" {
			err = nil
		} else {
			err = fmt.Errorf("ES index create error: %s", r["error"].(map[string]interface{})["type"])
		}
	}

	return err
}

type _es_log_item struct {
	Timestamp  int64  `json:"@timestamp"`
	Level      int    `json:"level"`
	Pid        int    `json:"pid"`
	Process    string `json:"process"`
	Identifier string `json:"identifier"`
	Message    string `json:"message"`
}

func LogElastic(level int, format string, v ...interface{}) {
	item := _es_log_item{
		Timestamp:  time.Now().UnixMilli(),
		Level:      level,
		Pid:        pid,
		Process:    processName,
		Identifier: Identifier,
		Message:    fmt.Sprintf(format, v...),
	}

	serialized, _ := json.Marshal(item)
	req := esapi.IndexRequest{
		Index: _index,
		Body:  strings.NewReader(string(serialized)),
	}

	res, err := req.Do(context.Background(), _elastic_client)
	if err != nil {
		fmt.Printf("ES index error: %s\n", err)
		LogStdout(level, format, v...)
		return
	}

	defer res.Body.Close()
	if res.IsError() {
		fmt.Printf("ES index error: %s\n", res.Status())
		LogStdout(level, format, v...)
	}
}
