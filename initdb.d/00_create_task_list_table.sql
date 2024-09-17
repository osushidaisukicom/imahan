-- create task_list table

-- テーブルがすでに存在する場合は削除
DROP TABLE IF EXISTS "task_list";

-- trigger を定義
-- 行が更新された時に updated_at に現在時刻を設定
CREATE FUNCTION set_updated_at() RETURNS TRIGGER AS $$
BEGIN
    IF (TG_OP = 'UPDATE') THEN
        NEW.updated_at := now();
        return NEW;
    END IF;
END;
$$ LANGUAGE plpgsql;

-- テーブルを作成
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE "task_list" (
    task_id UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
    display_name VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
