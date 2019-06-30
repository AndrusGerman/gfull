package gfull

import (
	"testing"
)

func TestSnakeCamelC(t *testing.T) {
	type args struct {
		valor string
	}
	type testModel struct {
		name      string
		args      args
		wantTotal string
	}
	tests := []testModel{
		testModel{name: "Name", args: args{valor: "user_model"}, wantTotal: "UserModel"},
		testModel{name: "Name", args: args{valor: "UserModel"}, wantTotal: "UserModel"},
		testModel{name: "Name", args: args{valor: "country_user_id"}, wantTotal: "CountryUserId"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotTotal := SnakeCamelC(tt.args.valor); gotTotal != tt.wantTotal {
				t.Errorf("SnakeCamelC() = %v, want %v", gotTotal, tt.wantTotal)
			}
		})
	}
}
