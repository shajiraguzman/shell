// Code generated by "stringer -type=Scheduler"; DO NOT EDIT.

package main

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[fcfs-1]
	_ = x[sjf-2]
	_ = x[sjfp-3]
	_ = x[rr-4]
}

const _Scheduler_name = "fcfssjfsjfprr"

var _Scheduler_index = [...]uint8{0, 4, 7, 11, 13}

func (i Scheduler) String() string {
	i -= 1
	if i >= Scheduler(len(_Scheduler_index)-1) {
		return "Scheduler(" + strconv.FormatInt(int64(i+1), 10) + ")"
	}
	return _Scheduler_name[_Scheduler_index[i]:_Scheduler_index[i+1]]
}