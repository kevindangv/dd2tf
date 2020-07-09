//go:generate go-bindata -o tpl.go tmpl

package main

import (
	datadog "github.com/DataDog/datadog-api-client-go/api/v1/datadog"
)

type Monitor struct {
}

func (m Monitor) getElement(client datadog.APIClient, id interface{}) (interface{}, error) {
	mon, err := client.GetMonitor(*datadog.Int(id.(int)))
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
	monitors, err := client.GetMonitors()
	if err != nil {
		return nil, err
	}
	for _, elem := range monitors {
		ids = append(ids, Item{id: *elem.Id, d: Monitor{}})
	}
	return ids, nil
}
