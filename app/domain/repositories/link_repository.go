package repositories

import (
	"context"

	"github.com/dwdarm/shortify/app/domain/models"
)

type LinkRepository interface {
	FindBySlug(ctx context.Context, slug string) (*models.Link, error)
	Save(ctx context.Context, link *models.Link) (*models.Link, error)
}
