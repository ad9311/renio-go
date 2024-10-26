package svc

import "github.com/ad9311/renio-go/internal/model"

func SignOutUser(allowedJWT model.AllowedJWT) error {
	if err := allowedJWT.Delete(); err != nil {
		return err
	}

	return nil
}
