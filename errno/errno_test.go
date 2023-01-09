package errno

import "testing"

func TestIsSuccess(t *testing.T) {
	t.Log(IsSuccess(Success), IsSuccess(nil), IsSuccess(Failure))
}
