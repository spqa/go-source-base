package media

import (
	"fmt"
	"mcm-api/config"
)

type ImageProxyService interface {
	GetLink(key string) string
}

type DarthsimImageProxyService struct {
	cfg *config.Config
}

func NewDarthsimImageProxyService(cfg *config.Config) ImageProxyService {
	return &DarthsimImageProxyService{cfg: cfg}
}

func (d DarthsimImageProxyService) GetLink(key string) string {
	return fmt.Sprintf("%v/a/fill/{width}/{height}/sm/0/plain/%v", d.cfg.ImageProxyService, key)
}
