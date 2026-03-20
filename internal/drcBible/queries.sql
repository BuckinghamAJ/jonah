-- name: GetBook :one
SELECT * FROM DRC_books
WHERE id = ? LIMIT 1;

-- name: GetBookFromTitle :one
SELECT * FROM DRC_books
WHERE name = ? LIMIT 1;

-- name: GetVerses :many
SELECT v.chapter, v.verse, v.text FROM DRC_verses as v
JOIN DRC_books as b ON b.id=v.book_id
WHERE v.book_id = ? and v.chapter = ? AND v.verse BETWEEN ? and ?
ORDER BY v.verse;

-- name: GetChapter :many
SELECT v.chapter, v.verse, v.text FROM DRC_verses as v
JOIN DRC_books as b ON b.id=v.book_id
WHERE v.book_id = ? and v.chapter = ?
ORDER BY v.verse;
