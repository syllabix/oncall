package schedule

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/syllabix/oncall/datastore/model"
	"github.com/volatiletech/null/v8"
)

func Test_arrange(t *testing.T) {
	type args struct {
		shifts model.ShiftSlice
	}
	type expected struct {
		active  *model.Shift
		ordered model.ShiftSlice
	}
	tests := []struct {
		name string
		args args
		want expected
	}{
		{
			name: "returns_shifts_properly_ordered",
			args: args{
				shifts: model.ShiftSlice{
					{
						SequenceID: 7,
						UserID:     100,
						ScheduleID: 999,
					},
					{
						SequenceID: 10,
						UserID:     101,
						ScheduleID: 999,
					},
					{
						SequenceID: 13,
						UserID:     102,
						ScheduleID: 999,
					},
					{
						SequenceID: 29,
						UserID:     103,
						ScheduleID: 999,
						Status:     null.StringFrom(model.ShiftStatusActive),
					},
					{
						SequenceID: 32,
						UserID:     104,
						ScheduleID: 999,
					},
				},
			},
			want: expected{
				active: &model.Shift{
					SequenceID: 29,
					UserID:     103,
					ScheduleID: 999,
					Status:     null.StringFrom(model.ShiftStatusActive),
				},
				ordered: model.ShiftSlice{
					{
						SequenceID: 32,
						UserID:     104,
						ScheduleID: 999,
					},
					{
						SequenceID: 7,
						UserID:     100,
						ScheduleID: 999,
					},
					{
						SequenceID: 10,
						UserID:     101,
						ScheduleID: 999,
					},
					{
						SequenceID: 13,
						UserID:     102,
						ScheduleID: 999,
					},
					{
						SequenceID: 29,
						UserID:     103,
						ScheduleID: 999,
						Status:     null.StringFrom(model.ShiftStatusActive),
					},
				},
			},
		},
		{
			name: "returns_shifts_properly_ordered",
			args: args{
				shifts: model.ShiftSlice{
					{
						SequenceID: 7,
						UserID:     100,
						ScheduleID: 999,
					},
					{
						SequenceID: 10,
						UserID:     101,
						ScheduleID: 999,
						Status:     null.StringFrom(model.ShiftStatusActive),
					},
					{
						SequenceID: 13,
						UserID:     102,
						ScheduleID: 999,
					},
					{
						SequenceID: 29,
						UserID:     103,
						ScheduleID: 999,
					},
					{
						SequenceID: 32,
						UserID:     104,
						ScheduleID: 999,
					},
				},
			},
			want: expected{
				active: &model.Shift{
					SequenceID: 10,
					UserID:     101,
					ScheduleID: 999,
					Status:     null.StringFrom(model.ShiftStatusActive),
				},
				ordered: model.ShiftSlice{
					{
						SequenceID: 13,
						UserID:     102,
						ScheduleID: 999,
					},
					{
						SequenceID: 29,
						UserID:     103,
						ScheduleID: 999,
					},
					{
						SequenceID: 32,
						UserID:     104,
						ScheduleID: 999,
					},
					{
						SequenceID: 7,
						UserID:     100,
						ScheduleID: 999,
					},
					{
						SequenceID: 10,
						UserID:     101,
						ScheduleID: 999,
						Status:     null.StringFrom(model.ShiftStatusActive),
					},
				},
			},
		},
		{
			name: "returns_shifts_properly_ordered",
			args: args{
				shifts: model.ShiftSlice{
					{
						SequenceID: 7,
						UserID:     100,
						ScheduleID: 999,
						Status:     null.StringFrom(model.ShiftStatusActive),
					},
					{
						SequenceID: 10,
						UserID:     101,
						ScheduleID: 999,
					},
					{
						SequenceID: 13,
						UserID:     102,
						ScheduleID: 999,
					},
					{
						SequenceID: 29,
						UserID:     103,
						ScheduleID: 999,
					},
					{
						SequenceID: 32,
						UserID:     104,
						ScheduleID: 999,
					},
				},
			},
			want: expected{
				active: &model.Shift{
					SequenceID: 7,
					UserID:     100,
					ScheduleID: 999,
					Status:     null.StringFrom(model.ShiftStatusActive),
				},
				ordered: model.ShiftSlice{
					{
						SequenceID: 10,
						UserID:     101,
						ScheduleID: 999,
					},
					{
						SequenceID: 13,
						UserID:     102,
						ScheduleID: 999,
					},
					{
						SequenceID: 29,
						UserID:     103,
						ScheduleID: 999,
					},
					{
						SequenceID: 32,
						UserID:     104,
						ScheduleID: 999,
					},
					{
						SequenceID: 7,
						UserID:     100,
						ScheduleID: 999,
						Status:     null.StringFrom(model.ShiftStatusActive),
					},
				},
			},
		},
		{
			name: "returns_shifts_properly_ordered",
			args: args{
				shifts: model.ShiftSlice{
					{
						SequenceID: 7,
						UserID:     100,
						ScheduleID: 999,
					},
					{
						SequenceID: 10,
						UserID:     101,
						ScheduleID: 999,
					},
					{
						SequenceID: 13,
						UserID:     102,
						ScheduleID: 999,
					},
					{
						SequenceID: 29,
						UserID:     103,
						ScheduleID: 999,
					},
					{
						SequenceID: 32,
						UserID:     104,
						ScheduleID: 999,
						Status:     null.StringFrom(model.ShiftStatusActive),
					},
				},
			},
			want: expected{
				active: &model.Shift{
					SequenceID: 32,
					UserID:     104,
					ScheduleID: 999,
					Status:     null.StringFrom(model.ShiftStatusActive),
				},
				ordered: model.ShiftSlice{
					{
						SequenceID: 7,
						UserID:     100,
						ScheduleID: 999,
					},
					{
						SequenceID: 10,
						UserID:     101,
						ScheduleID: 999,
					},
					{
						SequenceID: 13,
						UserID:     102,
						ScheduleID: 999,
					},
					{
						SequenceID: 29,
						UserID:     103,
						ScheduleID: 999,
					},
					{
						SequenceID: 32,
						UserID:     104,
						ScheduleID: 999,
						Status:     null.StringFrom(model.ShiftStatusActive),
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			active, ordered := arrange(tt.args.shifts)
			assert.EqualValues(t, tt.want.active, active)
			assert.EqualValues(t, tt.want.ordered, ordered)
		})
	}
}
