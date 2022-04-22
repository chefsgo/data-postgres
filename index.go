package data_postgres

import (
	"github.com/chefsgo/chef"
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
func Driver() chef.DataDriver {
	return &PostgresDriver{}
}

func init() {
	driver := Driver()
	chef.Register("postgres", driver)
	chef.Register("postgressql", driver)
	chef.Register("pgsql", driver)
}
