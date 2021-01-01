# Cozy Point of Sales

Cozy Point of Sales application written in Go (backend) and Vue 3 TS (frontend).

## Backend
1. `go get`
```
PASSWORD=mypassword #generated from /generate
```
3. `go build && ./cozypos-full`

## Frontend
1. Create `.env` file containing `VUE_APP_BASE_URL=http://yourapiurl:yourport`
2. `cd` to `frontend` dir
3. Run `npm i` to install dependencies
4. Run `npm run build`, folder will be in `dist`
