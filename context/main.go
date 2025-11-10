package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)

	defer cancel()

	go worker(ctx, "Worker-A")
	go worker(ctx, "Worker-B")
	go worker(ctx, "Worker-C")

	time.Sleep(3 * time.Second)

	fmt.Println("main: finalizado")

}

func worker(ctx context.Context, id string) {
	ticker := time.NewTicker(500 * time.Millisecond)

	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			fmt.Printf("[%s] cancelado: %v\n", id, ctx.Err())
			return
		case <-ticker.C:
			fmt.Printf("[%s] trabalhando...\n", id)
		}
	}
}
