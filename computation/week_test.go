package computation

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestWeek(t *testing.T) {
	assert := assert.New(t)

	// return the only Wednesday between 7 days
	starting := time.Now()
	ending := time.Now().AddDate(0, 0, 7)

	times := Weekly(WEDNESDAY, starting, ending)

	assert.Equal(1, len(times))
	for i := range times {
		assert.Equal("Wednesday", times[i].Weekday().String())
	}

	// return the both Wednesdays between 14 days

	ending = time.Now().AddDate(0, 0, 14)

	times = Weekly(WEDNESDAY, starting, ending)

	assert.Equal(2, len(times))
	for i := range times {
		assert.Equal("Wednesday", times[i].Weekday().String())
	}

	assert.True(true)
}
