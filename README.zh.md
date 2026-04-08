[![GitHub Workflow Status (branch)](https://img.shields.io/github/actions/workflow/status/yylego/cronzh/release.yml?branch=main&label=BUILD)](https://github.com/yylego/cronzh/actions/workflows/release.yml?query=branch%3Amain)
[![GoDoc](https://pkg.go.dev/badge/github.com/yylego/cronzh)](https://pkg.go.dev/github.com/yylego/cronzh)
[![Coverage Status](https://img.shields.io/coveralls/github/yylego/cronzh/main.svg)](https://coveralls.io/github/yylego/cronzh?branch=main)
[![Supported Go Versions](https://img.shields.io/badge/Go-1.25+-lightgrey.svg)](https://go.dev/)
[![GitHub Release](https://img.shields.io/github/release/yylego/cronzh.svg)](https://github.com/yylego/cronzh/releases)
[![Go Report Card](https://goreportcard.com/badge/github.com/yylego/cronzh)](https://goreportcard.com/report/github.com/yylego/cronzh)

# cronzh

围绕 `github.com/robfig/cron/v3` 的中文封装，提供直观的 API 进行定时任务管理

---

<!-- TEMPLATE (ZH) BEGIN: LANGUAGE NAVIGATION -->
## 英文文档

[ENGLISH README](README.md)
<!-- TEMPLATE (ZH) END: LANGUAGE NAVIGATION -->

## 免责声明

使用中文编写 Go 代码在技术上可行，但在实际工程中极不推荐。这种做法不适用于任何正式或商业场景，采用它的团队或公司可能面临同行的轻视和行业的负面评价，在商业公司中尤其可能成为被业内议论的对象。该项目仅供学习和研究，请勿在正式工程中采用。

## 主要特性

🎯 **中文函数名**: 基于 robfig/cron 的直观中文命名函数
⏰ **多定时表达式**: 支持单个任务配置多个 cron 表达式
📊 **计划预览**: 调试模式可视化未来执行时间
🔧 **灵活解析器**: 同时支持秒级精度（6 字段）和分钟级精度（5 字段）
📝 **详细日志**: 内置 zaplog 集成用于跟踪任务执行

## 安装

```bash
go get github.com/yylego/cronzh
```

## 快速开始

### 基础任务注册

使用 cron 表达式创建和注册定时任务。

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

⬆️ **源码:** [源码](internal/demos/demo1x/main.go)

### 单个任务多个表达式

单个任务可以为工作日和周末配置不同的计划。

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

⬆️ **源码:** [源码](internal/demos/demo2x/main.go)

### 表达式解析和预览

解析 cron 表达式并计算未来执行时间，无需运行任务。

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

⬆️ **源码:** [源码](internal/demos/demo3x/main.go)

## API 参考

### 核心模块

#### cronnextzh - Cron 表达式解析器

**预定义解析器:**

- `P带秒数的表达式解析器` - 秒级精度解析器（6 个字段：秒 分 时 日 月 周）
- `P只到分的表达式解析器` - 分钟级精度解析器（5 个字段：分 时 日 月 周）

**主要类型:**

```go
type P表达式解析器 cron.Parser
```

**主要方法:**

- `New(parser cron.Parser) *P表达式解析器` - 创建自定义解析器
- `Get获取未来N天内的执行时间(spec string, since time.Time, nDate int) ([]time.Time, error)` - 计算单个表达式的执行时间
- `Get计算未来N天内的执行时间(specs []string, since time.Time, nDate int) ([]time.Time, error)` - 计算多个表达式的执行时间（已排序）

#### crontaskzh - 任务列表管理

**主要类型:**

```go
type T定时任务 struct {
    S定时表达式列表 []string           // cron 表达式列表
    E任务名称      string             // 任务名称
    F执行函数      func(e任务名称 string) // 执行函数
}

type S定时任务列表 []*T定时任务
```

**主要方法:**

- `NewS定时任务列表(s定时任务列表 []*T定时任务) S定时任务列表` - 创建新任务列表
- `Set注册定时任务(cron *cron.Cron) error` - 注册所有任务到 cron 实例
- `Debug(p表达式解析器 *cronnextzh.P表达式解析器, nDate int) error` - 显示未来执行计划

## Cron 表达式格式

**6 字段格式（包含秒数）:**
```
┌─── 秒 (0-59)
│ ┌─── 分 (0-59)
│ │ ┌─── 时 (0-23)
│ │ │ ┌─── 日 (1-31)
│ │ │ │ ┌─── 月 (1-12)
│ │ │ │ │ ┌─── 周 (0-6, 周日=0)
│ │ │ │ │ │
* * * * * *
```

**5 字段格式（分钟精度）:**
```
┌─── 分 (0-59)
│ ┌─── 时 (0-23)
│ │ ┌─── 日 (1-31)
│ │ │ ┌─── 月 (1-12)
│ │ │ │ ┌─── 周 (0-6, 周日=0)
│ │ │ │ │
* * * * *
```

**常见示例:**

- `"0 30 8 * * 1-5"` - 工作日上午 8:30
- `"*/5 * * * * *"` - 每 5 秒
- `"0 0 2 * * *"` - 每天凌晨 2:00
- `"0 0 0 * * 0"` - 每周日午夜
- `"0 */30 * * * *"` - 每 30 分钟

## 设计理念

本包遵循以下原则：

1. **中文命名**: 函数使用符合 robfig/cron 概念的直观中文名称
2. **多定时表达式**: 单个任务可配置多个 cron 表达式
3. **调试友好**: 内置计划可视化功能，便于部署前测试
4. **类型安全**: 利用 Go 的类型系统实现安全的任务管理
5. **灵活解析**: 同时支持秒级精度和分钟级精度格式

---

<!-- TEMPLATE (ZH) BEGIN: STANDARD PROJECT FOOTER -->
<!-- VERSION 2025-11-20 04:26:32.402216 +0000 UTC -->

## 📄 许可证类型

MIT 许可证 - 详见 [LICENSE](LICENSE)。

---

## 💬 联系与反馈

非常欢迎贡献代码！报告 BUG、建议功能、贡献代码：

- 🐛 **问题报告？** 在 GitHub 上提交问题并附上重现步骤
- 💡 **新颖思路？** 创建 issue 讨论
- 📖 **文档疑惑？** 报告问题，帮助我们改进文档
- 🚀 **需要功能？** 分享使用场景，帮助理解需求
- ⚡ **性能瓶颈？** 报告慢操作，帮助我们优化性能
- 🔧 **配置困扰？** 询问复杂设置的相关问题
- 📢 **关注进展？** 关注仓库以获取新版本和功能
- 🌟 **成功案例？** 分享这个包如何改善工作流程
- 💬 **反馈意见？** 欢迎提出建议和意见

---

## 🔧 代码贡献

新代码贡献，请遵循此流程：

1. **Fork**：在 GitHub 上 Fork 仓库（使用网页界面）
2. **克隆**：克隆 Fork 的项目（`git clone https://github.com/yourname/repo-name.git`）
3. **导航**：进入克隆的项目（`cd repo-name`）
4. **分支**：创建功能分支（`git checkout -b feature/xxx`）
5. **编码**：实现您的更改并编写全面的测试
6. **测试**：（Golang 项目）确保测试通过（`go test ./...`）并遵循 Go 代码风格约定
7. **文档**：为面向用户的更改更新文档，并使用有意义的提交消息
8. **暂存**：暂存更改（`git add .`）
9. **提交**：提交更改（`git commit -m "Add feature xxx"`）确保向后兼容的代码
10. **推送**：推送到分支（`git push origin feature/xxx`）
11. **PR**：在 GitHub 上打开 Merge Request（在 GitHub 网页上）并提供详细描述

请确保测试通过并包含相关的文档更新。

---

## 🌟 项目支持

非常欢迎通过提交 Merge Request 和报告问题来为此项目做出贡献。

**项目支持：**

- ⭐ **给予星标**如果项目对您有帮助
- 🤝 **分享项目**给团队成员和（golang）编程朋友
- 📝 **撰写博客**关于开发工具和工作流程 - 我们提供写作支持
- 🌟 **加入生态** - 致力于支持开源和（golang）开发场景

**祝你用这个包编程愉快！** 🎉🎉🎉

<!-- TEMPLATE (ZH) END: STANDARD PROJECT FOOTER -->

---

## GitHub 标星点赞

[![Stargazers](https://starchart.cc/yylego/cronzh.svg?variant=adaptive)](https://starchart.cc/yylego/cronzh)
