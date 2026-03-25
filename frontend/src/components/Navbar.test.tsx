import { query } from "@solidjs/router";
import { MemoryRouter } from "@solidjs/router";
import { Route } from "@solidjs/router";
import { fireEvent, render, screen } from "@solidjs/testing-library";
import { beforeEach, describe, expect, it, vi } from "vitest";
import Navbar from "./Navbar";

const mockGetAllBooks = vi.fn();
const mockGetAllChapters = vi.fn();

vi.mock("../../wailsjs/go/main/App", () => ({
  GetAllBooks: (...args: unknown[]) => mockGetAllBooks(...args),
  GetAllChapters: (...args: unknown[]) => mockGetAllChapters(...args),
}));

vi.mock("./solid-ui/DropdownMenu", () => ({
  DropdownMenu: (props: { children: unknown }) => <div>{props.children}</div>,
  DropdownMenuTrigger: (props: {
    children: unknown;
    [key: string]: unknown;
  }) => <button {...props}>{props.children}</button>,
  DropdownMenuContent: (props: { children: unknown }) => (
    <div>{props.children}</div>
  ),
  DropdownMenuItem: (props: {
    children: unknown;
    onSelect?: (event: { preventDefault: () => void }) => void;
    [key: string]: unknown;
  }) => (
    <button
      onClick={() =>
        props.onSelect?.({
          preventDefault: () => undefined,
        })
      }
    >
      {props.children}
    </button>
  ),
}));

describe("Navbar", () => {
  const renderNavbar = () =>
    render(() => (
      <MemoryRouter>
        <Route path="/" component={Navbar} />
      </MemoryRouter>
    ));

  beforeEach(() => {
    query.clear();
    mockGetAllBooks.mockReset();
    mockGetAllChapters.mockReset();
  });

  it("renders the ReadChapters trigger icon", () => {
    mockGetAllBooks.mockResolvedValue([]);

    const { container } = renderNavbar();

    expect(screen.getByText("Jonah")).toBeInTheDocument();
    expect(container.querySelector("#lucide-book-marked")).toBeInTheDocument();
  });

  it("switches dropdown content to chapters on book selection", async () => {
    mockGetAllBooks.mockResolvedValue([
      { id: 43, name: "John" },
      { id: 19, name: "Psalms" },
    ]);
    mockGetAllChapters.mockResolvedValue([1, 2, 3]);

    renderNavbar();

    const johnItem = await screen.findByText("John");
    expect(screen.getByText("Psalms")).toBeInTheDocument();

    fireEvent.click(johnItem);

    expect(mockGetAllChapters).toHaveBeenCalledWith(43);
    expect(screen.getByText("Back to books")).toBeInTheDocument();
    expect(await screen.findByText("Chapter 1")).toBeInTheDocument();
    expect(screen.getByText("Chapter 2")).toBeInTheDocument();
    expect(screen.getByText("Chapter 3")).toBeInTheDocument();
    expect(screen.queryByText("Psalms")).not.toBeInTheDocument();
  });

  it("returns to books list when back is selected", async () => {
    mockGetAllBooks.mockResolvedValue([
      { id: 43, name: "John" },
      { id: 19, name: "Psalms" },
    ]);
    mockGetAllChapters.mockResolvedValue([1, 2, 3]);

    renderNavbar();

    const johnItem = await screen.findByText("John");
    fireEvent.click(johnItem);

    const backItem = await screen.findByText("Back to books");
    fireEvent.click(backItem);

    expect(screen.getByText("John")).toBeInTheDocument();
    expect(screen.getByText("Psalms")).toBeInTheDocument();
    expect(screen.queryByText("Chapter 1")).not.toBeInTheDocument();
  });

  it("reuses cached chapters for repeated book selection", async () => {
    mockGetAllBooks.mockResolvedValue([{ id: 43, name: "John" }]);
    mockGetAllChapters.mockResolvedValue([1, 2, 3]);

    renderNavbar();

    const johnItem = await screen.findByText("John");

    fireEvent.click(johnItem);
    fireEvent.click(johnItem);

    expect(mockGetAllChapters).toHaveBeenCalledTimes(1);
    expect(mockGetAllChapters).toHaveBeenCalledWith(43);
  });

});
