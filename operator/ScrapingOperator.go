package operator

import "fmt"

type ScrapingOperator struct {
}

func (o *ScrapingOperator) Operate() {
	fmt.Println("test")
}
