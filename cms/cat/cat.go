package cat

import (
	"context"
	"io"
	"log"

	"google.golang.org/grpc"

	protoCategory "grpc-category/proto/category"
)

type Client struct {
	client	protoCategory.CategoryServiceClient
}

func NewClient(conn grpc.ClientConnInterface) Client {
	return Client{
		client: protoCategory.NewCategoryServiceClient(conn),
	}
}

func (c *Client) GetCategory(id int64) (*protoCategory.GetCategoryResponse, error) {
	return c.client.GetCategory(context.Background(), &protoCategory.GetCategoryRequest{
		ID : id,
	})
}

func (c *Client) GetCategories() error {
		stream, err := c.client.GetCategories(context.Background(), &protoCategory.GetCategoriesRequest{})

		if err != nil {
			return err
		}

		for {
		res, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}
		log.Printf("Received ID: %d", res.GetID())	
	}
}