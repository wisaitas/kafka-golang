package pkg

import "github.com/segmentio/kafka-go"

func mappingHeader(headers []Header) []kafka.Header {
	headersKafka := make([]kafka.Header, len(headers))
	for i, header := range headers {
		headersKafka[i] = kafka.Header{
			Key:   header.Key,
			Value: header.Value,
		}
	}

	return headersKafka
}

func mappingMessage(message Message) kafka.Message {
	return kafka.Message{
		Key:           message.Key,
		Value:         message.Value,
		Partition:     message.Partition,
		Offset:        message.Offset,
		HighWaterMark: message.HighWaterMark,
		Headers:       mappingHeader(message.Headers),
		WriterData:    message.WriterData,
		Time:          message.Time,
	}
}
