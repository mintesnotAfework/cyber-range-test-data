# 🛡️ SQL Injection (SQLi) – CTF Exploitation Guide

## 📘 Overview

**SQL Injection (SQLi)** is one of the most prevalent and dangerous vulnerabilities in web applications. It allows attackers to execute arbitrary SQL code on a database by injecting malicious SQL statements through unsanitized user input. If exploited successfully, SQLi can lead to:

- Authentication bypass
- Data exfiltration
- Data modification or deletion
- Complete database takeover

SQLi vulnerabilities occur when input from a client (user, API consumer, etc.) is concatenated directly into a SQL query without adequate escaping or use of parameterized statements.

---

## 💡 How It Works

Consider this vulnerable login form:

```sql
SELECT * FROM users WHERE username = 'admin' AND password = '1234';
````

If an attacker inputs:

* `Username`: `admin' --`
* `Password`: `anything`

The resulting query becomes:

```sql
SELECT * FROM users WHERE username = 'admin' --' AND password = 'anything';
```

The `--` comments out the rest of the query, bypassing the password check.

---

## ⚙️ Vulnerable Example in Go (GORM)

### 🚨 Insecure Code

```go
// WARNING: vulnerable to SQL injection!
db.Raw("SELECT * FROM users WHERE username = '" + username + "' AND password = '" + password + "'").Scan(&user)
```

### ✅ Secure Code

```go
// SAFE: uses parameterized query
db.Raw("SELECT * FROM users WHERE username = ? AND password = ?", username, password).Scan(&user)
```

Even better, use GORM's built-in query methods:

```go
db.Where("username = ? AND password = ?", username, password).First(&user)
```

---

## 🧪 CTF Challenge Hints

**Target Areas**:

* Login forms
* Search fields
* URL query parameters (e.g., `?id=`)
* Admin panels or dashboards
* Any feature that filters or returns database results

**Payload Examples**:

| Goal                       | Payload                                                            |
| -------------------------- | ------------------------------------------------------------------ |
| Auth Bypass                | `' OR '1'='1`                                                      |
| Extract Table Info         | `' UNION SELECT null, table_name FROM information_schema.tables--` |
| Time-based Blind Injection | `' OR pg_sleep(5)--`                                               |
| Boolean-based Blind SQLi   | `' AND 1=1--` or `' AND 1=2--`                                     |
| Data Extraction via UNION  | `' UNION SELECT null, version()--`                                 |
| Comment Injection          | `';--` or `--`                                                     |

---

## 🔍 Manual Testing Strategy

1. **Map Attack Surface**: Locate all endpoints accepting input.

   * URLs with `?id=123`
   * Forms with username/password, search, filters

2. **Inject Basic Payloads**:

   * Start with `'`, `"`, or `--` to cause errors or logic changes
   * Observe for changes in application behavior or server errors

3. **Try Logical Operators**:

   * `' OR 1=1--`
   * `' AND 1=2--`

4. **Try UNION SELECT Payloads**:

   * Determine column count (e.g., `' UNION SELECT null, null--`)
   * Use version extraction: `' UNION SELECT null, version()--`

5. **Blind Injection (if no error feedback)**:

   * `' AND (SELECT 1 WHERE 1=1)--`
   * `' OR CASE WHEN (1=1) THEN pg_sleep(5) ELSE pg_sleep(0) END--`

---

## 🛡️ Prevention Techniques

### ✅ Use Parameterized Queries

Always use prepared statements or parameterized queries to separate logic from data.

**In Go (GORM)**:

```go
db.Where("email = ?", inputEmail).Find(&user)
```

### ✅ Input Validation

* Accept only expected formats (e.g., email regex for login)
* Whitelist acceptable input values when possible

### ✅ Limit DB Permissions

* Use a limited-privilege DB user for the application
* Avoid granting `DROP`, `ALTER`, or `SUPERUSER` access to app-level accounts

### ✅ Logging & Monitoring

* Log unexpected input or failed queries
* Monitor for patterns of enumeration or timed responses

### ✅ Web Application Firewall (WAF)

Use WAFs (e.g., ModSecurity) to detect common SQLi payloads.

---

## 🧰 Tools for SQLi Testing

* [sqlmap](https://sqlmap.org/) – Automated SQLi scanner and exploit tool
* Burp Suite – Manual interception and fuzzing
* OWASP ZAP – Open source proxy for scanning
* Postman or curl – For testing APIs manually

---

## 🎯 CTF Challenge Ideas (You Can Include)

| Challenge Name   | Description                                    | Goal                          |
| ---------------- | ---------------------------------------------- | ----------------------------- |
| **Blind Gate**   | Login page with blind SQLi vulnerability       | Extract admin password        |
| **Query Sniper** | Profile page vulnerable to union-based SQLi    | Dump user table               |
| **Sleep Agent**  | Endpoint reveals data via time-based injection | Detect and exploit time delay |
| **Broken Auth**  | Bypass login using `' OR '1'='1`               | Log in without password       |

---

## 📚 Real-World Examples

* **2017**: Equifax breach – SQLi led to theft of data from 147M users.
* **2020**: Microsoft subdomain bug – SQLi found in legacy endpoint.
* **Bug Bounty Programs**: SQLi is still one of the most reported and rewarded bugs.

---

## 🧠 Final Tip

In a CTF, always look for places where input is directly reflected in results or errors. SQLi is often the key to unlocking hidden admin functions, dumping tables, or pivoting to deeper vulnerabilities.

---

```

Would you like me to do a similar **deep-dive write-up** for **XSS**, **LFI**, or **Command Injection** as well?
```
