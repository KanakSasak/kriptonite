package core

import (
	"Kryptonite/internal/core/domain"
	"Kryptonite/internal/core/ports"
	"Kryptonite/internal/core/services"
)

type Container struct {
	CryptService ports.Crypt
}

type ContainerConfig struct {
}

func NewContainer(c *ContainerConfig) (*Container, error) {
	data := domain.Crypt{Data: ""}

	cryptSvc := services.NewCryptService(data)

	return &Container{
		CryptService: cryptSvc,
	}, nil
}
