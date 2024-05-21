package dir2

import "github.com/sgg/gomod-subdir-repro/dir1"

func SayMore() string {
	return dir1.Say()
}
