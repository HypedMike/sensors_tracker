# Backend

## REST API with Node.js (JS/TS) or Go
+ MongoDB as database (10m --> 10m)
- Must be containerized with Docker (Docker Compose recommended) (10m)
- Endpoints: (2h --> 2h)
  + POST /api/devices: create/register device
  + GET /api/devices: list devices (optional ?sort=)
  + POST /api/measurements: bulk insert measurements
  + GET /api/measurements?deviceId=...&from=...&to=...: measurements for device in time range
  - GET /api/measurements/{deviceId}: ~100 most recent measurements (last 3 days)
  - Bonus: GET /api/stats?deviceId=...&metric=dispmm (avg/min/max)
- Seed script to create 5–8 devices and ~100 measurements/device (last 3 days, metric: dispmm) (1h --> 1h)
- Single command to seed DB (e.g., npm run seed) (30m)
- Bonus: input validation, tests, aggregation/statistics service (2h)

# Frontend

- Vue 3 (Composition API) + Tailwind CSS
- Chart library: Plotly or AmCharts
- Table of sensors: id, name, location, lastValue, status ("Alarm" if lastValue > threshold, else "OK")
- Sortable on at least one column
- Line chart for selected sensor (last 3 days), with threshold line or marker/color for values above threshold
- Must consume real backend APIs (no mocks)
- Bonus: TypeScript, Pinia, full routing

# General/DevOps

- Backend must run in Docker
- Recommended: docker-compose.yml for MongoDB + API
- Frontend can run with npm run dev or in a separate container
- Handle CORS or use a reverse proxy for /api
- Git repo with:
  - Complete FE/BE code
  - Backend Dockerfile (and Compose if present)
  - README with setup/run/test instructions
  - .env.example with required variables
  - Seed script
  - Brief technical note (architecture, trade-offs, time spent)