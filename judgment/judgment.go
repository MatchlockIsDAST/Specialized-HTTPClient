package judgment

import (
	"strings"
	"time"
)

//TimeBase 経過時間の時間の判定を行う
func TimeBase(elapsedMin, elapsedMax, duration time.Duration) bool {
	return (elapsedMin <= duration && duration <= elapsedMax)
}

//DisplayBase 得られた文字列に該当の文字列が含まれているか判定を行う
func DisplayBase(display, included string) bool {
	return strings.Contains(display, included)
}

//Diffbase 取得したbodyの照合を行う
func DiffBase(responsbodys []string) bool {
	var flag bool
	for i := 1; i < len(responsbodys); i++ {
		if responsbodys[0] == responsbodys[1] {
			flag = true
		}
	}
	return flag
}
