package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Observe86/intake-api/internal/model"
)

type MetricsService struct {
	CollectorURL string
}

func (s *MetricsService) Process(metrics []model.Metric) error {
	for _, m := range metrics {
		if m.Name == "" {
			return fmt.Errorf("metric name cannot be empty")
		}
	}

	data, err := json.Marshal(metrics)
	if err != nil {
		return err
	}

	resp, err := http.Post(s.CollectorURL+"/collect-metrics", "application/json", bytes.NewBuffer(data))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("collector-gateway returned %d", resp.StatusCode)
	}

	return nil
}
