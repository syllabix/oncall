package schedule

import "testing"

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
