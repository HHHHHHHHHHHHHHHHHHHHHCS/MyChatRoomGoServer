package SQLHelper

import (
	"math/rand"
	"strconv"
	"testing"
	"time"
)

func Test_CreateAccount(t *testing.T) {
	CreateAccount("qq", "ee")
	rand.Seed(time.Now().UnixNano())
	CreateAccount(strconv.FormatInt(int64(rand.Intn(10000)),10), strconv.FormatInt(int64(rand.Intn(10000)),10))
}
