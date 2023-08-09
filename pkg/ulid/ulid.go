package ulid

import (
	"crypto/rand"
	"strings"
	"sync"
	"time"

	"github.com/oklog/ulid/v2"
)

const Length = 26

var (
	once     sync.Once
	instance *generator
)

type generator struct {
	timestamp uint64
	entropy   *ulid.MonotonicEntropy
}

func New() string {
	u := initGenerator()

	return strings.ToLower(
		ulid.MustNew(u.timestamp, u.entropy).String())
}

func Verify(s string) error {
	_, err := ulid.Parse(s)
	return err
}

func initGenerator() *generator {
	once.Do(func() {
		instance = &generator{
			timestamp: ulid.Timestamp(time.Now()),
			entropy:   ulid.Monotonic(rand.Reader, 0),
		}
	})

	return instance
}
