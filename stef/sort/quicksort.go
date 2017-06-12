package sort

//IsSorted check if sorted
func IsSorted(a []int) bool {
	for i := 0; i < len(a)-1; i++ {
		if a[i] > a[i+1] {
			return false
		}
	}
	return true
}

//QuickSort sort inpute int slice
func QuickSort(a []int) {
	if len(a) == 0 {
		return
	}
	quickSortLR(a, 0, len(a)-1)
	// quickSortBad(a)
}

//在输入已经部分排好序或者重复元素很多时，性能很差
func quickSortBad(a []int) {
	size := len(a)
	if size <= 1 {
		return
	}
	// a[0], a[size/2] = a[size/2], a[0]
	cmp := a[0]
	last := 0
	for i := 0; i < size; i++ {
		if a[i] < cmp {
			last++
			a[last], a[i] = a[i], a[last]
		}
	}
	a[0], a[last] = a[last], a[0]
	// fmt.Println(last)
	quickSortBad(a[:last])
	quickSortBad(a[last+1:])
	return
}

//针对各种输入，性能都有保障
func quickSortLR(a []int, left, right int) {
	if right <= left { //在调用前比较，跟快些
		return
	}
	mid := left + (right-left)/2
	a[left], a[mid] = a[mid], a[left] //避免数组已经排好序，切分到第一个或最后一个元素
	cmp := a[left]                    //特殊比较基准选择第一个元素
	i, j := left, right+1
	for {
		//相等元素也交换，避免重复元素太多时性能下降
		for i++; a[i] < cmp; i++ {
			if i == right {
				break
			}
		}
		//move to lefmost >cmp
		for j--; a[j] > cmp; j-- {
			if j == left {
				break
			}
		}
		if i >= j {
			break
		}
		a[j], a[i] = a[i], a[j]
	}
	a[j], a[left] = a[left], a[j]
	// if i - 1 > left {
	// }
	quickSortLR(a, left, j-1)
	// if i + 1 < right {
	// }
	quickSortLR(a, j+1, right)

	return
}
