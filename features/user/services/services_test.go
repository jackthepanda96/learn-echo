package services_test

import (
	"21-api/features/user"
	"21-api/features/user/services"
	"21-api/helper"
	"21-api/mocks"
	"testing"

	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func TestRegister(t *testing.T) {
	model := mocks.NewUserModel(t)
	pm := mocks.NewPasswordManager(t)
	md := mocks.NewJwtInterface(t)
	rPm := helper.NewPasswordManager()
	srv := services.NewService(model, pm, md)
	registerData := user.User{Nama: "jerry", Hp: "081234567890", Password: "alta1234"}
	hashedPassword, _ := rPm.HashPassword(registerData.Password)
	insertData := user.User{Nama: "jerry", Hp: "081234567890", Password: hashedPassword}
	t.Run("Success register", func(t *testing.T) {
		pm.On("HashPassword", registerData.Password).Return(hashedPassword, nil).Once()
		insertData.Password = hashedPassword
		model.On("InsertUser", insertData).Return(nil).Once()

		err := srv.Register(registerData)

		pm.AssertExpectations(t)
		model.AssertExpectations(t)

		assert.Nil(t, err)
	})

	t.Run("Error hash password", func(t *testing.T) {
		pm.On("HashPassword", registerData.Password).Return("", bcrypt.ErrHashTooShort).Once()

		err := srv.Register(registerData)

		pm.AssertExpectations(t)

		assert.Error(t, err)
	})

	t.Run("Error from model", func(t *testing.T) {
		pm.On("HashPassword", registerData.Password).Return(hashedPassword, nil).Once()
		insertData.Password = hashedPassword
		model.On("InsertUser", insertData).Return(gorm.ErrInvalidData).Once()

		err := srv.Register(registerData)

		pm.AssertExpectations(t)
		model.AssertExpectations(t)

		assert.Error(t, err)
		assert.EqualError(t, err, helper.ServerGeneralError)
	})

	t.Run("Error validation", func(t *testing.T) {
		registerData.Password = "alta123"
		err := srv.Register(registerData)

		assert.Error(t, err)
	})
}

func TestLogin(t *testing.T) {
	model := mocks.NewUserModel(t)
	pm := mocks.NewPasswordManager(t)
	md := mocks.NewJwtInterface(t)
	srv := services.NewService(model, pm, md)
	loginData := user.User{Hp: "081234567890", Password: "alta1234"}
	loginDBData := user.User{Hp: "081234567890", Password: "hashedPassword"}
	t.Run("Success Login", func(t *testing.T) {
		model.On("Login", loginData.Hp).Return(loginDBData, nil).Once()
		pm.On("ComparePassword", loginData.Password, "hashedPassword").Return(nil).Once()
		md.On("GenerateJWT", loginDBData.Hp).Return("resultToken", nil).Once()

		userData, token, err := srv.Login(loginData)

		model.AssertExpectations(t)
		pm.AssertExpectations(t)
		md.AssertExpectations(t)

		assert.NotEmpty(t, userData)
		assert.NotEmpty(t, token)
		assert.Nil(t, err)
	})

	t.Run("Gagal Login - generate jwt error", func(t *testing.T) {
		model.On("Login", loginData.Hp).Return(loginDBData, nil).Once()
		pm.On("ComparePassword", loginData.Password, "hashedPassword").Return(nil).Once()
		md.On("GenerateJWT", loginDBData.Hp).Return("", jwt.ErrTokenMalformed).Once()

		userData, token, err := srv.Login(loginData)

		model.AssertExpectations(t)
		pm.AssertExpectations(t)

		assert.Empty(t, userData)
		assert.Empty(t, token)
		assert.Error(t, err)
	})

	t.Run("Gagal Login - hp not found", func(t *testing.T) {
		model.On("Login", loginData.Hp).Return(user.User{}, gorm.ErrRecordNotFound).Once()

		userData, token, err := srv.Login(loginData)

		model.AssertExpectations(t)
		pm.AssertExpectations(t)

		assert.Empty(t, userData)
		assert.Empty(t, token)
		assert.Error(t, err)
	})

	t.Run("Gagal Login - salah password", func(t *testing.T) {
		model.On("Login", loginData.Hp).Return(loginDBData, nil).Once()
		pm.On("ComparePassword", loginData.Password, "hashedPassword").Return(bcrypt.ErrMismatchedHashAndPassword).Once()

		userData, token, err := srv.Login(loginData)

		model.AssertExpectations(t)
		pm.AssertExpectations(t)

		assert.Empty(t, userData)
		assert.Empty(t, token)
		assert.Error(t, err)
	})

	t.Run("Gagal Login - validasi", func(t *testing.T) {
		loginData.Password = "alta123"
		userData, token, err := srv.Login(loginData)

		assert.Empty(t, userData)
		assert.Empty(t, token)
		assert.Error(t, err)
	})
}
