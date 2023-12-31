package usecase;

import "github.com/motA1215/goLang/internal/entity";

type OrderInput struct{
	ID string;
	Price float64;
	Taxa float64;
}

type OrderOutput struct{
	ID string;
	Price float64;
	Taxa float64;
	FinalPrice float64;
}

type CalculateFinalPrice struct{
	OrderRepository entity.OrderRepositoryInterface;
}

func NewCalculateFinalPrice(orderRepository entity.OrderRepositoryInterface) *CalculateFinalPrice{
	return &CalculateFinalPrice{
        OrderRepository: orderRepository,
    }
}

func (c * CalculateFinalPrice) Execute(input OrderInput) (*OrderOutput, error){
	order, err := entity.NewOrder(input.ID, input.Price, input.Taxa);

	if err != nil{
		return nil, err;
	}

	err = order.CalculateFinalPrice();

	if err != nil{
		return nil, err;
	}

	err = c.OrderRepository.Save(order);
	if err != nil{
		return nil, err;
	}

	return &OrderOutput{
		ID: order.ID,
		Price: order.Price,
		Taxa: order.Taxa,
		FinalPrice: order.FinalPrice,
	},
	nil;
}

func (c *CalculateFinalPrice) GetTotalRowsTable() (int, error){
	totalRows, err := c.OrderRepository.GetTotalTransaction();
	if err != nil{
		return 0, err;
	}

	return totalRows, nil;
}