package main

import (
	"errors"
	"testing"
	"time"
)

const wrongHostMock = "0.beevik-ntp.pool.ntp.oRGG"

type TimeMock struct {
	curTime time.Time
	err     error
}

type Test struct {
	name           string
	input          TimeMock
	expectedOutput TimeMock
}

func Test_GetHostTime(t *testing.T) {
	hostTime, _ := getHostTime(hostName)
	_, err := getHostTime(wrongHostMock)
	sysTime := time.Now()

	okTest := Test{
		name: "Ok",
		input: TimeMock{
			curTime: hostTime,
			err:     nil,
		},
		expectedOutput: TimeMock{
			curTime: sysTime, // предполагаем, что ожидаемое время (системное время) будет правильным, от которого и будем отталкиваться
			err:     nil,
		},
	}

	t.Run(okTest.name, func(t *testing.T) {
		if okTest.input.curTime.Sub(okTest.expectedOutput.curTime) > time.Second*1 {
			t.Errorf("got %s, want %s", okTest.input.curTime, okTest.expectedOutput.curTime)
			t.Error(err)
		}
	})

	wrongHostTest := Test{
		name: "WrongHost",
		input: TimeMock{
			curTime: time.Time{},
			err:     err,
		},
		expectedOutput: TimeMock{
			curTime: time.Time{},
			err:     errors.New("lookup 0.beevik-ntp.pool.ntp.oRGG: no such host"),
		},
	}

	t.Run(wrongHostTest.name, func(t *testing.T) {
		if wrongHostTest.input.err.Error() != wrongHostTest.expectedOutput.err.Error() {
			t.Errorf("got %s, want %s", wrongHostTest.input.err.Error(), wrongHostTest.expectedOutput.err.Error())
			t.Error(err)
		}
	})

}
