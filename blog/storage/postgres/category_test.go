package postgres

import (
	"context"
	"grpc-blog/blog/storage"
	"log"
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

func TestShowCategories(t *testing.T) {

	s := newTestStorage(t)

	tests := []struct {
		name    string
		in      storage.Category
		want	storage.Category
		wantErr bool
	}{
		{
			name: "SHOW_CATEGORY_SUCCESS",
			in: storage.Category{},
			want: storage.Category{},
		},
		// {
		// 	name: "FAILED_TO_SHOW_CATEGORY_DOES_NOT_EXIST",
		// 	in: storage.Category{},
		// 	wantErr: true,
		// },
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			got, err := s.Show(context.TODO(), tt.in)
			log.Printf("========: %#v", got)
			if (err != nil) != tt.wantErr {
				t.Errorf("Storage.Show() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Storage.Show() = %v, want %v", got, tt.want)
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

func TestUpdateCategories(t *testing.T) {

	s := newTestStorage(t)

	tests := []struct {
		name    string
		in      storage.Category
		want    *storage.Category
		wantErr bool
	}{
		{
			name: "UPDATE_CATEGORY_SUCCESS",
			in: storage.Category{
				ID:           1,
				CategoryName: "This is category update",
				IsCompleted:  false,
			},
			want: &storage.Category{
				ID:           1,
				CategoryName: "This is category update",
				IsCompleted:  false,
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			got, err := s.Update(context.TODO(), tt.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("Storage.update() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !cmp.Equal(got, tt.want) {
				t.Errorf("Diff: got -, want + = %v", cmp.Diff(got, tt.want))
			}
		})
	}
}

func TestDeleteCategories(t *testing.T) {

	s := newTestStorage(t)

	tests := []struct {
		name    string
		in      int64
		wantErr bool
	}{
		{
			name: "DELETE_CATEGORY_SUCCESS",
			in: 1,
		},
		{
			name: "FAILED_TO_DELETE_CATEGORY_ID_DOES_NOT_EXIST",
			in: 100,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			err := s.Delete(context.TODO(), tt.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("Storage.update() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
