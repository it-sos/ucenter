package caches

import (
	"testing"
	_ "ucenter/s/tests/testsimple"
)

func TestTimesCache(t *testing.T) {
	NAuthTimes.Key("peng.yu").Incr()
	NAuthTimes.Key("peng.yu").Incr()
	NAuthTimes.Key("peng.yu").Incr()
	t.Log(NAuthTimes.Key("peng.yu").Incr())
	t.Log(NAuthTimes.Key("peng.yu").Get())
	NAuthTimes.Key("peng.yu").Clear()
	t.Log(NAuthTimes.Key("peng.yu").Decr())
}
