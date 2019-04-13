package main

// model
type Note struct {
    ID        string   `json:"id,omitempty"`
    Timestamp int      `json:"timestamp,omitempty"`
    Text      string   `json:"text,omitempty"`
}