import { A } from "@solidjs/router";
import ReadChapters from "./ReadChapters";

export default function Navbar() {
  return (
    <header class="border-b bg-card/50 backdrop-blur-sm sticky top-0 z-10 -mt-3">
      <nav class="container mx-auto flex items-center justify-between mt-4 p-8">
        <div class="flex items-center gap-2">
          <h1 class="text-2xl font-bold text-amber-50 absolute left-4">
            <A href="" class="hover:text-purple-400 transition-colors">
              Jonah
            </A>
          </h1>
        </div>
        <ReadChapters />
      </nav>
    </header>
  );
}
