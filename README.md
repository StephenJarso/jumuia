

# 📝 Jumuia – Full Project Documentation

## 1️⃣ Project Overview

**Project Name:** Jumuia
**Meaning:** “Community” (Swahili)
**Primary Users:** VSLA Leaders & NGO / Relief Officers

**Mission:**

> Jumuia is a digital platform for village savings and loan associations that combines core financial management with agricultural risk tracking and disaster / relief awareness.
> It enables communities to track savings, loans, and repayments, while allowing NGOs to monitor financial vulnerability and provide timely support.

---

## 2️⃣ Project Goals

1. **Core VSLA functionality**

   * Replace paper ledgers with a digital record
   * Manage savings, loans, and repayments
   * Enable transparency and accountability

2. **Agricultural Risk Management**

   * Track farming-related loans by crop and season
   * Identify members at risk due to failed harvests
   * Provide leaders with actionable insights

3. **Disaster / Relief Layer**

   * Flag members affected by shocks (floods, droughts, fire, etc.)
   * Pause or reschedule loans during emergencies
   * Provide NGOs with read-only access for targeted support

---

## 3️⃣ Core System Layers

```text id="3kz2xp"
Layer 1: Core VSLA (always active)
Layer 2: Agriculture (focus on farming loans and risk)
Layer 3: Relief/Disaster (optional, activated when shocks occur)
```

* Each layer **builds on the previous**, allowing incremental development.
* Agriculture layer comes first for initial use-case focus.

---

## 4️⃣ Users & Roles

| Role        | Can Edit                                      | Can View                                   | Notes                                     |
| ----------- | --------------------------------------------- | ------------------------------------------ | ----------------------------------------- |
| VSLA Leader | ✔ Groups, Members, Loans, Savings, Repayments | ✔ Reports                                  | Manages group finances                    |
| Member      | ❌                                             | Limited personal data                      | Optional login later                      |
| NGO Officer | ❌                                             | ✔ Groups, Members, Loans, Disaster reports | Read-only access to monitor vulnerability |

---

## 5️⃣ Database Design (Full Project)

### Core Tables

#### 1. `users`

* id (PK)
* name
* email
* password_hash (future auth)
* role (`leader`, `ngo`, `admin`)

#### 2. `groups`

* id (PK)
* name
* location
* created_at

#### 3. `members`

* id (PK)
* group_id (FK → groups)
* name
* phone
* livelihood_type (`farmer`, `trader`, `mixed`)
* role (`leader`, `member`)

---

### Financial Tables (Core VSLA)

#### 4. `savings`

* id (PK)
* member_id (FK → members)
* amount
* date

#### 5. `loans`

* id (PK)
* member_id (FK → members)
* season_id (FK → seasons)
* amount
* purpose (`seed`, `fertilizer`, `tools`, `trade`, `school fees`)
* crop (nullable, only for agriculture)
* status (`active`, `paid`, `rescheduled`, `at_risk`)

#### 6. `repayments`

* id (PK)
* loan_id (FK → loans)
* amount
* date

---

### Agriculture Layer

#### 7. `seasons`

* id (PK)
* name (`2026 Long Rains`)
* start_date
* end_date

---

### Relief / Disaster Layer

#### 8. `disasters`

* id (PK)
* type (`flood`, `drought`, `fire`, `other`)
* start_date
* end_date
* description

#### 9. `member_disaster_status`

* member_id (FK → members)
* disaster_id (FK → disasters)
* impact_level (`low`, `medium`, `high`)

---

## 6️⃣ MVP Feature Set (Full Project)

### Phase 1 – Agriculture Layer

* Create group
* Add members
* Create season
* Issue agricultural loans
* Mark loans “at risk”
* View loan list by season

### Phase 2 – Core VSLA Layer

* Record savings
* Record repayments
* Balance tracking per member & group
* Reports for leaders

### Phase 3 – Relief / Disaster Layer

* Flag affected members
* Pause / reschedule loans
* NGO read-only access
* Export reports (CSV/PDF)

---

## 7️⃣ Technical Stack

| Layer           | Technology                                                      |
| --------------- | --------------------------------------------------------------- |
| Backend         | Go (Golang) with `net/http` or Gin later                        |
| Frontend        | Phase 1: HTML/CSS + minimal JS <br> Phase 2: React + TypeScript |
| Database        | PostgreSQL                                                      |
| Version Control | Git + GitHub                                                    |
| Deployment      | Local / VPS / eventually cloud (Heroku, Railway, etc.)          |

---

## 8️⃣ Folder Structure (Go Project)

```text id="7uq2jm"
jumuia/
│
├── cmd/
│   └── server/
│       └── main.go
│
├── internal/
│   ├── handlers/      // HTTP handlers
│   ├── services/      // Business logic
│   ├── models/        // Database models
│   └── repository/    // DB queries
│
├── web/
│   ├── templates/     // HTML files
│   └── static/        // CSS / JS
│
├── migrations/        // SQL migration files
├── go.mod
└── README.md
```

---

## 9️⃣ Development Roadmap

### Week 1 – Agriculture Layer MVP

* Go server setup
* Database connection
* Groups, Members, Seasons
* Agricultural loans + risk flagging

### Week 2 – Core VSLA Layer

* Add savings
* Add repayments
* Simple balances & reports

### Week 3 – Relief Layer

* Add disasters
* Member disaster status
* NGO read-only dashboard
* Export reports

### Week 4 – Frontend & Polish

* Improve HTML/CSS
* Form validation
* Optional React migration

---

## 10️⃣ Funding & Pitch Notes

**Pitch for funders:**

> “Village savings groups are critical for rural communities. Jumuia digitizes VSLA operations and adds agricultural risk and disaster tracking. It helps leaders manage seasonal loans and allows NGOs to monitor financial vulnerability, enabling faster recovery after shocks.”

* Agriculture = climate risk + loan exposure
* Relief = disaster preparedness + targeted assistance
* Core VSLA = transparency + trust

This makes it **fundable, real-world, and scalable**.

---
