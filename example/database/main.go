package main

import (
	"fmt"
	"github.com/MR5356/health/database"
)

func main() {
	checker := database.NewChecker(database.DBDriverSQLite, "file::memory:?cache=shared")
	result := checker.Check()
	fmt.Printf("status: %s\nrtt: %dms\nresult: %+v\n", result.Status, result.GetRTT(), result.GetResult())
}
