# 🚀 Task Manager App (Full Stack)

A full-stack Task Manager application built using **Go (Gin), Angular 21, and PostgreSQL**, with complete CRUD functionality and cloud deployment.

---

## 🌐 Live Demo

* 🔗 Frontend: https://sunny-twilight-6f1d66.netlify.app
* 🔗 Backend API: https://task-manager-onr0.onrender.com

---

## 🧰 Tech Stack

### Frontend

* Angular 21
* TypeScript
* HTML & CSS

### Backend

* Go (Golang)
* Gin Framework

### Database

* PostgreSQL (Render Cloud)

### Deployment

* Frontend: Netlify
* Backend & DB: Render

---

## ✨ Features

* ✅ Create new tasks
* ✅ View all tasks
* ✅ Update task details
* ✅ Delete tasks
* ✅ Real-time UI updates
* ✅ REST API integration
* ✅ Cloud deployment

---

## 📁 Project Structure

```
task-manager/
│
├── backend/        # Go API (Gin)
│   ├── cmd/
│   ├── controllers/
│   ├── models/
│   ├── routes/
│   ├── config/
│   └── main.go
│
├── frontend/       # Angular App
│   ├── src/
│   └── angular.json
│
└── README.md
```

---

## ⚙️ Backend Setup (Go)

```bash
cd backend
go mod tidy
go run cmd/main.go
```

### Environment Variables (.env)

```
DB_HOST=your_host
DB_PORT=5432
DB_USER=your_user
DB_PASSWORD=your_password
DB_NAME=your_db
```

---

## 💻 Frontend Setup (Angular)

```bash
cd frontend
npm install
ng serve
```

---

## 🔗 API Endpoints

| Method | Endpoint   | Description   |
| ------ | ---------- | ------------- |
| GET    | /tasks     | Get all tasks |
| POST   | /tasks     | Create task   |
| PUT    | /tasks/:id | Update task   |
| DELETE | /tasks/:id | Delete task   |

---

## 🚀 Deployment

### Backend (Render)

* Configured Go service
* Connected PostgreSQL database
* Set environment variables
* Enabled CORS for frontend

### Frontend (Netlify)

* Built Angular app for production
* Deployed via drag-and-drop / GitHub
* Connected to live backend API

---

## 🧪 Challenges & Fixes

* ❌ CORS errors → ✅ Fixed using Gin CORS middleware
* ❌ Git submodule issue → ✅ Removed nested `.git`
* ❌ Deployment errors → ✅ Corrected build paths
* ❌ DB connection → ✅ Used cloud credentials

---

## 🔮 Future Improvements

* 🔐 User Authentication (JWT)
* 🔍 Search & Filter Tasks
* 📊 Dashboard Analytics
* 🎨 Improved UI/UX

---

## 👨‍💻 Author

**Manikandan**
Full Stack Developer (Java → Go + Angular)

---

## ⭐ If you like this project

Give it a ⭐ on GitHub and feel free to contribute!

---

