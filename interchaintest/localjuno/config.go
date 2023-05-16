package main

import (
	"encoding/json"
	"io/ioutil"
)

type MainConfig struct {
	Chains []struct {
		ChainID string `json:"chain-id"`
		Version struct {
			Type    string `json:"type"`
			Version string `json:"version"`
		} `json:"version"`
		GasPrices     string  `json:"gas-prices"`
		GasAdjustment float64 `json:"gas-adjustment"`
		NumberVals    int     `json:"number-vals"`
		NumberNode    int     `json:"number-node"`
		BlocksTTL     int     `json:"blocks-ttl"`
		Genesis       struct {
			Modify []struct {
				Key   string `json:"key"`
				Value string `json:"value"`
			} `json:"modify"`
			Accounts []struct {
				Name     string `json:"name"`
				Amount   string `json:"amount"`
				Address  string `json:"address"`
				Mnemonic string `json:"mnemonic"`
			} `json:"accounts"`
		} `json:"genesis"`
	} `json:"chains"`
}

type LogOutput struct {
	ChainID     string `json:"chain-id"`
	ChainName   string `json:"chain-name"`
	RPCAddress  string `json:"rpc-address"`
	GRPCAddress string `json:"grpc-address"`
}

func LoadConfig() (*MainConfig, error) {
	// read from current dir "config.json". Allow user ability for multiple configs
	bytes, err := ioutil.ReadFile("config.json")
	if err != nil {
		return nil, err
	}

	var config MainConfig
	err = json.Unmarshal(bytes, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}