package s3

import (
	"os"
	"testing"
)

func Test_SetFile(t *testing.T) {

	os.Setenv("AWS_ACCESS_KEY_ID", "asdas")
	os.Setenv("AWS_SECRET_ACCESS_KEY", "asdsada")

	st, err := New("s3://eu-west-1/sqstruckstore/")

	st.Set("file", []byte("Here is a string...."))

	t.Log("Result:", st, err)

}

func Test_LoadFile(t *testing.T) {

	os.Setenv("AWS_ACCESS_KEY_ID", "asdsad")
	os.Setenv("AWS_SECRET_ACCESS_KEY", "asdsadas")

	st, err := New("s3://eu-west-1/sqstruckstore/")

	file, rs := st.Get("file")

	t.Log("Result:", file, rs, err)

}
