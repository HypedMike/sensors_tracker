**Titolo:** Take-home Full-Stack – IoT Sensors Dashboard (Vue 3 + Tailwind, REST API con MongoDB e Docker)

**Obiettivo**
Realizza una mini-piattaforma per monitorare sensori IoT:

1. un **backend REST** che gestisca dispositivi e misurazioni su **MongoDB** (containerizzato con **Docker**);
2. una **dashboard web** (**Vue 3 + Tailwind**) che elenchi i sensori e mostri l’andamento del sensore selezionato con evidenza delle soglie.

---

### Stack richiesto

* **Backend:** Node.js (JS/TS) **oppure** Go; MongoDB; REST. **Docker obbligatorio.**
  *Bonus:* input validation, test, Docker Compose, servizio di aggregazione/statistiche.
* **Frontend:** **Vue 3 (Composition API)** + **Tailwind CSS**; libreria grafici (Plotly o AmCharts a scelta).
  *Bonus:* TypeScript, Pinia, routing completo.

---

### Dati iniziali (seed)

* Crea **5–8 device** (es. `SEN-001…SEN-008`) con nome, location e **threshold** a tua scelta.
* Genera \~**100 misurazioni per device** distribuite sugli **ultimi 3 giorni**.
* Usa la metrica **`dispmm`** (spostamento in mm).
* Fornisci un **comando unico** (es. `npm run seed` / `go run cmd/seed/main.go` / `make seed`) che inizializza il DB (crea device e inserisce le misure).

---

### API richieste

Implementa i seguenti endpoint minimi:

* **POST `/api/devices`** – crea/registri un device.
* **GET `/api/devices`** – lista devices (supporto a `?sort=` opzionale).
* **POST `/api/measurements`** – inserimento **bulk** di misurazioni.
* **GET `/api/measurements?deviceId=...&from=...&to=...`** – misure per device in un intervallo temporale.
* **Compatibilità FE:** **GET `/api/measurements/{deviceId}`** – restituisce \~**100** misure più recenti (copertura \~**ultimi 3 giorni**).

**Esempio risposta** `GET /api/measurements/SEN-001`:

```json
{
  "deviceId": "SEN-001",
  "measurements": [
    { "timestamp": "2025-05-25T06:00:00Z", "metric": "dispmm", "value": 0.2 },
    { "timestamp": "2025-05-25T06:30:00Z", "metric": "dispmm", "value": 0.4 }
  ]
}
```

**Bonus:** **GET `/api/stats?deviceId=...&metric=dispmm`** → restituisci **avg/min/max** nel range (o sugli ultimi 3 giorni) per la metrica dispmm.

---

### Requisiti frontend

1. **Tabella sensori** con: `id`, `name`, `location`, `lastValue`, `stato`, dove

   * **Stato:** “Alarm” se `lastValue > threshold`, altrimenti “OK”.
   * Consenti **ordinamento** almeno su una colonna.
2. **Grafico a linee** per il sensore selezionato (ultimi **3 giorni**), con:

   * **linea orizzontale** di soglia **oppure** marker/colore per i punti sopra soglia.
3. Il FE deve consumare **le API reali** del backend (**no mock**).
4. Usa **Vue 3 (Composition API)** e **Tailwind**.

---

### Esecuzione

* Backend **obbligatoriamente** in Docker.
* **Consigliato:** `docker-compose.yml` per avviare **MongoDB + API**.
* Il FE può girare con `npm run dev` o in un container separato.
* Gestisci **CORS** oppure un semplice **reverse proxy** per servire `/api`.

---

### Consegna richiesta

**Repository Git** con:

* codice completo FE/BE;
* **Dockerfile** del backend (e Compose se presente);
* **README** con istruzioni chiare di setup/run/test;
* `.env.example` con le variabili richieste;
* **script di seed**;
* breve **nota tecnica** con scelte architetturali, trade-off e tempi effettivi.

---

### Criteri di valutazione (indicativi)

* **Funzionalità base (40%)**: API conformi, FE funzionante (tabella + grafico + soglia).
* **Qualità del codice & architettura (25%)**: struttura moduli, naming, separazione responsabilità; eventuale validation/test.
* **UX/UI & responsiveness (15%)**: pulizia UI, leggibilità grafico, interazioni minime.
* **DevOps (10%)**: Docker/Compose, README.
* **Extra/Bonus (10%)**: stats API, test, Pinia/TS, microservizio di aggregazione, sicurezza base (API key/JWT), rate limiting.

---

### Tempi indicativi

**12–18 ore** complessive (BE 8–12h + FE 4–6h). Indica nel README l’effort effettivo.
Se hai dubbi o devi ridurre il perimetro, descrivi chiaramente le assunzioni nel README.