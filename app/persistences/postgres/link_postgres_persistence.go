package persistences

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/dwdarm/shortify/app/domain/models"
	"github.com/dwdarm/shortify/app/domain/repositories"
)

type LinkPostgresPersistence struct {
	db *sql.DB
}

func NewLinkPostgresPersistence(db *sql.DB) repositories.LinkRepository {
	return &LinkPostgresPersistence{
		db: db,
	}
}

func (p *LinkPostgresPersistence) FindBySlug(ctx context.Context, slug string) (*models.Link, error) {
	row := p.db.QueryRowContext(ctx, `SELECT id, slug, href, qr_code FROM link WHERE slug = $1;`, slug)

	link := models.Link{}
	err := row.Scan(&link.Id, &link.Slug, &link.Href, &link.QrCode)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}

	return &link, nil
}

func (p *LinkPostgresPersistence) Save(ctx context.Context, link *models.Link) (*models.Link, error) {
	if link.Id != "" {
		_, err := p.db.ExecContext(
			ctx,
			`UPDATE link SET slug='$1', href='$2', qr_code='$3' WHERE slug = $4;`,
			link.Slug, link.Href, link.QrCode, link.Id,
		)

		if err != nil {
			return nil, err
		}

		return link, nil

	} else {
		result, err := p.db.ExecContext(
			ctx,
			`INSERT INTO link(slug, href, qr_code) VALUES($1, $2, $3);`,
			link.Slug, link.Href, link.QrCode,
		)

		if err != nil {
			return nil, err
		}

		lastId, _ := result.LastInsertId()

		return &models.Link{
			Id:     fmt.Sprint(lastId),
			Href:   link.Href,
			Slug:   link.Slug,
			QrCode: link.QrCode,
		}, nil
	}
}
