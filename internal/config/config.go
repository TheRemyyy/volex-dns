package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	Upstreams         []string `json:"upstreams"`
	ServerName        string   `json:"server_name"`
	BindAddr          string   `json:"bind_addr"`
	UpstreamTimeoutMs int      `json:"upstream_timeout_ms"`
	CacheMaxTtlSec    uint32   `json:"cache_max_ttl_sec"`
	CacheMinTtlSec    uint32   `json:"cache_min_ttl_sec"`
}

func LoadConfig(path string) (*Config, error) {
	var cfg Config
	file, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(file, &cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}
