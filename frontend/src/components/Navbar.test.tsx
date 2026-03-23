import { render, screen } from "@solidjs/testing-library";
import { describe, expect, it } from "vitest";
import Navbar from "./Navbar";

describe("Navbar", () => {
  it("renders the ReadChaptersIcon with the lucide-book-marked id", () => {
    const { container } = render(() => <Navbar />);

    expect(screen.getByText("Jonah")).toBeInTheDocument();
    expect(container.querySelector("#lucide-book-marked")).toBeInTheDocument();
  });
});
