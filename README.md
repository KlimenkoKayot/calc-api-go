## Микросервис для подсчета простых математических выражений

Реализовано как API, все запросы обрабатываются в формате JSON. 

***

### HOW TO BUILD

#### Задать порт в env:
    export PORT="XXXX"

#### Команда для запуска сервера:
    go run ./cmd/.

***

### Входящие данные

Запросы обрабатываются по следующим URL:
 `/api/v1/calculate`


Запросы с методом `POST` в формате JSON вида:
```JSON 
    {
        "expression": "ваше-арифметическое-выражение"
    }
```

В случае успешной обработки возвращает:
```JSON
    {
        "result": "результат-арифметического-выражения"
    }
```

***

### Исключения

Ошибки возвращаются в формате JSON:
```JSON
    {
        "error": "ошибка-во-время-обработки"
    }
```

***

#### Неверный форма body JSON запроса:
    
    400 StatusBadRequest 
```JSON
    {
        "error": "invalid body request"
    }
```

#### Ошибка во время обработки:
    
    422 StatusUnprocessableEntity 
```JSON
    {
        "error": "ошибка-обработки"
    }
```

#### Ошибка неверного метода:

    405 StatusMethodNotAllowed 
```JSON
    {
        "error": "only post requests"
    }
```

#### Все возможные ошибки обработки запросов:
```go
    ErrInvalidExpression       = "invalid expression"
	ErrInvalidSymbolExpression = "invalid symbol in expression"
	ErrDevisionByZero          = "division by zero"
	ErrEmptyExpression         = "empty expression"
```

*** 

### Логгер

Используется middleware с логгированием всех запросов и обработкой статус кода 500 StatusInternalServerError
    
    /pkg/middlewares/

***

### Тестирование с покрытием

Покрытие тестами 100%, все возвращаемые ошибки в `pkg/calculation/errors.go`
```go
    go test -v -cover ./pkg/calculation
```

***

### Тестрирование с покрытием + web страничка

Web-страничка появляется на пути `pkg/calculation/cover_calculation.html`
```go
    go test ./pkg/calculation -coverprofile=pkg/calculation/cover_calculation.out
    go tool cover -html=pkg/calculation/cover_calculation.out -o pkg/calculation/cover_calculation.html
```