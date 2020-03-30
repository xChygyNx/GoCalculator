package solution

var DefaultConfig = Config{
	Accuracy: 5,
}

type Config struct {
	Accuracy int
}

func (c *Config) SetDefaults() {
loop:
	for {
		switch {
		default:
			break loop
		}
	}
}
func (c Config) Validate() error {
	return nil
}
