package protocol

// event types used across client ↔ server ↔ server communication
const (
	EventJoinRoom     = "join_room"
	EventLeaveRoom    = "leave_room"
	EventPublishEvent = "publish_event"
	EventPing         = "ping"
	EventResume       = "resume_session"

	EventRoomJoined = "room_joined"
	EventRoomLeft   = "room_left"
	EventBroadcast  = "event_broadcast"
	EventPresence   = "presence_update"
	EventError      = "error"

	EventPropagate = "event_propagate"
)

// payload for publish_event
type PublishPayload struct {
	Event string                 `json:"event"`
	Data  map[string]interface{} `json:"data"`
}