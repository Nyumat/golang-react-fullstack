# ∞ Go React Full Stack ∞ 

I haven't touched Vite since [NyumatFlix](https://github.com/Nyumat/NyumatFlix) so I thought I'd hop back on the wave. Although this project doesn't really have a scope yet. 

> Instead of using Node.js like I usually do for the server-side, we'll be using Go and specifically FiberV2. I really like how verbose Fiber is and it makes developing a web server very similar to the way in which it's done in Express.js and Python's Flask.

## Specs
- Go
- Fiber
- Vite
- TypeScript
- React
<!-- - Mantine -->

## Setup
In Development mode, use `make -j2` to run the frontend and backend servers concurrently.

## API Reference
```
// Get a user by their ObjectID
app.Get("/api/users/:id", controllers.GetUser) 

// Get a user by their name
app.Get("/api/user/:name", controllers.GetUserByName)

// Get all documents in the user collection
app.Get("/api/users", controllers.GetUsers)

// Create a new user document
app.Post("/api/users", controllers.CreateUser)
```
## Todo List:
- Graceful error handling on both ends
- Decide on a project scope? (Todos/School System/IDK Yet)
- Complete all routes in BE router
- Add npm script command to concurrently run both FE and BE

