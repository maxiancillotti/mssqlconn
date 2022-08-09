package mssqlconn

type DBConnectorBuilder interface {
	SetHostname(hostname string) DBConnectorBuilder

	SetPort(port int) DBConnectorBuilder

	SetInstance(instance string) DBConnectorBuilder

	SetDatabaseName(dbname string) DBConnectorBuilder

	SetCredentials(user string, password string) DBConnectorBuilder

	EnableDebug() DBConnectorBuilder

	Build() DBConnector
}

type dbConnBuilder struct {
	hostname string
	port     int
	intance  string
	user     string
	password string
	dbname   string
	debug    bool
}

// NewBuiler returns a DBConnecterBuilder that you can configure to
// build a database connector that can open a database connection.
func NewBuilder() DBConnectorBuilder {
	return &dbConnBuilder{
		port:  1433,
		debug: false,
	}
}

func (b *dbConnBuilder) Build() DBConnector {
	return &dbConn{
		dbconfig: b,
	}
}

func (b *dbConnBuilder) SetHostname(hostname string) DBConnectorBuilder {
	b.hostname = hostname
	return b
}

func (b *dbConnBuilder) SetPort(port int) DBConnectorBuilder {
	b.port = port
	return b
}

func (b *dbConnBuilder) SetInstance(instance string) DBConnectorBuilder {
	b.intance = instance
	return b
}

func (b *dbConnBuilder) SetDatabaseName(dbname string) DBConnectorBuilder {
	b.dbname = dbname
	return b
}

func (b *dbConnBuilder) SetCredentials(user, password string) DBConnectorBuilder {
	b.user = user
	b.password = password
	return b
}

func (b *dbConnBuilder) EnableDebug() DBConnectorBuilder {
	b.debug = true
	return b
}
