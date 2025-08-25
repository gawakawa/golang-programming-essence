package skip

import (
	"runtime"
	"testing"
)

func TestReadData(t *testing.T) {
	if runtime.GOOS != "windows" {
		t.Skipf("skipping on %v", runtime.GOOS)
	}
}
