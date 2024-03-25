package main

import (
	"fmt"
)

/**
Your task is to design and implement a meeting scheduler. The scheduler allows users to book time slots for incoming meetings. A meeting occupies one or more consecutive time slots, and never overlaps any other meetings. Initially there are 8 available time slots: 0, 1, 2, ..., 7.

The scheduler supports two operations:
(A) Scheduling a meeting - Given the meeting duration D (an integer within 1..8), the scheduler books D consecutive available time slots, or rejects the meeting if there are no suc
h time slots.

(B) Canceling a meeting - Given a scheduled meeting, the scheduler frees up all time slots which were booked for this meeting.
[id,]
*/

const MAX_SLOT = 8

type meeting struct {
	booked bool
	owner  string
}

var available [MAX_SLOT]meeting

func dumper() {
	for ndx := 0; ndx < MAX_SLOT; ndx++ {
		fmt.Printf("slot %d: %v %s\n", ndx, available[ndx].booked, available[ndx].owner)
	}
}

// assign a meeting, asssumes legal startNdx and duration within limits
func assign(startNdx, duration int, owner string) {
	//fmt.Printf("assigning %d %d %s\n", startNdx, duration, owner)

	for ndx := 0; ndx < duration; ndx++ {
		available[startNdx+ndx].booked = true
		available[startNdx+ndx].owner = owner
	}
}

// discover available meeting slot
func fitter(startNdx, duration int) bool {
	//fmt.Printf("entry:%d %d\n", startNdx, duration)

	for ndx := 0; ndx < duration; ndx++ {
		if startNdx+ndx >= MAX_SLOT {
			return false
		}

		if available[startNdx+ndx].booked == true {
			return false
		}
	}

	return true
}

func schedule(duration int) bool {
	for ndx := 0; ndx < MAX_SLOT; ndx++ {
		if available[ndx].booked == false {
			if duration == 1 {
				assign(ndx, duration, "test")
				return true
			} else {
				flag := fitter(ndx, duration)
				if flag == true {
					assign(ndx, duration, "test")
					return true
				}
			}
		}
	}

	return false
}

func reporter(duration int, retflag bool) {
	fmt.Printf("duration %d, return %v\n", duration, retflag)
	dumper()
}

func main() {
	//fmt.Println("Hello, World!")

	duration := 1
	retflag := schedule(duration)
	reporter(duration, retflag)

	duration = 2
	retflag = schedule(duration)
	reporter(duration, retflag)

	duration = 22
	retflag = schedule(duration)
	reporter(duration, retflag)
}
