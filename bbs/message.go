package main

import (
	"encoding/json"
	"io/ioutil"
	"time"
)

const logFile = "logs.json"

type message struct {
	ID   int
	Name string
	Body string
	When time.Time
}

func (m *message) save() error {
	logs, err := loadLogs()
	if err != nil {
		return err
	}
	m.ID = len(logs) + 1
	m.When = time.Now()
	logs = append(logs, *m)

	// JSONにエンコード
	bytes, err := json.Marshal(logs)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(logFile, bytes, 0644)
}

// loadLogs is ログファイルの読み込み
func loadLogs() ([]message, error) {
	// ファイルを開く
	text, err := ioutil.ReadFile(logFile)
	if err != nil {
		return make([]message, 0), nil
	}
	// JSONをパース
	var logs []message
	if err := json.Unmarshal(text, &logs); err != nil {
		return nil, err
	}
	return logs, nil
}
