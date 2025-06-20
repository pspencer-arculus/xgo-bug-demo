package mypackage

import (
	"context"
	"fmt"

	"github.com/pspencer-arculus/xgo-bug-demo/mypackage/myclient"
	"github.com/pspencer-arculus/xgo-bug-demo/tasker"
)

var MyTasker = tasker.Tasker[myclient.Client]{
	Client: myclient.Client{URL: "https://example.com"},
	F: func(ctx context.Context, client myclient.Client) error {
		fmt.Printf("task called: %s", client.URL)
		return nil
	},
}

func DoTask(ctx context.Context) error {
	return MyTasker.Do(ctx)
}
