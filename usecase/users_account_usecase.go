package usecase

import (
	"EEP/e-wallets/model"
	"EEP/e-wallets/model/dto/req"
	"EEP/e-wallets/repository"
	"EEP/e-wallets/utils/common"
	"EEP/e-wallets/utils/security"
	"fmt"
	"github.com/go-playground/validator/v10"
	"time"
)

type UsersAccountUseCase interface {
	Register(payload req.RegisterRequest) error
	Login()
	ForgotPassword()
	ChangePassword()
}

type usersAccountUseCase struct {
	repo      repository.UsersAccountRepository
	walletsUC WalletsUseCase
}

func (u *usersAccountUseCase) Register(payload req.RegisterRequest) error {
	// Validate the payload
	validate := validator.New()
	err := validate.Struct(payload)
	if err != nil {
		return err
	}

	hashPassword, err := security.HashPassword(payload.Password)
	if err != nil {
		return err
	}

	usersAccount := model.UsersAccount{
		Id:            common.GenerateID(),
		UserName:      payload.UserName,
		Password:      hashPassword,
		Email:         payload.Identifier.Email,
		PhoneNumber:   payload.Identifier.PhoneNumber,
		AccountStatus: model.Active,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Time{},
	}

	err = u.repo.Save(usersAccount)
	if err != nil {
		return fmt.Errorf("failed save user: %v", err.Error())
	}

	wallet := model.Wallets{
		Id:           common.GenerateID(),
		UserId:       usersAccount.Id,
		RekeningUser: common.GenerateRandomRekeningNumber(10),
		Balance:      0,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Time{},
	}

	err = u.walletsUC.CreateWallet(wallet)
	if err != nil {
		return fmt.Errorf("failed create wallet: %v", err.Error())
	}

	return nil

}

func (u *usersAccountUseCase) Login() {
	//TODO implement me
	panic("implement me")
}

func (u *usersAccountUseCase) ForgotPassword() {
	//TODO implement me
	panic("implement me")
}

func (u *usersAccountUseCase) ChangePassword() {
	//TODO implement me
	panic("implement me")
}

func NewUsersAccountUseCase(repo repository.UsersAccountRepository, walletsUC WalletsUseCase) UsersAccountUseCase {
	return &usersAccountUseCase{
		repo:      repo,
		walletsUC: walletsUC,
	}
}
