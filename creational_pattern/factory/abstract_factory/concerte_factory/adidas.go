package concerte_factory

import (
	"DesignPattern/creational_pattern/factory/abstract_factory/abstract_product"
	"DesignPattern/creational_pattern/factory/abstract_factory/concrete_product"
)

type Adidas struct {
}

func (ad *Adidas) MakeShoes() abstract_product.IShoe {
	shoe := abstract_product.Shoe{}
	shoe.SetSize(14)
	shoe.SetLogo("adidas")

	return &concrete_product.AdidasShoe{
		Shoe: shoe,
	}
}

func (ad *Adidas) MakeShirt() abstract_product.IShirt {
	shirt := abstract_product.Shirt{}
	shirt.SetLogo("adidas")
	shirt.SetSize(10)

	return &concrete_product.AdidasShirt{
		Shirt: shirt,
	}
}
