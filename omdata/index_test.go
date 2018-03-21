package wafdata

import "testing"

func TestCreateIdAllocator(t *testing.T) {
	IdAllocator := CreateIdAllocator("test", 1)
	if IdAllocator == nil {
		t.Errorf("CreateIdAllocator failed.")
	}
}
