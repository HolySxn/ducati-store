#!/bin/bash

# Navigate to the backend directory and start the backend server
echo "Starting backend server..."
cd backend
go run main.go &
BACKEND_PID=$!

# Navigate to the frontend directory and start the frontend server
echo "Starting frontend server..."
cd ../frontend
npx vite &
FRONTEND_PID=$!

# Function to stop both servers
function stop_servers {
  echo "Stopping servers..."
  kill $BACKEND_PID
  kill $FRONTEND_PID
  exit
}

# Trap SIGINT and SIGTERM to stop servers when the script is terminated
trap stop_servers SIGINT SIGTERM

# Wait for both servers to exit
wait $BACKEND_PID
wait $FRONTEND_PID