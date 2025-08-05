# Redis Scheduler

<div align="center">
  <a href="https://github.com/gyaan/redis-scheduler">
    <img src="images/logo.png" alt="Logo" width="80" height="80">
  </a>

  <h3 align="center">Redis Scheduler</h3>

  <p align="center">
    A Redis-based scheduler in Go.
    <br />
    <a href="https://github.com/gyaan/redis-scheduler"><strong>Explore the docs »</strong></a>
    <br />
    <br />
    <a href="https://github.com/gyaan/redis-scheduler">View Demo</a>
    ·
    <a href="https://github.com/gyaan/redis-scheduler/issues">Report Bug</a>
    ·
    <a href="https://github.com/gyaan/redis-scheduler/issues">Request Feature</a>
  </p>
</div>

<details>
  <summary>Table of Contents</summary>
  <ol>
    <li>
      <a href="#about-the-project">About The Project</a>
      <ul>
        <li><a href="#built-with">Built With</a></li>
      </ul>
    </li>
    <li>
      <a href="#getting-started">Getting Started</a>
      <ul>
        <li><a href="#prerequisites">Prerequisites</a></li>
        <li><a href="#installation">Installation</a></li>
      </ul>
    </li>
    <li><a href="#usage">Usage</a></li>
    <li><a href="#roadmap">Roadmap</a></li>
    <li><a href="#contributing">Contributing</a></li>
    <li><a href="#license">License</a></li>
    <li><a href="#contact">Contact</a></li>
    <li><a href="#acknowledgments">Acknowledgments</a></li>
  </ol>
</details>

## About The Project

This project is a demonstration of a Redis-based scheduler in Go. It's designed to handle a large number of records that need to be updated regularly in a circular fashion.

### Built With

* [Go](https://golang.org/)
* [Redis](https://redis.io/)
* [gocron](https://github.com/jasonlvhit/gocron)
* [go-redis](https://github.com/go-redis/redis)
* [testify](https://github.com/stretchr/testify)

## Getting Started

To get started with this project, you'll need to have Go and Redis installed on your machine.

### Prerequisites

* Go
* Redis

### Installation

1. Clone the repo
   ```sh
   git clone https://github.com/gyaan/redis-scheduler.git
   ```
2. Install Go packages
   ```sh
   go mod tidy
   ```

## Usage

Use this project to understand how to build a Redis-based scheduler in Go.

_For more examples, please refer to the [Documentation](https://github.com/gyaan/redis-scheduler)_

## Roadmap

- [ ] Add more tests
- [ ] Add more features

See the [open issues](https://github.com/gyaan/redis-scheduler/issues) for a full list of proposed features (and known issues).

## Contributing

Contributions are what make the open source community such an amazing place to learn, inspire, and create. Any contributions you make are **greatly appreciated**.

If you have a suggestion that would make this better, please fork the repo and create a pull request. You can also simply open an issue with the tag "enhancement".

Don't forget to give the project a star! Thanks again!

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

## License

Distributed under the MIT License. See `LICENSE` for more information.

## Contact

Your Name - [@your_twitter](https://twitter.com/your_twitter) - email@example.com

Project Link: [https://github.com/gyaan/redis-scheduler](https://github.com/gyaan/redis-scheduler)

## Acknowledgments

* [othneildrew/Best-README-Template](https://github.com/othneildrew/Best-README-Template)
