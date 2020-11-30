package player

import (
	"reflect"
	"testing"
)

func TestNewService(t *testing.T) {
	type args struct {
		r Repository
	}
	tests := []struct {
		name string
		args args
		want Service
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewService(tt.args.r); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_service_Login(t *testing.T) {
	type fields struct {
		r Repository
	}
	type args struct {
		player Player
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
		want1  string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &service{
				r: tt.fields.r,
			}
			got, got1 := s.Login(tt.args.player)
			if got != tt.want {
				t.Errorf("Login() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Login() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_service_SignUp(t *testing.T) {
	type fields struct {
		r Repository
	}
	type args struct {
		player Player
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &service{
				r: tt.fields.r,
			}
			if got := s.SignUp(tt.args.player); got != tt.want {
				t.Errorf("SignUp() = %v, want %v", got, tt.want)
			}
		})
	}
}
