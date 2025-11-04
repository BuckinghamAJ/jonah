import { defaultAllowedOrigins } from "vite";
import { reference } from "../../wailsjs/go/models";

interface VerseDisplayProps {
  verse: reference.BiblePassage;
}

export default function VerseDisplay(props: VerseDisplayProps) {
  return (
    <>
      <div>
        <p class="text-lg">{props.verse.Text}</p>
      </div>
    </>
  );
}
