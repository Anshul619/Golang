package main

import (
    "log"
    "net/http"
    "myservice/internal/api"
    "myservice/internal/service"
    "myservice/internal/repository"
    "myservice/internal/model"
    "context"
)

// Fake repo for demo
type fakeRepo struct{}

func (f *fakeRepo) FindByID(_ context.Context, id string) (*model.User, error) {
    return &model.User{ID: id, Name: "John Doe", Email: "john@example.com"}, nil
}

func main() {
    repo := &fakeRepo{}
    svc := service.NewUserService(repo)
    handler := api.NewUserHandler(svc)

    http.HandleFunc("/user", handler.GetUser)
    log.Println("Listening on :8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
