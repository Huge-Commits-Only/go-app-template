<div align="center">

# 🧨 `go-app-template`

**Szablon nowego projektu Huge Commits Only: Go, Docker i komplet CI od pierwszego commita.**

[![License](https://img.shields.io/github/license/Huge-Commits-Only/go-app-template?style=flat-square&color=A3BE8C)](../LICENSE)
[![Go](https://img.shields.io/badge/Go-1.27-00ADD8?style=flat-square&logo=go&logoColor=white)](../go.mod)
[![Docker](https://img.shields.io/badge/Docker-scratch-2496ED?style=flat-square&logo=docker&logoColor=white)](../Dockerfile)
[![Platform](https://img.shields.io/badge/Platform-Raspberry%20Pi-c51a4a?style=flat-square&logo=raspberrypi&logoColor=white)](https://www.raspberrypi.com/)

</div>

<br>

## 🚀 Nowy projekt z szablonu
1. **Use this template → Create a new repository**, właściciel `Huge-Commits-Only`. Zacznij od repozytorium **prywatnego**; publiczne dopiero po przejrzeniu pod kątem sekretów i danych osobowych.
2. Zmień nazwę modułu w `go.mod` (`module app`) i katalogu `cmd/app`, a w `Dockerfile` ścieżkę `./cmd/app`.
3. Podmień ten README: opis, odznaki (nazwa repozytorium w adresach), sposób uruchomienia.
4. **Settings → Rules → Rulesets**: ochrona `main` jak w `duet` (tylko PR ze squashem, bez force-pusha i usuwania, wymagane testy CI). Szablon nie przenosi ustawień repozytorium. Dla repozytoriów prywatnych wymaga planu płatnego.
5. Repozytorium publiczne: po pierwszej publikacji ustaw obraz w **Packages** jako publiczny (darmowe miejsce bez limitu).

<br>

## 💻 Co jest w środku
| Element | Co robi |
|---|---|
| 🎯 `cmd/app` | Minimalny serwer HTTP z `/healthz` i testami; `app -health` to health check kontenera |
| 🐳 `Dockerfile` | Kompilacja krzyżowa (amd64 / arm64), obraz `scratch`, użytkownik bez uprawnień, obraz bazowy przypięty do skrótu |
| 🔍 `.golangci.yml` | Konfiguracja `golangci-lint` |
| 🤖 `.github/dependabot.yml` | Co tydzień: akcje, obraz bazowy i moduły Go |

<br>

## 🔄 CI/CD
| Workflow | Plik | Kiedy | Co robi |
|---|---|---|---|
| Go Verification | `go-test.yml` | PR, `main` | `go mod verify`, `go mod tidy -diff`, `gofmt`, kompilacja |
| Unit Test Verification | `unit-test.yml` | PR, `main` | `go test -race`, pokrycie i próg minimalny (80%) |
| Lint Verification | `lint-test.yml` | PR, `main` | `golangci-lint` |
| Security Verification | `security-test.yml` | PR, `main` | `govulncheck` + Trivy (podatności, sekrety, konfiguracja) |
| CodeQL Verification | `codeql.yml` | PR, `main`, co tydzień | CodeQL `security-extended` dla Go i workflowów (darmowy w repo publicznych) |
| Container Verification | `container-test.yml` | PR | Buduje obraz, sprawdza `/healthz`, skanuje Trivy |
| Workflow Verification | `workflow-test.yml` | PR, `main` | `actionlint` |
| Scorecard Verification | `scorecard.yml` | `main`, co tydzień | OpenSSF Scorecard (repo publiczne) |
| PR Automation | `pr-automation.yml` | nowy PR | Autor, reviewer i etykiety z nazwy brancha |
| Release Draft | `release-draft.yml` | `main` | Szkic notatek wydania |
| Docker Image Publish | `docker-build.yml` | `main` | Obraz amd64 + arm64 w GHCR z atestacją pochodzenia i SBOM (atestacja: repo publiczne) |

Wszystkie akcje są przypięte do pełnego SHA, a workflowy dostają tylko potrzebne uprawnienia. Sam szablon niczego nie publikuje.

Etykiety z nazwy brancha: `feat/*` → `enhancement` · `fix/*` → `bug` · `docs/*` → `documentation` · `chore/*` → `chore` · `deps/*` → `dependencies`

<br>

## 🛠️ Praca lokalna
```bash
go vet ./... && go test ./...
PORT=18080 go run ./cmd/app   # http://localhost:18080/
```
