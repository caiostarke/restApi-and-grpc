package database

import "github.com/caiostarke/restApi-and-grpc/app/queries"

type Queries struct {
	*queries.BookQueries
}

func OpenDBConnection() (*Queries, error) {

	db, err := MongoDBConnection()
	if err != nil {
		return nil, err
	}

	return &Queries{
		BookQueries: &queries.BookQueries{DB: db},
	}, nil
}
