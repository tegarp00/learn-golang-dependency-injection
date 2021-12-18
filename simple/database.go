package simple

type Database struct {
	Name string
}

type DatabasePostgreSQL Database
type DatabaseMongoDB Database

func NewDatabaseMongoDB() *DatabaseMongoDB {
	return (*DatabaseMongoDB)(&DatabaseMongoDB{Name: "MongoDB"})
}

func NewDatabasePostgreSQL() *DatabasePostgreSQL {
	return (*DatabasePostgreSQL)(&DatabasePostgreSQL{Name: "PostgreSQL"})

}

type DatabaseRepository struct {
	DatabasePostgreSQL *DatabasePostgreSQL
	DatabaseMongoDB    *DatabaseMongoDB
}

// func NewDatabaseRepository(postgreSQL *DatabasePostgreSQL, mongodb *DatabaseMongoDB) *DatabaseRepository {
// 	return &DatabaseRepository{DatabasePostgreSQL: postgreSQL, DatabaseMongoDB: mongodb}
// }

func NewDatabaseRepository(mongodb *DatabaseMongoDB, postgreSQL *DatabasePostgreSQL) *DatabaseRepository {
	return &DatabaseRepository{DatabaseMongoDB: mongodb, DatabasePostgreSQL: postgreSQL}
}
