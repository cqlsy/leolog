/**
 * Created by angelina-zf on 17/2/25.
 */
package leolog

import (
	"testing"
)

func TestSprint(t *testing.T) {
	//str := Sprint(0, 1)
	//Equal(str, "[leogoDebug] at TestSprint() [debug_test.go:11]\n0\n1\n")
}

func TestPrint(t *testing.T) {
	Print(0, 1)
	SimpleColorPrint("dasd" + "dasda" + "fasfas")
	Sprint("dasda","dfasf")
}
