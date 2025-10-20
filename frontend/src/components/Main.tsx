import VerseOfDay from "./VerseOfDay";

export default function Main() {
  return (
    <div class="text-purple-300 justify-center mt-4 container mx-auto px-4 py-12 max-w-3xl">
      <div class="text-center mb-8">
        <h1 class="text-4xl font-extrabold">Verse of the Day</h1>
        <p class="text-gray-300">May God Bless You</p>
      </div>

      <VerseOfDay />
    </div>
  );
}
