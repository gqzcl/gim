package main

import (
	"flag"
	"os"
	"os/signal"
	"syscall"

	"arod-im/internal/job"
	"arod-im/internal/job/conf"

	"github.com/bilibili/discovery/naming"

	resolver "github.com/bilibili/discovery/naming/grpc"
	log "github.com/golang/glog"
)

var (
	ver = "0.1.0"
)

func main() {
	flag.Parse()
	if err := conf.Init(); err != nil {
		panic(err)
	}
	log.Infof("goim-job [version: %s env: %+v] start", ver, conf.Conf.Env)
	// grpc register naming
	dis := naming.New(conf.Conf.Discovery)
	resolver.Register(dis)
	// job
	j := job.New(conf.Conf)
	go j.Consume()
	// signal
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		s := <-c
		log.Infof("goim-job get a signal %s", s.String())
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			j.Close()
			log.Infof("goim-job [version: %s] exit", ver)
			log.Flush()
			return
		case syscall.SIGHUP:
		default:
			return
		}
	}
}
