package gross

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	units := map[string]int{
		"quarter_of_a_dozen": 3,
		"half_of_a_dozen":    6,
		"dozen":              12,
		"small_gross":        120,
		"gross":              144,
		"great_gross":        1728,
	}

	return units
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	return map[string]int{}
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	score, isUnitValid := units[unit]
	if !isUnitValid {
		return false
	}

	_, isItemPresent := bill[item]

	if isItemPresent {
		bill[item] += score
	} else {
		bill[item] = score
	}
	return true

}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	score, isUnitValid := units[unit]
	_, isItemPresent := bill[item]
	newQuantity := bill[item] - score
	if !isItemPresent || !isUnitValid || newQuantity < 0 {
		return false
	}
	if newQuantity == 0 {
		delete(bill, item)
	} else {
		bill[item] -= score
	}
	return true
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	qty, ok := bill[item]
	return qty, ok
}
