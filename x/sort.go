package x

import (
	"errors"
)

type IntMapper struct {
	Key       int `json:"key"`
	Occurance int `json:"occurance"`
}

func SortInt(data []int, lenght int) ([]IntMapper, []IntMapper, error) {
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

	sortedValues := sort(values)

	var sortedObj []IntMapper
	//check if the lenght of sorted value is not more than the number of the sorted value array
	if len(sortedValues) > lenght {
		// get the value fom the sorted values and retrieve the key from the obj
		for i := 0; i < lenght; i++ {
			for _, x := range dataObjArray {
				if x.Occurance == sortedValues[i] {
					sortedObj = append(sortedObj, x)
				}
			}
		}
	} else {
		return nil, nil, errors.New("sort times is greater than the numbers in array")
	}

	return dataObjArray, sortedObj, nil
}

type StringMapper struct {
	Key       string `json:"key"`
	Occurance int    `json:"occurance"`
}

func SortString(data []string, lenght int) ([]StringMapper, []StringMapper, error) {
	var occuranceCounter int
	var dataObjArray []StringMapper
	var dataObj StringMapper
	var values []int
	for _, v := range data {

		//loop through the array and check the number of occurance of the digit
		for _, occurance := range data {
			//check if the current value is same as the top value and increment the number
			if v == occurance {
				newOccurance := occuranceCounter + 1
				occuranceCounter = newOccurance
			}
		}

		if dataObjArray == nil {
			//create a new object for the digit and its number of occurance
			dataObj.Key = v
			dataObj.Occurance = occuranceCounter
			//append to the dataObjArray
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
				//append to the dataObjArray
				dataObjArray = append(dataObjArray, dataObj)

				// append it the key to the values arry for further sorting
				values = append(values, occuranceCounter)
			}
		}

		//reset the counter
		occuranceCounter = 0
	}

	sortedValues := sort(values)

	var sortedObj []StringMapper
	//check if the lenght of sorted value is not more than the number of the sorted value array
	if len(sortedValues) > lenght {
		// get the value fom the sorted values and retrieve the key from the obj
		for i := 0; i < lenght; i++ {
			for _, x := range dataObjArray {
				if x.Occurance == sortedValues[i] {
					sortedObj = append(sortedObj, x)
				}
			}
		}
	} else {
		return nil, nil, errors.New("sort times is greater than the numbers in array")
	}

	return dataObjArray, sortedObj, nil
}

func sort(arr []int) []int {
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
