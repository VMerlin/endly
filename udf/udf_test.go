package udf

import (
	"github.com/stretchr/testify/assert"
	"github.com/viant/endly/test/proto"
	"github.com/viant/toolbox"
	"github.com/viant/toolbox/data"
	"testing"
)

func Test_AsProtobufMessage(t *testing.T) {
	var input = "{\"id\":1, \"name\":\"abc\"}"
	encoded, err := AsProtobufMessage(input, nil, &proto.Message{})
	if assert.Nil(t, err) {
		assert.EqualValues(t, "base64:CAESA2FiYw==", encoded)
	}

	input = "{\"id\":101, \"name\":\"xyz\"}"
	encoded, err = AsProtobufMessage(input, nil, &proto.Message{})
	if assert.Nil(t, err) {
		assert.EqualValues(t, "base64:CGUSA3h5eg==", encoded)
	}

	var inputMap = map[string]interface{}{
		"id":   1,
		"name": "abc",
	}
	encoded, err = AsProtobufMessage(inputMap, nil, &proto.Message{})
	if assert.Nil(t, err) {
		assert.EqualValues(t, "base64:CAESA2FiYw==", encoded)
	}

	message, err := FromProtobufMessage(encoded, nil, &proto.Message{})
	if assert.Nil(t, err) {
		aMap := toolbox.AsMap(message)
		assert.EqualValues(t, 1, aMap["Id"])
		assert.EqualValues(t, "abc", aMap["Name"])
	}

}

func Test_AsProtobufMessage_Errors(t *testing.T) {
	{
		var input = "{id\":1, \"name\":\"abc\"}"
		_, err := AsProtobufMessage(input, nil, &proto.Message{})
		assert.NotNil(t, err)
	}
	{
		var input = ""
		_, err := AsProtobufMessage(input, nil, &proto.Message{})
		assert.NotNil(t, err)
	}
	{
		_, err := FromProtobufMessage("base64:CAErSA2FiYw==", nil, &proto.Message{})
		assert.NotNil(t, err)
	}
	{
		_, err := FromProtobufMessage("base64:12=", nil, &proto.Message{})
		assert.NotNil(t, err)
	}
}

func TestURLJoin(t *testing.T) {
	aMap := data.NewMap()
	aMap.Put("URLJoin", URLJoin)
	aMap.Put("baseURL", "mem://127.0.0.1/abc")
	aMap.Put("URI", "file1.txt")
	var expr = `$URLJoin(["$baseURL", "$URI"])`
	var expanded = aMap.Expand(expr)
	assert.EqualValues(t, "mem://127.0.0.1/abc/file1.txt", expanded)

}

func TestURLPath(t *testing.T) {
	aMap := data.NewMap()
	aMap.Put("URLPath", URLPath)
	aMap.Put("URL", "mem://127.0.0.1/abc")
	var expr = `$URLPath($URL)`
	var expanded = aMap.Expand(expr)
	assert.EqualValues(t, "/abc", expanded)
}
