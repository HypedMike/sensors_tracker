# Move

## Tech stack
- Go
- MongoDB
- Gin
- Docker
- Vue3 (Vite + Typescript con router)
- Pinia (per state management)
- TailwindCSS
- Autenticazione con JWT

## How to run

### Settare le variabili d'ambiente

- server:
  - `PORT`: porta su cui il server ascolterà (default: 8080)
  - `MONGO_URI`: URI di connessione a MongoDB
  - `JWT_SECRET`: chiave segreta per la generazione dei token JWT
  - `DATABASE_NAME`: nome del database MongoDB

- client:
  - `VITE_API_URL`: URL dell'API del server (default: http://localhost:8080/api)


### Eseguire il progetto
- back:
  - senza docker
    - `cd server`
    - per runnarlo con seeder `go run . seed` altriwise `go run .`
    - per runnarlo con Docker `docker compose up`
  - con docker
    - `docker compose up --build`
- front:
    - `cd client`
    - `npm install`
    - per runnarlo in modo sviluppo `npm run dev` altrimenti `npm run build` e poi `npm run preview`

#### Per accedere:
- username: `admin`
- password: `admin`