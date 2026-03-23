import VerseOfDay from "../components/VerseOfDay";
import SearchTheScriptures from "../components/SearchTheScriptures";
import Verses from "../components/Verses";
import {
  createEffect,
  createResource,
  createSignal,
  ErrorBoundary,
  For,
  Match,
  Show,
  Switch,
} from "solid-js";
import { SearchVerse } from "../../wailsjs/go/main/App";
import { reference } from "../../wailsjs/go/models";

export default function Main() {
  const [passages, setPassages] = createSignal<string | null>();
  const [fetchPassages] = createResource(passages, SearchVerse);
  const [verseResults, setVerseResults] =
    createSignal<reference.BibleReference | null>(null);

  const [errors, setErrors] = createSignal<string[] | null>();

  createEffect(() => {
    if (fetchPassages.state === "ready" && fetchPassages()) {
      setVerseResults(fetchPassages());
    }

    if (verseResults()?.Errors) {
      setErrors(verseResults()?.Errors);
    } else {
      setErrors([]);
    }

    console.log(verseResults());
  });

  return (
    <>
      <VerseOfDay />
      <SearchTheScriptures setPassages={setPassages} errors={errors} />
      <Show when={verseResults()}>{(data) => <Verses data={data()} />}</Show>
    </>
  );
}
