package schedule

import (
	"context"
	"embed"
	"io/fs"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/syllabix/oncall/common/db"
	"github.com/syllabix/oncall/datastore/model"
)

//go:embed test_scenarios/*.sql
var scenarios embed.FS

func TestStore_GetByID(t *testing.T) {

	testdb := db.NewTester()

	testfs, err := fs.Sub(scenarios, "test_scenarios")
	assert.NoError(t, err)

	testdb.Reset()

	type args struct {
		ctx context.Context
		id  int
	}
	tests := []struct {
		name     string
		scenario string
		args     args
		want     *model.Schedule
		wantErr  bool
	}{
		{
			name:     "shifts_returned_in_order",
			scenario: "test_case_1.sql",
			args: args{
				ctx: context.Background(),
				id:  1,
			},
			want: &model.Schedule{
				ID:             1,
				SlackChannelID: "C02FZ54HJP6",
				TeamSlackID:    "T022YEG9FGA",
				Name:           "Daily On Call",
				Interval:       "daily",
				IsEnabled:      true,
				WeekdaysOnly:   true,
				StartTime:      time.Date(0, 0, 0, 8, 0, 0, 0, time.UTC),
				EndTime:        time.Date(0, 0, 0, 16, 0, 0, 0, time.UTC),
			},
		},
		// TODO: more test cases
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sql, err := testfs.Open(tt.scenario)
			assert.NoError(t, err)

			defer testdb.Reset()

			err = testdb.Seed(sql)
			if err != nil {
				assert.NoError(t, err)
				return
			}

			s := &Store{db: testdb.GetDB()}
			got, err := s.GetByID(tt.args.ctx, tt.args.id)

			assert.Equal(t, tt.want.ID, got.ID)
			assert.Equal(t, tt.want.SlackChannelID, got.SlackChannelID)
			assert.Equal(t, tt.want.TeamSlackID, got.TeamSlackID)
			assert.Equal(t, tt.want.Name, got.Name)
			assert.Equal(t, tt.want.IsEnabled, got.IsEnabled)
			assert.Equal(t, tt.want.WeekdaysOnly, got.WeekdaysOnly)

			if len(got.R.Shifts) > 0 {
				// assert shifts are returned in asc order by sequence
				prevseq := 0
				for _, shift := range got.R.Shifts {
					assert.Greater(t, shift.SequenceID, prevseq)
					prevseq = shift.SequenceID
				}
			}

			err = testdb.Reset()
			assert.NoError(t, err)
		})
	}
}

func TestStore_GetByChannelID(t *testing.T) {
	testdb := db.NewTester()

	testfs, err := fs.Sub(scenarios, "test_scenarios")
	assert.NoError(t, err)

	testdb.Reset()

	type args struct {
		ctx     context.Context
		slackID string
	}
	tests := []struct {
		name     string
		scenario string
		args     args
		want     *model.Schedule
		wantErr  bool
	}{
		{
			name:     "shifts_returned_in_order",
			scenario: "test_case_1.sql",
			args: args{
				ctx:     context.Background(),
				slackID: "C02FZ54HJP6",
			},
			want: &model.Schedule{
				ID:             1,
				SlackChannelID: "C02FZ54HJP6",
				TeamSlackID:    "T022YEG9FGA",
				Name:           "Daily On Call",
				Interval:       "daily",
				IsEnabled:      true,
				WeekdaysOnly:   true,
				StartTime:      time.Date(0, 0, 0, 8, 0, 0, 0, time.UTC),
				EndTime:        time.Date(0, 0, 0, 16, 0, 0, 0, time.UTC),
			},
		},
		// TODO: more test cases
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sql, err := testfs.Open(tt.scenario)
			assert.NoError(t, err)

			defer testdb.Reset()

			err = testdb.Seed(sql)
			if err != nil {
				assert.NoError(t, err)
				return
			}

			s := &Store{db: testdb.GetDB()}
			got, err := s.GetByChannelID(tt.args.ctx, tt.args.slackID)
			if err != nil {
				assert.NoError(t, err)
				return
			}

			assert.Equal(t, tt.want.ID, got.ID)
			assert.Equal(t, tt.want.SlackChannelID, got.SlackChannelID)
			assert.Equal(t, tt.want.TeamSlackID, got.TeamSlackID)
			assert.Equal(t, tt.want.Name, got.Name)
			assert.Equal(t, tt.want.IsEnabled, got.IsEnabled)
			assert.Equal(t, tt.want.WeekdaysOnly, got.WeekdaysOnly)

			if len(got.R.Shifts) > 0 {
				// assert shifts are returned in asc order by sequence
				prevseq := 0
				for _, shift := range got.R.Shifts {
					assert.Greater(t, shift.SequenceID, prevseq)
					prevseq = shift.SequenceID
				}
			}

			err = testdb.Reset()
			assert.NoError(t, err)
		})
	}
}
