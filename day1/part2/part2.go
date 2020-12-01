package part2

/*
--- Part Two ---
The Elves in accounting are thankful for your help; one of them even offers you a starfish coin they had left over from a past vacation. They offer you a second one if you can find three numbers in your expense report that meet the same criteria.

Using the above example again, the three entries that sum to 2020 are 979, 366, and 675. Multiplying them together produces the answer, 241861950.

In your expense report, what is the product of the three entries that sum to 2020?

*/

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
