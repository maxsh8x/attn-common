package common

import "time"

type (
	// EventData is incoming data from site visitor
	EventData struct {
		OfferID     string  `json:"OfferId"`
		PlacementID string  `json:"PlacementId"`
		WidgetID    string  `json:"WidgetId"`
		TimeSpent   float32 `json:"timeSpent,omitempty"`
		Scrolled    *bool   `json:"scrolled,omitempty"`
		Latitude    float32 `json:"latitude,omitempty"`
		Lontitude   float32 `json:"lontitude,omitempty"`
	}

	// BaseEvent is common event required data which you can extract from request
	BaseEvent struct {
		Date     time.Time `json:"date"`
		Mobile   bool      `json:"mobile"`
		Platform string    `json:"platform"`
		OS       string    `json:"os"`
		Browser  string    `json:"browser"`
		Version  string    `json:"version"`
		IP       string    `json:"ip"`
	}

	// RabbitMSG is structs mixture which is AMQP event message
	RabbitMSG struct {
		EventData
		BaseEvent
	}
)
