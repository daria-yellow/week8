package orderbook

type Orderbook struct {
	Ask []*Order
	Bid []*Order
}

func New() *Orderbook {
	return &Orderbook{Ask: *&[]*Order{}, Bid: *&[]*Order{}}
}

func (orderbook *Orderbook) MinAsk() *Order {
	minask := uint64(1000000000000000000)
	minOrder := &Order{}
	flag := false
	for _, i := range orderbook.Ask {
		if i.Price <= minask {
			minask = i.Price
			minOrder = i
			flag = true
		}
	}
	if !flag {
		return nil
	}
	return minOrder
}

func (orderbook *Orderbook) MinLimitedAsk(price uint64) *Order {
	minask := uint64(1000000000000000000)
	minOrder := &Order{}
	flag := false
	for _, i := range orderbook.Ask {
		if i.Price < minask && i.Price <= price {
			minask = i.Price
			minOrder = i
			flag = true
		}
	}
	if !flag {
		return nil
	}
	return minOrder
}

func (orderbook *Orderbook) MaxBid() *Order {
	maxbid := uint64(0)
	maxOrder := &Order{}
	flag := false
	for _, i := range orderbook.Bid {
		if i.Price > maxbid {
			maxbid = i.Price
			maxOrder = i
			flag = true
		}
	}
	if !flag {
		return nil
	}
	return maxOrder
}

func (orderbook *Orderbook) MaxLimitedBid(price uint64) *Order {
	maxbid := uint64(0)
	maxOrder := &Order{}
	flag := false
	for _, i := range orderbook.Bid {
		if i.Price > maxbid && i.Price >= price {
			maxbid = i.Price
			maxOrder = i
			flag = true
		}
	}
	if !flag {
		return nil
	}
	return maxOrder
}

func (orderbook *Orderbook) Delete(order *Order) *Orderbook {
	new_orderbook := New()
	if order.Side == SideBid {
		for _, i := range orderbook.Bid {
			if i != order {
				new_orderbook.Bid = append(new_orderbook.Bid, i)
			}
		}
	}

	if order.Side == SideAsk {
		for _, i := range orderbook.Ask {
			if i != order {
				new_orderbook.Ask = append(new_orderbook.Ask, i)
			}
		}
	}

	return new_orderbook
}

func (orderbook *Orderbook) Match(order *Order) ([]*Trade, *Order) {
	var trade []*Trade
	if order.Side == SideBid {
		for order.Volume != 0 && len(orderbook.Ask) != 0 {
			if order.Kind == KindMarket {
				i := orderbook.MinAsk()
				if i == nil {
					orderbook.Bid = append(orderbook.Bid, order)
					return nil, nil
				}

				if i.Volume > order.Volume {
					trade = append(trade, &Trade{order, i, order.Volume, i.Price})
					i.Volume -= order.Volume
					return trade, nil
				} else if i.Volume == order.Volume {
					trade = append(trade, &Trade{order, i, order.Volume, i.Price})
					orderbook = orderbook.Delete(i)
					return trade, nil
				} else {
					trade = append(trade, &Trade{order, i, i.Volume, i.Price})
					order.Volume -= i.Volume
					orderbook = orderbook.Delete(i)
				}

			}

			if order.Kind == KindLimit {
				i := orderbook.MinLimitedAsk(order.Price)
				if i == nil {
					orderbook.Bid = append(orderbook.Bid, order)
					return nil, nil
				}
				if i.Volume > order.Volume {
					trade = append(trade, &Trade{order, i, order.Volume, i.Price})
					i.Volume -= order.Volume
					return trade, nil
				} else if i.Volume == order.Volume {
					orderbook = orderbook.Delete(i)
					trade = append(trade, &Trade{i, order, order.Volume, i.Price})
					orderbook = New()
					return trade, nil
				} else {
					trade = append(trade, &Trade{order, i, i.Volume, i.Price})
					order.Volume -= i.Volume
					orderbook = orderbook.Delete(i)
				}

			}

		}

		if len(orderbook.Ask) == 0 {
			orderbook.Bid = append(orderbook.Bid, order)
			if order.Kind == KindMarket {
				return trade, order
			}
		}
	}

	if order.Side == SideAsk {
		for order.Volume != 0 && len(orderbook.Bid) != 0 {

			if order.Kind == KindMarket {
				i := orderbook.MaxBid()
				if i == nil {
					orderbook.Ask = append(orderbook.Ask, order)
					return nil, nil
				}
				if i.Volume > order.Volume {
					trade = append(trade, &Trade{order, i, order.Volume, i.Price})
					i.Volume -= order.Volume
					return trade, nil
				} else if i.Volume == order.Volume {
					trade = append(trade, &Trade{order, i, order.Volume, i.Price})
					orderbook = orderbook.Delete(i)
					return trade, nil
				} else {
					trade = append(trade, &Trade{order, i, i.Volume, i.Price})
					order.Volume -= i.Volume
					orderbook = orderbook.Delete(i)
				}

			}

			if order.Kind == KindLimit {
				i := orderbook.MaxLimitedBid(order.Price)
				if i == nil {
					orderbook.Ask = append(orderbook.Ask, order)
					return nil, nil
				}

				if i.Volume > order.Volume {
					trade = append(trade, &Trade{order, i, order.Volume, i.Price})
					i.Volume -= order.Volume
					return trade, nil
				} else if i.Volume == order.Volume {
					trade = append(trade, &Trade{order, i, order.Volume, i.Price})
					orderbook = orderbook.Delete(i)
					return trade, nil
				} else {
					trade = append(trade, &Trade{order, i, i.Volume, i.Price})
					order.Volume -= i.Volume
					orderbook = orderbook.Delete(i)
				}

			}

		}

		if len(orderbook.Bid) == 0 {
			orderbook.Ask = append(orderbook.Ask, order)
			if order.Kind == KindMarket {
				return trade, order
			}
		}
	}
	return trade, nil
}
