package main

// go test -tags=integration ./test/...

func TestUserRepo_SaveAndGet(t *testing.T) {
    db := sql.Open("postgres", "...")
    repo := NewPostgresUserRepo(db)

    err := repo.Save(User{ID: "123", Name: "Alice"})
    if err != nil {
        t.Fatal(err)
    }

    user, err := repo.FindByID("123")
    if err != nil || user.Name != "Alice" {
        t.Fatal("failed to retrieve user")
    }
}