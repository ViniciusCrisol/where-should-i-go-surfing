package test

import "time"

func init() {
	time.Local = time.UTC
}
