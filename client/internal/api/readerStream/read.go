package readerstream

import (
	"client/internal/utils"
	"context"
	frv1 "github.com/dadmaramf/protos/gen/go/frequency_service"
)

type ReaderWriterStream interface {
	Recv() (*frv1.FrequencyResponse, error)
	WriteAnomalies(ctx context.Context, data *frv1.FrequencyResponse) (err error)
}

func ReadStream(client ReaderWriterStream, doneTicker chan bool, errorHandler utils.ErrorHandler) {
	for {
		resp, err := client.Recv()
		if err != nil {
			doneTicker <- true
			errorHandler.HandleStreamError(err)
			return
		}
		client.WriteAnomalies(context.Background(), resp)
	}
}
