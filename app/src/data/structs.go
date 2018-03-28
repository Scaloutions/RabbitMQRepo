package data

type (
	QueueConfig struct {
		Name       string
		Durable    bool
		AutoDelete bool
		Exclusive  bool
		NoWait     bool
	}

	Quote struct {
		Stock     string
		Price     float64
		CrytoKey  string
		Timestamp int64
		UserID    string
	}
)
