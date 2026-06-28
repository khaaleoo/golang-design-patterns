# Golang Design Patterns

This repository contains simple examples of design patterns implemented in Go (Golang). It is a resource to help developers understand and apply design patterns in real-world projects using Go.

## Structure

| Module | Patterns |
|--------|----------|
| [creational-patterns](./creational-patterns) | Singleton, Builder, Factory Method, Abstract Factory, Prototype |
| [structural-patterns](./structural-patterns) | Adapter, Bridge, Composite, Decorator, Facade, Flyweight, Proxy |
| [behavioral-patterns](./behavioral-patterns) | Strategy, Observer, Command, State, Chain of Responsibility |

## Implemented patterns

### Creational (5/5)

| Pattern | Example | Path |
|---------|---------|------|
| **Singleton** | Thread-safe counter with `sync.Once` | [`creational-patterns/singleton`](./creational-patterns/singleton) |
| **Builder** | Manufacturing director builds bicycles/motorbikes step by step | [`creational-patterns/builder`](./creational-patterns/builder) |
| **Factory Method** | Coffee bar creates drinks (Cappuccino, Latte) by name | [`creational-patterns/factory-method`](./creational-patterns/factory-method) |
| **Abstract Factory** | Bicycle and motorbike factories produce related vehicle families | [`creational-patterns/abstract-factory`](./creational-patterns/abstract-factory) |
| **Prototype** | Clone pages from shared layout templates (main, blank) | [`creational-patterns/prototype`](./creational-patterns/prototype) |

### Structural (7/7)

| Pattern | Example | Path |
|---------|---------|------|
| **Adapter** | Unify `Fetch` and `Axios` HTTP clients behind one interface | [`structural-patterns/adapter`](./structural-patterns/adapter) |
| **Bridge** | MacOS/Windows computers delegate printing to Epson/HP printers | [`structural-patterns/bridge`](./structural-patterns/bridge) |
| **Composite** | File/folder tree with uniform `Print` on leaves and containers | [`structural-patterns/composite`](./structural-patterns/composite) |
| **Decorator** | Logging and retry layers wrapped around an HTTP client | [`structural-patterns/decorator`](./structural-patterns/decorator) |
| **Facade** | Simplified storage API over the composite file system | [`structural-patterns/facade`](./structural-patterns/facade) |
| **Flyweight** | Forest of thousands of trees sharing a small pool of `TreeType` objects | [`structural-patterns/flyweight`](./structural-patterns/flyweight) |
| **Proxy** | Lazy HTTP client initialization and LRU cache for user lookup | [`structural-patterns/proxy`](./structural-patterns/proxy) |

### Behavioral (5/11)

| Pattern | Example | Path |
|---------|---------|------|
| **Strategy** | Coffee shop checkout with regular, happy hour, and VIP pricing | [`behavioral-patterns/strategy`](./behavioral-patterns/strategy) |
| **Observer** | Coffee shop notifies barista and cashier when an order is placed | [`behavioral-patterns/observer`](./behavioral-patterns/observer) |
| **Command** | Manufacturing remote executes and undoes build commands | [`behavioral-patterns/command`](./behavioral-patterns/command) |
| **State** | Order lifecycle transitions (pending → preparing → ready) | [`behavioral-patterns/state`](./behavioral-patterns/state) |
| **Chain of Responsibility** | HTTP middleware chain: auth → rate limit → fetch | [`behavioral-patterns/chain`](./behavioral-patterns/chain) |

## Run examples

```bash
cd creational-patterns && go run .
cd structural-patterns && go run .
cd behavioral-patterns && go run .
```

## Pattern map (GoF)

| Category | Implemented | Not yet in repo |
|----------|-------------|-----------------|
| Creational | 5/5 | — |
| Structural | 7/7 | — |
| Behavioral | 5/11 | Iterator, Mediator, Memento, Template Method, Visitor, Interpreter |

## Suggested next patterns

- **Template Method** — shared beverage preparation steps in the coffee shop theme
- **Iterator** — walk the composite file tree without exposing internals
- **Memento** — snapshot and restore manufacturing command state
- **Functional Options** — flexible constructors for builder/prototype

## Contributing

Contributions are welcome! Add examples, improve existing code, or expand explanations.

Happy coding!
