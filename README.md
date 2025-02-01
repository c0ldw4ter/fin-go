# fin-go

## API Documentation

### 1. Пополнение баланса пользователя

- **Эндпоинт:**
  POST /api/v1/users/:user_id/topup

- **Описание:**
  Позволяет пополнить баланс пользователя с указанным `user_id` на заданную сумму.

- **Параметры запроса:**
- **URL-параметр:**
  - `user_id` (integer) — идентификатор пользователя, чей баланс нужно пополнить.
- **Тело запроса (JSON):**

  ```json
  {
  	"amount": 100.0
  }
  ```

  - `amount` (float) — сумма для пополнения баланса. Должна быть положительным числом.

- **Ответы:**
- **Успешный ответ (HTTP 200):**
  ```json
  {
  	"message": "balance topped up successfully"
  }
  ```
- **Ошибка (HTTP 400 или 500):**

  ```json
  {
  	"error": "failed to update user balance: user not found"
  }
  ```

- **Пример использования:**

````bash
curl -X POST http://localhost:8080/api/v1/users/1/topup \
-H "Content-Type: application/json" \
-d '{"amount": 100.0}'


### 2. Перевод средств между пользователями

- **Эндпоинт:**
  POST /api/v1/transactions/transfer

- **Описание:**
Позволяет пополнить баланс пользователя с указанным `user_id` на заданную сумму.

- **Параметры запроса:**
- **URL-параметр:**
  - `user_id` (integer) — идентификатор пользователя, чей баланс нужно пополнить.
- **Тело запроса (JSON):**
  ```json
  {
    "amount": 100.0
  }
````

- `amount` (float) — сумма для пополнения баланса. Должна быть положительным числом.

- **Ответы:**
- **Успешный ответ (HTTP 200):**
  ```json
  {
  	"message": "balance topped up successfully"
  }
  ```
- **Ошибка (HTTP 400 или 500):**

  ```json
  {
  	"error": "failed to update user balance: user not found"
  }
  ```

- **Пример использования:**

````bash
curl -X POST http://localhost:8080/api/v1/users/1/topup \
-H "Content-Type: application/json" \
-d '{"amount": 100.0}'



### 3. Получение последних транзакций пользователя

- **Эндпоинт:**
  GET /api/v1/users/:user_id/transactions

- **Описание:**
Позволяет пополнить баланс пользователя с указанным `user_id` на заданную сумму.

- **Параметры запроса:**
- **URL-параметр:**
  - `user_id` (integer) — идентификатор пользователя, чей баланс нужно пополнить.
- **Тело запроса (JSON):**
  ```json
  {
    "amount": 100.0
  }
````

- `amount` (float) — сумма для пополнения баланса. Должна быть положительным числом.

- **Ответы:**
- **Успешный ответ (HTTP 200):**
  ```json
  {
  	"message": "balance topped up successfully"
  }
  ```
- **Ошибка (HTTP 400 или 500):**

  ```json
  {
  	"error": "failed to update user balance: user not found"
  }
  ```

- **Пример использования:**

```bash
curl -X POST http://localhost:8080/api/v1/users/1/topup \
-H "Content-Type: application/json" \
-d '{"amount": 100.0}'
```
