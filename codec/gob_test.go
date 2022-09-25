package codec

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

type TestObj struct {
	Name string
	Age  int
}

func TestGob(t *testing.T) {
	should := assert.New(t)
	obj := &TestObj{
		Name: "ldd",
		Age:  23,
	}
	gobBytes, err := GobEncode(obj)
	if should.NoError(err) {
		fmt.Println(gobBytes)
	}
	var testObj = &TestObj{}
	err = GobDecode(gobBytes, testObj)
	if should.NoError(err) {
		fmt.Println(testObj.Name, testObj.Age)
	}
}
