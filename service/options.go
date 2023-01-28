package service

type Option func(*Options)

type Options struct {
	WithoutCount bool
}

func optionWithoutCount(b bool) Option {
	return func(options *Options) {
		options.WithoutCount = b
	}
}
