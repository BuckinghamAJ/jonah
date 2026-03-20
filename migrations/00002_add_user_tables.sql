-- +goose Up

-- ============================================================
-- 1. Recreate the translations table with an INTEGER PRIMARY KEY
--    SQLite does not support ALTER TABLE ... ADD PRIMARY KEY,
--    so we use the create-copy-drop-rename pattern.
-- ============================================================

-- +goose StatementBegin
CREATE TABLE translations_new (
  id          INTEGER PRIMARY KEY AUTOINCREMENT,
  translation TEXT    NOT NULL UNIQUE,
  title       TEXT,
  license     TEXT
);
-- +goose StatementEnd

-- +goose StatementBegin
INSERT INTO translations_new (translation, title, license)
  SELECT translation, title, license FROM translations;
-- +goose StatementEnd

DROP TABLE translations;

ALTER TABLE translations_new RENAME TO translations;

-- ============================================================
-- 2. A canonical "passage" entity so favorites can point to
--    either a single verse or a verse range.
-- ============================================================

-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS user_passages (
  id                   INTEGER PRIMARY KEY,
  translation_id       INTEGER NOT NULL REFERENCES translations(id) ON DELETE RESTRICT,
  start_verse_id       INTEGER NOT NULL REFERENCES DRC_verses(id) ON DELETE CASCADE,
  end_verse_id         INTEGER NOT NULL REFERENCES DRC_verses(id) ON DELETE CASCADE,

  start_verse_ordinal  INTEGER NOT NULL,
  end_verse_ordinal    INTEGER NOT NULL,

  created_at           TEXT NOT NULL DEFAULT (datetime('now')),

  CHECK (start_verse_ordinal > 0),
  CHECK (end_verse_ordinal >= start_verse_ordinal),

  UNIQUE (translation_id, start_verse_id, end_verse_id)
);
-- +goose StatementEnd

CREATE INDEX IF NOT EXISTS idx_user_passages_range
  ON user_passages (translation_id, start_verse_ordinal, end_verse_ordinal);

-- ============================================================
-- 3. User favorites (single implicit local user -- no user_id)
-- ============================================================

-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS user_favorites (
  id            INTEGER PRIMARY KEY,
  passage_id    INTEGER NOT NULL UNIQUE REFERENCES user_passages(id) ON DELETE CASCADE,

  note          TEXT,
  color         TEXT,
  created_at    TEXT NOT NULL DEFAULT (datetime('now')),
  updated_at    TEXT NOT NULL DEFAULT (datetime('now'))
);
-- +goose StatementEnd

CREATE INDEX IF NOT EXISTS idx_user_favorites_created
  ON user_favorites (created_at DESC);

-- ============================================================
-- 4. Tags (single implicit local user -- no user_id)
-- ============================================================

-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS user_tags (
  id            INTEGER PRIMARY KEY,
  name          TEXT NOT NULL UNIQUE,
  created_at    TEXT NOT NULL DEFAULT (datetime('now'))
);
-- +goose StatementEnd

-- ============================================================
-- 5. Many-to-many: favorites <-> tags
-- ============================================================

-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS user_favorite_tags (
  favorite_id   INTEGER NOT NULL REFERENCES user_favorites(id) ON DELETE CASCADE,
  tag_id        INTEGER NOT NULL REFERENCES user_tags(id) ON DELETE CASCADE,
  PRIMARY KEY (favorite_id, tag_id)
);
-- +goose StatementEnd


-- +goose Down

-- Drop user tables in reverse dependency order
DROP TABLE IF EXISTS user_favorite_tags;
DROP TABLE IF EXISTS user_tags;
DROP TABLE IF EXISTS user_favorites;
DROP TABLE IF EXISTS user_passages;
DROP INDEX IF EXISTS idx_user_passages_range;
DROP INDEX IF EXISTS idx_user_favorites_created;

-- Restore translations table to original TEXT PRIMARY KEY schema

-- +goose StatementBegin
CREATE TABLE translations_old (
  translation TEXT PRIMARY KEY,
  title       TEXT,
  license     TEXT
);
-- +goose StatementEnd

-- +goose StatementBegin
INSERT INTO translations_old (translation, title, license)
  SELECT translation, title, license FROM translations;
-- +goose StatementEnd

DROP TABLE translations;

ALTER TABLE translations_old RENAME TO translations;
