package sort

type IntMapper struct {
	Key       int `json:"key"`
	Occurance int `json:"occurance"`
}

func SortInt(data []int, lenght int) []IntMapper {
	var occuranceCounter int
	var dataObjArray []IntMapper
	var dataObj IntMapper
	var values []int
	for _, v := range data {
		for _, occurance := range data {
			if v == occurance {
				occuranceCounter += 1
			}
		}

		if dataObjArray == nil {
			dataObj.Key = v
			dataObj.Occurance = occuranceCounter
			dataObjArray = append(dataObjArray, dataObj)

			// append it the key to the values arry for further sorting
			values = append(values, occuranceCounter)
		} else {
			var check bool
			for _, avail := range dataObjArray {
				if avail.Key == v {
					check = true
				}
			}

			if !check {
				dataObj.Key = v
				dataObj.Occurance = occuranceCounter
				dataObjArray = append(dataObjArray, dataObj)

				// append it the key to the values arry for further sorting
				values = append(values, occuranceCounter)
			}
		}

		//reset the counter
		occuranceCounter = 0
	}

	sortedValues := sorter(values)

	var sortedObj []IntMapper
	var newSortedObj []IntMapper
	//check if the lenght of sorted value is not more than the number of the sorted value array

	if lenght == 0 {

		for i := 0; i < len(sortedValues); i++ {
			for _, x := range dataObjArray {

				if sortedObj == nil {
					if x.Occurance == sortedValues[i] {
						sortedObj = append(sortedObj, x)
					}
				} else {
					var available bool
					for _, avail := range sortedObj {
						if avail.Key == x.Key {
							available = true
						}
					}

					if !available {
						if x.Occurance == sortedValues[i] {
							sortedObj = append(sortedObj, x)
						}
					}
				}

			}
		}

	} else if len(sortedValues) > lenght {

		for i := 0; i < lenght; i++ {

			for _, x := range dataObjArray {

				if newSortedObj == nil {
					if x.Occurance == sortedValues[i] {
						newSortedObj = append(newSortedObj, x)
					}
				} else {
					var available bool
					for _, avail := range newSortedObj {
						if avail.Key == x.Key {
							available = true
						}
					}

					if !available {
						if x.Occurance == sortedValues[i] {
							newSortedObj = append(newSortedObj, x)
						}
					}
				}

			}
		}

		for i := 0; i < lenght; i++ {
			sortedObj = append(sortedObj, newSortedObj[i])
		}

	}

	return sortedObj
}

type StringMapper struct {
	Key       string `json:"key"`
	Occurance int    `json:"occurance"`
}

func SortString(data []string, lenght int) []StringMapper {
	var occuranceCounter int
	var dataObjArray []StringMapper
	var dataObj StringMapper
	var values []int
	for _, v := range data {
		for _, occurance := range data {
			if v == occurance {
				occuranceCounter += 1
			}
		}

		if dataObjArray == nil {
			dataObj.Key = v
			dataObj.Occurance = occuranceCounter
			dataObjArray = append(dataObjArray, dataObj)

			// append it the key to the values arry for further sorting
			values = append(values, occuranceCounter)
		} else {
			var check bool
			for _, avail := range dataObjArray {
				if avail.Key == v {
					check = true
				}
			}

			if !check {
				dataObj.Key = v
				dataObj.Occurance = occuranceCounter
				dataObjArray = append(dataObjArray, dataObj)

				// append it the key to the values arry for further sorting
				values = append(values, occuranceCounter)
			}
		}

		//reset the counter
		occuranceCounter = 0
	}

	sortedValues := sorter(values)

	var sortedObj []StringMapper
	var newSortedObj []StringMapper
	//check if the lenght of sorted value is not more than the number of the sorted value array

	if lenght == 0 {

		for i := 0; i < len(sortedValues); i++ {
			for _, x := range dataObjArray {

				if sortedObj == nil {
					if x.Occurance == sortedValues[i] {
						sortedObj = append(sortedObj, x)
					}
				} else {
					var available bool
					for _, avail := range sortedObj {
						if avail.Key == x.Key {
							available = true
						}
					}

					if !available {
						if x.Occurance == sortedValues[i] {
							sortedObj = append(sortedObj, x)
						}
					}
				}

			}
		}

	} else if len(sortedValues) > lenght {

		for i := 0; i < lenght; i++ {

			for _, x := range dataObjArray {

				if newSortedObj == nil {
					if x.Occurance == sortedValues[i] {
						newSortedObj = append(newSortedObj, x)
					}
				} else {
					var available bool
					for _, avail := range newSortedObj {
						if avail.Key == x.Key {
							available = true
						}
					}

					if !available {
						if x.Occurance == sortedValues[i] {
							newSortedObj = append(newSortedObj, x)
						}
					}
				}

			}
		}

		for i := 0; i < lenght; i++ {
			sortedObj = append(sortedObj, newSortedObj[i])
		}

	}

	return sortedObj
}

func sorter(arr []int) []int {
	swapped := true

	for swapped {
		swapped = false
		for i := 0; i < len(arr)-1; i++ {
			if arr[i] < arr[i+1] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
				swapped = true
			}
		}
	}

	return arr
}
