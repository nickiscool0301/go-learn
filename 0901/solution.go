package main

type StockSpanner struct {
	prices []int
	spans  [][]int
}

func Constructor() StockSpanner {
	return StockSpanner{}
}

func (this *StockSpanner) Next(price int) int {
	span := 1
	for len(this.spans) != 0 && this.spans[len(this.spans)-1][0] <= price {
		span += this.spans[len(this.spans)-1][1]
		this.spans = this.spans[:len(this.spans)-1]
	}
	this.spans = append(this.spans, []int{price, span})
	return span
}

/**
 * Your StockSpanner object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Next(price);
 */
