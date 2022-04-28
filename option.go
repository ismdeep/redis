package redis

import "time"

// Option struct
type Option struct {
	ExpireDuration *time.Duration
	Encoder        *string // JSON, GOB
}

// EncoderJSON json encoder
const EncoderJSON = "json"

// EncoderGOB gob encoder
const EncoderGOB = "gob"

// WithExpireTime with expire time
func WithExpireTime(duration time.Duration) Option {
	return Option{
		ExpireDuration: &duration,
	}
}

// with encoder
func withEncoder(str string) Option {
	return Option{Encoder: &str}
}

// WithDefaultEncoder with default encoder
func WithDefaultEncoder() Option {
	return withEncoder(EncoderGOB)
}

// WithJSONEncoder with json encoder
func WithJSONEncoder() Option {
	return withEncoder(EncoderJSON)
}

// WithGOBEncoder with gob encoder
func WithGOBEncoder() Option {
	return withEncoder(EncoderGOB)
}

func getOption(opts ...Option) Option {
	o := Option{}
	for _, v := range opts {
		if v.ExpireDuration != nil {
			o.ExpireDuration = v.ExpireDuration
		}
		if v.Encoder != nil {
			o.Encoder = v.Encoder
		}
	}
	return o
}
