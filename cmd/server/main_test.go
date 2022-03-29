package main

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/murat/testing-practices/mocks"
)

func TestHandler_ServeHTTP(t *testing.T) {
	type args struct {
		path   string
		method string
	}
	tests := []struct {
		name           string
		args           args
		wantResp       string
		wantRespStatus int
	}{
		{"happy path 200", args{"/ping", http.MethodGet}, "pong", http.StatusOK},
		{"happy path 404", args{"/ping", http.MethodPost}, "", http.StatusNotFound},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := httptest.NewRequest(tt.args.method, tt.args.path, strings.NewReader(tt.wantResp))
			w := httptest.NewRecorder()

			h := &Handler{}
			h.ServeHTTP(w, r)

			if w.Code != tt.wantRespStatus {
				t.Errorf("Want status '%d', got '%d'", tt.wantRespStatus, w.Code)
			}

			if w.Body.String() != tt.wantResp {
				t.Errorf("Want '%s', got '%s'", tt.wantResp, w.Body)
			}
		})
	}
}

func TestUserHandler_ServeHTTP(t *testing.T) {
	mockedDatabase := new(mocks.CRUD)

	type args struct {
		path   string
		method string
	}
	tests := []struct {
		name           string
		args           args
		wantResp       string
		wantRespStatus int
		wantErr        bool
	}{
		{"list users", args{"/users", http.MethodGet}, "tests", http.StatusOK, false},
		{"could not list users", args{"/users", http.MethodGet}, "error occurred", http.StatusInternalServerError, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.wantErr {
				mockedDatabase.On("All").Return(nil, errors.New("error occurred")).Once()
			}

			req := httptest.NewRequest(tt.args.method, tt.args.path, strings.NewReader(tt.wantResp))
			rec := httptest.NewRecorder()

			h := &UserHandler{}
			h.ServeHTTP(rec, req)

			if rec.Code != tt.wantRespStatus {
				t.Errorf("Want status '%d', got '%d'", tt.wantRespStatus, rec.Code)
			}

			if strings.TrimSpace(rec.Body.String()) != tt.wantResp {
				t.Errorf("Want '%s', got '%s'", tt.wantResp, rec.Body)
			}
		})
	}
}
