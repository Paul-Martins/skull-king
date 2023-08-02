package game

import (
	"testing"

	"github.com/metalblueberry/skull-king/pkg/skullking"
)

func TestSendCommand(t *testing.T) {
	type args struct {
		command Command
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Command PlayCard",
			args: args{
				command: CommandPlayCard{
					Card: skullking.Card{Number: 57, Type: skullking.CardTypePirate},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := SendCommand(tt.args.command); (err != nil) != tt.wantErr {
				t.Errorf("SendCommand() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
