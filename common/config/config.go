package config

import (
	"encoding/json"
	"log"
	"os"
	"sync"

	"github.com/henkgo/chatgpt/common/logger"
)

var (
	once   sync.Once
	config *Configuration
)

// Configuration ...
type Configuration struct {
	APIKey           string  `json:"api_key"`
	Addr             string  `json:"addr"`
	BotDesc          string  `json:"bot_desc"`
	MaxTokens        int     `json:"max_tokens"`
	Model            string  `json:"model"`
	Temperature      float64 `json:"temperature"`
	TopP             float32 `json:"top_p"`
	PresencePenalty  float32 `json:"presence_penalty"`
	FrequencyPenalty float32 `json:"frequency_penalty"`
}

// LoadConfig ...
func LoadConfig() *Configuration {
	once.Do(func() {
		// default configuration
		config = &Configuration{
			MaxTokens:        60,
			Addr:             "8080",
			Model:            "gpt-3.5-turbo-0301",
			Temperature:      0.9,
			TopP:             1,
			FrequencyPenalty: 0.0,
			PresencePenalty:  0.6,
		}

		// check config file
		_, err := os.Stat("config.json")
		if err == nil {
			f, err := os.Open("config.json")
			if err != nil {
				log.Fatalf("open config err: %v", err)
				return
			}
			defer f.Close()
			encoder := json.NewDecoder(f)
			err = encoder.Decode(config)
			if err != nil {
				log.Fatalf("decode config err: %v", err)
				return
			}
		}

	})
	if config.APIKey == "" {
		logger.Error("config err: api key required")
	}

	return config
}
