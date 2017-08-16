package main

import (
	"log"
	"time"
	"github.com/zpatrick/go-config"
	"github.com/qframe/types/qchannel"
	"github.com/qframe/handler-influxdb"
	"github.com/qframe/types/metrics"
)

func Run(qChan qtypes_qchannel.QChan, cfg *config.Config, name string) {
	p, _ := qhandler_influxdb.New(qChan, cfg, name)
	p.Run()
}

func main() {
	qChan := qtypes_qchannel.NewQChan()
	qChan.Broadcast()
	cfgMap := map[string]string{
		"log.level": "trace",
		"handler.influxdb.inputs": "test",
		"handler.influxdb.batch-size": "1",
	}

	cfg := config.NewConfig(
		[]config.Provider{
			config.NewStatic(cfgMap),
		},
	)
	// handler
	p, err := qhandler_influxdb.New(qChan, cfg, "influxdb")
	if err != nil {
		log.Printf("[EE] Failed to create collector: %v", err)
		return
	}
	go p.Run()
	time.Sleep(time.Millisecond*time.Duration(250))
	m := qtypes_metrics.NewExt("test", "metric1", qtypes_metrics.Counter, 1.0, map[string]string{}, time.Now(), false)
	p.QChan.SendData(m)
	time.Sleep(time.Millisecond*time.Duration(50))
}
