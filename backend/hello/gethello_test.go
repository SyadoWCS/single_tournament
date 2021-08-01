package hello

import (
	"testing"
)

func TestHello(t *testing.T) {
	result := GetHello("山澤さん")
	expext := "こんにちは、山澤さん！!"
	if result != expext {
		t.Error("\n実際： ", result, "\n理想： ", expext)
	}

	t.Log("TestHello終了")
}
