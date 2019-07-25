package service

import (
	"context"

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
	s.sessionCollection.Doc(req.SessionId).Set(ctx, req)
	return &apis.CreateSessionResponse{}, nil
}

func (s *ExchangeService) GetSessionAnswer(ctx context.Context, req *apis.GetSessionRequest) (*apis.GetSessionResponse, error) {
	return s.GetSession(ctx, req, "answer")
}

func (s *ExchangeService) GetSessionOffer(ctx context.Context, req *apis.GetSessionRequest) (*apis.GetSessionResponse, error) {
	return s.GetSession(ctx, req, "offer")
}

func (s *ExchangeService) GetSession(ctx context.Context, req *apis.GetSessionRequest, sdpType string) (*apis.GetSessionResponse, error) {
	ss, err := s.sessionCollection.Doc(req.SessionId).Get(ctx)
	if err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Error(codes.NotFound, "no data")
		}
		return nil, status.Error(codes.Internal, "failed to get session")
	}
	res := &apis.GetSessionResponse{}
	if err := ss.DataTo(res); err != nil {
		return nil, status.Error(codes.Internal, "failed to read session")
	}
	if res.SessionDescription.GetType() != sdpType {
		return nil, status.Error(codes.NotFound, "no data")
	}

	return res, nil
}

func (s *ExchangeService) DeleteSession(ctx context.Context, req *apis.DeleteSessionRequest) (*apis.DeleteSessionResponse, error) {
	s.sessionCollection.Doc(req.SessionId).Delete(ctx)
	return &apis.DeleteSessionResponse{}, nil
}
