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
	times1 := cronnextzh.P带秒数的表达式解析器.Get获取未来N天内的执行时间(spec1, time.Now(), 7)

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
	times2 := cronnextzh.P带秒数的表达式解析器.Get计算未来N天内的执行时间(specs2, time.Now(), 3)

	fmt.Printf("Expressions: %v\n", specs2)
	fmt.Printf("Combined next %d execution times (sorted):\n", len(times2))
	for i, t := range times2 {
		fmt.Printf("  %2d. %s (Weekday: %d)\n", i+1, t.Format("2006-01-02 15:04:05"), t.Weekday())
	}

	// Example 3: Using minute-precision parser (示例3：使用分钟精度解析器)
	fmt.Println("\n=== Example 3: Minute-Precision Parser ===")
	spec3 := "15 10 * * 1-5" // 5-field format: Weekdays at 10:15 (5字段格式：工作日10:15)
	times3 := cronnextzh.P只到分的表达式解析器.Get获取未来N天内的执行时间(spec3, time.Now(), 5)

	fmt.Printf("Expression: %s (5-field format)\n", spec3)
	fmt.Printf("Next %d execution times:\n", len(times3))
	for i, t := range times3 {
		fmt.Printf("  %2d. %s (Weekday: %d)\n", i+1, t.Format("2006-01-02 15:04:05"), t.Weekday())
	}
}
