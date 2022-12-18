package txtstore

import "testing"

func TestRead(t *testing.T) {
	tests := []struct {
		id    string
		value string
	}{
		{value: "123"},
		{value: "456"},
		{value: "789"},
	}

	for i, tt := range tests {
		tests[i].id = Write(tt.value)
	}

	for _, tt := range tests {
		t.Run(tt.id, func(t *testing.T) {
			if got := Read(tt.id); got != tt.value {
				t.Errorf("Read() = got [%v], want [%v]", got, tt.value)
			}
		})
	}
}
