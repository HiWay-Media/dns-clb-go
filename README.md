# dns-clb-go
[![Go build](https://github.com/HiWay-Media/dns-clb-go/actions/workflows/go-build.yml/badge.svg)](https://github.com/HiWay-Media/dns-clb-go/actions/workflows/go-build.yml)
[![Go test workflow](https://github.com/HiWay-Media/dns-clb-go/actions/workflows/go-test.yml/badge.svg)](https://github.com/HiWay-Media/dns-clb-go/actions/workflows/go-test.yml)


Client-Side Load-Balancer for DNS based Service Discovery written in Golang. It allows you to distribute incoming traffic across multiple servers based on DNS resolution. This project aims to provide a simple and efficient solution for load balancing without the need for additional infrastructure components.

### Features 

- DNS-Based Load Balancing: Distribute traffic based on DNS resolution.
- Round Robin Algorithm: Utilizes a simple round-robin algorithm for load distribution.
- Configurable TTL: Set a custom Time-To-Live (TTL) for DNS responses.
- Easy Configuration: Simple configuration using a YAML file.
- Logging: Detailed logs for monitoring and troubleshooting.

### Getting Started

#### Prerequisites
- Go 1.16 or later

#### Installation

```bash
go get -u github.com/HiWay-Media/dns-clb-go
```

### Contributing
If you'd like to contribute to the project, please follow the standard GitHub Fork and Pull Request workflow. Make sure to adhere to the code style guidelines and include unit tests for any new features.

### License
This project is licensed under the MIT License - see the LICENSE file for details.

### Acknowledgments
Inspired by the need for a lightweight and easy-to-use DNS load balancer.

### Contact
For any questions or issues, please open an issue on the GitHub repository.

