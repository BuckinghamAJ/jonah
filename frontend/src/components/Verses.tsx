import { For } from "solid-js";
import { reference } from "../../wailsjs/go/models";
import VerseDisplay from "./VerseDisplay";

interface VerseProps {
  data: reference.BibleReference;
}

export default function Verses(props: VerseProps) {
  return (
    <>
      <For each={props.data.Passages}>
        {(verse) => <VerseDisplay verse={verse} />}
      </For>
    </>
  );
}
