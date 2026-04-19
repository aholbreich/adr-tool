package model

import "testing"

func TestIsFinalStatus(t *testing.T) {
	tests := []struct {
		name   string
		status ADRStatus
		want   bool
	}{
		{name: "unknown is not final", status: StatusUnknown, want: false},
		{name: "draft is not final", status: StatusDraft, want: false},
		{name: "proposed is not final", status: StatusProposed, want: false},
		{name: "accepted is final", status: StatusAccepted, want: true},
		{name: "rejected is final", status: StatusRejected, want: true},
		{name: "deprecated is final", status: StatusDeprecated, want: true},
		{name: "superseded is final", status: StatusSuperseded, want: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IsFinalStatus(tt.status)
			if got != tt.want {
				t.Fatalf("got %v, want %v", got, tt.want)
			}
		})
	}
}
