# Appointments System

An appointments system with a frontend, a backend and a database.

- The frontend uses HTML 5, Bootstap 5 and Javascript.
- The backend uses Fiber (Go framework) and JWT for authentication.
- The DB is Postgres.

## Run it with Docker Compose

A `.env` file with user and password credentials (`SENDER_EMAIL` and `SENDER_PASSWORD` in docker-compose.yaml) is required to send notification emails.

```bash
docker compose up -d
```

## ERD Diagram

![arch_basic](https://drive.google.com/uc?export=view&id=1E6kkXp-uOc_FyRFliFzuO0bEYyPjeKMj)

## Endpoints

Client Endpoints
- `GET /bookings` (main page)
- `GET /booked` (appointment booked page)
- `POST /bookings` (book appointment)

API Endpoints
- `GET /api/calendar` (get bootstrap-calendar data)
- `POST /api/attentionHours` (update attention hours)
- `POST /api/appointments` (update appointment)

Authentication Endpoints
- `GET /auth/login` (login page)
- `POST /auth/login` (login)
- `POST /auth/logout` (logout)

Admin Endpoints
- `GET /admin/profile` (profile page)
- `POST /admin/profile` (update profile)
- `GET /admin/appointments` (appointments page)
- `GET /admin/bookings` (bookings page)
- `POST /admin/bookings` (update booking)
- `GET /admin/appointmentTypes` (appointment types page)
- `POST /admin/appointmentTypes` (create appointment type)
- `POST /admin/appointmentTypes/:id` (update appointment type)
- `GET /admin/clients` (clients page)
- `POST /admin/clients` (update client)
- `GET /admin/metrics` (app metrics)