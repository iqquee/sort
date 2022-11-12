package x

type IntMapper struct {
	Key       int `json:"key"`
	Occurance int `json:"occurance"`
}

func SortInt(data []int) []IntMapper {
	var occuranceCounter int
	var dataObjArray []IntMapper
	var dataObj IntMapper
	for _, v := range data {
		for _, occurance := range data {
			if v == occurance {
				newOccurance := occuranceCounter + 1
				occuranceCounter = newOccurance
			}
		}

		if dataObjArray == nil {
			dataObj.Key = v
			dataObj.Occurance = occuranceCounter
			dataObjArray = append(dataObjArray, dataObj)
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
			}
		}

		//reset the counter
		occuranceCounter = 0
	}

	return dataObjArray

}

type StringMapper struct {
	Key       string `json:"key"`
	Occurance int    `json:"occurance"`
}

func SortString(data []string) []StringMapper {
	var occuranceCounter int
	var dataObjArray []StringMapper
	var dataObj StringMapper
	// var check []int
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
			}
		}

		//reset the counter
		occuranceCounter = 0
	}

	return dataObjArray
}
