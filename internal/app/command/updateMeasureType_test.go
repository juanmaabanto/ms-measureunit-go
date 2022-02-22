package command

import "testing"

func Test_NewUpdateMeasureTypeHandler(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("panic")
		}
	}()

	NewUpdateMeasureTypeHandler(nil)
}
