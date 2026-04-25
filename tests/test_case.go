package tests

import (
	"github.com/goravel/framework/testing"

	"dakhl/bootstrap"
)

func init() {
	bootstrap.Boot()
}

type TestCase struct {
	testing.TestCase
}
