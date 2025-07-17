# 🎯 Targeting Engine - GreedyGame Backend Assignment

This project implements a **targeting engine microservice** that serves active ad campaigns based on targeting rules for app, country, and OS dimensions. The service is optimized for **read-heavy performance**, supports **inclusion/exclusion targeting**, and is fully integrated with **Supabase PostgreSQL**.

---

## 🚀 Features

- ✅ Campaign targeting based on app, country, OS
- ✅ Inclusion and exclusion targeting rules
- ✅ Returns multiple matched campaigns
- ✅ Filters out inactive campaigns
- ✅ RESTful GET endpoint with proper status codes
- ✅ Built using Golang, Supabase PostgreSQL
- ✅ Logs rule evaluations for debugging

---

## 📦 Tech Stack

- **Language:** Go 1.21+
- **Database:** Supabase (PostgreSQL)
- **Dependencies:** 
  - `pgx` for DB interaction
  - `net/http` for API
  - `encoding/json` for response formatting

---

## 📄 API Spec

### `GET /v1/delivery`

#### 🔹 Query Parameters (required)
| Param   | Type   | Description                        |
|---------|--------|------------------------------------|
| `app`   | string | App ID (e.g. `com.gametion.ludokinggame`) |
| `country` | string | Country code (e.g. `us`, `india`) |
| `os`    | string | OS name (e.g. `android`, `ios`)    |

#### 🔹 Response

- `200 OK`: Returns list of matched campaigns.
- `204 No Content`: No campaign matched.
- `400 Bad Request`: Missing query parameters.

#### 🔸 Example

```http
GET /v1/delivery?app=com.gametion.ludokinggame&country=us&os=android
```

**Response:**

```json
[
  {
    "cid": "8f5bfc02-bcc1-41fc-9238-ecbe8392d2a2",
    "img": "https://image.spotify.com",
    "cta": "Listen Now"
  },
  {
    "cid": "6c01b5c6-90a9-45aa-b3d4-bde219624702",
    "img": "https://image.subwaysurfer.com",
    "cta": "Play Now"
  }
]
```

---

## 🗃️ Database Schema

### `campaigns`

| Column | Type | Description            |
| ------ | ---- | ---------------------- |
| id     | UUID | Primary Key            |
| name   | TEXT | Campaign name          |
| image  | TEXT | Image URL              |
| cta    | TEXT | Call To Action text    |
| status | TEXT | `ACTIVE` or `INACTIVE` |

### `targeting_rules`

| Column       | Type    | Description                      |
| ------------ | ------- | -------------------------------- |
| id           | UUID    | Primary Key                      |
| campaign\_id | UUID    | FK to `campaigns(id)`            |
| dimension    | TEXT    | `app`, `country`, `os`           |
| rule\_type   | TEXT    | `INCLUDE` or `EXCLUDE`           |
| values       | TEXT\[] | Array of values to match/exclude |

---

## 🛠️ Setup

### 1. Clone the Repo

```bash
git clone https://github.com/your-username/targeting-engine.git
cd targeting-engine
```

### 2. Configure Supabase

Update your `db/db.go` with the correct Supabase connection URL.

```go
conn, err := pgxpool.New(context.Background(), os.Getenv("SUPABASE_DB_URL"))
```

> 💡 Tip: Use `.env` for local secrets

---

### 3. Run the Server

```bash
go run main.go
```

> Server runs on: [http://localhost:8080](http://localhost:8080)

---

### 4. Test the API

Use Postman or cURL:

```bash
curl "http://localhost:8080/v1/delivery?app=com.gametion.ludokinggame&country=us&os=android"
```

---

## ✅ Done

* [x] Fully functional GET API
* [x] Matching logic for INCLUDE/EXCLUDE rules
* [x] Integration with Supabase
* [x] Logs with match debugging
* [x] 204 and 400 status handling

---

## 🧪 Future Improvements

* Add Prometheus metrics for performance tracking
* Integrate Redis for rule-level caching
* Build a UI for campaign/rule management
* Add tests using `testing` package

---

## 🔐 Confidential

This project is a **private submission for GreedyGame assignment** and is not to be made public.

---

## 👤 Author

**Vignesh Prabhu**  
📧 [vigneshprabhu1140@gmail.com](mailto:vigneshprabhu1140@gmail.com)  
🌐 [linkedin.com/in/vigneshprabhu](https://linkedin.com/in/vigneshprabhu)

```http
GET /v1/delivery?app=com.gametion.ludokinggame&country=us&os=android
