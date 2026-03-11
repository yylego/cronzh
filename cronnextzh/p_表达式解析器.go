package cronnextzh

import (
	"time"

	"github.com/robfig/cron/v3"
	"github.com/yylego/rese"
	"github.com/yylego/sortx"
	"github.com/yylego/timezh"
)

// P带秒数的表达式解析器 is a pre-configured cron-engine supporting second-precision (6 fields)
// Format: second, minute, hour, date, month, week-days
//
// P带秒数的表达式解析器 是预配置的 cron 引擎，支持秒级精度（6个字段）
// 格式：秒 分 时 日 月 星期
var P带秒数的表达式解析器 = New(cron.NewParser(cron.Second | cron.Minute | cron.Hour | cron.Dom | cron.Month | cron.Dow | cron.Descriptor))

// P只到分的表达式解析器 is a pre-configured cron-engine supporting minute-precision (5 fields)
// Format: minute, hour, date, month, week-days
//
// P只到分的表达式解析器 是预配置的 cron 引擎，支持分钟级精度（5个字段）
// 格式：分 时 日 月 星期
var P只到分的表达式解析器 = New(cron.NewParser(cron.Minute | cron.Hour | cron.Dom | cron.Month | cron.Dow | cron.Descriptor))

// P表达式解析器 extends cron-engine to provide Chinese-named methods
// which parse cron expressions and calculate execution times
//
// P表达式解析器 封装 cron 引擎，提供中文命名的方法
// 用来解析 cron 表达式和计算未来执行时间
type P表达式解析器 cron.Parser

// New creates a new cron-engine from the given input
// cronEngine: the underlying cron engine instance
// Returns a new P表达式解析器 instance
//
// New 从给定的输入创建新的 cron 引擎
// cronEngine: 底层的 cron 引擎实例
// 返回新的 P表达式解析器 实例
func New(parser cron.Parser) *P表达式解析器 {
	return (*P表达式解析器)(&parser)
}

// Get获取未来N天内的执行时间 calculates execution times within the next N days given a cron expression
// spec: cron expression text (e.g., "0 30 8,20 * * 1-5")
// since: starting time point
// nDate: count of days to check ahead
// Returns a slice of time.Time representing the execution times
//
// Get获取未来N天内的执行时间 计算 cron 表达式在未来 N 天内的执行时间
// spec: cron 表达式字符串（如 "0 30 8,20 * * 1-5"）
// since: 起始时间点
// nDate: 向前查找的天数
// 返回表示所有执行时间的 time.Time 切片
func (P *P表达式解析器) Get获取未来N天内的执行时间(spec string, since time.Time, nDate int) []time.Time {
	p解析器 := (*cron.Parser)(P)
	s时刻表 := rese.V1(p解析器.Parse(spec))
	v执行时间 := since
	e结束时间 := timezh.TS.D日期.Get转字符串(since.AddDate(0, 0, nDate))
	var result []time.Time
	for {
		v执行时间 = s时刻表.Next(v执行时间)
		if timezh.TS.D日期.Get转字符串(v执行时间) > e结束时间 {
			return result
		}
		result = append(result, v执行时间)
	}
}

// Get计算未来N天内的执行时间 calculates execution times within the next N days given multiple cron expressions
// specs: slice of cron expression texts
// since: starting time point
// nDate: count of days to check ahead
// Returns a sorted slice of time.Time representing the execution times from these expressions
//
// Get计算未来N天内的执行时间 计算多个 cron 表达式在未来 N 天内的执行时间
// specs: cron 表达式字符串切片
// since: 起始时间点
// nDate: 向前查找的天数
// 返回排序后的 time.Time 切片，包含所有表达式的所有执行时间
func (P *P表达式解析器) Get计算未来N天内的执行时间(specs []string, since time.Time, nDate int) []time.Time {
	var result []time.Time
	for _, spec := range specs {
		res := P.Get获取未来N天内的执行时间(spec, since, nDate)
		result = append(result, res...)
	}
	sortx.SortVStable(result, func(a, b time.Time) bool {
		return a.Before(b)
	})
	return result
}
