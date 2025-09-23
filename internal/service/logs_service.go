package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Observe86/intake-api/internal/model"
)

type LogsService struct {
	CollectorURL string
}

func (s *LogsService) Process(logs []model.Log) error {
	data, err := json.Marshal(logs)
	if err != nil {
		return err
	}

	resp, err := http.Post(s.CollectorURL+"/collect-logs", "application/json", bytes.NewBuffer(data))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("collector-gateway returned %d", resp.StatusCode)
	}

	return nil
}
