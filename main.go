package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"myscheduler/pkg/plugins"

	"k8s.io/component-base/logs"
	"k8s.io/kubernetes/cmd/kube-scheduler/app"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	fmt.Println("scheduler name:", plugins.Name)

	// 注册自定义 scheduler 插件
	command := app.NewSchedulerCommand(
		// plugins.Name 和 plugins.New 可以在
		// myscheduler/pkg/plugin/ 文件夹中的 plugins package 中获得
		app.WithPlugin(plugins.Name, plugins.New),
	)

	if err := command.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}

	logs.InitLogs()
	defer logs.FlushLogs()
}
