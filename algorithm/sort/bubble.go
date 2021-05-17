package sort

func BubbleSort(a []int) {
	for i, _ := range a {
		flag := false // 提前退出循环的标志位
		for j := 0; j < len(a)-i-1; j++ {
			if a[j] > a[j+1] { // 交换数据
				tmp := a[j]
				a[j] = a[j+1]
				a[j+1] = tmp
				flag = true // 表示有数据交换
			}
		}
		if !flag {
			return // 没有数据交换，提前退出
		}
	}
}
