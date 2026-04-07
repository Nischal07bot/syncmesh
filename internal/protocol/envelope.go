package protocol

type Envelope struct {
	ID  string  `json:"id"`
	Type string `json:"type"`
	RoomID string `json:"roomId"`
	UserID string `json:"userId"`
	ServerID string `json:"serverId"`
	Timestamp int64  `json:"timestamp"`
	Payload map[string]interface{} `json:"payload"`
	Meta    map[string]interface{} `json:"meta"`
}

func New(eventType string, roomID string, userID string, serverID string, payload map[string]interface{}) *Envelope 
{
	return &Envelope{
		ID:        generateUniqueID(),
		Type:      eventType,
		RoomID:    roomID,
		UserID:    userID,
		ServerID:  serverID,
		Timestamp: time.Now().Unix(),
		Payload:   payload,
		Meta:      make(map[string]interface{}),
	}
}

func Encode(e *Envelope) ([]byte, error) {
	return json.Marshal(e)
}