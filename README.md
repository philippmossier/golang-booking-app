# Bookings and Reservations APP

**This app focuses on a robust Go backend which is lightweight, fast and secure.**

As this project does not focus on frontend the client code is kinda framework diagnostic except of some bootstrap for the html.
The datepicker and alerts are achieved with vanilla-js.

## How Backend is build

- Built in Go version 1.16
- Uses the [chi-router](https://github.com/go-chi/chi)
- Uses [alex edwards SCS](https://github.com/alexedwards/scs) session management
- Uses CSRF protection middleware [nosurf](https://github.com/justinas/nosurf)
