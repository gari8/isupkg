package isupkg

import (
	"cmp"
	"github.com/grafana/pyroscope-go"
	"runtime"
	"time"
)

const (
	defaultMutexProfileRate = 1
	defaultBlockProfileRate = 1
	defaultAppName          = "isucon"
)

type Profile struct {
	profiler
	MutexProfileRate int
	BlockProfileRate int
	Hostname         string
	ServerAddress    string
	AppName          string
	Version          string
}

type profiler interface {
	Stop() error
	Flush(wait bool)
}

func (p *Profile) Run() error {
	runtime.SetMutexProfileFraction(cmp.Or(p.MutexProfileRate, defaultMutexProfileRate))
	runtime.SetBlockProfileRate(cmp.Or(p.BlockProfileRate, defaultBlockProfileRate))
	p.AppName = cmp.Or(p.AppName, defaultAppName)
	p.ServerAddress = cmp.Or(p.ServerAddress, "http://monitoring.tail10361b.ts.net:4040")
	p.Hostname = cmp.Or(p.Hostname, defaultAppName)
	p.Version = cmp.Or(p.Version, time.Now().Format("2006-01-02T15:04:05Z07:00"))

	var err error
	p.profiler, err = pyroscope.Start(pyroscope.Config{
		ApplicationName: p.AppName,
		ServerAddress:   p.ServerAddress,
		Logger:          pyroscope.StandardLogger,

		Tags: map[string]string{
			"hostname": p.Hostname,
			"version":  p.Version,
		},

		ProfileTypes: append([]pyroscope.ProfileType{
			pyroscope.ProfileGoroutines,
			pyroscope.ProfileMutexCount,
			pyroscope.ProfileMutexDuration,
			pyroscope.ProfileBlockCount,
			pyroscope.ProfileBlockDuration,
		}, pyroscope.DefaultProfileTypes...),
	})
	if err != nil {
		return err
	}
	return nil
}

func (p *Profile) Stop() error {
	return p.profiler.Stop()
}
