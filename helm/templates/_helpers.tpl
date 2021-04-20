{{/* vim: set filetype=mustache: */}}

{{/*
Expand the name of the chart.
*/}}
{{- define "app.name" -}}
{{- printf "%s-%s" .Release.Name .Chart.Name | trunc 63 | trimSuffix "-" -}}
{{- end -}}

{{/*
Create labels
*/}}
{{- define "app.labels" -}}
app.kubernetes.io/name: {{ include "app.name" . }}
helm.sh/chart: {{ .Chart.Name }}-{{ .Chart.Version }}
app.kubernetes.io/managed-by: {{ .Release.Service }}
app.kubernetes.io/instance: {{ .Release.Name }}
{{- end -}}


{{/*
Expand service account name.
*/}}
{{- define "app.serviceAccountName" -}}
{{- default (include "app.name" .) .Values.webapp.serviceAccount.name -}}
{{- end -}}

{{/*
Expand service name.
*/}}
{{- define "app.serviceName" -}}
{{- default (include "app.name" .) .Values.service.name }}
{{- end -}}



{{/*
Expand app name.
*/}}
{{- define "app.appName" -}}
{{- default (include "app.name" .) .Chart.Name -}}
{{- end -}}