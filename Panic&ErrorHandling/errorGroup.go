package main

import (
        "context"
        "fmt"
        "math/rand"
        "time"

        "golang.org/x/sync/errgroup"
)

func fetchAll(ctx context.Context) error {
        errs, ctx := errgroup.WithContext(ctx)

        // run all the http requests in parallel
        for i := 0; i < 4; i++ {
                errs.Go(func() error {
                        // pretend this does an http request and returns an error
                        time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
                        return fmt.Errorf("error in go routine, bailing")
                })
        }

        // Wait for completion and return the first error (if any)
        return errs.Wait()
}

func main() {
        fmt.Println(fetchAll(context.Background()))
}