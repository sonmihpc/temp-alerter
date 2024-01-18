// Package report @Author Zhan 2024/1/18 10:32:00
package report

import (
	"bytes"
	"embed"
	"html/template"
)

type Report struct {
	Temp1 float64
	Temp2 float64
}

func NewReport(temp1, temp2 float64) *Report {
	return &Report{
		Temp1: temp1,
		Temp2: temp2,
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
