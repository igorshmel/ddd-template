package domain

import (
	"default_ddd/app/pkg/ddo"
	"default_ddd/app/pkg/errs"
	"default_ddd/app/pkg/logger"
	"github.com/google/uuid"
	"time"
)

// User --
type User struct {
	uuid      string     // Уникальный идентификатор пользователя
	name      string     // Наименование пользователя
	balance   int64      // Баланс денежных средств пользователя
	updatedAt *time.Time // Время изменения записи реестра
	createdAt time.Time  // Время внесения записи в реестр
}

// newUser
func (ths *Domain) newUser() {
	ths.user.uuid = uuid.New().String()
	ths.user.balance = 100000
	ths.userCreatedAt()
	ths.userUpdatedAt()
}

// FillUser --
func (ths *Domain) FillUser() *ddo.UserDDO {
	return &ddo.UserDDO{
		UUID:      ths.user.uuid,
		Name:      ths.user.name,
		Balance:   ths.user.balance,
		UpdatedAt: ths.user.updatedAt,
		CreatedAt: ths.user.createdAt,
	}
}

// CreateUser --
func (ths *Domain) CreateUser(ddo *ddo.UserDDO) *ddo.UserDDO {
	ths.newUser()
	ths.user.name = ddo.Name
	return ths.FillUser()
}

// LoadUser --
func (ths *Domain) LoadUser(req *ddo.UserDDO) {
	ths.user.uuid = req.UUID
	ths.user.name = req.Name
	ths.user.balance = req.Balance
	ths.user.updatedAt = req.UpdatedAt
	ths.user.createdAt = req.CreatedAt
}

// UpdateUser --
func (ths *Domain) reduceUserBalance(l logger.Logger, price uint64) error {
	log := l.WithMethod("ReduceBalance domain")

	if int64(price) > ths.user.balance {
		log.Error("reduce user balance failed. (balance = %d) (totalPrice = %d)", ths.user.balance, price)
		return errs.ErrReduceUserBalanceFailed
	}

	ths.user.balance = ths.user.balance - int64(price)
	ths.userUpdatedAt()
	log.Info("USER_BALANCE: %d\n", ths.user.balance)
	return nil
}

// userUpdatedAt --
func (ths *Domain) userUpdatedAt() {
	t := time.Now()
	ths.user.updatedAt = &t
}

// userCreatedAt --
func (ths *Domain) userCreatedAt() {
	t := time.Now()
	ths.user.createdAt = t
}
