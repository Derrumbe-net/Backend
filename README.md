file structure plan
lo puse por feature en vez de layer pq go lo sugiere
might as well try it out

├── cmd/
│   └── api/
│       └── main.go       # Main file
├── internal/
│   ├── auth/             # JWT, Admin Login
│   ├── landslide/        # Reports, Validation Logic
│   ├── station/          # Station metadata and Sensor readings
│   ├── content/          # Projects, Publications, etc. (students)
│   └── email/            # Email service logic
└── go.mod                # Project dependencies

for docker

docker build -t backend:latest -f Dockerfile .
docker run -p 8080:8080 backend
