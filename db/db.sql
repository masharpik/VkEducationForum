CREATE EXTENSION IF NOT EXISTS citext WITH SCHEMA public;
CREATE EXTENSION IF NOT EXISTS ltree;

CREATE TABLE IF NOT EXISTS users(
   nickname CITEXT COLLATE "ucs_basic" PRIMARY KEY,
   fullname TEXT,
   about TEXT,
   email CITEXT UNIQUE NOT NULL
);

CREATE TABLE IF NOT EXISTS forums (
   slug CITEXT PRIMARY KEY,
   title TEXT NOT NULL,
   author CITEXT REFERENCES users (nickname) ON DELETE SET NULL ON UPDATE CASCADE,
   threads INT DEFAULT 0,
   posts INT DEFAULT 0
);

CREATE TABLE IF NOT EXISTS threads(
   id SERIAL PRIMARY KEY,
   slug CITEXT UNIQUE,
   title TEXT NOT NULL,
   author CITEXT REFERENCES users (nickname) ON DELETE SET NULL ON UPDATE CASCADE,
   forum CITEXT REFERENCES forums (slug) ON DELETE CASCADE ON UPDATE CASCADE NOT NULL,
   message TEXT NOT NULL,
   votes INT DEFAULT 0,
   created timestamp with time zone DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS posts(
   id SERIAL PRIMARY KEY,
   parent INT REFERENCES posts(id) ON DELETE CASCADE ON UPDATE CASCADE,
   author CITEXT NOT NULL REFERENCES users (nickname) ON DELETE SET NULL ON UPDATE CASCADE,
   message TEXT NOT NULL,
   isEdited BOOLEAN DEFAULT false,
   forum CITEXT REFERENCES forums (slug) ON DELETE CASCADE ON UPDATE CASCADE NOT NULL,
   thread INT REFERENCES threads (id) ON DELETE CASCADE ON UPDATE CASCADE NOT NULL,
   created TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
   num int[] DEFAULT ARRAY []::INTEGER[]
);

CREATE OR REPLACE FUNCTION update_is_edited() RETURNS TRIGGER AS $$
BEGIN
    IF OLD.message <> NEW.message THEN
        NEW.isEdited = true;
    END IF;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER trigger_update_is_edited
BEFORE UPDATE OF message ON posts
FOR EACH ROW
EXECUTE PROCEDURE update_is_edited();

-- Тригер на добавление поста => составляем его номер по вложенности
CREATE OR REPLACE FUNCTION update_parent()
RETURNS TRIGGER AS $$
DECLARE
BEGIN
  NEW.num = (SELECT num FROM posts WHERE id = NEW.parent) || NEW.id;
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER update_parent_trigger
BEFORE INSERT ON posts
FOR EACH ROW
EXECUTE PROCEDURE update_parent();

CREATE TABLE votes(
  id SERIAL PRIMARY KEY,
  thread INT REFERENCES threads (id) ON DELETE CASCADE ON UPDATE CASCADE NOT NULL,
  nickname CITEXT REFERENCES users (nickname) ON DELETE SET NULL ON UPDATE CASCADE,
  vote INT NOT NULL,
  CONSTRAINT unique_thread_nickname UNIQUE (thread, nickname)
);

-- ТРИГГЕР НА ДОБАВЛЕНИЕ ГОЛОСА - ИЗМЕНЯЕТСЯ СЧЕТЧИК КОЛИЧЕСТВА ГОЛОСОВ В ВЕТКЕ
CREATE OR REPLACE FUNCTION increase_vote_count() RETURNS TRIGGER AS $$
BEGIN
  UPDATE threads
  SET votes = votes + NEW.vote
  WHERE id = NEW.thread;
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER increase_vote_count_trigger
AFTER INSERT ON votes
FOR EACH ROW EXECUTE PROCEDURE increase_vote_count();

-- ТРИГГЕР НА УДАЛЕНИЕ ГОЛОСА - ИЗМЕНЯЕТСЯ СЧЕТЧИК КОЛИЧЕСТВА ГОЛОСОВ В ВЕТКЕ
CREATE OR REPLACE FUNCTION decrease_vote_count() RETURNS TRIGGER AS $$
BEGIN
  UPDATE threads
  SET votes = votes - OLD.vote
  WHERE id = OLD.thread;
  RETURN OLD;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER decrease_vote_count_trigger
AFTER DELETE ON votes
FOR EACH ROW EXECUTE PROCEDURE decrease_vote_count();

-- ТРИГГЕР НА ОБНОВЛЕНИЕ ГОЛОСА - ИЗМЕНЯЕТСЯ СЧЕТЧИК КОЛИЧЕСТВА ГОЛОСОВ В ВЕТКЕ
CREATE OR REPLACE FUNCTION update_vote_count() RETURNS TRIGGER AS $$
BEGIN
  UPDATE threads
  SET votes = votes - OLD.vote + NEW.vote
  WHERE id = OLD.thread;
  RETURN OLD;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER update_vote_count_trigger
AFTER UPDATE ON votes
FOR EACH ROW EXECUTE PROCEDURE update_vote_count();

-- ТРИГГЕР НА ДОБАВЛЕНИЕ ВЕТКИ - УВЕЛИЧЕНИЕ СЧЕТЧИКА ВЕТОК В ФОРУМЕ
CREATE OR REPLACE FUNCTION increase_thread_count() RETURNS TRIGGER AS $$
BEGIN
  UPDATE forums
  SET threads = threads + 1
  WHERE slug = NEW.forum;
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER increase_thread_count_trigger
AFTER INSERT ON threads
FOR EACH ROW EXECUTE PROCEDURE increase_thread_count();

-- ТРИГГЕР НА УДАЛЕНИЕ ВЕТКИ - УМЕНЬШЕНИЕ СЧЕТЧИКА ВЕТОК В ФОРУМЕ
CREATE OR REPLACE FUNCTION decrease_thread_count() RETURNS TRIGGER AS $$
BEGIN
  UPDATE forums
  SET threads = threads - 1
  WHERE slug = OLD.forum;
  RETURN OLD;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER decrease_thread_count_trigger
AFTER DELETE ON threads
FOR EACH ROW EXECUTE PROCEDURE decrease_thread_count();

-- ТРИГГЕР НА ДОБАВЛЕНИЕ ПОСТА - УВЕЛИЧЕНИЕ СЧЕТЧИКА ПОСТОВ В ФОРУМЕ
CREATE OR REPLACE FUNCTION increase_post_count() RETURNS TRIGGER AS $$
BEGIN
  UPDATE forums
  SET posts = posts + 1
  WHERE slug = NEW.forum;
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER increase_post_count_trigger
AFTER INSERT ON posts
FOR EACH ROW EXECUTE PROCEDURE increase_post_count();

-- ТРИГГЕР НА УДАЛЕНИЕ ПОСТА - УМЕНЬШЕНИЕ СЧЕТЧИКА ПОСТОВ В ФОРУМЕ
CREATE OR REPLACE FUNCTION decrease_post_count() RETURNS TRIGGER AS $$
BEGIN
  UPDATE forums
  SET posts = posts - 1
  WHERE slug = OLD.forum;
  RETURN OLD;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER decrease_post_count_trigger
AFTER DELETE ON posts
FOR EACH ROW EXECUTE PROCEDURE decrease_post_count();

CREATE UNLOGGED TABLE IF NOT EXISTS user_forums (
  nickname CITEXT COLLATE "ucs_basic" NOT NULL REFERENCES users (nickname),
  forum    CITEXT NOT NULL REFERENCES forums (slug),
  fullname TEXT,
  about    TEXT,
  email    CITEXT,
  CONSTRAINT user_forum_key UNIQUE (nickname, forum)
);

CREATE OR REPLACE FUNCTION update_user_forum()
RETURNS TRIGGER AS $$
DECLARE
    _nickname CITEXT;
    _fullname TEXT;
    _about    TEXT;
    _email    CITEXT;
BEGIN
    SELECT u.nickname, u.fullname, u.about, u.email
    FROM users u
    WHERE u.nickname = NEW.author
    INTO _nickname, _fullname, _about, _email;

    INSERT INTO user_forums (nickname, fullname, about, email, forum)
    VALUES (_nickname, _fullname, _about, _email, NEW.forum)
    ON CONFLICT DO NOTHING;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER trigger_insert_threads
AFTER INSERT ON threads
FOR EACH ROW EXECUTE PROCEDURE update_user_forum();

CREATE TRIGGER trigger_insert_posts
AFTER INSERT ON posts
FOR EACH ROW EXECUTE PROCEDURE update_user_forum();

CREATE INDEX IF NOT EXISTS idx_threads_slug ON threads USING hash (slug);
CREATE INDEX IF NOT EXISTS idx_threads_author ON threads USING hash (author);

CREATE INDEX IF NOT EXISTS idx_posts_id ON posts USING hash (id);
CREATE INDEX IF NOT EXISTS idx_posts_id_btree ON posts using btree (id);
CREATE INDEX IF NOT EXISTS post_thread ON posts USING hash (thread);
CREATE INDEX IF NOT EXISTS post_parent ON posts (thread, id, (num[1]), parent);
CREATE INDEX IF NOT EXISTS post_path_1_path ON posts ((num[1]), num);
CREATE INDEX IF NOT EXISTS post_thread_path ON posts (thread, num);

CREATE INDEX idx_posts_created ON posts using btree (created);
CREATE INDEX idx_posts_path ON posts using btree (num);

CREATE INDEX idx_forums_slug ON forums USING HASH (slug);
CREATE INDEX idx_threads_id_hash ON threads USING hash (id);
CREATE INDEX idx_threads_created ON threads USING btree (created);
CREATE INDEX idx_users_nickname ON users USING HASH (nickname);
CREATE INDEX idx_users_nickname_btree ON users USING btree (nickname varchar_pattern_ops);
CREATE INDEX idx_users_email ON users USING HASH (email);
