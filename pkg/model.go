package pkg

import "time"

type Header struct {
	Key   string
	Value []byte
}

type Message struct {
	Brokers       []string
	Topic         string
	Partition     int
	Offset        int64
	HighWaterMark int64
	Key           []byte
	Value         []byte
	Headers       []Header
	WriterData    interface{}
	Time          time.Time
}
