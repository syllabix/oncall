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
						SequenceID: 29,
						UserID:     103,
						ScheduleID: 999,
						Status:     null.StringFrom(model.ShiftStatusActive),
					},
					{
						SequenceID: 13,
						UserID:     102,
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
		{
			name: "returns_shifts_properly_ordered_when_override_is_present",
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
						SequenceID: 29,
						UserID:     103,
						ScheduleID: 999,
					},
					{
						SequenceID: 13,
						UserID:     102,
						ScheduleID: 999,
						Status:     null.StringFrom(model.ShiftStatusActive),
					},
					{
						SequenceID: 32,
						UserID:     104,
						ScheduleID: 999,
						Status:     null.StringFrom(model.ShiftStatusOverride),
					},
				},
			},
			want: expected{
				active: &model.Shift{
					SequenceID: 13,
					UserID:     102,
					ScheduleID: 999,
					Status:     null.StringFrom(model.ShiftStatusActive),
				},
				ordered: model.ShiftSlice{
					{
						SequenceID: 29,
						UserID:     103,
						ScheduleID: 999,
					},
					{
						SequenceID: 32,
						UserID:     104,
						ScheduleID: 999,
						Status:     null.StringFrom(model.ShiftStatusOverride),
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

func Test_nextShiftFrom(t *testing.T) {
	type args struct {
		schedule *model.Schedule
	}
	tests := []struct {
		name        string
		args        args
		wantCurrent *model.Shift
		wantNext    *model.Shift
	}{
		{
			name: "returns_the_active_and_next_shift_1",
			args: args{
				schedule: func() *model.Schedule {
					sched := new(model.Schedule)
					sched.ID = 999
					sched.R = sched.R.NewStruct()
					sched.R.Shifts = model.ShiftSlice{
						{
							SequenceID: 6,
							UserID:     111,
							ScheduleID: 999,
						},
						{
							SequenceID: 10,
							UserID:     222,
							ScheduleID: 999,
							Status:     null.StringFrom(model.ShiftStatusActive),
						},
					}
					return sched
				}(),
			},
			wantCurrent: &model.Shift{
				SequenceID: 10,
				UserID:     222,
				ScheduleID: 999,
				Status:     null.StringFrom(model.ShiftStatusActive),
			},
			wantNext: &model.Shift{
				SequenceID: 6,
				UserID:     111,
				ScheduleID: 999,
			},
		},
		{
			name: "returns_the_active_and_next_shift_2",
			args: args{
				schedule: func() *model.Schedule {
					sched := new(model.Schedule)
					sched.ID = 999
					sched.R = sched.R.NewStruct()
					sched.R.Shifts = model.ShiftSlice{
						{
							SequenceID: 6,
							UserID:     111,
							ScheduleID: 999,
							Status:     null.StringFrom(model.ShiftStatusActive),
						},
						{
							SequenceID: 10,
							UserID:     222,
							ScheduleID: 999,
						},
					}
					return sched
				}(),
			},
			wantCurrent: &model.Shift{
				SequenceID: 6,
				UserID:     111,
				ScheduleID: 999,
				Status:     null.StringFrom(model.ShiftStatusActive),
			},
			wantNext: &model.Shift{
				SequenceID: 10,
				UserID:     222,
				ScheduleID: 999,
			},
		},
		{
			name: "returns_the_active_and_next_shift_3",
			args: args{
				schedule: func() *model.Schedule {
					sched := new(model.Schedule)
					sched.ID = 999
					sched.R = sched.R.NewStruct()
					sched.R.Shifts = model.ShiftSlice{
						{
							SequenceID: 10,
							UserID:     222,
							ScheduleID: 999,
						},
						{
							SequenceID: 12,
							UserID:     333,
							ScheduleID: 999,
						},
						{
							SequenceID: 23,
							UserID:     444,
							ScheduleID: 999,
						},
						{
							SequenceID: 33,
							UserID:     555,
							ScheduleID: 999,
							Status:     null.StringFrom(model.ShiftStatusActive),
						},
						{
							SequenceID: 42,
							UserID:     666,
							ScheduleID: 999,
						},
					}
					return sched
				}(),
			},
			wantCurrent: &model.Shift{
				SequenceID: 33,
				UserID:     555,
				ScheduleID: 999,
				Status:     null.StringFrom(model.ShiftStatusActive),
			},
			wantNext: &model.Shift{
				SequenceID: 42,
				UserID:     666,
				ScheduleID: 999,
			},
		},
		{
			name: "returns_the_active_and_next_shift_1_when_there_is_no_active",
			args: args{
				schedule: func() *model.Schedule {
					sched := new(model.Schedule)
					sched.ID = 999
					sched.R = sched.R.NewStruct()
					sched.R.Shifts = model.ShiftSlice{
						{
							SequenceID: 6,
							UserID:     111,
							ScheduleID: 999,
						},
						{
							SequenceID: 10,
							UserID:     222,
							ScheduleID: 999,
						},
						{
							SequenceID: 12,
							UserID:     333,
							ScheduleID: 999,
						},
					}
					return sched
				}(),
			},
			wantNext: &model.Shift{
				SequenceID: 6,
				UserID:     111,
				ScheduleID: 999,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotCurrent, gotNext := nextShiftFrom(tt.args.schedule)
			assert.EqualValues(t, tt.wantCurrent, gotCurrent)
			assert.EqualValues(t, tt.wantNext, gotNext)
		})
	}
}
