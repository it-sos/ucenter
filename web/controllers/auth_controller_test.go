package controllers

import (
	"reflect"
	"testing"
	"ucenter/models/vo"
	"ucenter/repositories"
	_ "ucenter/s/tests/testsimple"
	"ucenter/services"
)

func TestAuthController_PostLogin(t *testing.T) {
	type fields struct {
		Service services.AuthService
	}
	var tests = []struct {
		name    string
		fields  fields
		want    vo.AuthTokenVO
		wantErr bool
	}{
		{
			"test",
			fields{services.NewAuthService(services.NewUserService(repositories.NewUserRepository()))},
			vo.AuthTokenVO{},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &AuthController{
				Service: tt.fields.Service,
			}
			got, err := c.PostLogin()
			if (err != nil) != tt.wantErr {
				t.Errorf("PostLogin() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PostLogin() got = %v, want %v", got, tt.want)
			}
		})
	}
}
