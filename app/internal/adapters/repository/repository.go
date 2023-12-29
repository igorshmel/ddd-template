package repository

import (
	"default_ddd/app/internal/adapters/port"
	"default_ddd/app/internal/adapters/repository/postgres"
	"default_ddd/app/pkg/config"
	"default_ddd/app/pkg/logger"
)

// RepoConfiguration is an alias for a function that will take in a pointer to an Repository and modify it
type RepoConfiguration func(r *Repository) error

// Repository is an implementation of the Repository
type Repository struct {
	persister port.Persister
	extractor port.Extractor
}

// GetPersister getting the persistence Repository
func (ths *Repository) GetPersister() port.Persister {
	return ths.persister
}

// GetExtractor getting the extractor Repository
func (ths *Repository) GetExtractor() port.Extractor {
	return ths.extractor
}

// NewRepository takes a variable amount of RepoConfiguration functions and returns a new Repository
// Each RepoConfiguration will be called in the order they are passed in
func NewRepository(configs ...RepoConfiguration) (*Repository, error) {
	repo := &Repository{}

	// Apply all Configurations passed in
	for _, cfg := range configs {
		// Pass the repository into the configuration function
		err := cfg(repo)
		if err != nil {
			return nil, err
		}
	}

	return repo, nil
}

// WithPostgresRepository applies a postgres document repository to the Repository
func WithPostgresRepository(cfg config.Config, log logger.Logger) RepoConfiguration {
	return func(ths *Repository) error {
		pr, err := postgres.NewPostgresRepository(cfg, log, true)
		if err != nil {
			return err
		}

		ths.persister = pr
		ths.extractor = pr
		return nil
	}
}
