package seo

import (
	"rusEGE/database/models"
	"rusEGE/repositories"
	"rusEGE/web/schemas"
)

func EditIndexSeo(
	sr *repositories.GormSeoRepository,
	data schemas.EditIndexSeoRequest,
) (*models.IndexSeo, error) {
	seo, err := sr.GetIndexSeo()
	if err != nil {
		return nil, err
	}

	if data.About != nil {
		seo.About = *data.About
	}
	if data.Title != nil {
		seo.Title = *data.Title
	}
	if data.Logo != nil {
		seo.Logo = *data.Logo
	}
	if data.Image != nil {
		seo.Image = *data.Image
	}
	if data.FipiLink != nil {
		seo.FipiLink = *data.FipiLink
	}

	seo, err = sr.EditIndexSeo(seo)

	if err != nil {
		return nil, err
	}

	return seo, nil
}
