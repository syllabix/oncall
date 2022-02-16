package mention

import (
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParser_Parse(t *testing.T) {
	type fields struct {
		regex *regexp.Regexp
	}
	type args struct {
		input string
	}
	tests := []struct {
		name    string
		input   string
		want    User
		wantErr bool
	}{
		{
			name:  "should match user",
			input: "<@U031RQ1L21K|jane.doe040>",
			want: User{
				ID:     "U031RQ1L21K",
				NameID: "jane.doe040",
			},
		},
		{
			name:    "should not match user",
			input:   "<@U031RQ1L21jane.doe040>",
			wantErr: true,
		},
		{
			name:    "should not match user",
			input:   "<@U031RQ1L21|jane.doe040",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			parser, err := NewParser()
			assert.NoError(t, err)

			got, err := parser.Parse(tt.input)
			if tt.wantErr {
				assert.Error(t, err)
			}

			assert.Equal(t, tt.want, got)
		})
	}
}

func TestParser_ParseList(t *testing.T) {

	tests := []struct {
		name    string
		input   string
		want    []User
		wantErr bool
	}{
		{
			name:  "should match 3 users",
			input: "<@U031CUA9T7Z|jane.doe123> <@U031RQ1L21K|john.smith812> <@U031RQ1PP1K|kanye.west89>",
			want: []User{
				{
					ID:     "U031CUA9T7Z",
					NameID: "jane.doe123",
				},
				{
					ID:     "U031RQ1L21K",
					NameID: "john.smith812",
				},
				{
					ID:     "U031RQ1PP1K",
					NameID: "kanye.west89",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			parser, err := NewParser()
			assert.NoError(t, err)

			got, err := parser.ParseList(tt.input)
			if tt.wantErr {
				assert.Error(t, err)
			}

			assert.Equal(t, tt.want, got)
		})
	}
}
