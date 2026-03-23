import { render, screen } from "@solidjs/testing-library";
import { createSignal } from "solid-js";
import { describe, expect, it } from "vitest";
import SearchTheScriptures from "./SearchTheScriptures";

describe("SearchTheScriptures", () => {
  it("does not render errors when errors signal is null", () => {
    const [errors] = createSignal<string[] | null>(null);
    render(() => (
      <SearchTheScriptures setPassages={() => {}} errors={errors} />
    ));

    expect(screen.queryByText(/error/i)).not.toBeInTheDocument();
  });

  it("does not render errors when errors signal is empty array", () => {
    const [errors] = createSignal<string[] | null>([]);
    render(() => (
      <SearchTheScriptures setPassages={() => {}} errors={errors} />
    ));

    expect(screen.queryByText(/error/i)).not.toBeInTheDocument();
  });

  it("renders a single error message", () => {
    const [errors] = createSignal<string[] | null>([
      "Book not found: FakeBook",
    ]);
    render(() => (
      <SearchTheScriptures setPassages={() => {}} errors={errors} />
    ));

    expect(screen.getByText("Book not found: FakeBook")).toBeInTheDocument();
  });

  it("renders multiple error messages", () => {
    const [errors] = createSignal<string[] | null>([
      "Book not found: FakeBook",
      "Invalid verse range: 1-999",
    ]);
    render(() => (
      <SearchTheScriptures setPassages={() => {}} errors={errors} />
    ));

    expect(screen.getByText("Book not found: FakeBook")).toBeInTheDocument();
    expect(
      screen.getByText("Invalid verse range: 1-999")
    ).toBeInTheDocument();
  });

  it("renders the search input and button", () => {
    const [errors] = createSignal<string[] | null>(null);
    render(() => (
      <SearchTheScriptures setPassages={() => {}} errors={errors} />
    ));

    expect(screen.getByPlaceholderText(/search verses/i)).toBeInTheDocument();
    expect(screen.getByText("Search")).toBeInTheDocument();
  });
});
