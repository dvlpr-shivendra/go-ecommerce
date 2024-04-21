The backend is developed using Go programming language. Payments are handled by Razorpay. Please note that this project is developed for exploration and learning purposes and is not meant to be used in production in its current state.

### Setup
- Make sure you have docker installed.
- ```cp .env.example .env```
- For payments create an account on Razorpay and get the key and secret and put them in `.env`
- ```docker compose up -d```
