import { reference } from "../../wailsjs/go/models";
import { For, Show } from "solid-js";

interface VerseDisplayProps {
  verse: reference.BiblePassage;
}

export default function VerseDisplay(props: VerseDisplayProps) {
  return (
    <Show when={props.verse.FullText.length > 0}>
      <div class="px-4 text-purple-300 mt-4">
        <div class="space-y-4 duration-500 rounded-xl hover:shadow-amber-400 transition-all bg-gray-800">
          <div class="p-6 flex justify-between items-start gap-4">
            <div class="flex-1">
              <p class="verse-reference text-base font-semibold pb-1">
                {props.verse.Book} {props.verse.Chapter}
                {props.verse.StartVerse !== 0
                  ? `:${props.verse.StartVerse}`
                  : ""}
                {props.verse.EndVerse &&
                props.verse.EndVerse !== props.verse.StartVerse
                  ? `-${props.verse.EndVerse}`
                  : ""}
              </p>
              <p class="text-lg/relaxed mb-3 font-serif text-amber-50">
                <For each={props.verse.FullText}>
                  {(singleV) => (
                    <span>
                      <sup class="mr-0.5">
                        <em>{singleV.Number}</em>
                      </sup>
                      {singleV.Text}{" "}
                    </span>
                  )}
                </For>
              </p>
            </div>
            <button
              class="inline-flex items-center justify-center
            gap-2 whitespace-nowrap rounded-md text-sm font-medium ring-offset-background transition-colors
            focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-2 hover:bg-amber-50 hover:text-purple-400 hover:cursor-pointer
            hover:bg-accent hover:text-accent-foreground h-10 w-10"
              type="button"
            >
              <svg
                xmlns="http://www.w3.org/2000/svg"
                fill="none"
                viewBox="0 0 24 24"
                stroke-width="1.5"
                stroke="currentColor"
                class="size-6"
              >
                <path
                  stroke-linecap="round"
                  stroke-linejoin="round"
                  d="M17.593 3.322c1.1.128 1.907 1.077 1.907 2.185V21L12 17.25 4.5 21V5.507c0-1.108.806-2.057 1.907-2.185a48.507 48.507 0 0 1 11.186 0Z"
                />
              </svg>
            </button>
          </div>
        </div>
      </div>
    </Show>
  );
}
