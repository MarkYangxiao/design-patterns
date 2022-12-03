package concerte_factory

import (
	"DesignPattern/creational_pattern/factory/abstract_factory/abstract_product"
	"DesignPattern/creational_pattern/factory/abstract_factory/concrete_product"
)

type Nike struct {
}

func (nike *Nike) MakeShoes() abstract_product.IShoe {
	shoe := abstract_product.Shoe{}
	shoe.SetSize(15)
	shoe.SetLogo("nike")

	return &concrete_product.NikeShoe{
		Shoe: shoe,
	}
}

func (nike *Nike) MakeShirt() abstract_product.IShirt {
	shirt := abstract_product.Shirt{}
	shirt.SetLogo("nike")
	shirt.SetSize(11)

	return &concrete_product.NikeShirt{
		Shirt: shirt,
	}
}
