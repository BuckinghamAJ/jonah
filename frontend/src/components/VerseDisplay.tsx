import { defaultAllowedOrigins } from "vite";
import { reference } from "../../wailsjs/go/models";

export default function VerseDisplay(props) {
  return (
    <>
      <div>
        <p class="text-lg">{props.verse.Text}</p>
      </div>
    </>
  );
}
