package course

import "fmt"

type course struct {
	name    string
	Price   float64
	IsFree  bool
	UserIDs []uint
	classes map[uint]string
}

func (c *course) SetName(name string) { c.name = name }
func (c *course) Name() string        { return c.name }

func (c *course) SetClasses(classes map[uint]string) {
	c.classes = classes
}

func NewCourse(name string, price float64, isFree bool) *course {
	if price == 0 {
		price = 30
	}

	return &course{
		name:   name,
		Price:  price,
		IsFree: isFree,
	}
}

func (c *course) changePrice(price float64) {
	c.Price = price
}

func (c *course) PrintClasses() {
	text := "Las clases son: "
	for _, class := range c.classes {
		text += class + ", "
	}

	fmt.Println(text[:len(text)-2])
}
