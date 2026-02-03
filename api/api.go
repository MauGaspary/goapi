package api

import (
	"net/http"
	"encoding/json"
	
)

type AccountBalanceParams struct {
	AccountID string `json:"account_id"`
}

type AccountBalanceResponse struct {
	Code	int     `json:"code"`
	Balance	float64 `json:"balance"`
}

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}