[![GitHub Workflow Status (branch)](https://img.shields.io/github/actions/workflow/status/yylego/cronzh/release.yml?branch=main&label=BUILD)](https://github.com/yylego/cronzh/actions/workflows/release.yml?query=branch%3Amain)
[![GoDoc](https://pkg.go.dev/badge/github.com/yylego/cronzh)](https://pkg.go.dev/github.com/yylego/cronzh)
[![Coverage Status](https://img.shields.io/coveralls/github/yylego/cronzh/main.svg)](https://coveralls.io/github/yylego/cronzh?branch=main)
[![Supported Go Versions](https://img.shields.io/badge/Go-1.25+-lightgrey.svg)](https://go.dev/)
[![GitHub Release](https://img.shields.io/github/release/yylego/cronzh.svg)](https://github.com/yylego/cronzh/releases)
[![Go Report Card](https://goreportcard.com/badge/github.com/yylego/cronzh)](https://goreportcard.com/report/github.com/yylego/cronzh)

# cronzh

Chinese-named package extending `github.com/robfig/cron/v3` with intuitive APIs to manage scheduled tasks

---

<!-- TEMPLATE (EN) BEGIN: LANGUAGE NAVIGATION -->
## CHINESE README

[中文说明](README.zh.md)
<!-- TEMPLATE (EN) END: LANGUAGE NAVIGATION -->

## DISCLAIMER

Writing Go code in Chinese is a viable technique, but something to avoid in production engineering. This approach should not be used in serious and business settings. Teams and companies that embrace it could face contempt from peers and negative judgment across the profession. In business companies, this practice is even more prone to becoming a target of public criticism. This project is dedicated to research and academic studies. Do not use this approach in production.

## Main Features

🎯 **Chinese Function Names**: Intuitive Chinese-named wrappers around robfig/cron
⏰ **Multiple Schedules**: Each task supports multiple cron expressions
📊 **Schedule Preview**: Debug mode to visualize future execution times
🔧 **Flexible Parsers**: Both second-precision (6-field) and minute-precision (5-field) support
📝 **Detailed Logging**: Built-in zaplog integration for tracking task execution

## Installation

```bash
go get github.com/yylego/cronzh
```

## Quick Start

### Basic Task Registration

Create scheduled tasks and add them to the cron instance with cron expressions.

```go
package main

import (
	"fmt"
	"time"

	"github.com/yylego/cronzh/cronnextzh"
	"github.com/yylego/cronzh/crontaskzh"
	cronv3 "github.com/robfig/cron/v3"
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
	if err := taskList.Debug(cronnextzh.P带秒数的表达式解析器, 7); err != nil {
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
```

⬆️ **Source:** [Source](internal/demos/demo1x/main.go)

### Multiple Expressions Within Single Task

A single task can have different schedules for weekdays and weekends.

```go
package main

import (
	"fmt"
	"time"

	"github.com/yylego/cronzh/cronnextzh"
	"github.com/yylego/cronzh/crontaskzh"
	cronv3 "github.com/robfig/cron/v3"
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
	taskList.Debug(cronnextzh.P带秒数的表达式解析器, 5)

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
```

⬆️ **Source:** [Source](internal/demos/demo2x/main.go)

### Expression Parsing and Preview

Parse cron expressions and calculate future execution times without running tasks.

```go
package main

import (
	"fmt"
	"time"

	"github.com/yylego/cronzh/cronnextzh"
)

func main() {
	// Parsing cron expressions and calculating future execution times (解析 cron 表达式并计算未来执行时间)
	// Useful for previewing schedules without running tasks (用于预览计划而无需运行任务)

	// Example 1: Single cron expression (示例1：单个 cron 表达式)
	fmt.Println("=== Example 1: Single Expression ===")
	spec1 := "0 15 10 * * 1-5" // Weekdays at 10:15 AM (工作日上午10:15)
	times1, err := cronnextzh.P带秒数的表达式解析器.Get获取未来N天内的执行时间(spec1, time.Now(), 7)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Expression: %s\n", spec1)
	fmt.Printf("Next %d execution times:\n", len(times1))
	for i, t := range times1 {
		fmt.Printf("  %2d. %s (Weekday: %d)\n", i+1, t.Format("2006-01-02 15:04:05"), t.Weekday())
	}

	// Example 2: Multiple cron expressions (示例2：多个 cron 表达式)
	fmt.Println("\n=== Example 2: Multiple Expressions ===")
	specs2 := []string{
		"0 30 9 * * 1-5",  // Weekdays at 9:30 AM (工作日上午9:30)
		"0 0 14 * * 1-5",  // Weekdays at 2:00 PM (工作日下午2:00)
		"0 30 18 * * 1-5", // Weekdays at 6:30 PM (工作日下午6:30)
	}
	times2, err := cronnextzh.P带秒数的表达式解析器.Get计算未来N天内的执行时间(specs2, time.Now(), 3)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Expressions: %v\n", specs2)
	fmt.Printf("Combined next %d execution times (sorted):\n", len(times2))
	for i, t := range times2 {
		fmt.Printf("  %2d. %s (Weekday: %d)\n", i+1, t.Format("2006-01-02 15:04:05"), t.Weekday())
	}

	// Example 3: Using minute-precision parser (示例3：使用分钟精度解析器)
	fmt.Println("\n=== Example 3: Minute-Precision Parser ===")
	spec3 := "15 10 * * 1-5" // 5-field format: Weekdays at 10:15 (5字段格式：工作日10:15)
	times3, err := cronnextzh.P只到分的表达式解析器.Get获取未来N天内的执行时间(spec3, time.Now(), 5)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Expression: %s (5-field format)\n", spec3)
	fmt.Printf("Next %d execution times:\n", len(times3))
	for i, t := range times3 {
		fmt.Printf("  %2d. %s (Weekday: %d)\n", i+1, t.Format("2006-01-02 15:04:05"), t.Weekday())
	}
}
```

⬆️ **Source:** [Source](internal/demos/demo3x/main.go)

## API Reference

### Core Modules

#### cronnextzh - Cron Expression Parser

**Predefined Parsers:**

- `P带秒数的表达式解析器` - Second-precision parser (6 fields: second, minute, hour, day, month, weekday)
- `P只到分的表达式解析器` - Minute-precision parser (5 fields: minute, hour, day, month, weekday)

**Main Type:**

```go
type P表达式解析器 cron.Parser
```

**Main Methods:**

- `New(parser cron.Parser) *P表达式解析器` - Create custom parser
- `Get获取未来N天内的执行时间(spec string, since time.Time, nDate int) ([]time.Time, error)` - Calculate execution times given a single expression
- `Get计算未来N天内的执行时间(specs []string, since time.Time, nDate int) ([]time.Time, error)` - Calculate execution times given multiple expressions (sorted)

#### crontaskzh - Task List Management

**Main Types:**

```go
type T定时任务 struct {
    S定时表达式列表 []string           // List of cron expressions
    E任务名称      string             // Task name
    F执行函数      func(e任务名称 string) // Execution function
}

type S定时任务列表 []*T定时任务
```

**Main Methods:**

- `NewS定时任务列表(s定时任务列表 []*T定时任务) S定时任务列表` - Create new task list
- `Set注册定时任务(cron *cron.Cron) error` - Add all tasks to the cron instance
- `Debug(p表达式解析器 *cronnextzh.P表达式解析器, nDate int) error` - Show the future execution schedule

## Cron Expression Format

**6-field format (with seconds):**
```
┌─── second (0-59)
│ ┌─── minute (0-59)
│ │ ┌─── hour (0-23)
│ │ │ ┌─── day (1-31)
│ │ │ │ ┌─── month (1-12)
│ │ │ │ │ ┌─── weekday (0-6, Sunday=0)
│ │ │ │ │ │
* * * * * *
```

**5-field format (minute precision):**
```
┌─── minute (0-59)
│ ┌─── hour (0-23)
│ │ ┌─── day (1-31)
│ │ │ ┌─── month (1-12)
│ │ │ │ ┌─── weekday (0-6, Sunday=0)
│ │ │ │ │
* * * * *
```

**Common Examples:**

- `"0 30 8 * * 1-5"` - Weekdays at 8:30 AM
- `"*/5 * * * * *"` - Every 5 seconds
- `"0 0 2 * * *"` - Every day at 2:00 AM
- `"0 0 0 * * 0"` - Every Sunday at midnight
- `"0 */30 * * * *"` - Every 30 minutes

## Design Concept

This package follows these principles:

1. **Chinese Naming**: Functions use intuitive Chinese names matching robfig/cron concepts
2. **Multiple Schedules**: Single tasks can have multiple cron expressions
3. **Debug-Friendly**: Built-in schedule visualization before deployment
4. **Type-Safe Operations**: Leverages Go's type system for safe task management
5. **Flexible Parsing**: Support both second-precision and minute-precision formats

---

<!-- TEMPLATE (EN) BEGIN: STANDARD PROJECT FOOTER -->
<!-- VERSION 2025-11-20 04:26:32.402216 +0000 UTC -->

## 📄 License

MIT License - see [LICENSE](LICENSE).

---

## 💬 Contact & Feedback

Contributions are welcome! Report bugs, suggest features, and contribute code:

- 🐛 **Mistake reports?** Open an issue on GitHub with reproduction steps
- 💡 **Fresh ideas?** Create an issue to discuss
- 📖 **Documentation confusing?** Report it so we can improve
- 🚀 **Need new features?** Share the use cases to help us understand requirements
- ⚡ **Performance issue?** Help us optimize through reporting slow operations
- 🔧 **Configuration problem?** Ask questions about complex setups
- 📢 **Follow project progress?** Watch the repo to get new releases and features
- 🌟 **Success stories?** Share how this package improved the workflow
- 💬 **Feedback?** We welcome suggestions and comments

---

## 🔧 Development

New code contributions, follow this process:

1. **Fork**: Fork the repo on GitHub (using the webpage UI).
2. **Clone**: Clone the forked project (`git clone https://github.com/yourname/repo-name.git`).
3. **Navigate**: Navigate to the cloned project (`cd repo-name`)
4. **Branch**: Create a feature branch (`git checkout -b feature/xxx`).
5. **Code**: Implement the changes with comprehensive tests
6. **Testing**: (Golang project) Ensure tests pass (`go test ./...`) and follow Go code style conventions
7. **Documentation**: Update documentation to support client-facing changes and use significant commit messages
8. **Stage**: Stage changes (`git add .`)
9. **Commit**: Commit changes (`git commit -m "Add feature xxx"`) ensuring backward compatible code
10. **Push**: Push to the branch (`git push origin feature/xxx`).
11. **PR**: Open a merge request on GitHub (on the GitHub webpage) with detailed description.

Please ensure tests pass and include relevant documentation updates.

---

## 🌟 Support

Welcome to contribute to this project via submitting merge requests and reporting issues.

**Project Support:**

- ⭐ **Give GitHub stars** if this project helps you
- 🤝 **Share with teammates** and (golang) programming friends
- 📝 **Write tech blogs** about development tools and workflows - we provide content writing support
- 🌟 **Join the ecosystem** - committed to supporting open source and the (golang) development scene

**Have Fun Coding with this package!** 🎉🎉🎉

<!-- TEMPLATE (EN) END: STANDARD PROJECT FOOTER -->

---

## GitHub Stars

[![Stargazers](https://starchart.cc/yylego/cronzh.svg?variant=adaptive)](https://starchart.cc/yylego/cronzh)
