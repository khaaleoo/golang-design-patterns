package main

import (
	"fmt"

	"github.com/behavioral-patterns/chain"
	"github.com/behavioral-patterns/command"
	"github.com/behavioral-patterns/observer"
	"github.com/behavioral-patterns/state"
	"github.com/behavioral-patterns/strategy"
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
}
