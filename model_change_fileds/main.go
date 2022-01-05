package main

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
)

type DBClusterHandler struct {
	ClusterName string `json:"cluster_name" validate:"required"`
	ClusterType string `json:"cluster_type" validate:"required"` // 集群类型，1、MM  2、MS  3、MyCat 4、DRDS  5ADG 6RAC
	ClusterDesc string `json:"cluster_desc"`                     // 描述信息
}

func (DBClusterHandler) TableName() string {
	return "db_cluster"
}
func (db DBClusterHandler) Print() {
	fmt.Println(db)
}

func (db DBClusterHandler) Set(ClusterName string, ClusterType string, ClusterDesc string) {
	db.ClusterName = ClusterName
	db.ClusterType = ClusterType
	db.ClusterDesc = ClusterDesc
}

func getStructJsonTagField(tableName string, obj interface{}) (string, error) {
	// 首先判断obj 是否为nil
	var (
		objError    = errors.New("接口对象不能为nil")
		objPtrError = errors.New("接口对象参数必须是是指针接口")
		strError    = errors.New("表名不能为空")
		result      string
	)
	if obj == nil {
		return "", objError
	}
	if tableName == "" {
		return "", strError
	}
	// 判断obj是否是指针
	objType := reflect.TypeOf(obj)
	if objType.Kind() == reflect.Ptr {
		s := reflect.TypeOf(obj).Elem()
		for i := 0; i < s.NumField(); i++ {
			jsonTag := s.Field(i).Tag.Get("json")
			newJosnTag := fmt.Sprintf("%s.%s,", tableName, jsonTag)
			result += newJosnTag
		}
	} else {
		return "", objPtrError
	}
	newResult := strings.Trim(result, ",")
	return newResult, nil
}

func main() {
	test := &DBClusterHandler{}
	reuslt, err := getStructJsonTagField("db_cluster", test)
	if err != nil {
		fmt.Println("error: %s", err)
		return
	}
	fmt.Println("result: %s", reuslt)

}
