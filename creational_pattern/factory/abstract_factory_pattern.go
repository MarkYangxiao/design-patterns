package main

import (
	"DesignPattern/creational_pattern/factory/abstract_factory/abstract_product"
	"DesignPattern/creational_pattern/factory/abstract_factory/concerte_factory"
	"errors"
	"fmt"
)

// 抽象工厂
type ISportsFactory interface {
	MakeShoes() abstract_product.IShoe
	MakeShirt() abstract_product.IShirt
}

func getSportsFactory(brand string) (ISportsFactory, error) {
	if brand == "adidas" {
		return &concerte_factory.Adidas{}, nil
	} else if brand == "nike" {
		return &concerte_factory.Nike{}, nil
	} else {
		return nil, errors.New("can not support this brand")
	}
}

func main() {
	nikeFactory, _ := getSportsFactory("nike")
	adidasFactory, _ := getSportsFactory("adidas")

	nikeShirt := nikeFactory.MakeShirt()
	nikeShoe := nikeFactory.MakeShoes()

	adidasShirt := adidasFactory.MakeShirt()
	adidasShoe := adidasFactory.MakeShoes()

	printShoe(nikeShoe)
	printShoe(adidasShoe)

	printShirt(nikeShirt)
	printShirt(adidasShirt)

}

func printShoe(s abstract_product.IShoe) {
	fmt.Println("shoe logo:", s.GetLogo())
	fmt.Println("shoe size:", s.GetSize())
}

func printShirt(s abstract_product.IShirt) {
	fmt.Println("shirt logo:", s.GetLogo())
	fmt.Println("shirt size:", s.GetSize())
}
