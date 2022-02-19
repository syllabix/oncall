package schedule

import (
	"reflect"
	"testing"

	"github.com/syllabix/oncall/datastore/model"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/types"
)

func Test_nextShift(t *testing.T) {
	type args struct {
		activeShift string
		shifts      []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "moves_to_next_shift",
			args: args{
				activeShift: "123",
				shifts: []string{
					"234",
					"345",
					"456",
					"262",
					"123",
					"897",
				},
			},
			want: "897",
		},
		{
			name: "moves_to_next_shift_at_beginning_of_shift_list",
			args: args{
				activeShift: "897",
				shifts: []string{
					"234",
					"345",
					"456",
					"262",
					"123",
					"897",
				},
			},
			want: "234",
		},
		{
			name: "only_one_soldier",
			args: args{
				activeShift: "234",
				shifts: []string{
					"234",
				},
			},
			want: "234",
		},
		{
			name: "nothing_setup_yet",
			args: args{
				activeShift: "",
				shifts: []string{
					"234",
				},
			},
			want: "234",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := nextShift(tt.args.activeShift, tt.args.shifts); got != tt.want {
				t.Errorf("nextShift() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_nextShifts(t *testing.T) {
	type args struct {
		count    int
		schedule model.Schedule
	}
	tests := []struct {
		name       string
		args       args
		wantShifts []string
	}{
		{
			name: "should_fetch_next_shifts",
			args: args{
				count: 5,
				schedule: model.Schedule{
					ActiveShift: null.StringFrom("55"),
					Shifts:      types.StringArray{"11", "22", "33", "44", "55", "66", "77"},
				},
			},
			wantShifts: []string{"66", "77", "11", "22", "33"},
		},
		{
			name: "should_fetch_next_shifts",
			args: args{
				count: 10,
				schedule: model.Schedule{
					ActiveShift: null.StringFrom("77"),
					Shifts:      types.StringArray{"11", "22", "33", "44", "55", "66", "77"},
				},
			},
			wantShifts: []string{"11", "22", "33", "44", "55", "66", "77", "11", "22", "33"},
		},
		{
			name: "should_fetch_next_shifts",
			args: args{
				count: 1,
				schedule: model.Schedule{
					ActiveShift: null.StringFrom("22"),
					Shifts:      types.StringArray{"11", "22", "33"},
				},
			},
			wantShifts: []string{"33"},
		},
		{
			name: "should_fetch_next_shifts",
			args: args{
				count: 5,
				schedule: model.Schedule{
					Shifts: types.StringArray{"11", "22", "33"},
				},
			},
			wantShifts: []string{"11", "22", "33", "11", "22"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotShifts := nextShifts(tt.args.count, tt.args.schedule); !reflect.DeepEqual(gotShifts, tt.wantShifts) {
				t.Errorf("nextShifts() = %v, want %v", gotShifts, tt.wantShifts)
			}
		})
	}
}
