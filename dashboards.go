//go:generate go-bindata -o tpl.go tmpl

package main

import (
	"context"
	"fmt"
	"os"
	datadog "github.com/DataDog/datadog-api-client-go/api/v1/datadog"
)

type Dashboard struct {
}

func (d Dashboard) getElement(client datadog.APIClient, id interface{}) (interface{}, error) {
	ctx := context.WithValue(
		context.Background(),
		datadog.ContextAPIKeys,
		map[string]datadog.APIKey{
			"apiKeyAuth": {
				Key: os.Getenv("DATADOG_API_KEY"),
			},
			"appKeyAuth": {
				Key: os.Getenv("DATADOG_APP_KEY"),
			},
		},
    	)
	idStr := fmt.Sprintf("%v", id)
	dash, _, err := client.DashboardsApi.GetDashboard(ctx, idStr).Execute()
	return dash, err
}

func (d Dashboard) getAsset() string {
	return "tmpl/dashboard.tmpl"
}

func (d Dashboard) getName() string {
	return "dashboards"
}

func (d Dashboard) String() string {
	return d.getName()
}

func (d Dashboard) getAllElements(client datadog.APIClient) ([]Item, error) {
	var ids []Item
	ctx := context.WithValue(
		context.Background(),
		datadog.ContextAPIKeys,
		map[string]datadog.APIKey{
			"apiKeyAuth": {
				Key: os.Getenv("DATADOG_API_KEY"),
			},
			"appKeyAuth": {
				Key: os.Getenv("DATADOG_APP_KEY"),
			},
		},
	)
	dashboards, _, err := client.DashboardsApi.ListDashboards(ctx).Execute()
	if err != nil {
		return ids, err
	}
	for _, elem := range *dashboards.Dashboards {
		fmt.Println("found dashboard", *elem.Id)
		ids = append(ids, Item{id: *elem.Id, d: Dashboard{}})
	}
	return ids, nil
}
