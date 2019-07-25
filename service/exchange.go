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
	fsClient          *firestore.Client
	sessionCollection *firestore.CollectionRef
}

func NewExchangeService(fsClient *firestore.Client) *ExchangeService {
	return &ExchangeService{
		fsClient:          fsClient,
		sessionCollection: fsClient.Collection("sessions"),
	}
}

func (s *ExchangeService) CreateSession(ctx context.Context, req *apis.CreateSessionRequest) (*apis.CreateSessionResponse, error) {
	switch req.SessionDescription.Type {
	case "offer":
		s.sessionCollection.Doc(session(req.SessionId, "answer")).Delete(ctx)
	case "answer":
		s.sessionCollection.Doc(session(req.SessionId, "offer")).Delete(ctx)
	default:
		return nil, status.Errorf(codes.InvalidArgument, "unknown SDP Type: %s", req.SessionDescription.Type)
	}
	s.sessionCollection.Doc(session(req.SessionId, req.SessionDescription.Type)).Set(ctx, req)
	return &apis.CreateSessionResponse{}, nil
}

func (s *ExchangeService) GetSessionAnswer(ctx context.Context, req *apis.GetSessionRequest) (*apis.GetSessionResponse, error) {
	res := &apis.GetSessionResponse{}
	ss, err := s.sessionCollection.Doc(session(req.SessionId, "answer")).Get(ctx)
	if err != nil {
		if status.Code(err) == codes.NotFound {
			return res, nil
		}
		return nil, status.Error(codes.Internal, "failed to get session")
	}
	if err := ss.DataTo(res); err != nil {
		return nil, status.Error(codes.Internal, "failed to read session")
	}
	return res, nil
}

func (s *ExchangeService) GetSessionOffer(ctx context.Context, req *apis.GetSessionRequest) (*apis.GetSessionResponse, error) {
	res := &apis.GetSessionResponse{}
	ss, err := s.sessionCollection.Doc(session(req.SessionId, "offer")).Get(ctx)
	if err != nil {
		if status.Code(err) == codes.NotFound {
			return res, nil
		}
		return nil, status.Error(codes.Internal, "failed to get session")
	}
	if err := ss.DataTo(res); err != nil {
		return nil, status.Error(codes.Internal, "failed to read session")
	}
	return res, nil
}

func (s *ExchangeService) DeleteSession(ctx context.Context, req *apis.DeleteSessionRequest) (*apis.DeleteSessionResponse, error) {
	s.sessionCollection.Doc(session(req.SessionId, "answer")).Delete(ctx)
	s.sessionCollection.Doc(session(req.SessionId, "offer")).Delete(ctx)
	return &apis.DeleteSessionResponse{}, nil
}

func session(sessionID string, sdpType string) string {
	return fmt.Sprintf("%s-%s", sessionID, sdpType)
}
