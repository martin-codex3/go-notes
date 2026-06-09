package database

// we will create the database pooling here
func Connection(DatabaseUrl string) (*pgxpool, error)