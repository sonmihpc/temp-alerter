// Package report @Author Zhan 2024/1/18 10:32:00
package report

import (
	"bytes"
	"embed"
	"html/template"
)

type Report struct {
	Position string
	Temps    []float64
}

func NewReport(position string, temps []float64) *Report {
	return &Report{
		position,
		temps,
	}
}

//go:embed report.tpl
var report embed.FS

func (r *Report) GetHtmlBody() (string, error) {
	var b bytes.Buffer
	t := template.Must(template.ParseFS(report, "report.tpl"))
	if err := t.ExecuteTemplate(&b, "report", r); err != nil {
		return "", err
	}
	return b.String(), nil
}
