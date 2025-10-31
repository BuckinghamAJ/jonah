import { For } from "solid-js";
import { reference } from "../../wailsjs/go/models";
import VerseDisplay from "./VerseDisplay";

export default function Verses(props) {
  return (
    <>
      <For each={props.data.Passages}>
        {(verse) => <VerseDisplay verse={verse} />}
      </For>
    </>
  );
}
