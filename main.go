package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"myscheduler/pkg/httpreq"
	"myscheduler/pkg/simplelog"

	"k8s.io/component-base/logs"
	"k8s.io/kubernetes/cmd/kube-scheduler/app"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	// 注册自定义 scheduler plugin
	command := app.NewSchedulerCommand(
		app.WithPlugin(httpreq.Name, httpreq.New),
		app.WithPlugin(simplelog.Name, simplelog.New),
	)

	if err := command.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}

	logs.InitLogs()
	defer logs.FlushLogs()
}
