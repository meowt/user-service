package storage

type Storage struct {
}

type Config struct {
}

func New(cfg *Config) (*Storage, error) {
	return &Storage{}, nil
}
