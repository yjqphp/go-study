package main

/**
	映射的文法（续）

	若顶级类型只是一个类型名，你可以再文法的元素中省略它
 */
import "fmt"

type Vertex struct {
	Lat,Long float64
}

var m = map[string]Vertex{
	"Bell Labs": {40.68433, -74.39967},
	"Google": {37.42202, -122.08408},
}

func main()  {
	fmt.Println(m)
}
