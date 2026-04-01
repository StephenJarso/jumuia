# Jumuia VSLA Platform - Setup Guide

## Overview
Jumuia is a Village Savings & Loan Association (VSLA) platform that enables communities to manage group-based savings, loans, and financial activities.

## Features
- **Group Management**: Create and manage VSLA groups with leaders
- **Member Management**: Add members with roles (Leader, Treasurer, Member)
- **Savings Tracking**: Record and track member savings
- **Loan Management**: Issue loans and track repayments
- **Member Accounts**: Individual member financial summaries
- **Disaster Relief**: Track disasters and distribute relief
- **Dashboard**: Overview of all financial activities

## Quick Start

### 1. Initialize Database
```bash
# Create database and schema
sqlite3 jumuia.db < migrations/schema.sql

# Load mock data (optional but recommended for testing)
sqlite3 jumuia.db < migrations/mock_data.sql
```

### 2. Run the Application
```bash
go run cmd/server/main.go
```

### 3. Access the Application
Open your browser and navigate to: `http://localhost:8080`

## Mock Data Included

The mock data includes realistic Kenyan VSLA data:

### Groups (5 groups)
1. **Tumaini VSLA** - Kisumu Village, Kisumu Central
2. **Upendo Group** - Busia Town, Busia County
3. **Amani Savings** - Siaya Center, Siaya County
4. **Umoja Women** - Mumias Village, Kakamega County
5. **Vijana Youth** - Bungoma Town, Bungoma County

### Members (32 members)
- Each group has 5-8 members
- Roles: Leader, Treasurer, Member
- Realistic Kenyan names and phone numbers

### Financial Data
- **Savings**: Monthly contributions (KES 300-1000)
- **Loans**: Various purposes (school fees, business, medical, farming)
- **Repayments**: Regular loan repayments
- **Relief**: Disaster relief distributions

## Key Workflows

### 1. View Groups
- Navigate to `/groups`
- See all groups with member counts and leaders
- Click "View Details" to see group members

### 2. View Member Account
- From group detail page, click "View Account" on any member
- See savings history, active loans, and financial summary

### 3. Add New Member
- From group detail page, use the "Add New Member" form
- Or navigate to `/members/new`

### 4. Record Savings
- Navigate to `/savings/new`
- Select member and enter amount

### 5. Issue Loan
- Navigate to `/loans/new`
- Select member, enter amount and purpose

## Database Schema

### Tables
- **groups**: VSLA groups with leader reference
- **members**: Group members with roles
- **savings**: Member savings records
- **loans**: Loan applications and status
- **repayments**: Loan repayment records
- **seasons**: Agricultural seasons
- **crops**: Crop types
- **disasters**: Disaster events
- **relief**: Relief distributions

## API Endpoints

### Groups
- `GET /groups` - List all groups
- `GET /groups/new` - Create group form
- `POST /groups/create` - Create new group
- `GET /groups/detail?id={id}` - Group detail page
- `POST /groups/update-leader` - Update group leader

### Members
- `GET /members` - List members
- `GET /members/new` - Add member form
- `POST /members/create` - Create new member
- `GET /members/account?id={id}` - Member account page

### Financial
- `GET /savings` - List savings
- `GET /savings/new` - Record savings form
- `POST /savings/create` - Record new savings
- `GET /loans` - List loans
- `GET /loans/new` - Issue loan form
- `POST /loans/create` - Issue new loan
- `GET /repayments` - List repayments
- `GET /repayments/new` - Record repayment form
- `POST /repayments/create` - Record new repayment

## Technology Stack
- **Backend**: Go (Golang)
- **Database**: SQLite
- **Frontend**: HTML, CSS, JavaScript
- **Templates**: Go html/template

## Project Structure
```
jumuia/
├── cmd/server/          # Application entry point
├── internal/
│   ├── db/             # Database initialization
│   ├── handlers/       # HTTP handlers
│   ├── models/         # Data models
│   └── repository/     # Database operations
├── migrations/         # SQL schema and mock data
└── web/templates/      # HTML templates
```

## Currency
All financial amounts are in **Kenyan Shillings (KES)**

## Support
For issues or questions, please refer to the project documentation.
