package common

func Pointer[C any](v C) *C {
	return &v
}
