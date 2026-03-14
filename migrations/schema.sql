-- GROUPS
CREATE TABLE IF NOT EXISTS groups (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL,
    village TEXT,
    district TEXT,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP
);

-- MEMBERS
CREATE TABLE IF NOT EXISTS members (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    group_id INTEGER NOT NULL,
    name TEXT NOT NULL,
    phone TEXT,
    role TEXT,
    joined_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (group_id) REFERENCES groups(id)
);

-- SAVINGS
CREATE TABLE IF NOT EXISTS savings (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    member_id INTEGER NOT NULL,
    amount REAL NOT NULL,
    meeting_date DATE,
    FOREIGN KEY (member_id) REFERENCES members(id)
);

-- SEASONS (Agriculture)
CREATE TABLE IF NOT EXISTS seasons (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL,
    year INTEGER
);

-- CROPS
CREATE TABLE IF NOT EXISTS crops (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL
);

-- LOANS
CREATE TABLE IF NOT EXISTS loans (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    member_id INTEGER NOT NULL,
    amount REAL NOT NULL,
    purpose TEXT,
    status TEXT,
    issued_date DATE,
    due_date DATE,
    season_id INTEGER,
    FOREIGN KEY (member_id) REFERENCES members(id),
    FOREIGN KEY (season_id) REFERENCES seasons(id)
);

-- REPAYMENTS
CREATE TABLE IF NOT EXISTS repayments (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    loan_id INTEGER NOT NULL,
    amount REAL NOT NULL,
    payment_date DATE,
    FOREIGN KEY (loan_id) REFERENCES loans(id)
);

-- DISASTERS
CREATE TABLE IF NOT EXISTS disasters (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    type TEXT,
    description TEXT,
    location TEXT,
    date DATE
);

-- RELIEF
CREATE TABLE IF NOT EXISTS relief (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    member_id INTEGER,
    disaster_id INTEGER,
    type TEXT,
    amount REAL,
    date_given DATE,
    FOREIGN KEY (member_id) REFERENCES members(id),
    FOREIGN KEY (disaster_id) REFERENCES disasters(id)
);