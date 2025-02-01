# fin-go
### API Documentation

- 1. Пополнение баланса пользователя

#### POST /api/v1/users/:user_id/topup

    - Позволяет пополнить баланс пользователя с указанным user_id на заданную сумму
    - Успешный ответ (HTTP 200):
      {

`"message": "balance topped up successfully"`
} - Ошибка (HTTP 400 или 500):
{
`"error": "failed to update user balance: user not found"`
} - Пример:
```
curl -X POST http://localhost:8080/api/v1/users/1/topup \
-H "Content-Type: application/json" \
-d '{"amount": 100.0}'

```
