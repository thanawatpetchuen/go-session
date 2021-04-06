package gos

type User struct {
	Username string `json:"username"`
}

type LoginBody struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type GetSessionResponse struct {
	SessionID string `json:"sessionId"`
	Data      User   `json:"data"`
}

func NewSessionResponse(user User, sessionId string) GetSessionResponse {
	return GetSessionResponse{
		Data:      user,
		SessionID: sessionId,
	}
}
