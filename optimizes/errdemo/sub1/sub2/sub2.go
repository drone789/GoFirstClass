package sub2

import "github.com/pkg/errors"

func Diff(foo int, bar int) error {
	return errors.New("diff error")
}
