# Jonah

A desktop Bible reader for the Douay-Rheims Catholic Bible.
Pet project to mess around with SolidJS, Go, and the Wails framework.

## Tech Stack

- **Backend:** Go 1.25, SQLite (go-sqlite3)
- **Frontend:** SolidJS, TypeScript, Tailwind CSS v4
- **Framework:** Wails v2 (Go <-> JS bridge, native desktop builds)

## Project Structure

```
jonah/
  app.go                 -- Wails app controller, exposes methods to frontend
  main.go                -- Entrypoint, Wails window config
  data/                  -- Embedded SQLite seed database (DRC.db)
  migrations/            -- Goose SQL migrations
  internal/
    db/                  -- Database setup, embedded extraction
    drcBible/dto/        -- sqlc-generated query layer (do not edit)
    parser/              -- Bible passage text parser (goparsec)
    reference/           -- Domain types: BibleReference, BiblePassage, Verse
    services/            -- Business logic layer
  frontend/              -- SolidJS + TypeScript + Tailwind CSS v4
    src/
      components/        -- UI components
      routes/            -- Page-level route components
      layouts/           -- Layout wrappers
      lib/               -- Wails RPC query wrappers
    wailsjs/             -- Auto-generated Wails JS bindings (do not edit)
  scripts/               -- Shell build scripts for various platforms
  .github/workflows/     -- CI/CD release workflow
```

## Getting Started

Prerequisites: Go 1.25+, Node 20+, Wails CLI

```bash
# install the wails cli
go install github.com/wailsapp/wails/v2/cmd/wails@latest

# run in dev mode with hot reload
wails dev
```

## License

[MIT](LICENSE)
