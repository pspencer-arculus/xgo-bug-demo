package mypackage

import (
	"context"
	"fmt"
	"io"
	"os"
)

type Client struct {
	url string
}

type Tasker[T any] struct {
	Client T
	F      func(context.Context, T) error
}

var MyTasker = Tasker[io.Writer]{
	Client: os.Stdout,
	F: func(ctx context.Context, w io.Writer) error {
		_, err := fmt.Fprintln(w, "task called")
		return err
	},
}

func DoTask(ctx context.Context) error {
	return MyTasker.F(ctx, MyTasker.Client)
}

func (t Tasker[T]) Do(ctx context.Context) error {
	return t.F(ctx, t.Client)
}
