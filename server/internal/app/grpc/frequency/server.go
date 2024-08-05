package frequency

import (
	"frequency_service/internal/services"
	"log"
	"time"

	frv1 "github.com/dadmaramf/protos/gen/go/frequency_service"
	"google.golang.org/grpc"
)

type server struct {
	frv1.UnimplementedStreamingServiceServer
	services *services.Services
}

func Register(gr *grpc.Server, s *services.Services) {
	frv1.RegisterStreamingServiceServer(gr, &server{services: s})
}

func (s *server) NewConnect(req *frv1.FrequencyRequest, stream frv1.StreamingService_NewConnectServer) error {
	s.services.ServiceFrequency.GenerateInit()

	for {
		select {
		case <-stream.Context().Done():
			log.Println("MEAN:  ", s.services.GetMean(), "STDDEV:  ", s.services.GetStddev())
			return nil

		default:

			resp := &frv1.FrequencyResponse{
				SessionId: s.services.GetUUID(),
				Timenow:   s.services.UpdateGenerateTime(),
				Frequency: s.services.UpdateGenerateFrequency(),
			}

			if err := stream.Send(resp); err != nil {
				log.Println("error generate response", err)
				return err
			}
			log.Println("The client is successfully connected", resp)
		}

		time.Sleep(time.Millisecond * 1)
	}
}
