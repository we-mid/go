package util

import "errors"

func IsErrorLike(left, right error) bool {
	return left == nil && right == nil ||
		left != nil && right != nil && left.Error() == right.Error() ||
		errors.Is(left, right)
}
