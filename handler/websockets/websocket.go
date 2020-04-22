package websockets

type Receive struct {
	Content string `validate:"required"`
	Ploy    int64  `validate:"required"`
}
