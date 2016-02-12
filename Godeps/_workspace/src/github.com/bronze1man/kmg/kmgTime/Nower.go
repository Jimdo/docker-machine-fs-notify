package kmgTime

import "time"

type Nower interface {
	Now() time.Time
}

type tDefaultNower struct{}

type FixedNower struct {
	Time time.Time
}

var NowTime Nower = tDefaultNower{}

func GetDefaultNower() Nower {
	return NowTime
}

func NowFromDefaultNower() time.Time {
	return NowTime.Now()
}

func MysqlNowFromDefaultNower() string {
	return NowTime.Now().Format(FormatMysql)
}

func MysqlUsNowFromDefaultNower() string {
	return NowTime.Now().Format(FormatMysqlUs)
}

func NewFixedNower(time time.Time) Nower {
	return FixedNower{time}
}

func SetFixNowFromString(s string) {
	NowTime = NewFixedNower(MustParseAutoInDefault(s))
}

func SetNowerToDefault() {
	NowTime = tDefaultNower{}
}

func (nower tDefaultNower) Now() time.Time {
	return time.Now().In(DefaultTimeZone)
}

func (nower FixedNower) Now() time.Time {
	return nower.Time
}
