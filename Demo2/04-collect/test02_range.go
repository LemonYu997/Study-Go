package main

import "fmt"

// 声明一个包含 2 的幂次方的切片
var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

// range 用在 for 循环中迭代数组(array)、切片(slice)、通道(channel)或集合(map)的元素
// 在数组和切片中它返回元素的索引和索引对应的值，在集合中返回 key-value 对
func main() {
	// 遍历 pow 切片，i 是索引，v 是值
	for i, v := range pow {
		// 打印 2 的 i 次方等于 v
		fmt.Printf("2**%d = %d\n", i, v)
	}

	// 遍历字符串
	for i, c := range "hello" {
		fmt.Printf("index: %d, char: %c\n", i, c)
	}

	// 创建一个空的 map，key 是 int 类型，value 是 float32 类型
	map1 := make(map[int]float32)

	// 向 map1 中添加 key-value 对
	map1[1] = 1.0
	map1[2] = 2.0
	map1[3] = 3.0
	map1[4] = 4.0

	// 遍历 map1，读取 key 和 value
	for key, value := range map1 {
		// 打印 key 和 value
		fmt.Printf("key is: %d - value is: %f\n", key, value)
	}

	// 忽略值 遍历 map1，只读取 key
	for key := range map1 {
		// 打印 key
		fmt.Printf("key is: %d\n", key)
	}

	// 忽略索引 遍历 map1，只读取 value
	for _, value := range map1 {
		// 打印 value
		fmt.Printf("value is: %f\n", value)
	}

	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	close(ch)

	// 遍历通道
	for v := range ch {
		fmt.Println(v)
	}

	//这是我们使用 range 去求一个 slice 的和。使用数组跟这个很类似
	nums := []int{2, 3, 4}
	sum := 0
	// 忽略索引
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)
	//在数组上使用 range 将传入索引和值两个变量。上面那个例子我们不需要使用该元素的序号，所以我们使用空白符"_"省略了。有时侯我们确实需要知道它的索引。
	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}
	//range 也可以用在 map 的键值对上。
	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	//range也可以用来枚举 Unicode 字符串。第一个参数是字符的索引，第二个是字符（Unicode的值）本身。
	for i, c := range "go" {
		fmt.Println(i, c)
	}

}
