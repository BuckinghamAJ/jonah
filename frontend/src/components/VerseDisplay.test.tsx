import { render, screen } from "@solidjs/testing-library";
import { describe, expect, it } from "vitest";
import VerseDisplay from "./VerseDisplay";

describe("VerseDisplay", () => {
  const baseVerse = {
    Book: "Jonah",
    Chapter: 1,
    StartVerse: 1,
    EndVerse: 1,
  };

  it("does not throw and does not render when FullText is null", () => {
    expect(() =>
      render(() => <VerseDisplay verse={{ ...baseVerse, FullText: null } as any} />)
    ).not.toThrow();

    expect(document.querySelector(".verse-reference")).toBeNull();
  });

  it("does not throw and does not render when FullText is undefined", () => {
    expect(() =>
      render(() => <VerseDisplay verse={{ ...baseVerse, FullText: undefined } as any} />)
    ).not.toThrow();

    expect(document.querySelector(".verse-reference")).toBeNull();
  });

  it("does not render when FullText is empty array", () => {
    render(() => <VerseDisplay verse={{ ...baseVerse, FullText: [] } as any} />);
    expect(document.querySelector(".verse-reference")).toBeNull();
  });

  it("renders when FullText has verse entries", () => {
    render(() => (
      <VerseDisplay
        verse={
          {
            ...baseVerse,
            FullText: [{ Number: 1, Text: "The word of the Lord came to Jonah." }],
          } as any
        }
      />
    ));

    expect(screen.getByText("Jonah 1:1")).toBeInTheDocument();
    expect(screen.getByText("The word of the Lord came to Jonah.")).toBeInTheDocument();
  });
});
