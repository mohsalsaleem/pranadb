package schema

import (
	"strings"
	"testing"

	"github.com/squareup/pranadb/conf"
	"github.com/squareup/pranadb/kafka"

	"github.com/squareup/pranadb/notifier"

	"github.com/squareup/pranadb/cluster"
	"github.com/squareup/pranadb/command"
	"github.com/squareup/pranadb/common"
	"github.com/squareup/pranadb/meta"
	"github.com/squareup/pranadb/pull"
	"github.com/squareup/pranadb/push"
	"github.com/squareup/pranadb/sharder"
	"github.com/stretchr/testify/require"
)

func TestLoader(t *testing.T) {
	type ddl struct {
		schema  string
		queries []string
	}
	tests := []struct {
		name string
		ddl  []ddl
	}{
		{
			name: "sources",
			ddl: []ddl{{
				schema: "location",
				queries: []string{`create source location(id bigint, x varchar, y varchar, primary key (id) )
					with (
						brokername = "testbroker",
						topicname = "testtopic",
		              headerencoding = "json",
						keyencoding = "json",
						valueencoding = "json",
						columnselectors = (
							"k.k0",
							"v.v1",
							"v.v2"
						)
						properties = (
							"prop1" = "val1",
							"prop2" = "val2"
						)
					)
					`},
			}, {
				schema: "hollywood",
				queries: []string{
					`create source actor(id bigint, name varchar, age int, primary key (id) )
					with (
						brokername = "testbroker",
						topicname = "testtopic",
		              headerencoding = "json",
						keyencoding = "json",
						valueencoding = "json",
						columnselectors = (
							"k.k0",
							"v.v1",
							"v.v2"
						)
						properties = (
							"prop1" = "val1",
							"prop2" = "val2"
						)
					)
				`,
					`create source movies(id bigint, title varchar, director varchar, year int, primary key (id))
					with (
						brokername = "testbroker",
						topicname = "testtopic",
		              headerencoding = "json",
						keyencoding = "json",
						valueencoding = "json",
						columnselectors = (
							"k.k0",
							"v.v1",
							"v.v2",
							"v.v3"
						)
						properties = (
							"prop1" = "val1",
							"prop2" = "val2"
						)
					)`,
				},
			}},
		},
		{
			name: "mvs",
			ddl: []ddl{{
				schema: "hollywood",
				queries: []string{
					`create source movies(id bigint, title varchar, director varchar, year int, primary key (id))
                     with (
						brokername = "testbroker",
						topicname = "testtopic",
                        headerencoding = "json",
						keyencoding = "json",
						valueencoding = "json",
						columnselectors = (
							"k.k0",
							"v.v1",
							"v.v2",
							"v.v3"
						)
						properties = (
							"prop1" = "val1",
							"prop2" = "val2"
						)
					)`,
					`create materialized view latest_movies as
						select director, max(year)
						from movies
						group by director`,
				},
			}},
		},
	}
	for _, test := range tests {
		// nolint: scopelint
		t.Run(test.name, func(t *testing.T) {
			clus := cluster.NewFakeCluster(1, 1)
			notifier := notifier.NewFakeNotifier()
			metaController, executor := runServer(t, clus, notifier)
			expectedSchemas := make(map[string]*common.Schema)
			for _, ddl := range test.ddl {
				numTables := 0
				session := executor.CreateSession()
				schema := metaController.GetOrCreateSchema(ddl.schema)
				session.UseSchema(schema)
				for _, query := range ddl.queries {
					_, err := executor.ExecuteSQLStatement(session, query)
					if strings.Contains(query, "create source") {
						numTables++
					} else if strings.Contains(query, "create materialized view") {
						numTables += 2
					}
					require.NoError(t, err)
				}
				require.Equal(t, schema.LenTables(), numTables)
				expectedSchemas[ddl.schema] = schema
			}

			// Restart the server
			_ = clus.Stop()
			metaController, executor = runServer(t, clus, notifier)
			_, ok := metaController.GetSchema("test")
			require.False(t, ok)

			loader := NewLoader(metaController, executor.GetPushEngine(), executor.GetPullEngine(), 0)
			require.NoError(t, loader.Start())

			for _, schema := range test.ddl {
				expected := expectedSchemas[schema.schema]
				actual, ok := metaController.GetSchema(schema.schema)
				require.True(t, ok)
				require.Equal(t, expected, actual)
			}
		})
	}
}

func runServer(t *testing.T, clus cluster.Cluster, notif *notifier.FakeNotifier) (*meta.Controller, *command.Executor) {
	t.Helper()
	fakeKafka := kafka.NewFakeKafka()
	_, err := fakeKafka.CreateTopic("testtopic", 10)
	require.NoError(t, err)
	metaController := meta.NewController(clus)
	shardr := sharder.NewSharder(clus)
	pullEngine := pull.NewPullEngine(clus, metaController)
	config := conf.NewTestConfig(fakeKafka.ID)
	pushEngine := push.NewPushEngine(clus, shardr, metaController, config, pullEngine)
	ce := command.NewCommandExecutor(metaController, pushEngine, pullEngine, clus, notif)
	notif.RegisterNotificationListener(notifier.NotificationTypeDDLStatement, ce)
	notif.RegisterNotificationListener(notifier.NotificationTypeCloseSession, pullEngine)
	clus.SetRemoteQueryExecutionCallback(pullEngine)
	clus.RegisterShardListenerFactory(pushEngine)
	err = clus.Start()
	require.NoError(t, err)
	err = metaController.Start()
	require.NoError(t, err)
	err = pushEngine.Start()
	require.NoError(t, err)
	err = pullEngine.Start()
	require.NoError(t, err)
	err = shardr.Start()
	require.NoError(t, err)

	return metaController, ce
}
