# Go Todo App 📝

A beautiful, modern todo application built with Go, Gin framework, and Tailwind CSS.

![Go](https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white)
![Tailwind CSS](https://img.shields.io/badge/Tailwind_CSS-38B2AC?style=for-the-badge&logo=tailwind-css&logoColor=white)
![SQLite](https://img.shields.io/badge/SQLite-07405E?style=for-the-badge&logo=sqlite&logoColor=white)

## ✨ Features

- **Beautiful UI** - Modern, responsive design with Tailwind CSS
- **Real-time Updates** - Dynamic todo management without page refresh
- **Filter Todos** - View all, active, or completed tasks
- **Statistics Dashboard** - Track your progress with visual stats
- **Smooth Animations** - Delightful user experience with CSS animations
- **RESTful API** - Full CRUD operations available via API endpoints

## 🚀 Getting Started

### Prerequisites

- Go 1.16 or higher
- SQLite (handled automatically by GORM)

### Installation

1. Clone the repository or navigate to the project directory:
```bash
cd /home/syedtayyabsagheer/Desktop/Office/go-todo
```

2. Install dependencies:
```bash
go mod download
```

3. Run the application:
```bash
go run main.go
```

4. Open your browser and visit:
```
http://localhost:8080
```

## 🎨 Screenshots

The application features:
- **Hero Section** - Beautiful gradient header with app branding
- **Add Todo Form** - Easy-to-use input with modern styling
- **Statistics Cards** - Visual overview of total, completed, and pending tasks
- **Filter Buttons** - Quick filtering between all, active, and completed todos
- **Todo Items** - Clean card design with hover effects and animations
- **Responsive Design** - Works perfectly on mobile, tablet, and desktop

## 📚 API Endpoints

### Web Routes
- `GET /` - Main web interface

### API Routes
- `GET /todos` - Get all todos
- `POST /todos` - Create a new todo
- `PUT /todos/:id` - Update a todo
- `DELETE /todos/:id` - Delete a todo

### Example API Requests

**Create a todo:**
```bash
curl -X POST http://localhost:8080/todos \
  -H "Content-Type: application/json" \
  -d '{"title":"Learn Go","completed":false}'
```

**Get all todos:**
```bash
curl http://localhost:8080/todos
```

**Update a todo:**
```bash
curl -X PUT http://localhost:8080/todos/1 \
  -H "Content-Type: application/json" \
  -d '{"completed":true}'
```

**Delete a todo:**
```bash
curl -X DELETE http://localhost:8080/todos/1
```

## 🏗️ Project Structure

```
go-todo/
├── assets/              # Static assets
├── controllers/         # Request handlers
│   └── todo_controller.go
├── database/           # Database configuration
│   └── database.go
├── models/             # Data models
│   └── todo.go
├── routes/             # Route definitions
│   └── routes.go
├── templates/          # HTML templates
│   ├── base.html      # Base layout with Tailwind CSS
│   └── index.html     # Main todo page
├── favicon.ico        # App favicon
├── go.mod             # Go module file
├── go.sum             # Go dependencies
└── main.go            # Application entry point
```

## 🛠️ Technologies Used

- **Backend**: Go with Gin framework
- **Database**: SQLite with GORM ORM
- **Frontend**: HTML5, Tailwind CSS (via CDN)
- **JavaScript**: Vanilla JS for dynamic interactions

## 🎯 Features Breakdown

### UI Components
- **Gradient backgrounds** - Beautiful blue to purple gradients
- **Card design** - Modern card-based layout
- **Smooth transitions** - All interactions have smooth animations
- **Responsive grid** - Adapts to any screen size
- **Icons** - SVG icons for visual clarity

### Functionality
- **Add todos** - Quick todo creation with Enter key support
- **Mark complete** - Toggle completion with animated checkboxes
- **Delete todos** - Remove tasks with confirmation
- **Filter views** - Switch between all/active/completed
- **Live stats** - Real-time count of total, completed, and pending tasks

## 🔧 Development

To run in development mode with live reload, you can use:
```bash
# Install air for live reload
go install github.com/cosmtrek/air@latest

# Run with air
air
```

## 📝 License

This project is open source and available for personal and commercial use.

## 🤝 Contributing

Contributions, issues, and feature requests are welcome!

## 👨‍💻 Author

Built with ❤️ using Go and Tailwind CSS

---

**Happy Task Managing! 🎉**
