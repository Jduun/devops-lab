{{- define "monitoring.name" -}}
{{ .Chart.Name }}
{{- end }}

{{- define "monitoring.fullname" -}}
{{ .Release.Name }}-{{ .Chart.Name }}
{{- end }}

{{- define "monitoring.labels" -}}
app.kubernetes.io/name: {{ include "monitoring.name" . }}
app.kubernetes.io/instance: {{ .Release.Name }}
app.kubernetes.io/version: {{ .Chart.AppVersion | quote }}
app.kubernetes.io/managed-by: {{ .Release.Service }}
helm.sh/chart: {{ .Chart.Name }}-{{ .Chart.Version | replace "+" "_" }}
{{- end }}

{{- define "monitoring.telegramSecretName" -}}
{{ include "monitoring.fullname" . }}-telegram
{{- end }}
