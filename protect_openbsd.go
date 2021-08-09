// +build openbsd

package hello_assert_bot

import "golang.org/x/sys/unix"

// Pledge on OpenBSD lets us "promise" to only run a subset of
// system calls: http://man.openbsd.org/pledge
func Pledge(s string) {
    must(unix.Pledge(s, ""))
}

func Unveil(path string, flags string) {
    must(unix.Unveil(path, flags))
}

func UnveilBlock() {
    must(unix.UnveilBlock())
}

func must(err error) {
    if err != nil {
        panic(err)
    }
}
