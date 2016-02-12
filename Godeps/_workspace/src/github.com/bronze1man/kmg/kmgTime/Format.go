package kmgTime

import "time"

const (
	FormatMysqlZero  = "0000-00-00 00:00:00"
	FormatMysql      = "2006-01-02 15:04:05"
	FormatMysqlUs    = "2006-01-02 15:04:05.999999"
	FormatFileName   = "2006-01-02_15-04-05" //适合显示在文件上面的日期格式 @deprecated
	FormatFileNameV2 = "2006-01-02-15-04-05" //版本2,更规整,方便使用正则取出
	FormatDateMysql  = "2006-01-02"
	Iso3339Hour      = "2006-01-02T15"
	Iso3339Minute    = "2006-01-02T15:04"
	Iso3339Second    = "2006-01-02T15:04:05"
	AppleJsonFormat  = "2006-01-02 15:04:05 Etc/MST" //仅解决GMT的这个特殊情况.其他不管,如果苹果返回的字符串换时区了就悲剧了

	FormatMysqlMinute       = "2006-01-02 15:04"
	FormatMysqlMouthAndDay  = "01-02"
	FormatMysqlYearAndMoney = "2006-01"
)

var ParseFormatGuessList = []string{
	FormatMysqlZero,
	FormatMysql,
	FormatDateMysql,
	Iso3339Hour,
	Iso3339Minute,
	Iso3339Second,
	time.RFC3339,
	time.RFC3339Nano,
}

var MysqlStart = "0000-01-01 00:00:00"
var MysqlEnd = "9999-12-31 23:59:59"

//输出成mysql的格式,并且使用默认时区,并且在0值的时候输出空字符串
func DefaultFormat(t time.Time) string {
	if t.IsZero() {
		return ""
	}
	return t.In(DefaultTimeZone).Format(FormatMysql)
}

func MonthAndDayFormat(t time.Time) string {
	if t.IsZero() {
		return ""
	}
	return t.In(DefaultTimeZone).Format(FormatMysqlMouthAndDay)
}

func NowWithFileNameFormatV2() string {
	return NowFromDefaultNower().Format(FormatFileNameV2)
}
