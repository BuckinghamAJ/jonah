import { query } from "@solidjs/router";
import {
  GetAllBooks,
  GetAllChapters,
  GetAllVerses,
} from "../../wailsjs/go/main/App";

export const getAllBooks = query(() => GetAllBooks(), "bible-books");

export const getBookChapters = query(
  (bookId: number) => GetAllChapters(bookId),
  "bible-chapters",
);

export const getVersesFrom = query(
  (bookId: number, chapterId: number) => GetAllVerses(bookId, chapterId),
  "get-verses",
);
