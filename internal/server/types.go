package server

import "time"

type WebhookMessage struct {
	Status      string            `json:"status"`
	Receiver    string            `json:"receiver"`
	Alerts      []Alert           `json:"alerts"`
	GroupLabels map[string]string `json:"groupLabels"`
}

type Alert struct {
	Status       string           `json:"status"`
	Labels       AlertLabels      `json:"labels"`
	Annotations  AlertAnnotations `json:"annotations"`
	StartsAt     time.Time        `json:"startsAt"`
	EndsAt       time.Time        `json:"endsAt"`
	GeneratorURL string           `json:"generatorURL"`
}

type AlertLabels struct {
	AlertName string `json:"alertname"`
	Job       string `json:"job"`
	Severity  string `json:"severity"`
	Instance  string `json:"instance"`
	Team      string `json:"team"`
	TaskID    string `json:"task_id"`
	DagID     string `json:"dag_id"`
}

type AlertAnnotations struct {
	Summary     string `json:"summary"`
	Description string `json:"description"`
	Runbook     string `json:"runbook"`
}

type AlertRouter interface {
	routeAlert(alert Alert)
}

type stdoutRouter struct{}
type slackRouter struct{}
type emailRouter struct{}
