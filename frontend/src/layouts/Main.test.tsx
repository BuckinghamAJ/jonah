import { render, screen, fireEvent } from "@solidjs/testing-library";
import { describe, expect, it, vi } from "vitest";
import { reference } from "../../wailsjs/go/models";

// Mock the Wails SearchVerse binding
const mockSearchVerse = vi.fn();
vi.mock("../../wailsjs/go/main/App", () => ({
  SearchVerse: (...args: unknown[]) => mockSearchVerse(...args),
}));

// Import Main after mock is set up
import Main from "../layouts/Main";

function makeRef(
  passages: Partial<reference.BiblePassage>[],
  errors?: string[]
): reference.BibleReference {
  return {
    Passages: passages.map(
      (p) =>
        ({
          Book: p.Book ?? "",
          Chapter: p.Chapter ?? 1,
          StartVerse: p.StartVerse ?? 0,
          EndVerse: p.EndVerse ?? 0,
          FullText: p.FullText ?? [],
          Error: p.Error ?? "",
        }) as reference.BiblePassage
    ),
    Errors: errors ?? [],
  } as reference.BibleReference;
}

describe("Main", () => {
  it("renders the search bar and verse of day on initial load", () => {
    mockSearchVerse.mockReset();
    render(() => <Main />);

    expect(screen.getByText("Verse of the Day")).toBeInTheDocument();
    expect(screen.getByText("Search the Scriptures")).toBeInTheDocument();
  });

  it("displays errors returned from SearchVerse", async () => {
    mockSearchVerse.mockReset();
    mockSearchVerse.mockResolvedValue(
      makeRef([], ["Book not found: FakeBook"])
    );

    render(() => <Main />);

    const input = screen.getByPlaceholderText(/search verses/i);
    fireEvent.input(input, { target: { value: "FakeBook 1:1" } });
    fireEvent.keyDown(input, { key: "Enter" });

    expect(
      await screen.findByText("Book not found: FakeBook")
    ).toBeInTheDocument();
  });

  it("displays verses and errors simultaneously", async () => {
    mockSearchVerse.mockReset();
    mockSearchVerse.mockResolvedValue(
      makeRef(
        [
          {
            Book: "John",
            Chapter: 3,
            StartVerse: 16,
            EndVerse: 16,
            FullText: [
              { Number: 16, Text: "For God so loved the world..." } as reference.Verse,
            ],
          },
        ],
        ["Book not found: FakeBook"]
      )
    );

    render(() => <Main />);

    const input = screen.getByPlaceholderText(/search verses/i);
    fireEvent.input(input, { target: { value: "John 3:16; FakeBook 1:1" } });
    fireEvent.keyDown(input, { key: "Enter" });

    // Both the verse content and the error should appear
    expect(
      await screen.findByText("For God so loved the world...")
    ).toBeInTheDocument();
    expect(screen.getByText("Book not found: FakeBook")).toBeInTheDocument();
  });

  it("clears errors when a new search returns no errors", async () => {
    mockSearchVerse.mockReset();

    // First search returns errors
    mockSearchVerse.mockResolvedValueOnce(
      makeRef([], ["Book not found: FakeBook"])
    );

    render(() => <Main />);

    const input = screen.getByPlaceholderText(/search verses/i);
    fireEvent.input(input, { target: { value: "FakeBook 1:1" } });
    fireEvent.keyDown(input, { key: "Enter" });

    expect(
      await screen.findByText("Book not found: FakeBook")
    ).toBeInTheDocument();

    // Second search returns clean results
    mockSearchVerse.mockResolvedValueOnce(
      makeRef(
        [
          {
            Book: "John",
            Chapter: 3,
            StartVerse: 16,
            EndVerse: 16,
            FullText: [
              { Number: 16, Text: "For God so loved the world..." } as reference.Verse,
            ],
          },
        ],
        []
      )
    );

    fireEvent.input(input, { target: { value: "John 3:16" } });
    fireEvent.keyDown(input, { key: "Enter" });

    expect(
      await screen.findByText("For God so loved the world...")
    ).toBeInTheDocument();
    expect(
      screen.queryByText("Book not found: FakeBook")
    ).not.toBeInTheDocument();
  });
});
