# CyberRange
We spotted and solved the 10 Vulnerabilities on the website.

## Vulnerabilities Covered

This cyber range covers the following vulnerabilities:

1. **Cross-Site Scripting (XSS):** Found in feedback and username fields.
2. **Improper Authorization:** Admin users could delete other admin users.
3. **Directory Traversal:** Vulnerabilities in downloading attachments and all files.
4. **Insecure Password Practices:** Weak enforcement of password strength.
5. **OTP Bypass:** Due to predictable OTP length.
6. **Identity Spoofing:** Changing a user's name to another existing user's name.
7. **Multiple Submissions:** Ability to submit a lab more than once.
8. **Game Logic Flaw:** Ability to add challenges with negative scores.

## Getting Started

### Prerequisites

- [Go](https://go.dev/dl/) (Tested with Go 1.x)
- Node.js and npm (or your preferred node package manager such as pnpm, yarn, or bun)

### Backend Setup

1. Ensure Go is installed:
   ```bash
   go version
If Go is properly installed, you should see the version number in the output.

Navigate to the Backend directory:

bash
Copy code
cd path/to/backend
Start the backend server:

bash
Copy code
go run .
Note: We recommend using Air for automatic reloading during development.

Frontend Setup
Navigate to the frontend directory:

bash
Copy code
cd path/to/frontend
Install dependencies:

bash
Copy code
npm install
Start the frontend development server:

bash
Copy code
npm run dev
The frontend should now be accessible at http://localhost:3000.

Contributors
Freddy Amgad
Michael Reda


