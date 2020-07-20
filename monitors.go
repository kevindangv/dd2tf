//go:generate go-bindata -o tpl.go tmpl

package main

import (
	"context"
	"os"
	datadog "github.com/DataDog/datadog-api-client-go/api/v1/datadog"
)

type Monitor struct {
}

func (m Monitor) getElement(client datadog.APIClient, id interface{}) (interface{}, error) {
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
        mon, _, err := client.MonitorsApi.GetMonitor(ctx, id.(int64)).Execute()
        return mon, err
}

func (m Monitor) getAsset() string {
	return "tmpl/monitor.tmpl"
}

func (m Monitor) getName() string {
	return "monitors"
}

func (m Monitor) String() string {
	return m.getName()
}

func (m Monitor) getAllElements(client datadog.APIClient) ([]Item, error) {
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
	monitors, _, err := client.MonitorsApi.ListMonitors(ctx).Execute()
        if err != nil {
                return nil, err
        }
        for _, elem := range monitors {
                ids = append(ids, Item{id: *elem.Id, d: Monitor{}})
        }
        return ids, nil
}

