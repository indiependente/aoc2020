package part2

func Sum2020(input []int) (int, error) {
	var (
		a, b, c int
	)
	for i := 0; i < len(input)/3; i++ {
		for j := i + 1; j < len(input)*2/3; j++ {
			for k := j + 1; k < len(input); k++ {
				if input[i]+input[j]+input[k] == 2020 {
					a = input[i]
					b = input[j]
					c = input[k]
				}
			}
		}
	}
	return a * b * c, nil
}
