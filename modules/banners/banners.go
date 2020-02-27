package banners

import (
	"fmt"
	"github.com/ArtemYeremeev/go-repository-pattern/interfaces"
	"github.com/ArtemYeremeev/go-repository-pattern/models"
)

// BannerRepository describes structure of banner repository
type BannerRepository struct {
	interfaces.IDbHandler
}

// GetAllBanners gets all banners
func (repository *BannerRepository) GetAllBanners() (interface{}) {
	row, err := repository.Query(fmt.Sprintf("select * from t_banners"))
	if err != nil {
		return nil
	}

	var banners []models.Banner
	if row.Next() {
		var b models.Banner
		if err := row.Scan(&b.ID, &b.Header, &b.Description); err != nil {
			return nil
		}
		banners = append(banners, b)
	}

	return banners
}

// GetBannerByID gets banner by received ID
func (repository *BannerRepository) GetBannerByID(id int) (interface{}, error) {
	row, err := repository.Query(fmt.Sprintf("select * from t_banners where name = %d", id))
	if err != nil {
		return nil, err
	}

	var b models.Banner
	row.Next()
	if err := row.Scan(&b.ID, &b.Header, &b.Description); err != nil {
		return nil, err
	}

	return b, nil
}

// GetBannerByName gets banner by received name
func (repository *BannerRepository) GetBannerByName(name string) (interface{}, error) {
	row, err := repository.Query(fmt.Sprintf("select * from t_banners where name = '%s'", name))
	if err != nil {
		return nil, err
	}

	var b models.Banner
	row.Next()
	if err := row.Scan(&b.ID, &b.Header, &b.Description); err != nil {
		return nil, err
	}

	return b, nil
}
