package main

import (
	"context"
	"fmt"

	inv "github.com/gemini-commerce/go-client-inventory"
)

func main() {
	cfg := inv.NewConfiguration()
	cfg.Servers = inv.ServerConfigurations{
		{
			URL:         "http://gem-inventory-gateway_go.docker:8081",
			Description: "No description provided",
		},
	}
	ctxWithAuth := context.WithValue(
		context.Background(),
		inv.ContextAPIKeys,
		map[string]inv.APIKey{
			"Authorization": {Key: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJyb2xlcyI6WyJpbnZlbnRvcnkucmVhZGVyIiwiaW52ZW50b3J5LndyaXRlciJdLCJ0ZW5hbnRJZCI6IioifQ.jWcjVI6Rt0NXF4b8LtPa2uFOmwdGWui5mWqdSRdmGGI"},
		},
	)
	client := inv.NewAPIClient(cfg)
	req := client.InventoryAPI.GetQtySalable(ctxWithAuth)
	body := inv.NewInventoryGetQtySalableRequest()
	tntId := "asas"
	sku := "asas"
	body.TenantId = &tntId
	body.Sku = &sku
	req = req.Body(*body)
	resp, httpResp, err := req.Execute()
	fmt.Printf("resp: %+v\n", resp)
	fmt.Printf("httpResp: %+v\n", httpResp)
	fmt.Printf("error: %+v\n", err)
}
