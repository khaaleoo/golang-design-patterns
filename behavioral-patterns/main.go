package main

import (
	"fmt"

	"github.com/behavioral-patterns/chain"
	"github.com/behavioral-patterns/command"
	"github.com/behavioral-patterns/interpreter"
	"github.com/behavioral-patterns/iterator"
	"github.com/behavioral-patterns/mediator"
	"github.com/behavioral-patterns/memento"
	"github.com/behavioral-patterns/observer"
	"github.com/behavioral-patterns/state"
	"github.com/behavioral-patterns/strategy"
	"github.com/behavioral-patterns/templatemethod"
	"github.com/behavioral-patterns/visitor"
)

func main() {
	/*
		Example Strategy
	*/
	fmt.Println("*** Example Strategy ***")
	checkout := strategy.NewCheckout(strategy.RegularPricing{})
	basePrice := 5.0
	fmt.Printf("Regular total: %.2f (%s)\n", checkout.Total(basePrice), checkout.StrategyName())

	checkout.SetStrategy(strategy.HappyHourPricing{})
	fmt.Printf("Happy hour total: %.2f (%s)\n", checkout.Total(basePrice), checkout.StrategyName())

	checkout.SetStrategy(strategy.VIPPricing{})
	fmt.Printf("VIP total: %.2f (%s)\n", checkout.Total(basePrice), checkout.StrategyName())
	fmt.Print("*** End of Strategy ***\n\n\n")

	/*
		Example Observer
	*/
	fmt.Println("*** Example Observer ***")
	shop := observer.NewCoffeeShop()
	shop.Subscribe(&observer.Barista{Name: "Anna"})
	shop.Subscribe(&observer.Cashier{Name: "Ben"})
	shop.PlaceOrder("ORD-001", "Latte")
	fmt.Print("*** End of Observer ***\n\n\n")

	/*
		Example Command
	*/
	fmt.Println("*** Example Command ***")
	ctx := command.NewManufacturingContext()
	remote := command.NewManufacturingRemote()

	fmt.Println(remote.Run(command.NewAddWheelsCommand(ctx)))
	fmt.Println(remote.Run(command.NewPaintFrameCommand(ctx)))
	fmt.Println("Current parts:", ctx.Snapshot())
	fmt.Println(remote.UndoLast())
	fmt.Println("Current parts:", ctx.Snapshot())
	fmt.Print("*** End of Command ***\n\n\n")

	/*
		Example State
	*/
	fmt.Println("*** Example State ***")
	order := state.NewOrder("ORD-002", "Cappuccino")
	fmt.Println("Status:", order.Status())
	fmt.Println(order.Advance())
	fmt.Println("Status:", order.Status())
	fmt.Println(order.Advance())
	fmt.Println("Status:", order.Status())
	fmt.Print("*** End of State ***\n\n\n")

	/*
		Example Chain of Responsibility
	*/
	fmt.Println("*** Example Chain of Responsibility ***")
	auth := chain.NewAuthHandler("secret-token")
	rateLimit := chain.NewRateLimitHandler(2)
	fetch := chain.NewFetchHandler()

	auth.SetNext(rateLimit).SetNext(fetch)

	response, err := auth.Handle("https://example.com")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Response:", response)
	}

	_, err = chain.NewAuthHandler("").Handle("https://example.com")
	if err != nil {
		fmt.Println("Auth error:", err)
	}
	fmt.Print("*** End of Chain of Responsibility ***\n\n\n")

	/*
		Example Iterator
	*/
	fmt.Println("*** Example Iterator ***")
	menu := iterator.NewMenu()
	menu.AddItem("Espresso", 3.00)
	menu.AddItem("Latte", 4.50)
	menu.AddItem("Croissant", 2.75)

	menuIterator := menu.CreateIterator()
	for menuIterator.HasNext() {
		item := menuIterator.Next()
		fmt.Printf("%s: %.2f\n", item.Name, item.Price)
	}
	fmt.Print("*** End of Iterator ***\n\n\n")

	/*
		Example Mediator
	*/
	fmt.Println("*** Example Mediator ***")
	coffeeShopMediator := mediator.NewCoffeeShopMediator()
	cashier := mediator.NewCashier("Mia", coffeeShopMediator)
	barista := mediator.NewBarista("Leo", coffeeShopMediator)
	coffeeShopMediator.SetCashier(cashier)
	coffeeShopMediator.SetBarista(barista)

	fmt.Println(cashier.TakeOrder("Americano"))
	fmt.Println(barista.FinishDrink("Americano"))
	fmt.Print("*** End of Mediator ***\n\n\n")

	/*
		Example Memento
	*/
	fmt.Println("*** Example Memento ***")
	editor := memento.NewOrderEditor("Latte", "Medium")
	history := memento.NewHistory()
	history.Backup(editor)

	editor.SetDrinkName("Mocha")
	editor.SetSize("Large")
	fmt.Println("Current order:", editor.Snapshot())
	fmt.Println(history.Undo(editor))
	fmt.Print("*** End of Memento ***\n\n\n")

	/*
		Example Template Method
	*/
	fmt.Println("*** Example Template Method ***")
	for _, step := range templatemethod.Prepare(templatemethod.Coffee{}) {
		fmt.Println(step)
	}
	for _, step := range templatemethod.Prepare(templatemethod.Tea{}) {
		fmt.Println(step)
	}
	fmt.Print("*** End of Template Method ***\n\n\n")

	/*
		Example Visitor
	*/
	fmt.Println("*** Example Visitor ***")
	menuItems := []visitor.MenuItem{
		visitor.NewDrink("Cappuccino", 4.25),
		visitor.NewPastry("Muffin", 3.10),
	}
	labelVisitor := visitor.LabelVisitor{}
	priceVisitor := visitor.PriceVisitor{}

	for _, item := range menuItems {
		fmt.Println(item.Accept(labelVisitor))
		fmt.Println(item.Accept(priceVisitor))
	}
	fmt.Print("*** End of Visitor ***\n\n\n")

	/*
		Example Interpreter
	*/
	fmt.Println("*** Example Interpreter ***")
	orderTotal := interpreter.NewSubtractExpression(
		interpreter.NewAddExpression(
			interpreter.NewNumberExpression(5.00),
			interpreter.NewNumberExpression(2.00),
		),
		interpreter.NewNumberExpression(1.50),
	)
	fmt.Printf("Order total: %.2f\n", orderTotal.Interpret())
	fmt.Print("*** End of Interpreter ***\n\n\n")
}
