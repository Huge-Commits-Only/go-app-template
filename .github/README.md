<p align="center">
  <img src="assets/banner.png" alt="go-app-template — nowy projekt Huge Commits Only" width="100%">
</p>

<p align="center">
  <a href="https://github.com/Huge-Commits-Only/go-app-template/generate"><img src="https://img.shields.io/badge/Use%20this%20template-238636?style=for-the-badge&logo=github&logoColor=white" alt="Use this template"></a>
</p>

<p align="center">
  <a href="https://github.com/Huge-Commits-Only/go-app-template/actions/workflows/unit-test.yml"><img src="https://img.shields.io/github/actions/workflow/status/Huge-Commits-Only/go-app-template/unit-test.yml?branch=main&style=flat-square&label=testy" alt="Testy"></a>
  <a href="https://github.com/Huge-Commits-Only/go-app-template/actions/workflows/codeql.yml"><img src="https://img.shields.io/github/actions/workflow/status/Huge-Commits-Only/go-app-template/codeql.yml?branch=main&style=flat-square&label=CodeQL" alt="CodeQL"></a>
  <a href="https://scorecard.dev/viewer/?uri=github.com/Huge-Commits-Only/go-app-template"><img src="https://img.shields.io/ossf-scorecard/github.com/Huge-Commits-Only/go-app-template?style=flat-square&label=OpenSSF%20Scorecard" alt="OpenSSF Scorecard"></a>
  <img src="https://img.shields.io/badge/Go-1.27-00ADD8?style=flat-square&logo=go&logoColor=white" alt="Go 1.27">
</p>

<p align="center">
  <img src="assets/demo.gif" alt="Od „Use this template” przez testy, CodeQL i skany do podpisanego obrazu" width="100%">
</p>

Start każdego nowego projektu w **Huge Commits Only**: aplikacja w Go z testami, obraz Dockera dla Raspberry Pi i komplet CI — na zielono od pierwszego commita.

## ✨ Co dostajesz
- **🎯 Go** — minimalny serwer HTTP z `/healthz` i testami (pokrycie powyżej 80%); `app -health` to health check kontenera.
- **🐳 Docker** — kompilacja krzyżowa amd64 / arm64, obraz `scratch`, użytkownik bez uprawnień, obraz bazowy przypięty do skrótu.
- **🔒 Bezpieczeństwo** — CodeQL `security-extended`, govulncheck, Trivy, OpenSSF Scorecard; akcje przypięte do SHA, minimalne uprawnienia.
- **🔏 Podpisane obrazy** — atestacja pochodzenia (SLSA) i SBOM przy każdym merge'u.
- **🤖 Porządek** — Dependabot co tydzień, etykiety z nazwy brancha, autor i reviewer w PR, szkic wydania.

## 🚀 Nowy projekt
- [ ] **[Use this template](https://github.com/Huge-Commits-Only/go-app-template/generate)**, właściciel `Huge-Commits-Only`. Zacznij od repozytorium **prywatnego**; publiczne dopiero po przejrzeniu pod kątem sekretów i danych osobowych.
- [ ] Zmień nazwę modułu w `go.mod` (`module app`) i katalogu `cmd/app`, a w `Dockerfile` ścieżkę `./cmd/app`.
- [ ] Nadaj aplikacji **własną identyfikację**: logo, okładka w `.github/assets/`, GIF z demo i ten README.
- [ ] **Settings → Rules → Rulesets**: ochrona `main` (tylko PR ze squashem, bez force-pusha i usuwania, wymagane testy CI). Szablon nie przenosi ustawień repozytorium; dla repo prywatnych wymaga planu płatnego.
- [ ] **Settings → Social preview**: obrazek 1280 × 640 w stylu aplikacji.
- [ ] Repozytorium publiczne: po pierwszej publikacji ustaw obraz w **Packages** jako publiczny.

```bash
go vet ./... && go test ./...
PORT=18080 go run ./cmd/app   # http://localhost:18080/
```

<details>
<summary><b>🔄 Workflowy</b></summary>
<br>

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
| Dependabot Auto-Merge | `dependabot-auto-merge.yml` | PR od Dependabota | Łatki i drobne wersje scalają się same po przejściu wymaganych checków; duże wersje czekają na człowieka. Merge tokenem workflowu nie uruchamia workflowów na `main` — obraz z nową zależnością powstaje przy następnym merge'u |
| PR Automation | `pr-automation.yml` | nowy PR | Autor, reviewer i etykiety z nazwy brancha |
| Release Draft | `release-draft.yml` | `main` | Szkic notatek wydania |
| Docker Image Publish | `docker-build.yml` | `main` | Obraz amd64 + arm64 w GHCR z atestacją pochodzenia i SBOM (atestacja: repo publiczne) |

Sam szablon niczego nie publikuje. Etykiety z nazwy brancha: `feat/*` → `enhancement` · `fix/*` → `bug` · `docs/*` → `documentation` · `chore/*` → `chore` · `deps/*` → `dependencies`

</details>

---

<p align="center"><sub>Część <a href="https://github.com/Huge-Commits-Only"><b>Huge Commits Only</b></a> · licencja <a href="../LICENSE">MIT</a></sub></p>
