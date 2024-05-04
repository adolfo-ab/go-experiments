package main

import "fmt"

type CoffeeMachine struct {
	beanAmount   float32
	grinderLevel int
	waterTemp    int
	waterAmt     float32
	milkAmount   float32
	addFoam      bool
}

func (c *CoffeeMachine) startCoffee(beanAmount float32, grind int) {
	c.beanAmount = beanAmount
	c.grinderLevel = grind
	fmt.Println("Starting coffee order with beands: ", beanAmount, "and grind level: ", grind)
}

func (c *CoffeeMachine) endCoffee() {
	fmt.Println("Ending coffee order")
}

func (c *CoffeeMachine) grindBeans() bool {
	fmt.Println("Grinding the beans: ", c.beanAmount, "at level", c.grinderLevel)
	return true
}

func (c *CoffeeMachine) useMilk(amount float32) float32 {
	fmt.Println("Adding milk:", amount)
	c.milkAmount = amount
	return amount
}

func (c *CoffeeMachine) setWaterTemp(temp int) {
	fmt.Println("Setting water temp: ", temp)
	c.waterTemp = temp
}

func (c *CoffeeMachine) useHotWater(amount float32) float32 {
	fmt.Println("Adding hot water: ", amount)
	c.waterAmt = amount
	return amount
}

func (c *CoffeeMachine) doFoam(foam bool) {
	c.addFoam = foam
}
