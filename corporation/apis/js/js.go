package js

import (
	"fmt"
	"github.com/fastwego/wxwork/corporation"
	"github.com/fastwego/wxwork/util"
)

// Config 返回给用户jssdk配置信息
type Config struct {
	AppID     string `json:"app_id"`
	Timestamp string `json:"timestamp"`
	NonceStr  string `json:"nonce_str"`
	Signature string `json:"signature"`
	Url       string `json:"url"`
}

// AgentConfig 返回给用户jssdk配置信息
type AgentConfig struct {
	CorpID    string `json:"corp_id"`
	AgentID   string `json:"agent_id"`
	Timestamp string `json:"timestamp"`
	NonceStr  string `json:"nonce_str"`
	Signature string `json:"signature"`
	Url       string `json:"url"`
}

func GetConfig(ctx *corporation.App, url string) (config *Config, err error) {
	config = new(Config)
	ticketStr, err := corporation.GetJsApiTicket(ctx)
	if err != nil {
		return
	}

	nonceStr := util.GetRandString(16)
	timestamp := util.GetCurrentTs()
	str := fmt.Sprintf("jsapi_ticket=%s&noncestr=%s&timestamp=%d&url=%s", ticketStr, nonceStr, timestamp, url)
	sigStr := util.Signature(str)

	config.AppID = ctx.Corporation.Config.Corpid
	config.NonceStr = nonceStr
	config.Timestamp = timestamp
	config.Signature = sigStr
	return
}

func GetAgentConfig(ctx *corporation.App, url string) (config *AgentConfig, err error) {
	config = new(AgentConfig)
	ticketStr, err := corporation.GetAgentJsApiTicket(ctx)
	if err != nil {
		return
	}

	nonceStr := util.GetRandString(16)
	timestamp := util.GetCurrentTs()
	str := fmt.Sprintf("jsapi_ticket=%s&noncestr=%s&timestamp=%d&url=%s", ticketStr, nonceStr, timestamp, url)
	sigStr := util.Signature(str)

	config.CorpID = ctx.Corporation.Config.Corpid
	config.AgentID = ctx.Config.AgentId
	config.NonceStr = nonceStr
	config.Timestamp = timestamp
	config.Signature = sigStr
	return
}
