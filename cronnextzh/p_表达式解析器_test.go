package cronnextzh_test

import (
	"testing"
	"time"

	"github.com/gorhill/cronexpr"
	"github.com/stretchr/testify/require"
	"github.com/yylego/cronzh/cronnextzh"
	"github.com/yylego/timezh"
)

func TestP表达式解析器_Get获取未来N天内的执行时间(t *testing.T) {
	times, err := cronnextzh.P带秒数的表达式解析器.Get获取未来N天内的执行时间("30 25/10 9 * * 1-5", time.Now(), 15)
	require.NoError(t, err)
	require.NotEmpty(t, times)
	for i, v := range times {
		t.Logf("%02d %s %d", i, timezh.TS.T时间.Get转字符串(v), int(v.Weekday()))
	}
}

func TestP表达式解析器_Get获取未来N天内的执行时间_周末(t *testing.T) {
	times, err := cronnextzh.P带秒数的表达式解析器.Get获取未来N天内的执行时间("0 5 0,12 * * 0/6", time.Now(), 15)
	require.NoError(t, err)
	require.NotEmpty(t, times)
	for i, v := range times {
		t.Logf("%02d %s %d", i, timezh.TS.T时间.Get转字符串(v), int(v.Weekday()))
	}
}

func TestP表达式解析器_Get获取未来N天内的执行时间_错误表达式(t *testing.T) {
	_, err := cronnextzh.P带秒数的表达式解析器.Get获取未来N天内的执行时间("bad expression", time.Now(), 15)
	require.Error(t, err)
	t.Log("expected:", err)
}

func TestP表达式解析器_Get计算未来N天内的执行时间(t *testing.T) {
	specs := []string{
		"30 25/10 9 * * 1-5",
		"30 5/10 10,13,14 * * 1-5",
		"30 5-35/10 11 * * 1-5",
		"30 3 15 * * 1-5",
	}
	times, err := cronnextzh.P带秒数的表达式解析器.Get计算未来N天内的执行时间(specs, time.Now(), 15)
	require.NoError(t, err)
	require.NotEmpty(t, times)
	// Validate the times are sorted
	for i := 1; i < len(times); i++ {
		require.False(t, times[i].Before(times[i-1]), "times should be sorted")
	}
	for i, v := range times {
		t.Logf("%02d %s %d", i, timezh.TS.T时间.Get转字符串(v), int(v.Weekday()))
	}
}

func TestP表达式解析器_Get计算未来N天内的执行时间_错误表达式(t *testing.T) {
	specs := []string{"30 25/10 9 * * 1-5", "bad expression"}
	_, err := cronnextzh.P带秒数的表达式解析器.Get计算未来N天内的执行时间(specs, time.Now(), 15)
	require.Error(t, err)
	t.Log("expected:", err)
}

func TestX其它工具的对照结果(t *testing.T) {
	times := cronexpr.MustParse("* * 10 * * * *").NextN(time.Now(), 5)
	require.Len(t, times, 5)
	for _, v := range times {
		t.Log(v)
	}
}
