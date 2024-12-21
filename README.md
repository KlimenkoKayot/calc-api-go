## Микросервис для подсчета простых математических выражений

Реализовано как API, все запросы обрабатываются в формате JSON. 

## `/api/v1/calculate`

Обрабатывает `POST` запросы в формате:
    
    {
        "expression": "ваше-арифметическое-выражение"
    }

В случае успешной обработки возвращает:
    
    {
        "result": "результат-арифметического-выражения"
    }

В случае ошибки возвращает:
    
    {
        "error": "ошибка-во-время-обработки"
    }

### Тестирование с покрытием

Покрытие тестами 100%, все возвращаемые ошибки в `pkg/calculation/errors.go`
    
    go test -v -cover ./pkg/calculation

***

### Тестрирование с покрытием + web страничка

Web-страничка появляется на пути `pkg/calculation/cover_calculation.html`

    go test ./pkg/calculation -coverprofile=pkg/calculation/cover_calculation.out
    go tool cover -html=pkg/calculation/cover_calculation.out -o pkg/calculation/cover_calculation.html
