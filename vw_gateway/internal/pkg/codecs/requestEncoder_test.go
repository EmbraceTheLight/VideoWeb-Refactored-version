package codecs

import (
	"github.com/spewerspew/spew"
	"github.com/stretchr/testify/require"
	"testing"
)

// test bindFormDataByFieldName，根据结构体字段来绑定form表单数据
func TestBindForm(t *testing.T) {

	// 定义一个测试结构体
	type TestStruct struct {
		Username string `json:"userName"`
		Age      int    `json:"aGe"`
		Active   bool   `json:"aCtiVe"`
	}

	testStructFieldName := formDataValue{

		"Username": []string{"test_fieldName"},
		"Age":      []string{"30"},
		"Active":   []string{"false"},
	}

	testStructFieldJSONTag := formDataValue{
		"userName": []string{"test_jsonTag"},
		"aGe":      []string{"20"},
		"aCtiVe":   []string{"true"},
	}
	ret := &TestStruct{}
	err := bindFormDataByFieldName(testStructFieldName, ret)
	require.NoError(t, err)
	spew.Dump(ret)
	require.Equal(t, ret, &TestStruct{
		Username: "test_fieldName",
		Age:      30,
		Active:   false,
	})

	ret = &TestStruct{}
	err = bindFormDataByFieldName(testStructFieldJSONTag, ret)
	require.NoError(t, err)
	spew.Dump(ret)
	require.NotEqual(t, ret, &TestStruct{
		Username: "test_jsonTag",
		Age:      20,
		Active:   true,
	})

}

// test bindFormDataByJSONTag，根据结构体json标签来绑定form表单数据
func TestBindForm2(t *testing.T) {

	// 定义一个测试结构体
	type TestStruct struct {
		Username string `json:"userName,omitempty"`
		Age      int    `json:"aGe,omitempty"`
		Active   bool   `json:"aCtiVe,omitempty"`
	}
	testStructFieldName := formDataValue{

		"Username": []string{"test_fieldName"},
		"Age":      []string{"30"},
		"Active":   []string{"false"},
	}

	testStructFieldJSONTag := formDataValue{
		"userName": []string{"test_jsonTag"},
		"aGe":      []string{"20"},
		"aCtiVe":   []string{"true"},
	}
	ret := &TestStruct{}
	err := bindFormDataByJSONTag(testStructFieldName, ret)
	require.NoError(t, err)
	spew.Dump(ret)
	require.NotEqual(t, ret, &TestStruct{
		Username: "test_fieldName",
		Age:      30,
		Active:   false,
	})

	ret = &TestStruct{}
	err = bindFormDataByJSONTag(testStructFieldJSONTag, ret)
	require.NoError(t, err)
	spew.Dump(ret)
	require.Equal(t, ret, &TestStruct{
		Username: "test_jsonTag",
		Age:      20,
		Active:   true,
	})

}
