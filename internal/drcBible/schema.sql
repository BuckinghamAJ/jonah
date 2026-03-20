CREATE TABLE DRC_books (
	"id"	INTEGER PRIMARY KEY AUTOINCREMENT,
	"name"	TEXT
);

CREATE TABLE DRC_verses (
	"id"	INTEGER PRIMARY KEY AUTOINCREMENT,
	"book_id"	INTEGER,
	"chapter"	INTEGER,
	"verse"	INTEGER,
	"text"	TEXT,
	FOREIGN KEY("book_id") REFERENCES "DRC_books"("id")
);

CREATE TABLE cross_references (
	"id"	INTEGER PRIMARY KEY AUTOINCREMENT,
	"from_book"	TEXT,
	"from_chapter"	INTEGER,
	"from_verse"	INTEGER,
	"to_book"	TEXT,
	"to_chapter"	INTEGER,
	"to_verse_start"	INTEGER,
	"to_verse_end"	INTEGER,
	"votes"	INTEGER,
	FOREIGN KEY("to_book_id") REFERENCES "DRC_books"("id"),
	FOREIGN KEY("from_book_id") REFERENCES "DRC_books"("id")
);
