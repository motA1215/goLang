package entity;

import "errors";

type Order struct{
	ID string
	Price float64
	Taxa float64
	FinalPrice float64
}

func NewOrder(id string, price float64, taxa float64) (*Order, error){
	order:= &Order{
		ID: id,
		Price: price,
		Taxa: taxa,
	}

	err := order.Validate();

	if err != nil{
		return nil, err;
	}

	return order, nil;
}

func (o *Order) Validate() error{
	if (o.ID == ""){
		return errors.New("ID is required");
	}

	if (o.Price <= 0){
		return errors.New("Price must be greater than zero");
	}

	if (o.Taxa <= 0){
		return errors.New("Invalid Taxa");
	}

	return nil;
}

func (o *Order) CalculateFinalPrice() error{
	o.FinalPrice = (o.Price + o.Taxa);
	err := o.Validate();

	if (err != nil){
		return err;
	}
	
	return nil;
}