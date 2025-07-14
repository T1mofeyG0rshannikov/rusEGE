package seo

import (
	"errors"
	"rusEGE/database/models"
	"rusEGE/exceptions"
	"rusEGE/repositories"
	"rusEGE/web/schemas"
)

func CreateIndexSeo(
	sr *repositories.GormSeoRepository,
	data schemas.CreateIndexSeoRequest,
) (*models.IndexSeo, error) {
	existingSeo, err := sr.GetIndexSeo()
	if err != nil && !errors.Is(err, exceptions.ErrIndexSeoNotFound) {
		return nil, err
	}

	if existingSeo != nil {
		return nil, exceptions.ErrIndexSeoAlreadyExists
	}

	seo, err := sr.CreateIndexSeo(&models.IndexSeo{
		Title:    data.Title,
		Image:    data.Image,
		About:    data.About,
		Logo:     data.Logo,
		FipiLink: data.FipiLink,
	})

	if err != nil {
		return nil, err
	}

	return seo, nil
}
