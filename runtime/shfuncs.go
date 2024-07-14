package runtime

import (
	"context"
	"errors"
	"fmt"
	"io"

	"mvdan.cc/sh/v3/interp"
)

// write writes a string to a module context's stdout (linked in ctx).
func write(ctx context.Context, s string) error {
	mod := interp.HandlerCtx(ctx)
	_, err := io.WriteString(mod.Stdout, s)
	return err
}

func shVoptIf(ctx context.Context, args []string) error {
	var opt string
	v := false
	ifTrue, ifFalse := "", ""

	switch len(args) {
	case 4:
		ifFalse = args[3]
		fallthrough
	case 3:
		opt = args[1]
		ifTrue = args[2]
	default:
		return errors.New("missing argument")
	}

	switch x := ctx.Value("options").(type) {
	case Options:
		var ok bool
		if v, ok = x[opt]; !ok {
			return fmt.Errorf("invalid option: %q", opt)
		}
	}

	if v {
		write(ctx, ifTrue)
	} else {
		write(ctx, ifFalse)
	}
	return nil
}
