package main

import (
	"net"
	"os"
	"time"

	messagingProtos "github.com/FlutterGoChatApp/chat_app_backend/services/messaging/protos"
	messagingService "github.com/FlutterGoChatApp/chat_app_backend/services/messaging/server"
	"github.com/rs/zerolog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const SERVER_PROTO = "tcp"
const SERVER_ADDR = "localhost:50000"

type ChatAppServer struct {
	grpcServer *grpc.Server
	log        zerolog.Logger

	/* Services */
	messaging *messagingService.MessagingServer
}

func (s *ChatAppServer) createServer() {
	s.log.Info().Msg("creating gRPC server")

	s.grpcServer = grpc.NewServer()

	//TODO: Add authentication + encryption
}

func (s *ChatAppServer) registerServices() {
	s.log.Info().Msg("registering gRPC services to server")

	if s.grpcServer == nil {
		s.log.Fatal().Msg("gRPC server is not initialized")
	}

	s.messaging = messagingService.NewMessagingServer(s.log)
	messagingProtos.RegisterMessagingServer(s.grpcServer, s.messaging)

	reflection.Register(s.grpcServer)
}

func (s *ChatAppServer) startServer() {
	s.log.Info().Msg("starting gRPC server")

	if s.grpcServer == nil {
		s.log.Fatal().Msg("gRPC server is not initialized")
	}

	listen, err := net.Listen(SERVER_PROTO, SERVER_ADDR)
	if err != nil {
		s.log.Fatal().Msgf("failed to create listener, %v", err)
	}

	err = s.grpcServer.Serve(listen)
	if err != nil {
		s.log.Fatal().Msgf("failed to start GRPC server, %v", err)
	}
}

func main() {
	server := ChatAppServer{}

	server.log = zerolog.New(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC3339}).
		Level(zerolog.TraceLevel).
		With().
		Timestamp().
		Caller().
		Logger()

	server.createServer()
	server.registerServices()
	server.startServer()
}
