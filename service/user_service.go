package service

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/luqxus/spaces/storage"
	"github.com/luqxus/spaces/tokens"
	"github.com/luqxus/spaces/types"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	storage storage.Storage
}

func NewUserService(storage storage.Storage) *UserService {
	return &UserService{
		storage: storage,
	}
}

func (s *UserService) CreateUser(ctx context.Context, reqData *types.RegisterReqData) (string, error) {
	user := new(types.User)

	count, err := s.storage.CountEmail(ctx, reqData.Email)
	if err != nil {
		return "", err
	}

	if count > 0 {
		return "", errors.New("email already in use")
	}

	user.UID = uuid.NewString()
	user.Email = reqData.Email
	user.Username = reqData.Username
	user.Password = hash(reqData.Password)

	err = s.storage.CreateUser(ctx, user)
	if err != nil {
		return "", err
	}

	return tokens.GenerateJwt(user.Username, user.UID)
}

func (s *UserService) Login(
	ctx context.Context,
	reqData *types.LoginReqData) (*types.ResponseUser, string, error) {
	user, err := s.storage.GetUserByEmail(ctx, reqData.Email)
	if err != nil {
		return nil, "", err
	}

	if err := verifyPassword(user.Password, reqData.Password); err != nil {
		return nil, "", errors.New("wrong email or password")
	}

	token, _ := tokens.GenerateJwt(user.Username, user.UID)

	return user.ResponseUser(), token, nil
}

func hash(plaintext string) string {
	b, _ := bcrypt.GenerateFromPassword([]byte(plaintext), 14)
	return string(b)
}

func verifyPassword(hashedtext, plaintext string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedtext), []byte(plaintext))
}
