package postgres

import (
	"context"
	"grpc-blog/blog/storage"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestCreateCategory(t *testing.T) {

	s := newTestStorage(t)

	tests := []struct {
		name    string
		in      storage.Category
		want    int64
		wantErr bool
	}{
		{
			name: "CREATE_CATEGORY_SUCCESS",
			in: storage.Category{
				CategoryName: "This is category",
			},
			want: 1,
		},
		{
			name: "FAILED_DUPLICATE_CATEGORY",
			in: storage.Category{
				CategoryName: "This is category",
			},
			//want: 2,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			got, err := s.Create(context.TODO(), tt.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("Storage.Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Storage.Create() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetCategory(t *testing.T) {

	s := newTestStorage(t)

	tests := []struct {
		name    string
		in      int64
		want    *storage.Category
		wantErr bool
	}{
		{
			name: "GET_CATEGORY_SUCCESS",
			in: 1,
			want: &storage.Category{
				ID:           1,
				CategoryName: "This is category",
				IsCompleted:  false,
			},
		},
		{
			name: "FAILED_CATEGORY_DOES_NOT_EXIST",
			in: 100,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			got, err := s.Get(context.TODO(), tt.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("Storage.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !cmp.Equal(got, tt.want) {
				t.Errorf("Diff = %v", cmp.Diff(got, tt.want))
			}
		})
	}
}
