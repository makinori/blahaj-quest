package data

import (
	"database/sql"
	"errors"

	"git.ran.cafe/maki/foxlib/foxdb"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/sqlitedialect"
	"github.com/uptrace/bun/driver/sqliteshim"
)

var (
	Database *bun.DB
)

func Init() error {
	sqldb, err := sql.Open(sqliteshim.ShimName, "file:database.db")
	if err != nil {
		return errors.New("failed to open database: " + err.Error())
	}

	Database = bun.NewDB(sqldb, sqlitedialect.New())

	return foxdb.InitDataCache(Database, "cache", []foxdb.DataCache{
		Blahaj, GitHubStars,
	})
}
