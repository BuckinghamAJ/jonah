import {
  createAsync,
  createAsyncStore,
  Params,
  RouteDefinition,
  RoutePreloadFuncArgs,
  useParams,
} from "@solidjs/router";
import { createEffect, createSignal, For } from "solid-js";
import VerseDisplay from "~/components/VerseDisplay";
import { getVersesFrom } from "~/lib/bibleQueries";
import { reference } from "../../wailsjs/go/models";
import {
  createVisibilityObserver,
  withOccurrence,
} from "@solid-primitives/intersection-observer";
import { on } from "solid-js";

interface ChapterParams extends Params {
  bookId: string;
  chapterId: string;
}

export const route = {
  preload: ({ params }: RoutePreloadFuncArgs) => {
    getVersesFrom(Number(params.bookId), Number(params.chapterId));
  },
} satisfies RouteDefinition;

export default function Chapter() {
  const params = useParams<ChapterParams>();
  let el: HTMLDivElement | undefined;
  // let ready = false;

  const [bookId, setBookId] = createSignal(Number(params.bookId));
  const [chapterId, setChapterId] = createSignal(Number(params.chapterId));
  const [isReady, setIsReady] = createSignal(false);

  const verse = createAsync(() => getVersesFrom(bookId(), chapterId()), {
    initialValue: new reference.BiblePassage(),
  });

  const [verses, setVerses] = createSignal<reference.BiblePassage[]>([]);

  createEffect(
    on(
      () => [params.bookId, params.chapterId],
      () => {
        setBookId(Number(params.bookId));
        setChapterId(Number(params.chapterId));
        setVerses([]);
        setIsReady(false);
      },
    ),
  );

  createEffect(() => {
    setVerses((prev) => [...(prev ?? []), verse()]);
  });

  createEffect(() => {
    if (verses() && verses().length > 0) {
      setIsReady(true);
    }
  });

  const newVerseThreshold = createVisibilityObserver(
    { threshold: 0.8, initialValue: false },
    withOccurrence((entry, { occurrence }) => {
      if (entry.isIntersecting && occurrence == "Entering" && isReady()) {
        setChapterId((prev) => prev + 1);
      }
      return entry.isIntersecting;
    }),
  )(() => el);

  return (
    <div class="px-32 pt-16">
      <div>
        <p class="text-xl/loose mb-3 font-serif text-amber-50">
          <For each={verses()}>{(verse) => <VerseDisplay verse={verse} />}</For>
        </p>
        <div ref={el} class="h-0.5">
          {newVerseThreshold()}
        </div>
      </div>
    </div>
  );
}
