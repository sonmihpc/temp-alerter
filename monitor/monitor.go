// @Author Zhan 2024/1/18 10:40:00
package monitor

import (
	"github.com/goburrow/serial"
	"github.com/things-go/go-modbus"
	"log"
	"sync"
	"time"
	"wytfy.xyz/temp-alerter/mail"
	"wytfy.xyz/temp-alerter/report"
)

type Monitor struct {
	interval int

	maxTemp float64
	minTemp float64

	receiver  []string
	mailDelay int
	sensorNum int
	position  string
	status    bool
	statusMu  sync.Mutex

	client     modbus.Client
	mailClient *mail.Client
	closeCh    chan interface{}
}

func NewMonitor(address string, interval int, maxTemp, minTemp float64, receiver []string, mailDelay, sensorNum int, position string, mailClient *mail.Client) *Monitor {
	p := modbus.NewRTUClientProvider(
		modbus.WithSerialConfig(serial.Config{
			Address:  address,
			BaudRate: 9600,
			DataBits: 8,
			StopBits: 1,
			Parity:   "N",
			Timeout:  modbus.SerialDefaultTimeout,
		}))
	client := modbus.NewClient(p)
	return &Monitor{
		interval:   interval,
		maxTemp:    maxTemp,
		minTemp:    minTemp,
		receiver:   receiver,
		client:     client,
		mailDelay:  mailDelay,
		sensorNum:  sensorNum,
		position:   position,
		mailClient: mailClient,
	}
}

func (m *Monitor) connect() {
	m.statusMu.Lock()
	defer m.statusMu.Unlock()
	if err := m.client.Connect(); err != nil {
		log.Println(err)
		return
	}
	m.status = true
}

func (m *Monitor) close() {
	m.statusMu.Lock()
	defer m.statusMu.Unlock()
	if err := m.client.Close(); err != nil {
		log.Println(err)
		return
	}
	m.status = false
}

func (m *Monitor) Run() {
	go m.runInBackground()
}

func (m *Monitor) runInBackground() {
	m.connect()
	defer m.close()

	if m.status == false {
		log.Println("fail connect to modbus device")
		return
	}

	ticker := time.NewTicker(time.Duration(m.interval) * time.Second)

	for {
		select {
		case <-ticker.C:
			result, err := m.temperatureSample()
			if err != nil {
				log.Println(err)
				continue
			}
			log.Printf("environment temperature: %v °C\n", result)
			if m.tempOutOfRange(result) {
				mailBody, err := m.generateMailBody(result)
				if err != nil {
					log.Println(err)
					continue
				}
				// 以下为发送邮件逻辑
				if err := m.mailClient.Send(m.receiver, "Temperature Warning!", mailBody); err != nil {
					log.Println(err)
					continue
				}
				log.Println("send email")
				// 发送成功则进行休眠，防止邮件滥发
				time.Sleep(time.Duration(m.mailDelay) * time.Minute)
				continue
			}
		case <-m.closeCh:
			log.Println("stop temperature monitor")
			return
		default:
			time.Sleep(time.Duration(m.interval) * time.Second)
		}
	}
}

func (m *Monitor) temperatureSample() ([]float64, error) {
	result, err := m.client.ReadHoldingRegisters(1, 0x0000, uint16(m.sensorNum))
	if err != nil {
		log.Println(err)
		return nil, err
	}
	resInt := make([]float64, 0)
	for _, i := range result {
		resInt = append(resInt, float64(i)/10.0)
	}
	return resInt, err
}

func (m *Monitor) tempOutOfRange(res []float64) bool {
	flag := false
	for _, i := range res {
		if i >= m.maxTemp || i <= m.minTemp {
			flag = true
			break
		}
	}
	return flag
}

func (m *Monitor) generateMailBody(res []float64) (string, error) {
	rep := report.NewReport(m.position, res)
	return rep.GetHtmlBody()
}
