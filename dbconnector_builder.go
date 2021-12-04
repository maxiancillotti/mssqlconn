package mssqlconn

type DBConnecterBuilder interface {
	SetServer(server string) DBConnecterBuilder

	SetPort(port int) DBConnecterBuilder

	SetDatabaseName(dbname string) DBConnecterBuilder

	SetCredentials(user string, password string) DBConnecterBuilder

	Build() DBConnecter
}

type dbConnBuilder struct {
	server   string
	port     int
	user     string
	password string
	dbname   string
	debug    bool
}

// NewBuiler returns a DBConnecterBuilder that you can configure to
// build a database connector that can open a database connection.
func NewBuilder() DBConnecterBuilder {
	return &dbConnBuilder{
		port:  1433,
		debug: false,
	}
}

func (b *dbConnBuilder) Build() DBConnecter {
	return &dbConn{
		dbconfig: b,
	}
}

func (b *dbConnBuilder) SetServer(server string) DBConnecterBuilder {
	b.server = server
	return b
}

func (b *dbConnBuilder) SetPort(port int) DBConnecterBuilder {
	b.port = port
	return b
}

func (b *dbConnBuilder) SetDatabaseName(dbname string) DBConnecterBuilder {
	b.dbname = dbname
	return b
}

func (b *dbConnBuilder) SetCredentials(user, password string) DBConnecterBuilder {
	b.user = user
	b.password = password
	return b
}
