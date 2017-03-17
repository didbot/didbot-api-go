package models

import (
	"gopkg.in/mgutz/dat.v2/sqlx-runner"
	"database/sql"
	"gopkg.in/mgutz/dat.v2/dat"
	"time"
	"os"
)

// global database (pooling provided by SQL driver)
var DB *runner.DB

func LoadDB(){
	// Redis: namespace is the prefix for keys and should be unique
	//store, err := kvs.NewRedisStore("didbot:", ":6379", "")
	//runner.SetCache(store)

	// create a normal database connection through database/sql
	db, err := sql.Open("postgres", os.Getenv("PG_DSN"))
	if err != nil {
		panic(err)
	}

	// ensures the database can be pinged with an exponential backoff (15 min)
	runner.MustPing(db)

	// set to reasonable values for production
	db.SetMaxIdleConns(4)
	db.SetMaxOpenConns(16)

	// set this to enable interpolation
	dat.EnableInterpolation = true

	// set to check things like sessions closing.
	// Should be disabled in production/release builds.
	dat.Strict = false

	// Log any query over 10ms as warnings.
	runner.LogQueriesThreshold = 10 * time.Millisecond

	DB = runner.NewDB(db, "postgres")
}
