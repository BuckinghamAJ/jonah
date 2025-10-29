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
          <svg
            xmlns="http://www.w3.org/2000/svg"
            width="24"
            height="24"
            viewBox="0 0 24 24"
            fill="none"
            stroke="currentColor"
            stroke-width="2"
            stroke-linecap="round"
            stroke-linejoin="round"
            class="lucide lucide-book-marked h-5 w-5"
          >
            <path d="M10 2v8l3-3 3 3V2"></path>
            <path d="M4 19.5v-15A2.5 2.5 0 0 1 6.5 2H19a1 1 0 0 1 1 1v18a1 1 0 0 1-1 1H6.5a1 1 0 0 1 0-5H20"></path>
          </svg>
        </button>
      </nav>
    </header>
  );
}
