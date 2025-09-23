package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Observe86/intake-api/internal/model"
)

type TracesService struct {
	CollectorURL string
}

func (s *TracesService) Process(traces []model.Trace) error {
	data, err := json.Marshal(traces)
	if err != nil {
		return err
	}

	resp, err := http.Post(s.CollectorURL+"/collect-traces", "application/json", bytes.NewBuffer(data))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("collector-gateway returned %d", resp.StatusCode)
	}

	return nil
}
