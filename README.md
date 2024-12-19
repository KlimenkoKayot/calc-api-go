# calc-api-go
Микросервис для подсчета простых математических выражений

Тестирование с покрытием
    go test -v -cover ./pkg/calculation

Тестрирование с покрытием + web страничка
    go test ./pkg/calculation -coverprofile=pkg/calculation/cover_calculation.out
    go tool cover -html=pkg/calculation/cover_calculation.out -o pkg/calculation/cover_calculation.html