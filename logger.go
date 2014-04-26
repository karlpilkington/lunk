package lunk

import (
	"encoding/json"
	"io"
)

// An EventLogger logs events and their metadata.
type EventLogger interface {
	// Log adds the given event to the log stream, returning the logged event's
	// unique ID.
	Log(root, parent ID, e Event) ID
}

// NewJSONEventLogger returns an EventLogger which writes entries as streaming
// JSON to the given writer.
func NewJSONEventLogger(w io.Writer) EventLogger {
	return jsonEventLogger{json.NewEncoder(w)}
}

// RawJSONEntry allows entries to be parsed without any knowledge of the schemas.
type RawJSONEntry struct {
	Metadata

	// Event is the unparsed event data.
	Event json.RawMessage `json:"event"`
}

type jsonEventLogger struct {
	*json.Encoder
}

func (l jsonEventLogger) Log(root, parent ID, e Event) ID {
	entry := NewEntry(root, parent, e)
	if err := l.Encode(entry); err != nil {
		panic(err)
	}
	return entry.ID
}