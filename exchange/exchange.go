package exchange

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/pion/webrtc/v4"
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
	return getSessionWithRetry(ctx, sessionID, webrtc.SDPTypeAnswer)
}

func GetSessionOffer(ctx context.Context, sessionID string) (*webrtc.SessionDescription, error) {
	return getSessionWithRetry(ctx, sessionID, webrtc.SDPTypeOffer)
}

func getSessionWithRetry(ctx context.Context, sessionID string, sdpType webrtc.SDPType) (*webrtc.SessionDescription, error) {
	for {
		res, err := getSession(ctx, sessionID, sdpType)
		if err != nil {
			return nil, err
		}

		// retry
		if res == nil {
			select {
			case <-ctx.Done():
				return nil, ctx.Err()
			case <-time.After(3 * time.Second):
			}
		}
		return res, nil
	}
}

func getSession(ctx context.Context, sessionID string, sdpType webrtc.SDPType) (*webrtc.SessionDescription, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("https://webrtc-sdp-exchanger.appspot.com/sessions/%s/%s", sessionID, sdpType.String()), nil)
	if err != nil {
		return nil, err
	}
	req.WithContext(ctx)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("http get: %s", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		if resp.StatusCode != http.StatusNotFound {
			b, _ := ioutil.ReadAll(resp.Body)
			return nil, fmt.Errorf("invalid status %s: %s", resp.Status, string(b))
		}
		return nil, nil
	}

	res := sessionResponse{}
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, nil
	}

	if res.SessionDescription == nil {
		return nil, fmt.Errorf("not found: session_description")
	}

	return res.SessionDescription, nil
}
