package timetable

import (
	assert "github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestListOfInit(t *testing.T) {
	assert := assert.New(t)
	tt := ListOf(7)

	// Set Number of Maximum output
	assert.Equal(7, tt.count)

	// Set Start date as today to start with
	// assert.Equal(time.Now().Day(), tt.start.Day())

	assert.True(true)
}

func TestStartingFrom(t *testing.T) {
	assert := assert.New(t)

	// Set custom start date
	tt := ListOf(7).StartingFrom(time.Now().AddDate(0, 0, 2))
	assert.Equal(time.Now().AddDate(0, 0, 2).Day(), tt.start.time.Day())
}
