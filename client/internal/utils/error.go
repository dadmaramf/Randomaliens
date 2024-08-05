package utils

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"io"
	"log"
)

type ErrorHandler interface {
	HandleStreamError(err error)
}

type DefaultErrorHandler struct{}

func (e *DefaultErrorHandler) HandleStreamError(err error) {
	if err == io.EOF {
		log.Println("Поток закрыт сервером")
	} else if st, ok := status.FromError(err); ok && st.Code() == codes.DeadlineExceeded {
		log.Println("Срок действия контекста истек")
	} else {
		log.Println("Ошибка при получении из потока:", err)
	}
}
