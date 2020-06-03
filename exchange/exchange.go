package exchange

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/pion/webrtc/v2"
)

type createSessionRequest struct {
	SessionDescription *webrtc.SessionDescription `json:"session_description"`
	SessionID          string                     `json:"session_id"`
}

type sessionResponse struct {
	SessionDescription *webrtc.SessionDescription `json:"session_description"`
}

func CreateSession(ctx context.Context, offer *webrtc.SessionDescription, sessionID string) error {
	req := createSessionRequest{
		SessionID:          sessionID,
		SessionDescription: offer,
	}

	buf := new(bytes.Buffer)
	err := json.NewEncoder(buf).Encode(req)
	if err != nil {
		return fmt.Errorf("failed to encode: %s", err)
	}

	resp, err := http.Post(fmt.Sprintf("https://webrtc-sdp-exchanger.appspot.com/sessions/%s", sessionID), "application/json; charset=utf-8", buf)
	if err != nil {
		return fmt.Errorf("http post: %s", err)
	}
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("invalid status: %s", resp.Status)
	}

	return nil
}

func GetSessionAnswer(ctx context.Context, sessionID string) (*webrtc.SessionDescription, error) {
	return getSession(ctx, sessionID, webrtc.SDPTypeAnswer)
}

func GetSessionOffer(ctx context.Context, sessionID string) (*webrtc.SessionDescription, error) {
	return getSession(ctx, sessionID, webrtc.SDPTypeOffer)
}

func getSession(ctx context.Context, sessionID string, sdpType webrtc.SDPType) (*webrtc.SessionDescription, error) {
	res := sessionResponse{}
	for {
		resp, err := http.Get(fmt.Sprintf("https://webrtc-sdp-exchanger.appspot.com/sessions/%s/%s", sessionID, sdpType.String()))
		if err != nil {
			return nil, fmt.Errorf("http get: %s", err)
		}
		if resp.StatusCode != http.StatusOK {
			//return nil, fmt.Errorf("invalid status: %s", resp.Status)
			resp.Body.Close()
			continue
		}

		if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
			resp.Body.Close()
			continue
		}

		if res.SessionDescription != nil {
			return res.SessionDescription, nil
		}

		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case <-time.After(3 * time.Second):
		}
	}
}
