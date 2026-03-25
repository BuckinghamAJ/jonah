-- name: GetBook :one
SELECT * FROM DRC_books
WHERE id = ? LIMIT 1;

-- name: GetBookFromTitle :one
SELECT * FROM DRC_books
WHERE name = ? LIMIT 1;

-- name: GetVerses :many
SELECT v.chapter, v.verse, v.text FROM DRC_verses as v
JOIN DRC_books as b ON b.id=v.book_id
WHERE v.book_id = sqlc.arg(book_id) AND v.chapter = sqlc.arg(chapter) AND v.verse >= sqlc.arg(start_verse) AND v.verse <= sqlc.arg(end_verse)
ORDER BY v.verse;

-- name: GetChapter :many
SELECT v.chapter, v.verse, v.text FROM DRC_verses as v
JOIN DRC_books as b ON b.id=v.book_id
WHERE v.book_id = ? and v.chapter = ?
ORDER BY v.verse;

-- name: GetAllBooks :many
SELECT name, id FROM DRC_books;

-- name: GetChaptersOfBook :many
SELECT DISTINCT v.chapter FROM DRC_verses as v
WHERE v.book_id = ?
ORDER BY v.chapter;
