package main

import (
	"errors"
	"fmt"
)

// 产品接口

type Phone interface {
	Call(number string)
}

// 具体产品

type Apple struct {
}

func (a *Apple) Call(number string) {
	fmt.Println("apple call:", number)
}

type HuaWei struct {
}

func (hw *HuaWei) Call(number string) {
	fmt.Println("huawei call:", number)
}

// 工厂
func getPhone(company string) (Phone, error) {
	if company == "apple" {
		return &Apple{}, nil
	} else if company == "huawei" {
		return &HuaWei{}, nil
	} else {
		return nil, errors.New("can not support company")
	}
}

func main() {
	phone, err := getPhone("apple")
	if err == nil {
		phone.Call("11111")
	} else {
		fmt.Println(err.Error())
	}

	phone, err = getPhone("huawei")
	if err == nil {
		phone.Call("33333")
	} else {
		fmt.Println(err.Error())
	}

	phone, err = getPhone("xiaomi")
	if err == nil {
		phone.Call("00000")
	} else {
		fmt.Println(err.Error())
	}
}
