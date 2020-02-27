package controllers

import (
	"encoding/json"
	"github.com/ArtemYeremeev/go-repository-pattern/interfaces"
	"net/http"
)

// BannerController descrybes structure of banner responses
type BannerController struct {
	interfaces.IBannerRepository
}

// ReadAllBanners wraps data for all banners response
func (controller *BannerController) ReadAllBanners(res http.ResponseWriter) {
	json.NewEncoder(res).Encode(controller.GetAllBanners)
}
