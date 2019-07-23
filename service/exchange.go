package service

import (
	"context"
	"fmt"

	"cloud.google.com/go/firestore"
	"github.com/atotto/webrtc-sdp-exchanger/apis"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var _ apis.ExchangeServiceServer = (*ExchangeService)(nil)

type ExchangeService struct {
	fsClient *firestore.Client
}

func NewExchangeService(fsClient *firestore.Client) *ExchangeService {
	return &ExchangeService{fsClient: fsClient}
}

func (s *ExchangeService) CreateSession(ctx context.Context, req *apis.CreateSessionRequest) (*apis.CreateSessionResponse, error) {
	s.fsClient.Collection("sessions").Doc(session(req.SessionId, req.SessionDescription.Type)).Set(ctx, req)
	return &apis.CreateSessionResponse{}, nil
}

func (s *ExchangeService) GetSessionAnswer(ctx context.Context, req *apis.GetSessionRequest) (*apis.GetSessionResponse, error) {
	ss, err := s.fsClient.Collection("sessions").Doc(session(req.SessionId, "answer")).Get(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to get session")
	}
	res := &apis.GetSessionResponse{}
	if err := ss.DataTo(res); err != nil {
		return nil, status.Error(codes.Internal, "failed to read session")
	}
	return res, nil
}

func (s *ExchangeService) GetSessionOffer(ctx context.Context, req *apis.GetSessionRequest) (*apis.GetSessionResponse, error) {
	ss, err := s.fsClient.Collection("sessions").Doc(session(req.SessionId, "offer")).Get(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to get session")
	}
	res := &apis.GetSessionResponse{}
	if err := ss.DataTo(res); err != nil {
		return nil, status.Error(codes.Internal, "failed to read session")
	}
	return res, nil
}

func session(sessionID string, sdpType string) string {
	return fmt.Sprintf("%s-%s", sessionID, sdpType)
}
