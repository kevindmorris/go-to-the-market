# 📑 Go to the Market

> A simple and lightweight CLI to-do list manager written in Go. Easily add, list, and mark tasks as complete with persistent JSON storage and minimal dependencies.

## ⚙️ Requirements

- Go 1.18 or later

Check your Go version:

```bash
go version
```

## 📥 Installation

Clone this repository:

```bash
git clone https://github.com/kevindmorris/go-to-the-market.git
cd go-to-the-market
```

## 🚀 Usage

- Add a task

```bash
go run . add "Buy groceries"
```

- List all tasks

```bash
go run . ls
```

- Mark a task as complete

```bash
go run . complete 1
```

## 🏭 Production

Build the executable:

```bash
go build -o todo
```

Use the application:

```bash
./todo add "Buy groceries"
./todo ls
./todo complete 1
```

## 📝 License

MIT License © 2025 Kevin Morris
