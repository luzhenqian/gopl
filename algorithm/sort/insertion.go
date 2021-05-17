package sort

func InsertionSort(a []int) {
	for i := 1; i < len(a); i++ {
		v := a[i]
		j := i - 1
		// 查找插入的位置
		for ; j >= 0; j-- {
			if v < a[j] {
				a[j+1] = a[j] // 移动数据
			} else {
				break
			}
		}
		a[j+1] = v //插入数据
	}
}
