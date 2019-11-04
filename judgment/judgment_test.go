package judgment

import (
	"testing"
	"time"
)

func TestTimeBase(t *testing.T) {
	testMin := 5 * time.Second
	testMax := 7 * time.Second
	duration := 6 * time.Second
	if !TimeBase(testMin, testMax, duration) {
		t.Fatal("結果が間違っています")
	}
}

func TestDisplayBase(t *testing.T) {
	display := "aaaaaaaaaaaaaaaaa<s>zzzz</s>aaaaaaaaaaaaaaaaaaa<s>f</s>aaaaaaaaaaaa"
	incl := "<s>f</s>"
	if !DisplayBase(display, incl) {
		t.Fatal("結果が間違っています")
	}
}
