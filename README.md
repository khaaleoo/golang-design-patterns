# Golang Design Patterns

This repository contains simple examples of design patterns implemented in Go (Golang). It is a resource to help developers understand and apply design patterns in real-world projects using Go.

## Structure

| Module | Patterns |
|--------|----------|
| [creational-patterns](./creational-patterns) | Singleton, Builder, Factory Method, Abstract Factory, Prototype |
| [structural-patterns](./structural-patterns) | Adapter, Bridge, Composite, Decorator, Facade, Proxy |
| [behavioral-patterns](./behavioral-patterns) | Strategy, Observer, Command, State, Chain of Responsibility |

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
| Structural | 6/7 | Flyweight |
| Behavioral | 5/11 | Iterator, Mediator, Memento, Template Method, Visitor, Interpreter |

## Suggested next patterns

- **Flyweight** — share layout/theme data across prototype pages
- **Template Method** — shared beverage preparation steps
- **Iterator** — walk composite file tree
- **Functional Options** — flexible constructors for builder/prototype

## Contributing

Contributions are welcome! Add examples, improve existing code, or expand explanations.

Happy coding!
