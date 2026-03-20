package naijafakergo

import (
	"math/rand"
	"time"
)

type Faker struct {
	language string
	gender   string
	network  string
	rng      *rand.Rand
}

type Config struct {
	Language string
	Gender   string
	Network  string
}

var languages = []string{"yoruba", "igbo", "hausa"}
var genders = []string{"male", "female"}

func New() *Faker {
	return &Faker{
		rng: rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

func (f *Faker) Configure(cfg Config) {
	if cfg.Language != "" {
		f.language = cfg.Language
	}
	if cfg.Gender != "" {
		f.gender = cfg.Gender
	}
	if cfg.Network != "" {
		f.network = cfg.Network
	}
}

func (f *Faker) Seed(seed int64) {
	if seed == 0 {
		f.rng = rand.New(rand.NewSource(time.Now().UnixNano()))
		return
	}
	f.rng = rand.New(rand.NewSource(seed))
}

// random returns a random int in [0, n)
func (f *Faker) random(n int) int {
	return f.rng.Intn(n)
}

// pick returns a random element from a string slice
func (f *Faker) pick(slice []string) string {
	return slice[f.random(len(slice))]
}

func (f *Faker) resolveLanguage(lang string) string {
	if lang != "" {
		return lang
	}
	if f.language != "" {
		return f.language
	}
	return f.pick(languages)
}

func (f *Faker) resolveGender(gender string) string {
	if gender != "" {
		return gender
	}
	if f.gender != "" {
		return f.gender
	}
	return f.pick(genders)
}
