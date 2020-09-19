CREATE TABLE IF NOT EXISTS mail_text (
    id UUID PRIMARY KEY,
    created_at TIMESTAMP,
    content TEXT
);

CREATE TABLE IF NOT EXISTS mail_send_mail (
    id UUID PRIMARY KEY,
    created_at TIMESTAMP,
    mail VARCHAR(100),
    send_time TIMESTAMP,
    send_status BOOLEAN,
    text_id UUID REFERENCES mail_text(id)
);

CREATE TABLE IF NOT EXISTS auth_user_type (
    id INT PRIMARY KEY,
    actor VARCHAR(20) NOT NULL
);

INSERT INTO auth_user_type(id, actor) VALUES(1, 'client'), (2, 'club') ON CONFLICT DO NOTHING;

CREATE TABLE IF NOT EXISTS auth_users (
    id UUID PRIMARY KEY,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT NULL,
    mail VARCHAR(100) UNIQUE NOT NULL,
    password VARCHAR(255),
    access_token VARCHAR(255) NOT NULL,
    refresh_token VARCHAR(255) NOT NULL,
    user_type_id INT REFERENCES auth_user_type(id) NOT NULL,
    is_verified BOOLEAN
);
