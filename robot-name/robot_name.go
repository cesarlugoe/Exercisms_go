package robotname

import (
	"fmt"
	"math/rand"
	"regexp"
	"time"
)

// Robot struct represents a Robot with a name
type Robot struct {
	name string
}

const abc = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
const maxNumberCode = 999
const minNumberCode = 100

var namePat = regexp.MustCompile(`^[A-Z]{2}\d{3}$`)

// Name creates and assings the name of the robot
func (r *Robot) Name() string {
	if !namePat.MatchString(r.name) {
		r.name = fmt.Sprintf("%v%v%d",
			string(abc[getRandomNumber(len(abc))]),
			string(abc[getRandomNumber(len(abc))]),
			getRandomNumberBetweenTwoNumbers(maxNumberCode, minNumberCode),
		)
	}

	return r.name
}

// Reset clears the name of the robot
func (r *Robot) Reset() {
	r.name = ""
}

func getRandomNumberBetweenTwoNumbers(max, min int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min) + min
}

func getRandomNumber(max int) int {
	return getRandomNumberBetweenTwoNumbers(max-1, 0)
}
