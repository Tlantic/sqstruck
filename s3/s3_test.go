package s3

import (
	"os"
	"testing"
)

func Test_SetFile(t *testing.T) {

	os.Setenv("AWS_ACCESS_KEY_ID", "AKIAIPCCIXIL4Q64V63A")
	os.Setenv("AWS_SECRET_ACCESS_KEY", "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")

	st, err := New("s3://eu-west-1/sqstruckstore/")

	st.Set("file", []byte("Here is a string...."))

	t.Log("Result:", st, err)

}

func Test_LoadFile(t *testing.T) {

	os.Setenv("AWS_ACCESS_KEY_ID", "AKIAIPCCIXIL4Q64V63A")
	os.Setenv("AWS_SECRET_ACCESS_KEY", "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")

	st, err := New("s3://eu-west-1/sqstruckstore/")

	file, rs := st.Get("file")

	t.Log("Result:", file, rs, err)

}
