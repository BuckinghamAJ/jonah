import ReadChaptersIcon from "./ReadChaptersIcon";

export default function Navbar() {
  return (
    <header class="border-b bg-card/50 backdrop-blur-sm sticky top-0 z-10 -mt-3">
      <nav class="container mx-auto px-4 py-4 flex items-center justify-between mt-4">
        <div class="flex items-center gap-2">
          <h1 class="text-2xl font-bold text-amber-50">Jonah</h1>
        </div>
        <button
          class="inline-flex items-center justify-center
          gap-2 whitespace-nowrap rounded-md text-sm font-medium ring-offset-background transition-colors
          focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-2 hover:bg-amber-50 hover:text-purple-400 hover:cursor-pointer
          hover:bg-accent hover:text-accent-foreground h-10 w-10"
          type="button"
        >
          <ReadChaptersIcon />
        </button>
      </nav>
    </header>
  );
}
