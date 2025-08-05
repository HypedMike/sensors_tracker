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
  - `cd server`
  - diversi modi per lanciare il server:
    - per lanciarlo con seeder `go run . seed` altrimenti `go run .`
    - per lanciarlo con Docker `docker compose up` (questo avvierà il seeder di default)
  - per rilanciare il seeder manualmente basta avviare il client e lanciare il comando `npm run seed` (assicurati che il server sia in esecuzione sull'url http://localhost:8080)
- front:
    - `cd client`
    - `npm install`
    - per lanciarlo in modo sviluppo `npm run dev` altrimenti `npm run build` e poi `npm run preview`

#### Per accedere:
- username: `admin`
- password: `admin`

## Considerazioni
- Il progetto è stato sviluppato in Vue 3 con Vite e Typescript, utilizzando Pinia per la gestione dello stato dell'utente loggato e TailwindCSS per lo stile.
- Il server è stato sviluppato in Go utilizzando il framework Gin e MongoDB come database. Ho utilizzato la libreria standard di Mongo per la connessione e le operazioni CRUD.
- Ho utilizzato JWT che durano 1 giorno per l'autenticazione degli utenti senza refresh token.
- Il progetto è containerizzato con Docker, con un file `docker-compose.yaml` che definisce i servizi per il server e MongoDB.
- Ho implementato un seeder per popolare il database con dati di esempio all'avvio del server.
- Non ho avuto tempo per implementare delle animazioni per rendere il front più fluido, ho però iniziato un'implementazione responsive per mobile.
- per facilità le statistiche sono all'endpoint `/api/stats`

## Timesheet: 15h 30m
- Autenticazione: **30m**
- Endpoint Backend: **6h**
- Frontend: **4h**
  - Layout: 1h
  - Grafici: 1h
  - Modelli, API e gestione stato: 2h
- Docker, docker compose e seeder: **2h**
- Bug fix e miglioramenti: **1h**
- Documentazione: **1h**
- Deploy su fly.io: **1h**