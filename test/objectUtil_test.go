package test

import (
	"fmt"
	"github.com/simonalong/tools/util"
	"testing"
)

type ValueInnerEntity1 struct {
	Name string
	Age  int
}

func TestObjectToYaml1(t *testing.T) {
	inner1 := map[string]interface{}{}
	inner1["name"] = "inner_1"
	inner1["age"] = 1

	var targetObj ValueInnerEntity1
	util.MapToObject(inner1, &targetObj)
	Equal(t, "{\"Name\":\"inner_1\",\"Age\":1}", util.ToJsonString(targetObj))
}

type ValueInnerEntity2 struct {
	Name   string
	Age    int
	Inner1 ValueInnerEntity1
}

func TestObjectToYaml2(t *testing.T) {
	inner1 := map[string]interface{}{}
	inner1["name"] = "inner_1"
	inner1["age"] = 1

	inner2 := map[string]interface{}{}
	inner2["name"] = "inner_2"
	inner2["age"] = 2
	inner2["inner1"] = inner1

	var targetObj ValueInnerEntity2
	util.MapToObject(inner2, &targetObj)
	Equal(t, "{\"Name\":\"inner_2\",\"Age\":2,\"Inner1\":{\"Name\":\"inner_1\",\"Age\":1}}", util.ToJsonString(targetObj))
}

type ValueInnerEntity3 struct {
	Name   string
	Age    int
	Inner2 ValueInnerEntity2
}

func TestObjectToYaml3(t *testing.T) {
	inner1 := map[string]interface{}{}
	inner1["name"] = "inner_1"
	inner1["age"] = 1

	inner2 := map[string]interface{}{}
	inner2["name"] = "inner_2"
	inner2["age"] = 2
	inner2["inner1"] = inner1

	inner3 := map[string]interface{}{}
	inner3["name"] = "inner_3"
	inner3["age"] = 3
	inner3["inner2"] = inner2

	var targetObj ValueInnerEntity3
	util.MapToObject(inner3, &targetObj)
	Equal(t, "{\"Name\":\"inner_3\",\"Age\":3,\"Inner2\":{\"Name\":\"inner_2\",\"Age\":2,\"Inner1\":{\"Name\":\"inner_1\",\"Age\":1}}}", util.ToJsonString(targetObj))
}

type ValueInnerEntity4 struct {
	Name    string
	Age     int
	DataMap map[string]string
}

func TestObjectToYaml4(t *testing.T) {
	kvMap := map[string]interface{}{}
	kvMap["k1"] = "name1"
	kvMap["k2"] = "name2"

	inner1 := map[string]interface{}{}
	inner1["name"] = "inner_1"
	inner1["age"] = 1
	inner1["dataMap"] = kvMap

	var targetObj ValueInnerEntity4
	util.MapToObject(inner1, &targetObj)
	Equal(t, "{\"Name\":\"inner_1\",\"Age\":1,\"DataMap\":{\"k1\":\"name1\",\"k2\":\"name2\"}}", util.ToJsonString(targetObj))
}

type ValueInnerEntity5 struct {
	Name    string
	Age     int
	DataMap map[string]ValueInnerEntity1
}

func TestObjectToYaml5(t *testing.T) {
	v1 := map[string]interface{}{}
	v1["name"] = "inner_1"
	v1["age"] = 1

	v2 := map[string]interface{}{}
	v2["name"] = "inner_2"
	v2["age"] = 2

	kvMap := map[string]interface{}{}
	kvMap["k1"] = v1
	kvMap["k2"] = v2

	inner1 := map[string]interface{}{}
	inner1["name"] = "inner_1"
	inner1["age"] = 1
	inner1["dataMap"] = kvMap

	var targetObj ValueInnerEntity5
	util.MapToObject(inner1, &targetObj)
	Equal(t, "{\"Name\":\"inner_1\",\"Age\":1,\"DataMap\":{\"k1\":{\"Name\":\"inner_1\",\"Age\":1},\"k2\":{\"Name\":\"inner_2\",\"Age\":2}}}", util.ToJsonString(targetObj))
}

type ValueInnerEntity6 struct {
	Name    string
	Age     int
	DataMap map[string][]int
}

func TestObjectToYaml6(t *testing.T) {
	var dataList []int
	dataList = append(dataList, 12)
	dataList = append(dataList, 13)

	kvMap := map[string]interface{}{}
	kvMap["k1"] = dataList
	kvMap["k2"] = dataList

	inner1 := map[string]interface{}{}
	inner1["name"] = "inner_1"
	inner1["age"] = 1
	inner1["dataMap"] = kvMap

	var targetObj ValueInnerEntity6
	util.MapToObject(inner1, &targetObj)
	Equal(t, "{\"Name\":\"inner_1\",\"Age\":1,\"DataMap\":{\"k1\":[12,13],\"k2\":[12,13]}}", util.ToJsonString(targetObj))
}

type ValueInnerEntity7 struct {
	Name    string
	Age     int
	DataMap map[string][]ValueInnerEntity1
}

func TestObjectToYaml7(t *testing.T) {
	var dataList []ValueInnerEntity1
	dataList = append(dataList, ValueInnerEntity1{Name: "name1", Age: 1})
	dataList = append(dataList, ValueInnerEntity1{Name: "name2", Age: 2})

	kvMap := map[string]interface{}{}
	kvMap["k1"] = dataList
	kvMap["k2"] = dataList

	inner1 := map[string]interface{}{}
	inner1["name"] = "inner_1"
	inner1["age"] = 1
	inner1["dataMap"] = kvMap

	var targetObj ValueInnerEntity7
	util.MapToObject(inner1, &targetObj)
	Equal(t, "{\"Name\":\"inner_1\",\"Age\":1,\"DataMap\":{\"k1\":[{\"Name\":\"name1\",\"Age\":1},{\"Name\":\"name2\",\"Age\":2}],\"k2\":[{\"Name\":\"name1\",\"Age\":1},{\"Name\":\"name2\",\"Age\":2}]}}", util.ToJsonString(targetObj))
}

type ValueInnerEntity1Tem struct {
	Name    string
	Address string
}

type ValueInnerEntity8 struct {
	Name    string
	Age     int
	DataMap map[string][]ValueInnerEntity1Tem
}

func TestObjectToYaml8(t *testing.T) {
	var dataList []ValueInnerEntity1
	dataList = append(dataList, ValueInnerEntity1{Name: "name1", Age: 1})
	dataList = append(dataList, ValueInnerEntity1{Name: "name2", Age: 2})

	kvMap := map[string]interface{}{}
	kvMap["k1"] = dataList
	kvMap["k2"] = dataList

	inner1 := map[string]interface{}{}
	inner1["name"] = "inner_1"
	inner1["age"] = 1
	inner1["dataMap"] = kvMap

	var targetObj ValueInnerEntity8
	util.MapToObject(inner1, &targetObj)
	fmt.Println(util.ToJsonString(targetObj))
}