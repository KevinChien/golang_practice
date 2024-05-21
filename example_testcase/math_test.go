package mymath

/*
package <name>, src and test package name must be the same, or go test cannot reference.

By the way, before execute go test, you need execute "go mod init <path/to/module>"
ex:
   go mod init example_testcase
*/

import (
	"os"
	"testing"
)

// TestAdd tests the Add function
func TestAdd(t *testing.T) {
	result := Add(2, 3)
	expected := 5
	if result != expected {
		t.Errorf("Add(2, 3) = %d; want %d", result, expected)
	}

	result = Add(-1, 1)
	expected = 0
	if result != expected {
		t.Errorf("Add(-1, 1) = %d; want %d", result, expected)
	}
}

func TestAdd2(t *testing.T) {
	/*
		under windows, you need :
		$env:SKIP_TEST_ADD2 = "1"
		go test -v ./...

		#remove key
		Remove-Item Env:\SKIP_TEST_ADD2

		under unix, you can :
		$ SKIP_TEST_ADD2=1 go test -v ./...
	*/

	if os.Getenv("SKIP_TEST_ADD2") == "1" {
		t.Skip("Skipping TestAdd2")
	}
	// correct case
	/*
		tests := []struct {
			a, b     int
			expected int
		}{
			{2, 3, 5},
			{-1, 1, 0},
			{0, 0, 0},
			{100, 200, 300},
		}
	*/

	//error case

	tests := []struct {
		a, b     int
		expected int
	}{
		{1, 2, 3},  // 此处预期结果是正确的
		{3, 4, 8},  // 此处预期结果是错误的
		{5, 6, 11}, // 此处预期结果是正确的
	}

	for _, tt := range tests {
		result := Add(tt.a, tt.b)
		if result != tt.expected {
			t.Errorf("Add(%d, %d) = %d; want %d", tt.a, tt.b, result, tt.expected)
		}
	}
}
