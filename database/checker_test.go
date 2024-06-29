package database

import (
	"github.com/MR5356/health"
	"testing"
)

func TestNewChecker(t *testing.T) {
	dbDriverType := DBDriverSQLite
	dsn := "file::memory:?cache=shared"
	checker := NewChecker(dbDriverType, dsn)
	if checker.dbDriverType != dbDriverType {
		t.Errorf("Expected dbDriverType %s, got %s", dbDriverType, checker.dbDriverType)
	}
	if checker.dbDsn != dsn {
		t.Errorf("Expected dsn %s, got %s", dsn, checker.dbDsn)
	}
}

func TestChecker_Check_Success(t *testing.T) {
	t.Run("test-sqlite", func(t *testing.T) {
		dbDriverType := DBDriverSQLite
		dsn := "file::memory:?cache=shared"
		checker := NewChecker(dbDriverType, dsn)

		result := checker.Check()
		if result.Status != health.StatusUp {
			t.Errorf("Expected status %s, got %s", health.StatusUp, result.Status)
		}
		if !result.HasResult() {
			t.Errorf("Expected to have result, but got %v", result.GetResult())
		}

		// from dbPool
		checker = NewChecker(dbDriverType, dsn)
		result = checker.Check()
		if result.Status != health.StatusUp {
			t.Errorf("Expected status %s, got %s", health.StatusUp, result.Status)
		}
		if !result.HasResult() {
			t.Errorf("Expected to have result, but got %v", result.GetResult())
		}
	})

	// TODO: add postgresql
	// TODO: add mysql
}

func TestChecker_Check_Failure(t *testing.T) {
	t.Run("test-mysql", func(t *testing.T) {
		dbDriverType := DBDriverMySQL
		dsn := "root:password@tcp(127.0.0.1:3306)/test"
		checker := NewChecker(dbDriverType, dsn)

		result := checker.Check()
		if result.Status != health.StatusDown {
			t.Errorf("Expected status %s, got %s", health.StatusDown, result.Status)
		}
		if !result.HasResult() {
			t.Errorf("Expected to have result, but got %v", result.GetResult())
		}
	})
	t.Run("test-postgresql", func(t *testing.T) {
		dbDriverType := DBDriverPostgreSQL
		dsn := "postgres://postgres:password@localhost:5432/test"
		checker := NewChecker(dbDriverType, dsn)

		result := checker.Check()
		if result.Status != health.StatusDown {
			t.Errorf("Expected status %s, got %s", health.StatusDown, result.Status)
		}
		if !result.HasResult() {
			t.Errorf("Expected to have result, but got %v", result.GetResult())
		}
	})
}
