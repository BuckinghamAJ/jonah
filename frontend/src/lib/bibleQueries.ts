import { query } from "@solidjs/router";
import { GetAllBooks, GetAllChapters } from "../../wailsjs/go/main/App";

export const getAllBooks = query(() => GetAllBooks(), "bible-books");

export const getBookChapters = query(
  (bookId: number) => GetAllChapters(bookId),
  "bible-chapters",
);
