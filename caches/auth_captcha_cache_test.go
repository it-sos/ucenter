package caches

import (
	"testing"
	_ "ucenter/s/tests/testsimple"
)

func TestCaptchaCache(t *testing.T) {
	NAuthCaptcha.Key("peng.yu").Set("hello world.")
	t.Log(NAuthCaptcha.Key("peng.yu").Get())
}
