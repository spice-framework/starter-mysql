//go:build integration

package mysql

import (
	"context"
	"database/sql"
	"errors"
	"net/url"
	"os"
	"strings"
	"testing"
	"time"
)

func TestRealMySQLPoolIsolationCancellationRecoveryAndSecrecy(t *testing.T) {
	connectionURL := os.Getenv("SPICE_MYSQL_TEST_URL")
	if connectionURL == "" {
		t.Fatal("SPICE_MYSQL_TEST_URL is required")
	}
	ctx, cancel := context.WithTimeout(t.Context(), 30*time.Second)
	defer cancel()

	first := openTestDatabase(t, connectionURL, "spice-mysql-first")
	second := openTestDatabase(t, connectionURL, "spice-mysql-second")
	firstClosed := false
	t.Cleanup(func() {
		if !firstClosed {
			closeTestDatabase(t, first)
		}
		closeTestDatabase(t, second)
	})

	if err := Ping(ctx, first); err != nil {
		t.Fatalf("Ping(first) error = %v", err)
	}
	if err := Ping(ctx, second); err != nil {
		t.Fatalf("Ping(second) error = %v", err)
	}
	var one int
	if err := first.QueryRowContext(ctx, "SELECT 1").Scan(&one); err != nil || one != 1 {
		t.Fatalf("SELECT 1 = %d, %v", one, err)
	}

	slowCtx, stopSlow := context.WithTimeout(ctx, 100*time.Millisecond)
	started := time.Now()
	var slept int
	err := first.QueryRowContext(slowCtx, "SELECT SLEEP(5)").Scan(&slept)
	stopSlow()
	if !errors.Is(err, context.DeadlineExceeded) {
		t.Fatalf("canceled query error = %v", err)
	}
	if elapsed := time.Since(started); elapsed > 2*time.Second {
		t.Fatalf("canceled query took %s", elapsed)
	}
	if err := Ping(ctx, first); err != nil {
		t.Fatalf("Ping(first after cancellation) error = %v", err)
	}

	if err := first.Close(); err != nil {
		t.Fatalf("Close(first) error = %v", err)
	}
	firstClosed = true
	if err := Ping(ctx, first); err == nil {
		t.Fatal("Ping(closed first) unexpectedly succeeded")
	}
	if err := Ping(ctx, second); err != nil {
		t.Fatalf("Ping(second after first closed) error = %v", err)
	}

	assertAuthenticationFailureRedactsSecret(t, ctx, connectionURL)
}

func openTestDatabase(t *testing.T, connectionURL, applicationName string) *sql.DB {
	t.Helper()
	database, err := Open(Options{
		URL:             connectionURL,
		ApplicationName: applicationName,
		AllowInsecure:   true,
	})
	if err != nil {
		t.Fatalf("Open() error = %v", err)
	}
	return database
}

func closeTestDatabase(t *testing.T, database *sql.DB) {
	t.Helper()
	if err := database.Close(); err != nil {
		t.Errorf("Close() error = %v", err)
	}
}

func assertAuthenticationFailureRedactsSecret(
	t *testing.T,
	ctx context.Context,
	connectionURL string,
) {
	t.Helper()
	parsed, err := url.Parse(connectionURL)
	if err != nil {
		t.Fatalf("parse test URL: %v", err)
	}
	const secret = "standalone-mysql-secret-must-not-leak"
	parsed.User = url.UserPassword(parsed.User.Username(), secret)
	badURL := parsed.String()
	database := openTestDatabase(t, badURL, "spice-mysql-secrecy")
	defer closeTestDatabase(t, database)
	err = Ping(ctx, database)
	if err == nil {
		t.Fatal("Ping() with invalid password unexpectedly succeeded")
	}
	if strings.Contains(err.Error(), secret) || strings.Contains(err.Error(), badURL) {
		t.Fatalf("authentication error exposed credentials: %v", err)
	}
}
