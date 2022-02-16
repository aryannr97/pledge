package pledge

import (
	"testing"
)

func Test_newJWTAuth(t *testing.T) {
	//defer cleanup()
	setup()
	type args struct {
		publisher bool
	}
	tests := []struct {
		name string
		args args
		want *jwtAuth
	}{
		{
			"Positive case",
			args{
				true,
			},
			nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newJWTAuth(tt.args.publisher); got == tt.want {
				t.Errorf("newJWTAuth() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_jwtAuth_GenerateIdentitiy(t *testing.T) {
	defer cleanup()
	setup()
	auth := &jwtAuth{}
	auth.loadKeys(true)
	type args struct {
		data interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"Positive case",
			args{
				data: "data",
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := auth.GenerateIdentitiy(tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("jwtAuth.GenerateIdentitiy() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func Test_jwtAuth_VerifyIdentity(t *testing.T) {
	defer cleanup()
	setup()
	auth := &jwtAuth{}
	auth.loadKeys(true)

	token, _ := auth.GenerateIdentitiy("data")

	type args struct {
		args []interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"Positive case",
			args{
				args: []interface{}{
					token,
				},
			},
			false,
		},
		{
			"Negative case",
			args{
				args: []interface{}{
					string(token.(string)[7:]),
				},
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := auth.VerifyIdentity(tt.args.args...)
			if (err != nil) != tt.wantErr {
				t.Errorf("jwtAuth.VerifyIdentity() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
