package jwtutil

import (
	"testing"

	"github.com/dgrijalva/jwt-go"
)

const (
	id          = "AAAABBBBBCCCCCDDDDEEEEFFFFGGGGAAAABBBBBCCCCCDDDDEEEEFFFFGGGGAAAABBBBBCCCCCDDDDEEEEFFFFGGGGAAAABBBBBCCCCCDDDDEEEEFFFFGGGGAAAABBBBBCCCCCDDDDEEEEFFFFGGGGAAAABBBBBCCCCCDDDDEEEEFFFFGGGGAAAABBBBBCCCCCDDDDEEEEFFFFGGGGAAAABBBBBCCCCCDDDDEEEEFFFFGGGGAAAABBBBBCCCCCDDDDEEEEFFFFGGGGAAAABBBBBCCCCCDDDDEEEEFFFFGGGG"
	issuedAt    = 1593350443
	expiresAt   = 1593360443
	signedToken = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1OTMzNjA0NDMsImp0aSI6IkFBQUFCQkJCQkNDQ0NDREREREVFRUVGRkZGR0dHR0FBQUFCQkJCQkNDQ0NDREREREVFRUVGRkZGR0dHR0FBQUFCQkJCQkNDQ0NDREREREVFRUVGRkZGR0dHR0FBQUFCQkJCQkNDQ0NDREREREVFRUVGRkZGR0dHR0FBQUFCQkJCQkNDQ0NDREREREVFRUVGRkZGR0dHR0FBQUFCQkJCQkNDQ0NDREREREVFRUVGRkZGR0dHR0FBQUFCQkJCQkNDQ0NDREREREVFRUVGRkZGR0dHR0FBQUFCQkJCQkNDQ0NDREREREVFRUVGRkZGR0dHR0FBQUFCQkJCQkNDQ0NDREREREVFRUVGRkZGR0dHR0FBQUFCQkJCQkNDQ0NDREREREVFRUVGRkZGR0dHRyIsImlhdCI6MTU5MzM1MDQ0M30.wIJgsg-sNXCDd_ax_H27vH1yeh6-QMw9E4-ji_WtNgE"
)

var secret = func() []byte { return []byte("TEST") }

func TestNews(t *testing.T) {
	type args struct {
		claims jwt.StandardClaims
	}
	Secret = secret
	tests := []struct {
		name            string
		args            args
		wantSignedToken string
		wantErr         bool
	}{
		{"100000 times", args{jwt.StandardClaims{Id: id, IssuedAt: issuedAt, ExpiresAt: expiresAt}}, signedToken, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for i := 0; i < 100000; i++ {
				gotSignedToken, err := New(tt.args.claims)
				if (err != nil) != tt.wantErr {
					t.Errorf("New() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				if gotSignedToken != tt.wantSignedToken {
					t.Errorf("New() = %v, want %v", gotSignedToken, tt.wantSignedToken)
				}
			}
		})
	}
}

func TestParses(t *testing.T) {
	type args struct {
		signedToken string
		claimsOut   jwt.StandardClaims
	}
	Secret = secret
	tests := []struct {
		name      string
		args      args
		wantToken *jwt.Token
		wantErr   bool
	}{
		{"100000 times", args{signedToken: signedToken}, nil, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for i := 0; i < 100000; i++ {
				_, err := Parse(tt.args.signedToken, &(tt.args.claimsOut))
				if (err != nil) != tt.wantErr {
					t.Errorf("Parse() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				if tt.args.claimsOut.Id != id || tt.args.claimsOut.IssuedAt != issuedAt || tt.args.claimsOut.ExpiresAt != expiresAt {
					t.Errorf("Parse() error = %v", tt.args.claimsOut)
					return
				}
			}
		})
	}
}
