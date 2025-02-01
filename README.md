# 📸 Photo Salon – A Hub for Photography Clubs & Competitions

Photo Salon is a modern web app designed for photographers who want to organize photo clubs, host competitions, and conduct critiques in a structured and engaging way. Inspired by the rich history of art salons, where artists gathered to showcase and refine their work, Photo Salon brings this tradition into the digital age—creating a space where photographers can connect, learn, and grow through community-driven feedback and friendly competition.

## 🚀 Why Photo Salon?
- Inspired by Tradition – Reviving the spirit of classic art salons in a modern, digital space.
- Built for Photographers – A tool by photographers, for photographers, designed to enhance skills through critique and competition.
- Seamless & Intuitive – An easy-to-use platform to organize, compete, and connect in the photography world.

![Postgres](https://img.shields.io/badge/postgres-%23316192.svg?style=for-the-badge&logo=postgresql&logoColor=white)
![Next JS](https://img.shields.io/badge/Next-black?style=for-the-badge&logo=next.js&logoColor=white)
![PNPM](https://img.shields.io/badge/pnpm-%234a4a4a.svg?style=for-the-badge&logo=pnpm&logoColor=f69220)
![React](https://img.shields.io/badge/react-%2320232a.svg?style=for-the-badge&logo=react&logoColor=%2361DAFB)
![React Hook Form](https://img.shields.io/badge/React%20Hook%20Form-%23EC5990.svg?style=for-the-badge&logo=reacthookform&logoColor=white)
![TailwindCSS](https://img.shields.io/badge/tailwindcss-%2338B2AC.svg?style=for-the-badge&logo=tailwind-css&logoColor=white)
![Zod](https://img.shields.io/badge/zod-%233068b7.svg?style=for-the-badge&logo=zod&logoColor=white)
![TypeScript](https://img.shields.io/badge/typescript-%23007ACC.svg?style=for-the-badge&logo=typescript&logoColor=white)
![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white)

## Table of contents
* [Installation Quick Guide](#installation-quick-guide)
* [Quick Start](#quick-start)
* [Helpful VSCode Extensions](#vscode-extensions)
* [Installation Links](#installation-links)

## Installation Quick Guide
Assuming you already have Node, go, and Postgresql installed
> &nbsp;
> NOTE: Use `brew list` to check your casks and and formulae already installed
> &nbsp;

#### Install [golang-migrate cli](https://github.com/golang-migrate/migrate?tab=readme-ov-file#migrate), [pnpm](https://pnpm.io/), [buf](https://buf.build), [sqlc](https://docs.sqlc.dev/en/latest/), and [minio](https://min.io/docs/minio/macos/index.html)
```shell
brew install golang-migrate
brew install pnpm
brew install bufbuild/buf/buf
brew install sqlc
brew install minio/stable/minio
```

## Quick Start
#### 1. Install server go modules
```shell
go mod tidy
```
#### 2. Next, install protoc-gen-connect-go
```shell
go install connectrpc.com/connect/cmd/protoc-gen-connect-go@latest
[ -n "$(go env GOBIN)" ] && export PATH="$(go env GOBIN):${PATH}"
[ -n "$(go env GOPATH)" ] && export PATH="$(go env GOPATH)/bin:${PATH}"
```
#### 3. Ensure the following is included in you `.zshrc` or related profile
```
# .zshrc or .bash_profile or whatever
alias air='~/go/bin/air'
export PATH=$PATH:$(go env GOPATH)/bin
```
#### 4. Now install client node modules
```shell
pnpm install --dir client
```
#### 5. You'll need to create a database for **salon** and grant permissions to admin

```shell
CREATE DATABASE salon;
GRANT ALL PRIVILEGES ON DATABASE salon TO admin;
GRANT ALL ON SCHEMA public TO admin;
```
Depending on the name you choose for the database you'll need to update the reference in 2 places:
```shell
# /server/.env
export POSTGRES_URL=postgresql://admin:admin@localhost:5432/<DB_NAME>?sslmode=disable
```

```Makefile
# Makefile
models:
	pg_dump --schema-only <DB_NAME> > server/schema.sql
	sqlc generate -f server/sqlc.yaml
```
#### 6. Create account with Mailtrap that you'll use to test emails (important for sign-up/in flows)
1. [Sign-up](https://mailtrap.io/register/signup) with Mailtrap
2. Navigate to [email testing](https://mailtrap.io/inboxes)
3. Add a project
4. Add an inbox (Not sure if this happens automatically)
5. Copy username and password for your project and set them to `SMTP_USERNAME` and `SMTP_PASSWORD` respectively inside your `/server/.env`
#### 7. Export config env file for minio
```shell
export MINIO_CONFIG_ENV_FILE=/etc/default/minio
```
#### 8. Run the make commands for setting things up
> &nbsp;
> NOTE: depending on where you want your minio store to be you'll need to update the Makefile
> &nbsp;

```Makefile
minio:
  minio server  --console-address :9001 ./data
```
```shell
make minio
make migrate-up
make models
make codegen
```
#### 9. Run the application 🎉
```
make dev
```

# Other things

## VSCode Extensions
- [Code Spell Checker](https://marketplace.visualstudio.com/items?itemName=streetsidesoftware.code-spell-checker)
- [Tailwind IntelliSense](https://marketplace.visualstudio.com/items?itemName=bradlc.vscode-tailwindcss)
- [ESLint](https://marketplace.visualstudio.com/items?itemName=dbaeumer.vscode-eslint)
- [Go](https://marketplace.visualstudio.com/items?itemName=golang.go)
- [Prettier](https://marketplace.visualstudio.com/items?itemName=esbenp.prettier-vscode)
- [Github Style Markdown](https://marketplace.visualstudio.com/items?itemName=bierner.markdown-preview-github-styles)

## Installation Links
You'll probably need to install a couple of tools before running the app:
- go -- [Install steps](https://go.dev/doc/install)
- pnpm -- [Install steps](https://pnpm.io/installation)
- postgreSQL -- [Install steps](https://www.postgresql.org/download/)
- psql -- should be included with postgreSQL installation
- bufbuild/buf/buf -- [Install steps](https://buf.build/docs/installation/)
- sqlc -- [Install steps](https://docs.sqlc.dev/en/latest/overview/install.html)
- golang-migrate -- [Install steps](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate#installation)
