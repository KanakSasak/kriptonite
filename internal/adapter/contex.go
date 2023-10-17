package adapter

import (
	"Kryptonite/internal/adapter/controller"
	"Kryptonite/internal/core"
	"Kryptonite/internal/core/ports"
	"fmt"
)

type HandlerContext struct {
	CryptController ports.Crypt
}

func createContainer() *core.Container {
	containerConf := &core.ContainerConfig{}

	cont, err := core.NewContainer(containerConf)

	if err != nil {
		fmt.Println("failed container", err)
		panic("failed to create container")
	}

	return cont
}

func NewContext() *HandlerContext {
	container := createContainer()

	cryptCtrl := controller.NewCryptController(container)

	return &HandlerContext{
		CryptController: cryptCtrl,
	}
}
