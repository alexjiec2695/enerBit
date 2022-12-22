package enerbit

import (
	"errors"
	"fmt"
	"rest_app/domain/model/entities/entities"
	"rest_app/domain/model/entities/gateways"
)

type UseCase struct {
	gateway gateways.Gateway
}

func NewUseCase(gateway gateways.Gateway) *UseCase {
	return &UseCase{gateway}
}

func (c *UseCase) Create(data entities.EnerBitEntities) error {
	existSerialBrand, err := c.gateway.ExistSerialBrand(data.Serial, data.Brand)
	if err != nil {
		return fmt.Errorf("consulting database %s", err.Error())
	}

	if existSerialBrand {
		return errors.New("business error, serial or brand exists")
	}

	existMeterInProperty, err := c.gateway.ExistMeterInProperty(data.Address)

	if err != nil {
		return fmt.Errorf("consulting database %s", err.Error())
	}

	if existMeterInProperty {
		return errors.New("business error, serial or brand exists")
	}

	return c.gateway.Create(data)
}

func (c *UseCase) Update(data entities.EnerBitEntities) error {

	return c.gateway.Update(data)
}

func (c *UseCase) Delete(id string) error {

	isActive, err := c.gateway.IsActive(id)

	if err != nil {
		return err
	}

	if isActive {
		return errors.New("business error, the meter is active")
	}

	if err = c.gateway.Delete(id); err != nil {
		return err
	}
	return nil
}

func (c *UseCase) GetDisabledMeters() (*[]entities.EnerBitEntities, error) {
	return c.gateway.GetDisabledMeters()
}

func (c *UseCase) Filter(data entities.EnerBitEntities) (*entities.EnerBitEntities, error) {
	return c.gateway.Filter(data.Serial, data.Brand)
}
