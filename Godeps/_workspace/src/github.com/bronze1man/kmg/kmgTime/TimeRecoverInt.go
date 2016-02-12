package kmgTime

import (
	"math"
	"time"
)

//类似于体力一样的,每隔一定时间就会恢复一点(加一点)的值,
//没有负值,不会超过最大值
//如果执行时间向前退,不会报错
type TimeRecoverInt struct {
	Num             int
	Max             int
	LastRecoverTime time.Time
	AddDuration     time.Duration
}

func (t *TimeRecoverInt) Sync(now time.Time) {
	timeDuring := now.Sub(t.LastRecoverTime)
	staminaTimes := float64(timeDuring) / float64(t.AddDuration)
	addStamina := int(math.Floor(staminaTimes))
	//添加体力后,体力满了
	if (addStamina + t.Num) >= t.Max {
		t.LastRecoverTime = now
		t.Num = t.Max
		return
	}
	//可能是数据错误,重置
	if addStamina < 0 {
		t.LastRecoverTime = now
		return
	}
	//没有添加体力
	if addStamina == 0 {
		return
	}
	t.Num = addStamina + t.Num
	t.LastRecoverTime = t.LastRecoverTime.Add(t.AddDuration * time.Duration(addStamina))
	return
}

func (t *TimeRecoverInt) Full(now time.Time) {
	t.Num = t.Max
	t.LastRecoverTime = now
}
