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

func Decode(data []byte) (*Envelope, error) {
	var e Envelope
	if err := json.Unmarshal(data, &e); err != nil {
		return nil, err
	}
	return &e, nil
}

func Validate(e *Envelope) error {
	if e.Type == "" {
		return errors.New("missing type")
	}
	if e.RoomID == "" {
		return errors.New("missing room_id")
	}
	if e.UserID == "" {
		return errors.New("missing user_id")
	}
	return nil
}

func generateID() string {
	return time.Now().Format("20060102150405.000000000")
}