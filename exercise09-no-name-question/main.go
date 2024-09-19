package main

func AddElement(numbers *[]int, element int) {
	*numbers = append(*numbers, element)
}

func FindMin(numbers *[]int) int {
	min := 0
	if len(*numbers) == 0 {
		return min
	} else {
		min = (*numbers)[0]
		for _, i := range *numbers {
			if i < min {
				min = i
			}
		}
	}
	return min
}

func ReverseSlice(numbers *[]int) {
	for i, j := 0, len(*numbers)-1; i < j; i, j = i+1, j-1 {
		(*numbers)[i], (*numbers)[j] = (*numbers)[j], (*numbers)[i]
	}
}

func SwapElements(numbers *[]int, i, j int) {
    if i < 0 || j < 0 || i >= len(*numbers) || j >= len(*numbers) {
        return
    }
    (*numbers)[i], (*numbers)[j] = (*numbers)[j], (*numbers)[i]
}

