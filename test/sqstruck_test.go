package sqstruck

import (
	"os"
	"testing"

	"github.com/Tlantic/sqstruck"
	strucks3 "github.com/Tlantic/sqstruck/s3"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sqs"
)

func Test_PrepareWrite(t *testing.T) {

	data := &sqs.SendMessageBatchRequestEntry{
		Id:                aws.String("1112222"),
		MessageBody:       aws.String(string("ola tudo bem")),
		MessageAttributes: make(map[string]*sqs.MessageAttributeValue),
	}

	data.MessageAttributes["_aggregateId"] = &sqs.MessageAttributeValue{
		DataType:    aws.String("String"),
		StringValue: aws.String("event.AggregateId222224444"),
	}

	data.MessageAttributes["_id"] = &sqs.MessageAttributeValue{
		DataType:    aws.String("String"),
		StringValue: aws.String("asdfghjkl"),
	}

	os.Setenv("AWS_ACCESS_KEY_ID", "asdsads")
	os.Setenv("AWS_SECRET_ACCESS_KEY", "asdas")

	st, _ := strucks3.New("s3://eu-west-1/sqstruckstore/")

	struck := sqstruck.New(10, st)
	struck.PrepareWrite(data)

	dataMsg := &sqs.Message{
		MessageAttributes: make(map[string]*sqs.MessageAttributeValue),
	}

	dataMsg.MessageAttributes["_s3"] = &sqs.MessageAttributeValue{
		DataType:    aws.String("String"),
		StringValue: aws.String("asdfghjkl"),
	}

	struck.PrepareWrite(data)

	msg, _ := struck.PrepareRead(dataMsg)

	t.Log(msg)
}
