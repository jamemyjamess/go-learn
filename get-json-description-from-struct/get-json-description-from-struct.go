package main

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"
	"time"
)

type ChildReqLogin struct {
	ChildA string `json:"child_a" info:"type:'string' description:'ลูก A'"`
	ChildB string `json:"child_b" info:"type:'string' description:'ลูก B'"`
}

type ReqLogin struct {
	Username       string          `json:"username" info:"type:'string' description:'บัญชีผู้ใช้งาน'"`
	Password       string          `json:"password" info:"type:'string' description:'รหัสผ่านผู้ใช้งาน'"`
	No             int             `json:"no" info:"type:'int' description:'ลำดับ'"`
	Option         []string        `json:"option" info:"type:'array-string' description:'ตัวเลือก'"`
	CreatedAt      *time.Time      `json:"created_at" info:"type:'datetime' description:'วันที่สร้าง'"`
	ChildReqLogin1 []ChildReqLogin `json:"child_req_login1" info:"type:'array-object' description:'ลูกของ ReqLogin'"`
	ChildReqLogin  []ChildReqLogin `json:"child_req_login2" info:"type:'array-object' description:'ลูกของ ReqLogin'"`
}

type BodyJson struct {
	ReqLogin
}

func assignStruct(t reflect.Type, mapBodyJson *(map[string]interface{})) {
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		keyJson := field.Tag.Get("json")
		info := field.Tag.Get("info")
		if keyJson == "" {
			keyJson = field.Name
		}

		fmt.Println("keyjson: ", keyJson, "kind: ", field.Type.Kind(), "type:", field.Type.String(), "field.Type.Elem().String(): ", field.Type.Elem().String())
		if field.Type.Kind() == reflect.Struct {
			mapBodyJsonChild := make(map[string]interface{})
			mapBodyJsonChild["info"] = fmt.Sprintf("%s %s", field.Type.String(), info)
			assignStruct(field.Type, &mapBodyJsonChild)
			(*mapBodyJson)[keyJson] = mapBodyJsonChild
			continue
		}
		//fmt.Println("field: ", fiel	d.Name, "info: ", info)
		(*mapBodyJson)[keyJson] = fmt.Sprintf("%s %s", field.Type.String(), info)
	}
}

func main2() {
	//reflect.ValueOf(&bodyJson).Elem().FieldByName("B").Set(reflect.ValueOf(a))
	// fmt.Println(*t.B)
	mapBodyJson := make(map[string]interface{})
	assignStruct(reflect.TypeOf(&ReqLogin{}).Elem(), &mapBodyJson)
	// fmt.Println(mapBodyJson)
	b, err := json.Marshal(mapBodyJson)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(b))

}
