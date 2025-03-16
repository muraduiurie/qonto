### HTTP Webhook Handler (Golang):

HTTP server that exposes an endpoint (`/webhook`) to handle POST requests from Prometheus Alertmanager.
Parses incoming JSON payloads into structured Go objects.

#### Alert Processing and Routing Logic:

Processes the payload and filters alerts based on:
- Team (`labels.team`)
- Severity (`labels.severity`)
- Routes the alerts according to defined rules.


#### Structured Logging and Output:

Simple, clear log formatting that highlights critical information:
- Job name (`dag_id`, `task_id`)
- Severity
- Team responsible
- Timestamps (alert start and end time)
- Description and Runbook URL for easy troubleshooting.
