package allergies

var allergicValues = []struct {
	description string
	value       uint
}{
	{
		description: "eggs",
		value:       1,
	},
	{
		description: "peanuts",
		value:       2,
	},
	{
		description: "shellfish",
		value:       4,
	},
	{
		description: "strawberries",
		value:       8,
	},
	{
		description: "tomatoes",
		value:       16,
	},
	{
		description: "chocolate",
		value:       32,
	},
	{
		description: "pollen",
		value:       64,
	},
	{
		description: "cats",
		value:       128,
	},
}

// Allergies comment
func Allergies(s uint) []string {
	var res []string
	var analizedValue uint
	var valueIncluded bool
	var validAllergen bool
	analizedValue = 1024

	if s == 0 {
		return []string{}
	}

	for {
		valueIncluded = s >= analizedValue
		if valueIncluded {
			validAllergen = analizedValue/s <= 1
			if validAllergen {
				for _, v := range allergicValues {
					if v.value == analizedValue {
						res = append([]string{v.description}, res...)
					}
				}
			}
			s -= analizedValue
		}

		if s == 0 {
			return res
		}
		analizedValue = analizedValue / 2
	}
}

// AllergicTo comment
func AllergicTo(s uint, subs string) bool {
	allergies := Allergies(s)
	for _, a := range allergies {
		if a == subs {
			return true
		}
	}
	return false
}
