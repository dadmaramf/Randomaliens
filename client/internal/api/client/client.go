package client

import (
	"client/internal/config"
	"client/internal/services"
	"context"

	frv1 "github.com/dadmaramf/protos/gen/go/frequency_service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Client struct {
	conn   *grpc.ClientConn
	stream frv1.StreamingService_NewConnectClient
	writer services.WriterAnomaliesService
}

func ConnectClient(ctx context.Context, writer services.WriterAnomaliesService, cfg *config.Config) (*Client, error) {
	conn, err := grpc.NewClient(cfg.HTTPClient.Host+cfg.HTTPClient.Port, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		return nil, err
	}
	client := frv1.NewStreamingServiceClient(conn)

	stream, err := client.NewConnect(ctx, &frv1.FrequencyRequest{Id: "1"})

	if err != nil {
		return nil, err
	}

	return &Client{
		conn:   conn,
		stream: stream,
		writer: writer,
	}, nil

}

func (c *Client) Recv() (*frv1.FrequencyResponse, error) {
	return c.stream.Recv()
}

func (c *Client) Close() {
	c.conn.Close()
}

func (c *Client) WriteAnomalies(ctx context.Context, data *frv1.FrequencyResponse) (err error) {
	return c.writer.WriteAnomalies(ctx, data)
}
