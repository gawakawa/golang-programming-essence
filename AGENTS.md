# Repository Guidelines

## Project Structure & Module Organization
- Root Go module: `go.mod` (Go 1.24.x toolchain). Submodule in `chapter09/` with its own `go.mod`.
- Examples by chapter: `chapter04`–`chapter09`. Assets live under `static/` (e.g., `chapter05/fileserve/static/`).
- Binaries/examples: run packages under each chapter (e.g., `chapter05/net-http`).

## Build, Test, and Development Commands
- Environment: `direnv` + `.envrc` auto-loads the Nix dev shell; just `cd` into the repo and `direnv allow` once.
- Format: `go fmt ./...` or `gofmt -s -w .` (must be clean before commit/CI).
- Lint: `go vet ./...` (primary CI gate; fix all findings).
- Test: `go test ./...` • Chapter 09: `cd chapter09 && go test ./...`
- Run example: `go run ./chapter05/net-http` • Chapter 09 apps: `cd chapter09 && go run ./cmd/aozora-search`

## CI Checklist
- Format: `go fmt ./...` then verify clean tree: `git diff --quiet`.
- Vet: `go vet ./...` • Chapter 09: `cd chapter09 && go vet ./...`.
- Test: `go test -race -cover ./...` • Chapter 09: `cd chapter09 && go test -race -cover ./...`.
- Optional (local): `go test -bench=. ./chapter08/benchmark` for performance checks.
- Commit: start with a gitmoji and a one-line English summary.

## Coding Style & Naming Conventions
- Formatting: run `gofmt -s -w .` (CI expects formatted code). Lint basics: `go vet ./...`.
- Indentation: tabs (Go default). Line length: be reasonable; wrap if needed.
- Naming: packages lower-case (no caps), exported identifiers `CamelCase`, unexported `camelCase`. File names use underscores if helpful (e.g., `calc_test.go`).

## Testing Guidelines
- Framework: standard `testing` package; benchmarks in `*_test.go` (see `chapter08/benchmark`).
- Naming: test files `*_test.go`; functions `TestXxx`, `BenchmarkXxx`.
- Coverage (local): `go test -cover ./...`.

## Commit & Pull Request Guidelines
- Commit messages: prefix with a gitmoji emoji + one concise English sentence on one line. Example: `✨ Add parallel calc and tests`.
- PRs: clear description, link issues, include run/usage steps; for UI/assets, add screenshots; keep diffs focused.

## Security & Configuration Tips
- Secrets: never commit credentials. Use env vars or local files ignored by Git.
- Environment: rely on `direnv` and `flake.nix` for reproducible tools; avoid committing local configs.

## Gitmoji Reference (gitmoji -l)
🎨 - :art: - Improve structure / format of the code.
⚡️ - :zap: - Improve performance.
🔥 - :fire: - Remove code or files.
🐛 - :bug: - Fix a bug.
🚑️ - :ambulance: - Critical hotfix.
✨ - :sparkles: - Introduce new features.
📝 - :memo: - Add or update documentation.
🚀 - :rocket: - Deploy stuff.
💄 - :lipstick: - Add or update the UI and style files.
🎉 - :tada: - Begin a project.
✅ - :white_check_mark: - Add, update, or pass tests.
🔒️ - :lock: - Fix security or privacy issues.
🔐 - :closed_lock_with_key: - Add or update secrets.
🔖 - :bookmark: - Release / Version tags.
🚨 - :rotating_light: - Fix compiler / linter warnings.
🚧 - :construction: - Work in progress.
💚 - :green_heart: - Fix CI Build.
⬇️ - :arrow_down: - Downgrade dependencies.
⬆️ - :arrow_up: - Upgrade dependencies.
📌 - :pushpin: - Pin dependencies to specific versions.
👷 - :construction_worker: - Add or update CI build system.
📈 - :chart_with_upwards_trend: - Add or update analytics or track code.
♻️ - :recycle: - Refactor code.
➕ - :heavy_plus_sign: - Add a dependency.
➖ - :heavy_minus_sign: - Remove a dependency.
🔧 - :wrench: - Add or update configuration files.
🔨 - :hammer: - Add or update development scripts.
🌐 - :globe_with_meridians: - Internationalization and localization.
✏️ - :pencil2: - Fix typos.
💩 - :poop: - Write bad code that needs to be improved.
⏪️ - :rewind: - Revert changes.
🔀 - :twisted_rightwards_arrows: - Merge branches.
📦️ - :package: - Add or update compiled files or packages.
👽️ - :alien: - Update code due to external API changes.
🚚 - :truck: - Move or rename resources (e.g.: files, paths, routes).
📄 - :page_facing_up: - Add or update license.
💥 - :boom: - Introduce breaking changes.
🍱 - :bento: - Add or update assets.
♿️ - :wheelchair: - Improve accessibility.
💡 - :bulb: - Add or update comments in source code.
🍻 - :beers: - Write code drunkenly.
💬 - :speech_balloon: - Add or update text and literals.
🗃️ - :card_file_box: - Perform database related changes.
🔊 - :loud_sound: - Add or update logs.
🔇 - :mute: - Remove logs.
👥 - :busts_in_silhouette: - Add or update contributor(s).
🚸 - :children_crossing: - Improve user experience / usability.
🏗️ - :building_construction: - Make architectural changes.
📱 - :iphone: - Work on responsive design.
🤡 - :clown_face: - Mock things.
🥚 - :egg: - Add or update an easter egg.
🙈 - :see_no_evil: - Add or update a .gitignore file.
📸 - :camera_flash: - Add or update snapshots.
⚗️ - :alembic: - Perform experiments.
🔍️ - :mag: - Improve SEO.
🏷️ - :label: - Add or update types.
🌱 - :seedling: - Add or update seed files.
🚩 - :triangular_flag_on_post: - Add, update, or remove feature flags.
🥅 - :goal_net: - Catch errors.
💫 - :dizzy: - Add or update animations and transitions.
🗑️ - :wastebasket: - Deprecate code that needs to be cleaned up.
🛂 - :passport_control: - Work on code related to authorization, roles and permissions.
🩹 - :adhesive_bandage: - Simple fix for a non-critical issue.
🧐 - :monocle_face: - Data exploration/inspection.
⚰️ - :coffin: - Remove dead code.
🧪 - :test_tube: - Add a failing test.
👔 - :necktie: - Add or update business logic.
🩺 - :stethoscope: - Add or update healthcheck.
🧱 - :bricks: - Infrastructure related changes.
🧑‍💻 - :technologist: - Improve developer experience.
💸 - :money_with_wings: - Add sponsorships or money related infrastructure.
🧵 - :thread: - Add or update code related to multithreading or concurrency.
🦺 - :safety_vest: - Add or update code related to validation.
