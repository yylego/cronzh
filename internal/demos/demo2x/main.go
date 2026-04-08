package main

import (
	"fmt"
	"time"

	cronv3 "github.com/robfig/cron/v3"
	"github.com/yylego/cronzh/cronnextzh"
	"github.com/yylego/cronzh/crontaskzh"
)

func main() {
	// Multiple cron expressions for a single task (单个任务使用多个 cron 表达式)
	// Different schedules for weekdays and weekends (工作日和周末使用不同的计划)

	// Define task execution functions (定义任务执行函数)
	backupTask := func(taskName string) {
		fmt.Printf("[%s] Running backup: %s\n", time.Now().Format("15:04:05"), taskName)
	}

	monitorTask := func(taskName string) {
		fmt.Printf("[%s] Running monitor: %s\n", time.Now().Format("15:04:05"), taskName)
	}

	// Create task list with multiple expressions per task (创建每个任务带多个表达式的任务列表)
	taskList := crontaskzh.NewS定时任务列表([]*crontaskzh.T定时任务{
		{
			E任务名称: "Database Backup",
			S定时表达式列表: []string{
				"0 0 2 * * 1-5", // Weekdays at 2:00 AM (工作日凌晨2点)
				"0 0 3 * * 0,6", // Weekends at 3:00 AM (周末凌晨3点)
			},
			F执行函数: backupTask,
		},
		{
			E任务名称: "System Monitor",
			S定时表达式列表: []string{
				"0 */30 * * * *", // Every 30 minutes (每30分钟)
			},
			F执行函数: monitorTask,
		},
	})

	// Display future execution schedule (显示未来执行计划)
	fmt.Println("=== Scheduled Tasks for Next 5 Days ===")
	if err := taskList.Debug(cronnextzh.P带秒数的表达式解析器, 5); err != nil {
		panic(err)
	}

	// Register and run the cron scheduler (注册并运行定时调度器)
	cron := cronv3.New(cronv3.WithSeconds())
	if err := taskList.Set注册定时任务(cron); err != nil {
		panic(err)
	}
	cron.Start()

	// Run for 10 seconds to demonstrate (演示运行10秒)
	fmt.Println("\nCron scheduler running... (will stop after 10 seconds)")
	time.Sleep(10 * time.Second)

	// Stop the scheduler (停止调度器)
	ctx := cron.Stop()
	<-ctx.Done()
	fmt.Println("Cron scheduler stopped")
}
