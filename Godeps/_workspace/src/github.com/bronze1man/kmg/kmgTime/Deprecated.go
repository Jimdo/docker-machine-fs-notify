package kmgTime

import "time"

//@deprecated
var DefaultNower tDefaultNower

//@deprecated
func ParseAutoInLocal(sTime string) (t time.Time, err error) {
	return ParseAutoInLocation(sTime, time.Local)
}

//@deprecated
func MustParseAutoInLocal(sTime string) (t time.Time) {
	t, err := ParseAutoInLocation(sTime, time.Local)
	if err != nil {
		panic(err)
	}
	return t
}
