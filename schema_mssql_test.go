package schema_test

import (
	"fmt"

	_ "github.com/denisenkom/go-mssqldb" // mssql
	// _ "github.com/minus5/gofreetds" // mssql

	. "github.com/onsi/ginkgo"
	// . "github.com/onsi/gomega"
)

var _ = XDescribe("schema", func() {
	Context("using github.com/denisenkom/go-mssqldb (Microsoft SQL-Server)", func() {

		const (
			user = "test_user"
			pass = "aNRV!^5-WCe4hz$3"
			host = "localhost"
			port = "32769"
		)

		var mssql = &testParams{
			DriverName: "mssql",
			ConnStr:    fmt.Sprintf("user id=%s;password=%s;server=%s;port=%s", user, pass, host, port),
			// ConnStr: fmt.Sprintf("user id=%s;password=%s;server=%s:%s", user, pass, host, port), // gofreetds

			CreateDDL: []string{`
				CREATE TABLE web_resource (
					id				INTEGER NOT NULL,
					url				NVARCHAR NOT NULL UNIQUE,
					content			VARBINARY,
					compressed_size	INTEGER NOT NULL,
					content_length	INTEGER NOT NULL,
					content_type	NVARCHAR NOT NULL,
					etag			NVARCHAR NOT NULL,
					last_modified	NVARCHAR NOT NULL,
					created_at		DATETIME NOT NULL,
					modified_at		DATETIME,
					PRIMARY KEY (id)
				)`,
				`CREATE INDEX idx_web_resource_url ON web_resource(url)`,
				`CREATE INDEX idx_web_resource_created_at ON web_resource (created_at)`,
				`CREATE INDEX idx_web_resource_modified_at ON web_resource (modified_at)`,
				`CREATE VIEW web_resource_view AS SELECT id, url FROM web_resource`,
			},
			DropDDL: []string{
				`DROP VIEW IF EXISTS web_resource_view`,
				`DROP INDEX IF EXISTS idx_web_resource_modified_at ON web_resource`,
				`DROP INDEX IF EXISTS idx_web_resource_created_at ON web_resource`,
				`DROP INDEX IF EXISTS idx_web_resource_url ON web_resource`,
				`DROP TABLE web_resource`,
			},

			TableExpRes: []string{
				"id INT",
				"url NVARCHAR",
				"content VARBINARY",
				"compressed_size INT",
				"content_length INT",
				"content_type NVARCHAR",
				"etag NVARCHAR",
				"last_modified NVARCHAR",
				"created_at DATETIME",
				"modified_at DATETIME",
			},
			ViewExpRes: []string{
				"id INT",
				"url NVARCHAR",
			},

			TableNameExpRes: "web_resource",
			ViewNameExpRes:  "web_resource_view",
		}

		SchemaTestRunner(mssql)
	})
})
