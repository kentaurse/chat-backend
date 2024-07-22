package server

import (
	"context"

	protos "github.com/FlutterGoChatApp/chat_app_backend/services/messaging/protos"
	"google.golang.org/protobuf/types/known/wrapperspb"

	"github.com/rs/zerolog"
)

type MessagingServer struct {
	log zerolog.Logger

	protos.UnimplementedMessagingServer
}

func NewMessagingServer(logger zerolog.Logger) *MessagingServer {
	server := MessagingServer{}
	server.log = logger.With().Str("service", "messaging").Logger()
	return &server
}

func (h *MessagingServer) SendMessage(ctx context.Context, msg *protos.MessageSend) (*wrapperspb.BoolValue, error) {
	h.log.Info().Msgf("I got message: %s", msg.Message)
	return &wrapperspb.BoolValue{Value: true}, nil
}
