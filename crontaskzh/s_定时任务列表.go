package crontaskzh

import (
	"fmt"
	"time"

	"github.com/robfig/cron/v3"
	"github.com/yylego/cronzh/cronnextzh"
	"github.com/yylego/neatjson/neatjsons"
	"github.com/yylego/rese"
	"github.com/yylego/sortx"
	"github.com/yylego/timezh"
	"github.com/yylego/zaplog"
	"go.uber.org/zap"
)

// T定时任务 represents a scheduled task with one or multiple cron expressions
// A single task can have multiple activation times defined via different cron expressions
//
// T定时任务 表示带有一个或多个 cron 表达式的定时任务
// 单个任务可以有多个由不同 cron 表达式定义的触发时间
type T定时任务 struct {
	S定时表达式列表 []string           // List of crontab expressions (spec-list) / crontab 的表达式列表
	E任务名称    string             // Task name / 任务名称
	F执行函数    func(e任务名称 string) `json:"-"` // Execution function / 执行函数
}

// S定时任务列表 is a collection of scheduled tasks
// Provides methods to set-up tasks and debug execution schedules
//
// S定时任务列表 是定时任务的集合
// 提供注册任务和调试执行计划的方法
type S定时任务列表 []*T定时任务

// NewS定时任务列表 creates a new task list from a slice of tasks
// s定时任务列表: slice of T定时任务 pointers
// Returns a new S定时任务列表 instance
//
// NewS定时任务列表 从任务切片创建新的任务列表
// s定时任务列表: T定时任务 指针切片
// 返回新的 S定时任务列表 实例
func NewS定时任务列表(s定时任务列表 []*T定时任务) S定时任务列表 {
	return s定时任务列表
}

// Set注册定时任务 registers the tasks in the list to a cron instance
// Each task's cron expressions get registered with the associated execution function
// Important: Variable capture is handled as expected in Go 1.22+ (see inline comment)
//
// Set注册定时任务 将列表中的所有任务注册到 cron 实例
// 每个任务的 cron 表达式都与其关联的执行函数一起注册
// 重要：在 Go 1.22+ 中正确处理变量捕获（见行内注释）
func (cs S定时任务列表) Set注册定时任务(cron *cron.Cron) {
	zaplog.LOG.Debug("新建定时任务")
	zaplog.SUG.Debugln("定时任务列表", neatjsons.S(cs))
	for idx, one := range cs {
		zaplog.LOG.Debug("注册定时任务", zap.Int("idx", idx), zap.Int("spec_list", len(one.S定时表达式列表)))

		for _, spec := range one.S定时表达式列表 {
			// Extract variables at this point - needed in Go <= 1.21 due to goroutine closure semantics
			// 在此提取变量 - 在 Go <= 1.21 中由于 goroutine 闭包行为而至关重要
			e任务名称 := one.E任务名称
			f执行函数 := one.F执行函数
			zaplog.LOG.Debug("注册定时条件", zap.String("spec", spec), zap.String("event_name", e任务名称))

			rese.C1(cron.AddFunc(spec, func() {
				zaplog.LOG.Debug("开始执行", zap.String("spec", spec), zap.String("event_name", e任务名称))
				f执行函数(e任务名称)
				zaplog.LOG.Debug("执行完毕", zap.String("spec", spec), zap.String("event_name", e任务名称))
			}))
		}
	}
	zaplog.LOG.Debug("任务注册完毕")
}

// Debug shows the scheduled execution times within the next N days
// Helps check task schedules before deploying to production
// p表达式解析器: cron-engine that calculates execution times
// nDate: count of days to show
//
// Debug 显示未来 N 天内的所有计划执行时间
// 帮助在部署到生产环境前验证任务计划
// p表达式解析器: 计算执行时间的 cron 引擎
// nDate: 要显示的天数
func (cs S定时任务列表) Debug(p表达式解析器 *cronnextzh.P表达式解析器, nDate int) {
	type T任务和时间 struct {
		E任务名称 string
		T执行时间 time.Time
	}
	var s任务列表 []*T任务和时间
	now := time.Now()
	for _, sa := range cs {
		for _, spec := range sa.S定时表达式列表 {
			for _, v执行时间 := range p表达式解析器.Get获取未来N天内的执行时间(spec, now, nDate) {
				s任务列表 = append(s任务列表, &T任务和时间{
					E任务名称: sa.E任务名称,
					T执行时间: v执行时间,
				})
			}
		}
	}

	sortx.SortVStable(s任务列表, func(a, b *T任务和时间) bool {
		return a.T执行时间.Before(b.T执行时间)
	})

	zaplog.SUG.Debugln("显示未来的执行时间")
	zaplog.SUG.Debugln("----------")
	for idx, tsk := range s任务列表 {
		zaplog.SUG.Debugln(
			fmt.Sprintf("%04d", idx),
			timezh.TS.T时间.Get转字符串(tsk.T执行时间),
			int(tsk.T执行时间.Weekday()),
			tsk.E任务名称,
		)
	}
	zaplog.SUG.Debugln("----------")
}
