// +build !openbsd

package hello_assert_bot

// Pledge on any other system than OpenBSD doesn't do anything
func Pledge(s string) {
    return
}

func Unveil(path string, flags string) {
}

func UnveilBlock() {
}
