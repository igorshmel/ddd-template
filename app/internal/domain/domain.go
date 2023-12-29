package domain

import (
	"default_ddd/app/internal/adapters/port"
	"default_ddd/app/pkg/logger"
)

// Domain --
type Domain struct {
	log       logger.Logger
	cartItem  CartItem
	cartItems []CartItem
	order     Order
	user      User
	product   Product
}

// NewDefaultDomain - инициализация домена
func NewDefaultDomain(log logger.Logger) *Domain {
	log = log.WithMethod("domain")
	return &Domain{log: log}
}

// DomConfiguration --
type DomConfiguration func(dr *Ports) error

// Ports --
type Ports struct {
	CartItem port.CartItemPort
	Order    port.OrderPort
	Product  port.ProductPort
	User     port.UserPort
}

// GetCartItemPort --
func (ths *Ports) GetCartItemPort() port.CartItemPort {
	return ths.CartItem
}

// GetOrderPort --
func (ths *Ports) GetOrderPort() port.OrderPort {
	return ths.Order
}

// GetUserPort --
func (ths *Ports) GetUserPort() port.UserPort {
	return ths.User
}

// GetProductPort --
func (ths *Ports) GetProductPort() port.ProductPort {
	return ths.Product
}

// NewPorts --
func NewPorts(configs ...DomConfiguration) (*Ports, error) {
	ports := &Ports{}

	// Apply all Configurations passed in
	for _, cfg := range configs {
		// Pass the repository into the configuration function
		err := cfg(ports)
		if err != nil {
			return nil, err
		}
	}

	return ports, nil
}

// WithDefaultDomain --
func WithDefaultDomain(log logger.Logger) DomConfiguration {
	return func(ths *Ports) error {
		dr := NewDefaultDomain(log)
		ths.User = dr
		ths.Product = dr
		ths.Order = dr
		ths.CartItem = dr
		return nil
	}
}
