import VerseOfDay from "./VerseOfDay";
import SearchTheScriptures from "./SearchTheScriptures";
import Verses from "./Verses";
import { createEffect, createResource, createSignal, Show } from "solid-js";
import { SearchVerse } from "../../wailsjs/go/main/App";
import { reference } from "../../wailsjs/go/models";

export default function Main() {
  const [passages, setPassages] = createSignal<string | null>();
  const [fetchPassages] = createResource(passages, SearchVerse);
  const [verseResults, setVerseResults] =
    createSignal<reference.BibleReference | null>(null);

  createEffect(() => {
    if (fetchPassages.state === "ready" && fetchPassages()) {
      setVerseResults(fetchPassages());
    }
  });

  return (
    <>
      <VerseOfDay />
      <SearchTheScriptures setPassages={setPassages} />
      <Show when={verseResults()} fallback={null}>
        {(data) => <Verses data={data()} />}
      </Show>
    </>
  );
}
