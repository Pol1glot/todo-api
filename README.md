## TODO-LIST
## 🚀 Возможности

- Создание задачи  
- Получение списка всех задач  
- Обновление задачи по ID  
- Удаление задачи по ID  

## ⚙️ Работа сервиса

### 1. Клонировать репозиторий

```bash
git clone https://github.com/your-username/todo-api.git
cd todo-api
````

### 2. Создать `.env` файл в корне проекта

```env
DB_HOST=localhost
DB_PORT=5432
DB_USER=your_db_user
DB_PASSWORD=your_db_password
DB_NAME=your_db_name
APP_PORT=your_port
```

### 3. Запустить PostgreSQL (через Docker)

```bash
docker-compose up
```

### 4. Запустить приложение

```bash
go run main.go
# или, если используешь taskfile
task dev
```

Сервис будет доступен по адресу: `http://localhost:3000`

---

## 📌 Эндпоинты

| Метод  | URL          | Описание             |
| ------ | ------------ | -------------------- |
| POST   | `/tasks`     | Создать новую задачу |
| GET    | `/tasks`     | Получить все задачи  |
| PUT    | `/tasks/:id` | Обновить задачу      |
| DELETE | `/tasks/:id` | Удалить задачу       |

Пример запроса:

```json
POST /tasks
{
  "title": "Название задачи",
  "description": "Описание задачи",
  "status": "new"
}
```

## ✅ Проверка работы

📷 **Postman**:
Скриншот успешного запроса POST `/tasks`
![изображение](https://github.com/user-attachments/assets/ef3eef11-a383-4a3e-9337-cf2e6df2a929)



📷 **PostgreSQL**:
Скриншот таблицы `tasks` с добавленной записью
![изображение](https://github.com/user-attachments/assets/e782b1b3-91d2-4547-95d9-7457ac427a71)

