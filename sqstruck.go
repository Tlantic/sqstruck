package sqstruck

import (
	"encoding/binary"
	"errors"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sqs"
)

//SQSTruck ...
type SQSTruck struct {
	size  int
	store Store
}

//New ...
func New(size int, store Store) *SQSTruck {
	return &SQSTruck{
		size:  size,
		store: store,
	}
}

//PrepareWrite ...
func (st *SQSTruck) PrepareWrite(data *sqs.SendMessageBatchRequestEntry) (*sqs.SendMessageBatchRequestEntry, error) {
	sizeMessage := binary.Size([]byte(*data.MessageBody))
	hsize := 0
	for x, k := range data.MessageAttributes {

		hsize = hsize + binary.Size([]byte(x))
		hsize = hsize + binary.Size(k.BinaryValue)
	}

	total := hsize + sizeMessage
	id := data.MessageAttributes["_id"].GoString()
	if total > st.size {
		err := st.store.Set(id, []byte(*data.MessageBody))
		if err != nil {
			data.MessageAttributes[st.store.GetName()] = &sqs.MessageAttributeValue{
				DataType:    aws.String("String"),
				StringValue: aws.String(id),
			}

			data.SetMessageBody("")
		} else {
			return data, err
		}
	}

	return data, nil
}

//PrepareRead ...
func (st *SQSTruck) PrepareRead(message *sqs.Message) (*sqs.Message, error) {

	if s3Id, ok := message.MessageAttributes[st.store.GetName()]; ok {

		data, success := st.store.Get(s3Id.GoString())
		if !success {
			return message, errors.New("File not found")
		}

		message.SetBody(string(data))
	}

	return message, nil
}
