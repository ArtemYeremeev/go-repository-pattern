package interfaces

import (
	"go-repository-pattern/models"
)

// IBannerRepository descrybes structure of banners repository
type IBannerRepository interface {
	GetAllBanners() (interface{}, error)
	GetBannerByID(id int) (models.Banner, error)
	GetBannerByHeader(header string) (models.Banner, error)
}
