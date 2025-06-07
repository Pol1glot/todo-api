CREATE TABLE IF NOT EXISTS tasks (
    id SERIAL PRIMARY KEY,
    title TEXT NOT NULL,
    description TEXT,
    status TEXT CHECK (status IN ('new', 'in_progress', 'done')) DEFAULT 'new',
    created_at TIMESTAMP DEFAULT now(),
    updated_at TIMESTAMP DEFAULT now()
    );

-- Создаем функцию для обновления updated_at
CREATE OR REPLACE FUNCTION update_timestamp_column()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = now();
    IF TG_OP = 'INSERT' THEN
        NEW.created_at = NEW.updated_at;
END IF;
RETURN NEW;
END;
$$ language 'plpgsql';

-- Создаем триггер, который будет вызывать функцию при вставке и обновлении записи
CREATE TRIGGER update_tasks_timestamp
    BEFORE INSERT OR UPDATE ON tasks
                         FOR EACH ROW
                         EXECUTE FUNCTION update_timestamp_column();