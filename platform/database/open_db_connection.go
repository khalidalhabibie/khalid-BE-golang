package database

import (
	fakesRepo "gokes/app/fakes/repository/postgres"
	userRepo "gokes/app/user/repository/postgres"
)

// Queries struct for collect all app queries.
type Repository struct {
	*userRepo.UserRepository
	*fakesRepo.FakesRepository
}

// OpenDBConnection func for opening database connection.
func OpenDBConnection() (*Repository, error) {
	// Define a new PostgreSQL connection.
	db, err := PostgreSQLConnection()
	if err != nil {
		return nil, err
	}

	return &Repository{
		// Set queries from models:
		UserRepository:  &userRepo.UserRepository{DB: db},
		FakesRepository: &fakesRepo.FakesRepository{DB: db},
	}, nil
}
