package gross

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	return map[string]int{
		"quarter_of_a_dozen": 3,
		"half_of_a_dozen":    6,
		"dozen":              12,
		"small_gross":        120,
		"gross":              144,
		"great_gross":        1728,
	}
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	return map[string]int{}
}

func hasKey(m map[string]int, key string) bool {
	_, exist := m[key]
	return exist
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	if !hasKey(units, unit) {
		return false
	}

	if hasKey(bill, item) {
		bill[item] += units[unit]
	} else {
		bill[item] = units[unit]
	}
	return true
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	if !hasKey(bill, item) {
		return false
	}
	if !hasKey(units, unit) {
		return false
	}

	if bill[item] < units[unit] {
		return false
	} else if bill[item] == units[unit] {
		delete(bill, item)
	} else {
		bill[item] -= units[unit]
	}

	return true
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	v, exist := bill[item]

	return v, exist
}
