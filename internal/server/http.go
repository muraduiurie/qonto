package server

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"time"
)

type httpServer struct {
	AR AlertRouter
}

func detectAlertRouter() AlertRouter {
	ar := os.Getenv("ALERT_ROUTER")
	switch ar {
	case "stdout":
		return &stdoutRouter{}
	case "slack":
		return &slackRouter{}
	case "email":
		return &emailRouter{}
	default:
		return &stdoutRouter{}
	}
}

func newHTTPServer() *httpServer {
	return &httpServer{
		AR: detectAlertRouter(),
	}
}

func NewHTTPServer(addr string) *http.Server {
	httpsrv := newHTTPServer()
	r := mux.NewRouter()

	r.HandleFunc("/webhook", httpsrv.handleWebhook).Methods("POST")
	return &http.Server{
		Addr:    addr,
		Handler: r,
	}
}

func (s *httpServer) handleWebhook(w http.ResponseWriter, r *http.Request) {
	var req WebhookMessage
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Iterate through alerts and route them based on team and severity
	for _, alert := range req.Alerts {
		s.AR.routeAlert(alert)
	}

	w.WriteHeader(http.StatusOK)
}

// routeAlert routes alerts based on team and severity.
func (s *stdoutRouter) routeAlert(alert Alert) {
	switch alert.Labels.Severity {
	case "critical":
		// Critical alerts handling logic
		log.Printf("[CRITICAL] [%s] Job '%s' Task '%s' failed at %s\nDetails: %s\nRunbook: %s\n",
			alert.Labels.Team,
			alert.Labels.DagID,
			alert.Labels.TaskID,
			alert.StartsAt.Format(time.RFC3339),
			alert.Annotations.Description,
			alert.Annotations.Runbook,
		)
	default:
		// Non-critical alerts or default handling logic
		log.Printf("[ALERT] [%s] Severity: %s Job '%s' Task '%s' failed at %s\n",
			alert.Labels.Team,
			alert.Labels.Severity,
			alert.Labels.DagID,
			alert.Labels.TaskID,
			alert.StartsAt.Format(time.RFC3339),
		)
	}
}

func (s *slackRouter) routeAlert(alert Alert) {
	// Slack alerts handling logic
}

func (s *emailRouter) routeAlert(alert Alert) {
	// Email alerts handling logic
}
