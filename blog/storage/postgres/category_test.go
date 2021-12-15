package postgres

import (
	"context"
	"grpc-blog/blog/storage"
	"testing"
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