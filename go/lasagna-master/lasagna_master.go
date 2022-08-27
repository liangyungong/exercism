package lasagna

func PreparationTime(layers []string, minute_per_layer int) int {

	if minute_per_layer == 0 {
		minute_per_layer = 2
	}

	return len(layers) * minute_per_layer
}

func Quantities(layers []string) (int, float64) {
	noodle_gram_per_layer := 50
	source_volume_per_layer := 0.2

	noodle_layer := 0
	source_layer := 0.0

	for _, layer := range layers {
		if layer == "noodles" {
			noodle_layer += 1
		} else if layer == "sauce" {
			source_layer += 1.0
		}
	}

	return noodle_gram_per_layer * noodle_layer,
		source_volume_per_layer * source_layer
}

func AddSecretIngredient(secret_ingredient []string, my_ingredient []string) {
	my_ingredient[len(my_ingredient)-1] = secret_ingredient[len(secret_ingredient)-1]
}

func ScaleRecipe(quantities []float64, count int) []float64 {
	var scaledQuantities []float64 = make([]float64, len(quantities), len(quantities))

	for i, v := range quantities {
		scaledQuantities[i] = v / 2 * float64(count)
	}

	return scaledQuantities
}
