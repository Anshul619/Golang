package main

type Addidas struct {
}

func (a *Addidas) makeShirt() IShirt {
	return &AdidasShirt{
		Shirt: Shirt{
			logo: "addidas",
			size: 14,
		},
	}
}

func (a *Addidas) makeShoe() IShoe {
	return &AdidasShoe{
		Shoe: Shoe{
			logo: "addidas",
			size: 12,
		},
	}
}
