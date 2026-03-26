import { createAsync, reload } from "@solidjs/router";
import { For, Setter, Show, createSignal } from "solid-js";
import { useNavigate } from "@solidjs/router";
import { getAllBooks, getBookChapters } from "../lib/bibleQueries";
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuTrigger,
} from "./solid-ui/DropdownMenu";
import { createEffect } from "solid-js";
import { Accessor } from "solid-js/types/server/reactive.js";

export default function ReadChapters() {
  const books = createAsync(() => getAllBooks());
  const [selectedBookId, setSelectedBookId] = createSignal<number | null>(null);
  const [selectedChapter, setSelectedChapter] = createSignal<number | null>(
    null,
  );

  const chapters = createAsync(() => {
    const bookId = selectedBookId();

    if (bookId === null) {
      return Promise.resolve([]);
    }

    return getBookChapters(bookId);
  });

  createEffect(() => {
    console.log("Selected book ID:", selectedBookId());
    console.log("Chapters:", chapters());
  });

  return (
    <DropdownMenu>
      <DropdownMenuTrigger
        class="inline-flex items-center justify-center
            gap-2 whitespace-nowrap rounded-md text-sm font-medium ring-offset-background transition-colors
            focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-2 hover:text-purple-400 hover:cursor-pointer
           h-10 w-10 absolute right-4"
        type="button"
        aria-label="Read chapters"
      >
        <svg
          xmlns="http://www.w3.org/2000/svg"
          width="24"
          height="24"
          viewBox="0 0 24 24"
          fill="none"
          stroke="currentColor"
          stroke-width="2"
          stroke-linecap="round"
          stroke-linejoin="round"
          class="lucide lucide-book-marked h-5 w-5"
          id="lucide-book-marked"
        >
          <path d="M10 2v8l3-3 3 3V2"></path>
          <path d="M4 19.5v-15A2.5 2.5 0 0 1 6.5 2H19a1 1 0 0 1 1 1v18a1 1 0 0 1-1 1H6.5a1 1 0 0 1 0-5H20"></path>
        </svg>
      </DropdownMenuTrigger>
      <DropdownMenuContent class="w-56 mt-6 max-h-96 h-fit overflow-auto absolute z-50 bg-slate-800">
        <BookDisplay
          books={books}
          selectedBookId={selectedBookId}
          setSelectedBookId={setSelectedBookId}
        />

        <ChapterDisplay
          chapters={chapters}
          selectedBookId={selectedBookId}
          setSelectedBookId={setSelectedBookId}
        />
      </DropdownMenuContent>
    </DropdownMenu>
  );
}

interface BookDisplayProps {
  books: Accessor<{ id: number; name: string }[] | null | undefined>;
  selectedBookId: Accessor<number | null>;
  setSelectedBookId: Setter<number | null>;
}

function BookDisplay({
  books,
  selectedBookId,
  setSelectedBookId,
}: BookDisplayProps) {
  return (
    <Show
      when={books() !== undefined}
      fallback={<DropdownMenuItem disabled>Loading books...</DropdownMenuItem>}
    >
      <Show when={(books() ?? []).length > 0 && selectedBookId() === null}>
        <For each={books() ?? []}>
          {(book) => (
            <DropdownMenuItem
              class="hover:border-l-2 hover:bg-amber-100 hover:text-purple-600"
              closeOnSelect={false}
              onSelect={() => {
                setSelectedBookId(book.id);
              }}
            >
              {book.name}
            </DropdownMenuItem>
          )}
        </For>
      </Show>
    </Show>
  );
}

interface ChapterDisplayProps {
  chapters: Accessor<number[] | null | undefined>;
  selectedBookId: Accessor<number | null>;
  setSelectedBookId: Setter<number | null>;
}

function ChapterDisplay({
  chapters,
  selectedBookId,
  setSelectedBookId,
}: ChapterDisplayProps) {
  const navigate = useNavigate();

  const chapterNav = (path: string) => {
    navigate(path, { replace: true });
    reload();
  };

  return (
    <Show when={selectedBookId() !== null}>
      {/* Back to Book Selections*/}
      <DropdownMenuItem
        class="hover:border-l-2 hover:bg-amber-100 hover:text-purple-600 sticky top-0 bg-slate-800 z-10"
        closeOnSelect={false}
        onSelect={() => {
          setSelectedBookId(null);
        }}
      >
        &larr; Back to Books
      </DropdownMenuItem>
      <For each={chapters() ?? []}>
        {(chapter) => (
          <DropdownMenuItem
            class="hover:border-l-2 hover:bg-amber-100 hover:text-purple-600"
            onSelect={() => {
              chapterNav(`/read/${selectedBookId()}/${chapter}`);
              setSelectedBookId(null);
            }}
          >
            Chapter {chapter}
          </DropdownMenuItem>
        )}
      </For>
    </Show>
  );
}
