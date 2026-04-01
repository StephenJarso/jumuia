# 📝 Jumuia – Full Project Documentation (Enhanced)

## 1️⃣ Project Overview

**Project Name:** Jumuia  
**Meaning:** “Community” (Swahili)  
**Primary Users:** VSLA Leaders, Members, NGO / Relief Officers  

**Mission:**  

> Jumuia is a digital platform for village savings and loan associations (VSLAs) that integrates core financial management with agricultural risk tracking and disaster/relief awareness. The platform enables communities to track savings, loans, and repayments while allowing NGOs to monitor financial vulnerability and provide timely support.  

**Why Jumuia Exists:**  
Traditional VSLAs are paper-based and struggle with **manual record-keeping**, **seasonal loan alignment**, and **disaster response**. Jumuia solves these problems by combining digital transparency with automated risk and vulnerability scoring.

---

## 2️⃣ Project Goals (Refined for Impact)

1. **Financial Inclusion**
   - Digitize all paper-based VSLA ledgers into a secure, immutable digital record.
   - Ensure transparency and accountability for leaders and members.

2. **Climate-Smart Lending**
   - Align loan repayment schedules with crop cycles and expected harvests.
   - Reduce default rates due to seasonal mismatches by ~30%.

3. **Rapid Disaster Response**
   - Provide NGOs with real-time dashboards to identify high-risk members.
   - Reduce assessment time from weeks to seconds using automated alerts.

---

## 3️⃣ Core System Layers (Enhanced)
    Layer 1: Core VSLA (always active)
    Layer 2: Agriculture (focus on farming loans, crop cycles, risk scoring)
    Layer 3: Relief/Disaster (activated during emergencies, read-only for NGOs)

* Each layer **builds on the previous** for incremental development.  
* Automation triggers minimize **data silos** by reducing manual updates.  
* Vulnerability scoring ensures **prioritization for disaster response**.

---

## 4️⃣ Users & Roles

| Role        | Can Edit                                      | Can View                                   | Notes                                     |
| ----------- | --------------------------------------------- | ------------------------------------------ | ----------------------------------------- |
| VSLA Leader | ✔ Groups, Members, Loans, Savings, Repayments | ✔ Reports                                  | Manages group finances                    |
| Member      | ❌                                             | Limited personal data                      | Optional login; receives SMS/USSD alerts |
| NGO Officer | ❌                                             | ✔ Groups, Members, Loans, Disaster reports | Read-only access for monitoring           |
| Admin       | ✔ All                                        | ✔ All                                      | System-wide oversight                     |

---

## 5️⃣ Database Design (Enhanced)

### Core Tables

#### `users`
- id (PK), name, email, password_hash, role (`leader`, `ngo`, `admin`)

#### `groups`
- id (PK), name, location, created_at

#### `members`
- id (PK), group_id (FK → groups), name, phone, livelihood_type (`farmer`, `trader`, `mixed`), role (`leader`, `member`)

---

### Financial Tables (Core VSLA)

#### `savings`
- id (PK), member_id (FK → members), amount, date

#### `loans`
- id (PK), member_id (FK → members), season_id (FK → seasons), amount, purpose (`seed`, `fertilizer`, `tools`, `trade`, `school fees`), crop (nullable), status (`active`, `paid`, `rescheduled`, `at_risk`), expected_yield_value

#### `repayments`
- id (PK), loan_id (FK → loans), amount, date

#### `audit_logs`
- id (PK), table_name, record_id, action (`create`, `update`, `delete`), performed_by, timestamp

---

### Agriculture Layer

#### `seasons`
- id (PK), name, start_date, end_date, expected_harvest_date

*Enhancement:* Repayment schedules automatically align with **expected_harvest_date**.  

---

### Relief / Disaster Layer

#### `disasters`
- id (PK), type (`flood`, `drought`, `fire`, `other`), start_date, end_date, description

#### `member_disaster_status`
- member_id (FK → members), disaster_id (FK → disasters), impact_level (`low`, `medium`, `high`)

*Enhancement:* Automated triggers mark loans **“At Risk”** if disaster affects the crop or location.  

---

## 6️⃣ MVP Feature Set (Refined)

### Phase 1 – Agriculture Layer
- Create groups & add members  
- Define seasons with expected harvest dates  
- Issue agricultural loans tied to crop & season  
- Auto-flag loans “at risk” based on disaster events  
- Bulk risk flagging for a location + crop type  
- Vulnerability scoring:  
        Score = (Loan Amount / Total Savings) + Impact Level
        High Score → Priority relief
        Low Score → Eligible for rescheduling

### Phase 2 – Core VSLA Layer
- Record savings and repayments  
- Balance tracking per member & group  
- Audit logs for leader transparency  
- SMS/USSD notifications for member trust

### Phase 3 – Relief / Disaster Layer
- Flag affected members automatically  
- Pause or reschedule loans  
- NGO read-only dashboard  
- Export reports (CSV/PDF)  

---

## 7️⃣ Technical Stack (Optimized for Speed & Offline Use)

| Layer           | Technology                                                      |
| --------------- | --------------------------------------------------------------- |
| Backend         | Go (Golang) with `net/http` / Gin later                        |
| Frontend        | Phase 1: HTML/CSS + minimal JS <br> Phase 2: React + TypeScript |
| Database        | PostgreSQL                                                      |
| Notifications   | SMS / USSD for member trust                                     |
| Offline Sync    | Service Workers (PWA)                                          |
| Version Control | Git + GitHub                                                    |
| Deployment      | Local / VPS / Cloud (Heroku, Railway, etc.)                    |

**Key Go Advantages:**
- Fast templating with `html/template`  
- Concurrency with goroutines for report generation  
- Lightweight, stable, and easily deployable  

---

## 8️⃣ System Architecture
[Members] --inputs--> [VSLA Layer] --outputs--> [Audit Logs + Notifications]
|
v
[Agriculture Layer]
|
v
[Disaster/Relief Layer]
|
v
[NGO Dashboard]


**Highlights:**
- Data silo risk eliminated via triggers  
- Loan-crop alignment reduces default rates  
- Transparency built via notifications and audit logs  

---

## 9️⃣ Development Roadmap

### Week 1 – Agriculture Layer MVP
- Go server setup, DB connection  
- Groups, members, seasons  
- Agricultural loans & risk flagging

### Week 2 – Core VSLA Layer
- Record savings & repayments  
- Balances & audit logs  
- SMS/USSD notifications

### Week 3 – Relief Layer
- Disasters & member disaster status  
- Bulk loan risk updates  
- NGO read-only dashboard  

### Week 4 – Frontend & Polish
- Improved UI/UX  
- Form validation  
- Optional React migration for advanced interactivity  

---

## 🔟 Funding & Pitch Notes

**Pitch for Funders:**

> “Jumuia digitizes VSLA operations while integrating climate-smart lending and rapid disaster response. Leaders can manage loans according to crop cycles, members gain trust via transparent notifications, and NGOs get real-time dashboards to support the most vulnerable. This is scalable, fundable, and directly impacts community resilience.”

**Impact Points:**
- Agriculture = climate risk + loan exposure  
- Relief = disaster preparedness + targeted assistance  
- Core VSLA = transparency + trust  