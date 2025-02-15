package init

import (
	"context"
	"fmt"

	"WMS/config"

	"github.com/omniful/go_commons/db/sql/postgres"
)

func InitializeDB() {
	db_cluster := postgres.InitializeDBInstance(config.Config, &[]postgres.DBConfig{config.Config})
	ctx := context.Background()
	config.DB = db_cluster.GetMasterDB(ctx)
	fmt.Println("DB connected successfully", config.DB)
	// migrate.MigrateDB(config.DB)
}
