package httpboomer

import (
	"testing"
)

func TestHttpBoomer(t *testing.T) {
	testcase1 := &TestCase{
		Config: TConfig{
			Name:   "TestCase1",
			Weight: 2,
		},
	}
	testcase2 := &TestCase{
		Config: TConfig{
			Name:   "TestCase2",
			Weight: 3,
		},
	}

	Run(testcase1, testcase2)
}