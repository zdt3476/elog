package elog

import "fmt"

type LogLevel uint8

const (
	DebugLvl LogLevel = iota
	InfoLvl
	WarnLvl
	ErrorLvl
	PanicLvl // panic
	FatalLvl // os.Exit()
)

const (
	DebugStr = "DEBUG"
	InfoStr  = "INFO"
	WarnStr  = "WARN"
	ErrorStr = "ERROR"
	PanicStr = "PANIC"
	FatalStr = "FATAL"
)

var (
	lvl2StrMap = map[LogLevel]string{
		DebugLvl: DebugStr,
		InfoLvl:  InfoStr,
		WarnLvl:  WarnStr,
		ErrorLvl: ErrorStr,
		PanicLvl: PanicStr,
		FatalLvl: FatalStr,
	}

	str2LvlMap = map[string]LogLevel{
		DebugStr: DebugLvl,
		InfoStr:  InfoLvl,
		WarnStr:  WarnLvl,
		ErrorStr: ErrorLvl,
		PanicStr: PanicLvl,
		FatalStr: FatalLvl,
	}
)

// 配置的时候使用
func Str2LogLvl(v string) (LogLevel, bool) {
	lvl, found := str2LvlMap[v]
	return lvl, found
}

func (l LogLevel) String() string {
	lvl, found := lvl2StrMap[l]
	if found {
		return lvl
	}

	return fmt.Sprintf("Level(%d)", l)
}
