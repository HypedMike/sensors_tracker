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
  - diversi modi per runnare il server:
    - per runnarlo con seeder `go run . seed` altrimenti `go run .`
    - per runnarlo con Docker `docker compose up` (questo avvierà il seeder di default)
  - per rilanciare il seeder manualmente basta avviare il client e lanciare il comando `npm run seed` (assicurati che il server sia in esecuzione sull'url http://localhost:8080)
- front:
    - `cd client`
    - `npm install`
    - per runnarlo in modo sviluppo `npm run dev` altrimenti `npm run build` e poi `npm run preview`

#### Per accedere:
- username: `admin`
- password: `admin`