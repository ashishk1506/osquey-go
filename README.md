# Osquey-Go Project

This project is a simple web application built with Go (Golang) that serves as a backend server. It integrates with a MySQL Docker container and provides a static HTML interface.

## Features

- **Backend**: Golang server :8081
- **Database**: MySQL running in a Docker container 3310:3306
- **Frontend**: Static HTML page for displaying data
- **Endpoint**: `/latest_data`

The `/latest_data` endpoint returns the following information:

1. OS version
2. Osquery version
3. List of installed applications on a Windows PC

## Prerequisites

Before running this project, ensure you have the following:

- Docker
- Go (Golang) installed
- MySQL
- **Osquery** installed on your Windows machine (you can download it from [here](https://osquery.io/download))

## Setup Instructions

1. **Clone the repository**:

   ```
   git clone https://github.com/ashishk1506/osquey-go.git
   cd osquey-go


2. **start sql container**:

  ```
  cd server
  docker-compose up
  ```

3. **create .env file**:

   ```env
   DB_USERNAME=
   DB_PASSWORD=
   DB_HOSTNAME=
   DB_NAME=
   DB_PORT=
  ```

4. **start server**:

  ```
  go run main.go
  ```


