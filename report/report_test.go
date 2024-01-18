// @Author Zhan 2024/1/18 20:34:00
package report

import (
	"fmt"
	"testing"
)

func TestNewReport(t *testing.T) {
	temps := []float64{20.0, 20.1, 20.3}
	report := NewReport(temps)
	fmt.Println(report.GetHtmlBody())
}
