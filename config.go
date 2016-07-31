package experiment

import "golang.org/x/net/context"

var (
	TestMode = false
)

type (
	// Config is the configuration used to set up an experiment. Config is not
	// safe for concurrent use.
	Config struct {
		Name       string  `json:"name"`
		Percentage float32 `json:"percentage"`

		BeforeFilters []BeforeFilter
		Publishers    []ResultPublisher
		Comparison    ComparisonMethod
	}

	// BeforeFilter is a wrapper around a method that is purely used to take a
	// context, adjust it and return a new context with the adjusted values.
	BeforeFilter func(context.Context) context.Context

	// ComparisonMethod is used as an interface for creating a method in which we
	// want to compare the observations of a test. This being the Control and a
	// random Test case.
	ComparisonMethod func(Observation, Observation) bool
)

// DefaultConfig sets up a default configuration where the Percentage is 100.
func DefaultConfig(name string) *Config {
	return &Config{
		Name:       name,
		Percentage: 100,
	}
}

func (c *Config) AddBeforeFilter(bf BeforeFilter) {
	if c.BeforeFilters == nil {
		c.BeforeFilters = []BeforeFilter{bf}
	} else {
		c.BeforeFilters = append(c.BeforeFilters, bf)
	}
}

func (c *Config) AddPublisher(rp ResultPublisher) {
	if c.Publishers == nil {
		c.Publishers = []ResultPublisher{rp}
	} else {
		c.Publishers = append(c.Publishers, rp)
	}
}

func (c *Config) SetComparisonMethod(cm ComparisonMethod) {
	c.Comparison = cm
}