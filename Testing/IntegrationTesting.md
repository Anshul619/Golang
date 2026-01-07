# Integration Testing
- Test how components work together — includes external systems (DB, API, file system, etc.)
- Use `test containers` or `dockertest` for DBs in integration tests.

# Same dockertest 

````go
package dockertest_test

import (
    "database/sql"
    "fmt"
    "log"
    "testing"
    _ "github.com/go-sql-driver/mysql"
    "github.com/ory/dockertest/v3"
)

var db *sql.DB

func TestMain(m *testing.M) {
    pool, err := dockertest.NewPool("")
    if err != nil {
        log.Fatalf("Could not construct pool: %s", err)
    }

    err = pool.Client.Ping()
    if err != nil {
        log.Fatalf("Could not connect to Docker: %s", err)
    }

    resource, err := pool.Run("mysql", "5.7", []string{"MYSQL_ROOT_PASSWORD=secret"})
    if err != nil {
        log.Fatalf("Could not start resource: %s", err)
    }

    if err := pool.Retry(func() error {
        var err error
        db, err = sql.Open("mysql", fmt.Sprintf("root:secret@(localhost:%s)/mysql", resource.GetPort("3306/tcp")))
        if err != nil {
            return err
        }
        return db.Ping()
    }); err != nil {
        log.Fatalf("Could not connect to database: %s", err)
    }

    defer func() {
        if err := pool.Purge(resource); err != nil {
            log.Fatalf("Could not purge resource: %s", err)
        }
    }()

    m.Run()
}

func TestSomething(t *testing.T) {
    // db.Query() can now be used safely
}   
````

# Characteristics

| Title                                                   |
|---------------------------------------------------------|
| Slower                                                  |
| May be non-deterministic (e.g. network latency)         |
| Real dependencies (Postgres, Redis, Kafka etc.)         |
| Often lives in `/test/` or `_integration_test.go` files |
