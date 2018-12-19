package service

import "fmt"

func (p *TypePSQL) NewConn() *gorm.DB {
	connStr := fmt.Sprintf("sslmode=disable host=%s port=%s dbname=%s user=%s password=%s",
		p.PGHOST, p.PGPORT, p.PGNAME, p.PGUSER, p.PGPASS)

	connection, err := gorm.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}

	err = connection.DB().Ping()
	if err != nil {
		panic(err)
	}

	//p.MakeMigrations(connection)

	return connection
}
