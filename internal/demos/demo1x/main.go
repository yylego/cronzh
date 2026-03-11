package main

import (
	"fmt"
	"time"

	cronv3 "github.com/robfig/cron/v3"
	"github.com/yylego/cronzh/cronnextzh"
	"github.com/yylego/cronzh/crontaskzh"
)

func main() {
	// Basic cron task registration (基础定时任务注册)
	// Create a task that runs on weekdays at specific times (创建在工作日特定时间运行的任务)

	// Define task execution function (定义任务执行函数)
	taskFunction := func(taskName string) {
		fmt.Printf("[%s] Executing task: %s\n", time.Now().Format("15:04:05"), taskName)
	}

	// Create task list with cron expressions (使用 cron 表达式创建任务列表)
	taskList := crontaskzh.NewS定时任务列表([]*crontaskzh.T定时任务{
		{
			E任务名称:    "Morning Report",
			S定时表达式列表: []string{"0 30 8 * * 1-5"}, // Weekdays at 8:30 AM (工作日上午8:30)
			F执行函数:    taskFunction,
		},
		{
			E任务名称:    "Evening Summary",
			S定时表达式列表: []string{"0 0 20 * * 1-5"}, // Weekdays at 8:00 PM (工作日晚上8:00)
			F执行函数:    taskFunction,
		},
	})

	// Display future execution schedule (显示未来执行计划)
	fmt.Println("=== Scheduled Tasks for Next 7 Days ===")
	taskList.Debug(cronnextzh.P带秒数的表达式解析器, 7)

	// Register and run the cron scheduler (注册并运行定时调度器)
	cron := cronv3.New(cronv3.WithSeconds())
	taskList.Set注册定时任务(cron)
	cron.Start()

	// Run for 10 seconds to demonstrate (演示运行10秒)
	fmt.Println("\nCron scheduler running... (will stop after 10 seconds)")
	time.Sleep(10 * time.Second)

	// Stop the scheduler (停止调度器)
	ctx := cron.Stop()
	<-ctx.Done()
	fmt.Println("Cron scheduler stopped")
}
