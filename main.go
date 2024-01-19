// @Author Zhan 2024/1/17 20:41:00
package main

import (
	"fmt"
	"github.com/coreos/go-systemd/daemon"
	"log"
	"os"
	"os/signal"
	"syscall"
	"wytfy.xyz/temp-alerter/config"
	"wytfy.xyz/temp-alerter/mail"
	"wytfy.xyz/temp-alerter/monitor"
)

func main() {
	cfg := config.Viper()
	mailInstance := mail.NewMailClient(fmt.Sprintf("%s:%v", cfg.SmtpHost, cfg.SmtpPort), cfg.SmtpEmail, cfg.SmtpUsername, cfg.SmtpPassword)
	tempMonitor := monitor.NewMonitor(cfg.SerialPort, cfg.SampleInterval, cfg.MaxTemp, cfg.MinTemp, cfg.MailReceiver, cfg.MailDelay, cfg.SensorNum, cfg.Position, mailInstance)
	tempMonitor.Run()

	// daemon
	if _, err := daemon.SdNotify(false, "READY=1"); err != nil {
		log.Printf("notification supported, but failed.")
	}
	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGHUP, syscall.SIGILL, syscall.SIGTERM, syscall.SIGQUIT)
	go func() {
		for s := range ch {
			switch s {
			case syscall.SIGHUP, syscall.SIGILL, syscall.SIGTERM, syscall.SIGQUIT:
				log.Println("temp-alerter exit.")
				ch <- s
			default:
				log.Printf("received other signal: %s, temp-alerter exit.", s)
			}
		}
	}()
	<-ch
}
