## TODO-LIST
## 🚀 Возможности

- Создание задачи  
- Получение списка всех задач  
- Обновление задачи по ID  
- Удаление задачи по ID  

## ⚙️ Работа сервиса

### 1. Клонировать репозиторий

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
# через taskfile
task dev
```

Сервис будет доступен по адресу: `http://localhost:your_port`

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
![изображение](https://github.com/user-attachments/assets/aa33a820-c960-44b5-8bc3-574bea796c82)




📷 **PostgreSQL**:
Скриншот таблицы `tasks` с добавленной записью
![изображение](https://github.com/user-attachments/assets/157e54b1-281c-4bdb-9370-08c202a3a2f9)


