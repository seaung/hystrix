package user

import (
	"context"

	"github.com/seaung/hystrix/internal/hystrix/store"
	v1 "github.com/seaung/hystrix/pkg/api/hystrix/v1"
	"github.com/seaung/hystrix/pkg/token"
)

type UserBiz interface {
	ChangePassword(ctx context.Context, username string, r *v1.ChangePasswordForm) error
	Login(ctx context.Context, r *v1.LoginFormRequest) (*v1.LoginTokenResponse, error)
}

type userBiz struct {
	ds store.IsStore
}

var _ UserBiz = (*userBiz)(nil)

func New(ds store.IsStore) *userBiz {
	return &userBiz{ds: ds}
}

func (u *userBiz) ChangePassword(ctx context.Context, username string, r *v1.ChangePasswordForm) error {
	userModel, err := u.ds.Users().GetUser(ctx, username)
	if err != nil {
		return err
	}

	if err := auth.Compare(userModel.Password, r.OldPassword); err != nil {
		return errno.PasswordIncorrectError
	}

	userModel.Password, _ = auth.Encrypt(r.NewPassword)
	if err := u.ds.Users().UpdateUser(ctx, userModel); err != nil {
		return err
	}

	return nil
}

func (u *userBiz) Login(ctx context.Context, r *v1.LoginFormRequest) (*v1.LoginTokenResponse, error) {
	user, err := u.ds.Users().GetUser(ctx, r.Username)
	if err != nil {
		return nil, errno.UserNotFoundError

	}

	if err := auth.Compare(user.Password, r.Password); err != nil {
		return nil, PasswordIncorrectError
	}

	tokenString, err := token.SignToken(r.Username)
	if err != nil {
		return errno.SignTokenError
	}

	return &v1.LoginTokenResponse{Token: tokenString}, nil
}
