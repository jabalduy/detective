package game

import (
	"testing"
	"time"
)

func TestClockRunTime(t *testing.T) {
	clock := NewClock(time.Now())

	go clock.RunTime()
	defer clock.StopTime()

	before := clock.GetTime()

	time.Sleep(2 * time.Second)

	after := clock.GetTime()

	if !after.After(before) {
		t.Fatalf("expected clock to move forward: before=%v, after=%v", before, after)
	}
}

func TestClockPause(t *testing.T) {
	clock := NewClock(time.Now())
	go clock.RunTime()
	defer clock.StopTime()

	time.Sleep(1 * time.Second)

	clock.Pause()
	pausedAt := clock.GetTime()

	time.Sleep(2 * time.Second)

	after := clock.GetTime()

	if !pausedAt.Equal(after) {
		t.Fatalf(
			"expected clock to stay paused: pausedAt=%v, after=%v",
			pausedAt,
			after,
		)
	}
}

func TestClockAdd(t *testing.T) {
	clock := NewClock(time.Now())
	go clock.RunTime()
	defer clock.StopTime()

	// Ставим на паузу, чтобы фоновый тик не вмешался
	// между двумя GetTime().
	clock.Pause()

	before := clock.GetTime()

	clock.Add(30 * time.Second)

	after := clock.GetTime()

	diff := after.Sub(before)

	if diff != 30*time.Second {
		t.Fatalf("expected +30s, got %v", diff)
	}
}

func TestClockResume(t *testing.T) {
	clock := NewClock(time.Now())
	go clock.RunTime()
	defer clock.StopTime()

	clock.Pause()

	before := clock.GetTime()

	time.Sleep(1 * time.Second)

	// Пока стоит на паузе — время не должно двигаться.
	if !before.Equal(clock.GetTime()) {
		t.Fatalf("clock moved while paused")
	}

	clock.Resume()

	time.Sleep(2 * time.Second)

	after := clock.GetTime()

	if !after.After(before) {
		t.Fatalf(
			"expected clock to move after resume: before=%v, after=%v",
			before,
			after,
		)
	}
}
