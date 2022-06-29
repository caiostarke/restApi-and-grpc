package caching

import "github.com/caiostarke/restApi-and-grpc/app/queries"

type Queries struct {
	*queries.CachingQueries
}

func OpenRedisConnection() (*Queries, error) {
	client, err := RedisConnection()
	if err != nil {
		return nil, err
	}

	return &Queries{
		&queries.CachingQueries{Redis: client},
	}, nil
}
