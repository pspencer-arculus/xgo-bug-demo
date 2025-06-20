package mypackage_test

import (
	"context"
	"testing"

	"github.com/pspencer-arculus/xgo-bug-demo/mypackage"
	"github.com/xhd2015/xgo/runtime/mock"
)

func TestDo(t *testing.T) {
	mock.Patch(mypackage.DoTask, func(ctx context.Context) error {
		t.Log("patched DoTask")
		return nil
	})
	err := mypackage.DoTask(context.Background())
	if err != nil {
		t.Errorf("mypackage failed: %s", err)
	}
}
