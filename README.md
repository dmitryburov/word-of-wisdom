# "Word of Wisdom" tcp server (protected from DDOS attacks with the Proof of Work).

## Task

Design and implement "Word of Wisdom" tcp server:

- TCP server should be protected from DDOS attacks with the Prof of Work (https://en.wikipedia.org/wiki/Proof_of_work), the challenge-response protocol should be used.
- The choice of the POW algorithm should be explained.
- After Prof Of Work verification, server should send one of the quotes from “word of wisdom” book or any other collection of the quotes.
- Docker file should be provided both for the server and for the client that solves the POW challenge.

## Getting started

Project based on [Clean architecture](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html) principles.

Requirements:

- Go 1.19+ installed (to run tests, start server or client without Docker)
- Docker installed (to run docker-compose)
- Environment file `.env` (see example in [env.example](env.example))

```
# Run server and client by docker-compose
make run-compose

# Run only server
make run-server

# Run only client
make run-client

# other command - call help
make help
```

## Resources

- [Word of Wisdom](https://en.wikipedia.org/wiki/Word_of_Wisdom)
- [Proof of work](https://en.wikipedia.org/wiki/Proof_of_work)
- [Hashcash](https://en.wikipedia.org/wiki/Hashcash)
- [Distributed Consensus – Proof-of-Work](https://oliverjumpertz.com/distributed-consensus-proof-of-work/)
- [The Clean Architecture](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html)
