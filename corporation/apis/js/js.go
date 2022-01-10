package js

import (
	"fmt"
	"github.com/fastwego/wxwork/corporation"
	"github.com/fastwego/wxwork/util"
)

// AgentConfig 返回给用户jssdk配置信息
type AgentConfig struct {
	CorpID    string `json:"corp_id"`
	AgentID   string `json:"agent_id"`
	Timestamp string `json:"timestamp"`
	NonceStr  string `json:"nonce_str"`
	Signature string `json:"signature"`
	Url       string `json:"url"`
}

func GetConfig(ctx *corporation.App, url string) (config map[string]string, err error) {
	ticketStr, err := corporation.GetJsApiTicket(ctx)
	if err != nil {
		return
	}

	nonceStr := util.GetRandString(16)
	timestamp := util.GetCurrentTs()
	str := fmt.Sprintf("jsapi_ticket=%s&noncestr=%s&timestamp=%d&url=%s", ticketStr, nonceStr, timestamp, url)
	sigStr := util.Signature(str)

	config["appId"] = ctx.Corporation.Config.Corpid
	config["nonceStr"] = nonceStr
	config["timestamp"] = timestamp
	config["signature"] = sigStr
	config["url"] = url
	return
}

func GetAgentConfig(ctx *corporation.App, url string) (config map[string]string, err error) {
	ticketStr, err := corporation.GetAgentJsApiTicket(ctx)
	if err != nil {
		return
	}

	nonceStr := util.GetRandString(16)
	timestamp := util.GetCurrentTs()
	str := fmt.Sprintf("jsapi_ticket=%s&noncestr=%s&timestamp=%d&url=%s", ticketStr, nonceStr, timestamp, url)
	sigStr := util.Signature(str)

	config["corpId"] = ctx.Corporation.Config.Corpid
	config["agentId"] = ctx.Config.AgentId
	config["nonceStr"] = nonceStr
	config["timestamp"] = timestamp
	config["signature"] = sigStr
	config["url"] = url
	return
}
