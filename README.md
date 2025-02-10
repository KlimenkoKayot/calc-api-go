# Микросервис для подсчета простых математических выражений

Микросервис предоставляет API для вычисления арифметических выражений. Все запросы и ответы обрабатываются в формате JSON.

---

## Содержание

1. [Как собрать и запустить](#как-собрать-и-запустить)
2. [Использование API](#использование-api)
3. [Формат запросов и ответов](#формат-запросов-и-ответов)
4. [Обработка ошибок](#обработка-ошибок)
5. [Логгирование](#логгирование)
6. [Тестирование](#тестирование)
7. [Лицензия](#лицензия)

---

<a name="как-собрать-и-запустить"></a>
## Как собрать и запустить

### Установка порта

Перед запуском сервера необходимо задать порт через переменную окружения:

```bash
export PORT="XXXX"
```

### Запуск сервера

Для запуска сервера выполните команду:

```bash
go run ./cmd/.
```

Сервер будет доступен на `http://localhost:XXXX`.

---

<a name="использование-api"></a>
## Использование API

### Основной эндпоинт

Все запросы отправляются на эндпоинт:

```
POST /api/v1/calculate
```

<a name="формат-запросов-и-ответов"></a>
### Формат запроса

Запрос должен быть в формате JSON:

```json
{
    "expression": "ваше-арифметическое-выражение"
}
```

Пример:

```json
{
    "expression": "2 + 2 * 2"
}
```

### Формат ответа

В случае успешного выполнения запроса, сервер вернет:

```json
{
    "result": 6
}
```

---

<a name="обработка-ошибок"></a>
## Обработка ошибок

### Формат ошибок

Все ошибки возвращаются в формате JSON:

```json
{
    "error": "описание-ошибки"
}
```

### Возможные ошибки

1. **Неверный формат тела запроса**  
   Код: `400 StatusBadRequest`  
   Ответ:
   ```json
   {
       "error": "invalid body request"
   }
   ```

2. **Ошибка обработки выражения**  
   Код: `422 StatusUnprocessableEntity`  
   Ответ:
   ```json
   {
       "error": "описание-ошибки"
   }
   ```

3. **Неверный метод запроса**  
   Код: `405 StatusMethodNotAllowed`  
   Ответ:
   ```json
   {
       "error": "only post requests"
   }
   ```

4. **Внутренняя ошибка сервера**  
   Код: `500 StatusInternalServerError`  
   Ответ:
   ```json
   {
       "error": "internal server error"
   }
   ```

### Список возможных ошибок обработки выражений

```go
ErrInvalidExpression       = "invalid expression"
ErrInvalidSymbolExpression = "invalid symbol in expression"
ErrDevisionByZero          = "division by zero"
ErrEmptyExpression         = "empty expression"
```

---

<a name="логгирование"></a>
## Логгирование

Для логгирования всех запросов и ошибок используется middleware, расположенный в:

```
/pkg/middlewares/
```

Middleware автоматически логирует:
- Метод запроса
- URL
- Статус ответа
- Время выполнения

---

<a name="тестирование"></a>
## Тестирование

### Запуск тестов

Для запуска тестов с покрытием выполните команду:

```bash
go test -v -cover ./pkg/calculation
```

### Генерация отчета о покрытии

Чтобы сгенерировать HTML-отчет о покрытии тестами, выполните:

```bash
go test ./pkg/calculation -coverprofile=pkg/calculation/cover_calculation.out
go tool cover -html=pkg/calculation/cover_calculation.out -o pkg/calculation/cover_calculation.html
```

Отчет будет доступен по пути:

```
pkg/calculation/cover_calculation.html
```

---

<a name="лицензия"></a>
## Лицензия

Этот проект распространяется под лицензией MIT. Подробности см. в файле [LICENSE](LICENSE).
