## v0.4.15-alpha-cs1.1.2
- Removed remaining "CasaOS" references from code and documentation
- Updated README to reflect CassetteOS naming

## v0.4.15-alpha1-cs1.1.1
This patch version updates the Go environment to 1.20 to ensure consistent behavior in CI and go-licenses tooling.

## v0.4.15-alpha1-cs1.1.0

### Added
- Add GitHub Action `ci.yml` to run tests on push to `main` and `develop` branches.

### Changed
- Transferred repository ownership from `mayumigit` to organization `BeesNestInc`.
- Renamed repository from `CasaOS-Common` to `CassetteOS-Common`.
- Updated module paths to reflect the new repository location.
  - e.g., `github.com/mayumigit/CasaOS-Common` → `github.com/BeesNestInc/CassetteOS-Common`

## v0.4.15-alpha1-cs1.0.0
- Based on CasaOS v0.4.15
- - Replaced module paths to use our own GitHub fork instead of the original IceWhaleTech repository.
  (e.g., `github.com/IceWhaleTech/CasaOS-Common` → `github.com/mayumigit/CasaOS-Common`)
