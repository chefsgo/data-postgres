package data_postgres

import (
	"github.com/chefsgo/data"
	_ "github.com/lib/pq" //此包自动注册名为postgres的sql驱动
)

var (
	DRIVERS = []string{
		"postgresql", "postgres", "pgsql", "pgdb", "pg",
		"cockroachdb", "cockroach", "crdb",
		"timescaledb", "timescale", "tsdb",
	}
)

//返回驱动
func Driver() data.Driver {
	return &PostgresDriver{}
}

func init() {
	driver := Driver()
	data.Register("postgres", driver)
	data.Register("postgressql", driver)
	data.Register("pgsql", driver)
}
