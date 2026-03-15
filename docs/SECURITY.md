# Password Security Documentation

## Overview

This document describes the dual-layer password encryption system implemented in NectarPin, designed to provide secure password handling from client input to database storage.

## Encryption Architecture

```
┌─────────────────────────────────────────────────────────────────────────────┐
│                         PASSWORD ENCRYPTION FLOW                             │
├─────────────────────────────────────────────────────────────────────────────┤
│                                                                              │
│   [User Input]          [Frontend]           [Network]         [Backend]    │
│                                                                              │
│   "password123"  ───►  MD5 Hash    ───►  HTTPS/TLS  ───►  bcrypt Hash      │
│                         (32 chars)          Encrypted       (60 chars)      │
│                                              Transport                       │
│                                                                              │
│                                              ▼                               │
│                                                                         [Database]    │
│                                                                              │
│                                                                        bcrypt hash  │
│                                                                        stored       │
└─────────────────────────────────────────────────────────────────────────────┘
```

## Layer 1: Frontend MD5 Encryption

### Purpose
- **NOT for secure storage** - MD5 is cryptographically broken
- **Purpose**: Prevent plaintext passwords from appearing in:
  - Browser developer tools
  - Network logs
  - Server access logs
  - Proxy/cache systems

### Implementation
```typescript
// Frontend: nectarpin-web/src/utils/index.ts
import md5 from 'blueimp-md5'

export function encryptPassword(password: string): string {
  return md5(password)  // Returns 32-char lowercase hex string
}
```

### Security Considerations
| Aspect | Status | Notes |
|--------|--------|-------|
| Collision Resistance | Weak | MD5 collisions are practical to find |
| Pre-image Resistance | Moderate | Still computationally expensive to reverse |
| Salt | None | MD5 alone has no built-in salt |
| Purpose | Transport | Only for transmission protection |

## Layer 2: Backend bcrypt Encryption

### Purpose
- **Primary security layer** for password storage
- Industry-standard password hashing algorithm
- Designed specifically for password storage

### Implementation
```go
// Backend: internal/utils/password.go
const BcryptCost = 12  // ~250ms computation time

func HashPassword(md5Password string) (string, error) {
    if len(md5Password) != 32 {
        return "", ErrPasswordTooShort
    }
    hashedBytes, err := bcrypt.GenerateFromPassword([]byte(md5Password), BcryptCost)
    if err != nil {
        return "", ErrPasswordHashFail
    }
    return string(hashedBytes), nil
}

func VerifyPassword(md5Password string, hashedPassword string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(md5Password))
    return err == nil
}
```

### bcrypt Parameters

| Parameter | Value | Rationale |
|-----------|-------|-----------|
| Cost Factor | 12 | ~250ms on modern hardware, balances security and UX |
| Output Length | 60 chars | Fixed-length bcrypt hash format |
| Salt | Auto-generated | 128-bit random salt per password |
| Algorithm | Blowfish | Key derivation function |

### bcrypt Hash Format
```
$2a$12$N9qo8uLOickgx2ZMRZoMy.MrqD3G4aE7VWQj3z5K8vLtP7ZPxJx1K
 │  │  │                                                        │
 │  │  │                                                        └── Hash (31 chars)
 │  │  └── Salt (22 chars)
 │  └── Cost factor (12)
 └── Algorithm version (2a)
```

## Transmission Security

### HTTPS Requirements
- **MANDATORY** for all production deployments
- TLS 1.2+ required
- Proper certificate validation
- HSTS headers recommended

### API Request Format
```json
POST /api/public/user/v1/register
Content-Type: application/json

{
  "username": "testuser",
  "password": "e10adc3949ba59abbe56e057f20f883e",  // MD5 hash (32 chars)
  "email": "user@example.com"
}
```

## Database Storage

### Schema
```sql
CREATE TABLE users (
    id BIGSERIAL PRIMARY KEY,
    username VARCHAR(50) NOT NULL UNIQUE,
    password VARCHAR(60) NOT NULL,  -- bcrypt hash (60 chars)
    email VARCHAR(255) NOT NULL UNIQUE,
    -- ... other fields
);
```

### Password Field Changes
| Version | Field Size | Algorithm |
|---------|------------|-----------|
| Before | 32 chars | MD5 only |
| After | 60 chars | bcrypt |

## Security Best Practices

### ✅ Implemented
1. **Dual-layer encryption** - MD5 for transport, bcrypt for storage
2. **Automatic salting** - bcrypt generates unique salt per password
3. **Configurable work factor** - Cost parameter adjustable for hardware
4. **Constant-time comparison** - bcrypt prevents timing attacks
5. **Password format validation** - Backend validates 32-char MD5 input

### 🔒 Additional Recommendations
1. **Rate limiting** - Implement login attempt throttling
2. **Password policies** - Enforce minimum complexity requirements
3. **Breached password check** - Check against known breach databases
4. **Password history** - Prevent reuse of recent passwords
5. **Account lockout** - Temporary lock after failed attempts

## Migration Guide

### For Existing Users (MD5-only passwords)
If migrating from MD5-only storage:

```go
// Migration script
func migratePassword(md5Hash string) (string, error) {
    // The MD5 hash becomes the "password" for bcrypt
    return HashPassword(md5Hash)
}
```

### Database Migration
```sql
-- Extend password column for bcrypt hashes
ALTER TABLE users ALTER COLUMN password TYPE VARCHAR(60);
```

## Testing

### Unit Tests
```bash
# Run password utility tests
go test ./internal/utils/... -v -run TestHash
```

### Test Coverage
- Hash generation
- Password verification
- Invalid input handling
- Cost factor validation
- Salt uniqueness

## Error Handling

| Error | HTTP Code | Message |
|-------|-----------|---------|
| ErrPasswordTooShort | 400 | 密码格式错误 |
| ErrPasswordHashFail | 500 | 注册失败 |
| ErrInvalidPassword | 401 | 用户不存在或密码错误 |

## References

- [OWASP Password Storage Cheat Sheet](https://cheatsheetseries.owasp.org/cheatsheets/Password_Storage_Cheat_Sheet.html)
- [bcrypt Wikipedia](https://en.wikipedia.org/wiki/Bcrypt)
- [NIST SP 800-63B](https://pages.nist.gov/800-63-3/sp800-63b.html)

## Version History

| Version | Date | Changes |
|---------|------|---------|
| 1.0 | 2026-03-15 | Initial implementation with MD5 + bcrypt |
