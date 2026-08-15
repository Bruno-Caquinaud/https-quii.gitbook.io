package mocking

import (
	"bytes"
	"os"
	"testing"
	"time"
)

const (
	sleepCall = "sleep"
	writeCall = "write"
)

type DefaultSleeper struct {
	Calls int
}

func (d *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
	d.Calls++
}

type SpySleeper struct {
	Calls                    int
	register                 *bytes.Buffer
	durationMillisecondSleep int
}

func (s *SpySleeper) Sleep() {
	s.Calls++
	time.Sleep(time.Duration(s.durationMillisecondSleep) * time.Millisecond)
	s.register.Write([]byte(sleepCall))
}

type SpyWritter struct {
	Buffer   bytes.Buffer
	Calls    int
	register *bytes.Buffer
}

func (s *SpyWritter) Write(p []byte) (n int, err error) {
	n, err = s.Buffer.Write(p)
	s.Calls++
	s.register.Write([]byte(writeCall))
	return
}

func TestCountDown(t *testing.T) {
	t.Run("Test CountDown 1 : Print a value on current output", func(t *testing.T) {
		testBufferRegister := bytes.Buffer{}
		spySleeper := SpySleeper{durationMillisecondSleep: 200, register: &testBufferRegister}
		spyWritter := SpyWritter{register: &testBufferRegister}
		countdownStarting := 3
		countdownWritting := countdownStarting + 1

		CountDown(&spyWritter, &spySleeper, countdownStarting)
		got := spyWritter.Buffer.String()
		want := "3\n2\n1\nGo!\n"

		assertInt(t, spySleeper.Calls, countdownStarting)
		assertInt(t, spyWritter.Calls, countdownWritting)
		assertString(t, got, want)

		got = testBufferRegister.String()
		initTestCallMethodCountdownRegister(&want, countdownStarting)

		assertString(t, got, want)
	})

	t.Run("Test Coutdown 2 : Use Stdout as Output", func(t *testing.T) {
		timeSleeper := DefaultSleeper{}
		countdownStarting := 3
		caracterNumber, _ := CountDown(os.Stdout, &timeSleeper, countdownStarting)

		want := len("3\n2\n1\nGo!\n")
		assertInt(t, timeSleeper.Calls, countdownStarting)
		assertInt(t, caracterNumber, want)
	})
}

func assertInt(t *testing.T, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("want %d != got %d", want, got)
	}
}

func assertString(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("want %s != got %s", want, got)
	}
}

func initTestCallMethodCountdownRegister(expectedCallbuffer *string, countdownStarting int) {
	*expectedCallbuffer = ""
	for i := countdownStarting; i > 0; i-- {
		*expectedCallbuffer += writeCall + sleepCall
	}
	*expectedCallbuffer += writeCall
}
