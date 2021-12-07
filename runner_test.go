package hrp

import (
	"testing"
)

func TestHttpRunner(t *testing.T) {
	testcase1 := &TestCase{
		Config: TConfig{
			Name:    "TestCase1",
			BaseURL: "http://httpbin.org",
		},
		TestSteps: []IStep{
			NewStep("headers").
				GET("/headers").
				Validate().
				AssertEqual("status_code", 200, "check status code").
				AssertEqual("headers.\"Content-Type\"", "application/json", "check http response Content-Type"),
			NewStep("user-agent").
				GET("/user-agent").
				Validate().
				AssertEqual("status_code", 200, "check status code").
				AssertEqual("headers.\"Content-Type\"", "application/json", "check http response Content-Type"),
			NewStep("TestCase3").CallRefCase(&TestCase{Config: TConfig{Name: "TestCase3"}}),
		},
	}
	testcase2 := &TestCase{
		Config: TConfig{
			Name:   "TestCase2",
			Weight: 3,
		},
	}
	testcase3 := &TestCasePath{demoTestCaseJSONPath}

	err := NewRunner(t).Run(testcase1, testcase2, testcase3)
	if err != nil {
		t.Fatalf("run testcase error: %v", err)
	}
}
