package singleton

import (
	"testing"
)

// TestGetInstance kiểm tra xem hàm gọi GetInstace trả về gì
func TestGetInstance(t *testing.T) {
	counter1 := GetInstance() 

	if counter1 == nil {
		t.Error("muốn có 1 địa chỉ pointer trỏ đến Singleton sau khi gọi GetInstance(), chứ không phải nil")
	} 
	expectedCounter := counter1 

	currentCount := counter1.AddOne()
	if currentCount != 1 {
		t.Errorf("After calling for the first time to count, the count must be 1 but it is %d\n", currentCount)
	}

	counter2 := GetInstance()
	if counter2 != expectedCounter {
		//Test 2 failed
		t.Error("Expected same instance in counter2 but it got a different instance")
	}

	currentCount = counter2.AddOne()
	if currentCount != 2 {
		t.Errorf("After calling 'AddOne' using the second counter, the current count must be 2 but was %d\n", currentCount)
	}


}