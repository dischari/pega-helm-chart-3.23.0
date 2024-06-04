{{- define  "backingservices.gke.backendConfig" -}}
apiVersion: cloud.google.com/v1
kind: BackendConfig
metadata:
  name: {{ .name }}
spec:
  timeoutSec: 40
  connectionDraining:
    drainingTimeoutSec: 60
  healthCheck:
    checkIntervalSec: 5
    healthyThreshold: 1
    port: {{ .healthCheckPort }}
    requestPath: {{ .requestPath }}
    timeoutSec: 5
    type: HTTP
    unhealthyThreshold: 2
---
{{ end }}