package migrations

import (
	"context"
	"fmt"
	"github.com/scylladb/gocqlx/v2"
	"github.com/scylladb/gocqlx/v2/migrate"
)

func InitCassandra(session *gocqlx.Session) error {
	metadata, err := session.KeyspaceMetadata("novabankapp")
	if err != nil {
		return err
	}
	fmt.Println(metadata.Name)
	log := func(ctx context.Context, session gocqlx.Session, ev migrate.CallbackEvent, name string) error {
		return nil
	}

	reg := migrate.CallbackRegister{}
	reg.Add(migrate.BeforeMigration, "cassandra_migrate.cql", log)
	reg.Add(migrate.AfterMigration, "cassandra_migrate.cql", log)

	migrate.Callback = reg.Callback
	ctx := context.Background()
	er := migrate.FromFS(ctx, *session, Files)

	list, err := migrate.List(ctx, *session)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	for i := range list {
		fmt.Println(fmt.Sprintf("%s migrated", list[i].Name))
	}

	return er

}
