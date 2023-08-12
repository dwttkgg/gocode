package main

import (
	"context"
	"fmt"
)

func process(ctx context.Context) {
	fmt.Println("trace_id is :", ctx.Value("trace_id").(string))
	fmt.Println("session is :", ctx.Value("session").(string))
}
func main() {
	ctx := context.WithValue(context.Background(), "trace_id", "ac13121649893165")
	ctx = context.WithValue(ctx, "session", "dasfsafasdfsdfwewrwerwerds-dfasdfksajf")
	process(ctx)
}
