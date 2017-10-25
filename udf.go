package neatly

import (
	"bytes"
	"crypto/md5"
	"fmt"
	"github.com/viant/toolbox"
	"github.com/viant/toolbox/data"
	"io"
	"path"
	"strings"
)

//AsMap converts source into map
func AsMap(source interface{}, state data.Map) (interface{}, error) {
	if source == nil || toolbox.IsMap(source) {
		return source, nil
	}
	if toolbox.IsString(source) {
		buf := new(bytes.Buffer)
		err := toolbox.NewJSONEncoderFactory().Create(buf).Encode(toolbox.AsString(source))
		if err != nil {
			return nil, err
		}
		aMap := make(map[string]interface{})
		err = toolbox.NewJSONDecoderFactory().Create(buf).Decode(aMap)
		if err != nil {
			return nil, err
		}
		return aMap, nil

	}
	return source, nil
}

//AsInt converts source into int
func AsInt(source interface{}, state data.Map) (interface{}, error) {
	return toolbox.AsInt(source), nil
}

//AsFloat converts source into float64
func AsFloat(source interface{}, state data.Map) (interface{}, error) {
	return toolbox.AsFloat(source), nil
}

//AsBool converts source into bool
func AsBool(source interface{}, state data.Map) (interface{}, error) {
	return toolbox.AsBoolean(source), nil
}

//Md5 computes source md5
func Md5(source interface{}, state data.Map) (interface{}, error) {
	hash := md5.New()
	_, err := io.WriteString(hash, toolbox.AsString(source))
	if err != nil {
		return nil, err
	}
	var result =  fmt.Sprintf("%x", hash.Sum(nil))
	return result, nil
}

//HasResource check if patg/url to external resource exists
func HasResource(source interface{}, state data.Map) (interface{}, error) {
	var parentDirecotry = ""
	if state.Has(OwnerURL) {
		var workflowPath = strings.Replace(state.GetString(OwnerURL), toolbox.FileSchema, "", 1)
		parentDirecotry, _ = path.Split(workflowPath)
	}



	filename := path.Join(parentDirecotry, toolbox.AsString(source))
	var result = toolbox.FileExists(filename)
	return result, nil
}