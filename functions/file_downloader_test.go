package functions

import "testing"

func TestDownload(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{"standard", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Download(tt.name); (err != nil) != tt.wantErr {
				t.Errorf("Download() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
