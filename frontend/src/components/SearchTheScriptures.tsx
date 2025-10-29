export default function SearchTheScriptures() {
  return (
    <div class="text-purple-300 justify-center mt-4 container mx-auto px-4 py-12 max-w-3xl">
      <div class="text-center">
        <h1 class="text-4xl font-extrabold">Search the Scriptures</h1>
        <p class="text-gray-300">Find any passage</p>
      </div>
      <div class="mt-2">
        <div class="flex items-center rounded-md h-10 pl-3 bg-gray-800 outline-1 -outline-offset-1 outline-gray-600 has-[input:focus-within]:outline-2 has-[input:focus-within]:-outline-offset-2 has-[input:focus-within]:outline-indigo-500">
          <svg
            xmlns="http://www.w3.org/2000/svg"
            fill="none"
            viewBox="0 0 24 24"
            stroke-width="1.5"
            stroke="currentColor"
            class="size-6"
          >
            <path
              stroke-linecap="round"
              stroke-linejoin="round"
              d="m15.75 15.75-2.489-2.489m0 0a3.375 3.375 0 1 0-4.773-4.773 3.375 3.375 0 0 0 4.774 4.774ZM21 12a9 9 0 1 1-18 0 9 9 0 0 1 18 0Z"
            />
          </svg>

          <input
            class="block min-w-0 grow bg-gray-800 py-1.5 pr-3 text-base text-white placeholder:text-gray-500 focus:outline-none sm:text-sm/6 pl-3"
            placeholder="Search verses (e.g. John 3:16; Psalm 23:1-6)"
            required
          ></input>

          <button class="px-5 bg-gray-800 outline-1 outline-gray-500 rounded-md block py-1.5 justify-center hover:bg-gray-600 hover:cursor-pointer hover:font-semibold transition-colors">
            Search
          </button>
        </div>
      </div>
    </div>
  );
}
