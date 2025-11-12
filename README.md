````markdown
# Do Not Fall For It: AI-Powered Honeypot System

**Do not fall for it** is a modular honeypot framework written in Go.  
It simulates multiple network services to study attacker behavior in a controlled environment.  
Supported protocols include SSH, HTTP, MySQL, and TCP.  
It integrates with RabbitMQ for message handling and Prometheus for monitoring, and runs entirely through Docker.

---

## Directory Layout

honeypot_clean/  
│  
├── configurations/  
│   ├── honeypot.yaml  
│   ├── honey.yaml  
│   └── services/  
│       ├── http-80.yaml  
│       ├── http-8080.yaml  
│       ├── http-8081.yaml  
│       ├── mysql-3306.yaml  
│       ├── ssh-22.yaml  
│       ├── ssh-2222.yaml  
│       └── tcp-3306.yaml  
│  
├── protocols/  
│   ├── protocol_manager.go  
│   └── strategies/  
│       ├── HTTP/  
│       ├── MySQL/  
│       ├── SSH/  
│       └── TCP/  
│  
├── plugins/  
│   ├── honeypot-cloud.go  
│   └── llm-integration.go  
│  
├── builder/  
│   ├── builder.go  
│   └── director.go  
│  
├── parser/  
│   └── configurations_parser.go  
│  
├── tracer/  
│   └── tracer.go  
│  
├── historystore/  
│   └── history_store.go  
│  
├── integration_test/  
│  
├── honeypot-chart/  
│  
├── docker-compose.yml  
├── Dockerfile  
├── main.go  
├── go.mod  
├── go.sum  
└── Makefile  

---

## Prerequisites

- Docker Desktop  
- Git  
- Minimum 2 GB RAM  

---

## Setup

### 1. Clone Repository

```bash
git clone <repository-url>
cd honeypot_clean
````

### 2. Configure Environment

Create a `.env` file in the root directory:

```
RABBITMQ_URI=amqp://guest:guest@rabbitmq:5672/
OPEN_AI_SECRET_KEY=your_openai_key_here
```

If no OpenAI key is provided, AI features will remain disabled.

### 3. Launch Containers

```bash
docker-compose up -d
docker ps
```

Expected containers:

* `honeypot` – main application
* `rabbitmq` – message broker

---

## Configuration

**File:** `configurations/honeypot.yaml`

```yaml
core:
  logging:
    debug: false
    debugReportCaller: false
    logDisableTimestamp: true
    logsPath: ./logs
  tracings:
    rabbit-mq:
      enabled: true
      uri: "amqp://guest:guest@rabbitmq:5672/"
    prometheus:
      path: /metrics
      port: ":2112"
    honeypot-cloud:
      enabled: false
      uri: ""
      auth-token: ""
```

---

## Services and Ports

| Service             | Protocol | Port  | Description                  |
| ------------------- | -------- | ----- | ---------------------------- |
| SSH Interactive     | SSH      | 2222  | Command-line SSH honeypot    |
| SSH AI Session      | SSH      | 2223  | SSH with AI-driven responses |
| HTTP Apache 401     | HTTP     | 8080  | Simulates 401 Unauthorized   |
| HTTP Test           | HTTP     | 8081  | Simple HTTP test endpoint    |
| HTTP WordPress      | HTTP     | 8082  | Fake WordPress site          |
| MySQL               | MySQL    | 3306  | Simulated MySQL server       |
| TCP Generic         | TCP      | 3307  | Generic TCP listener         |
| Prometheus          | HTTP     | 2112  | Metrics endpoint             |
| RabbitMQ            | AMQP     | 5672  | Message broker               |
| RabbitMQ Management | HTTP     | 15672 | Web admin console            |

---

## Testing Services

### SSH

```bash
ssh -p 2222 root@localhost
```

Common passwords:

```
root
toor
qwerty
123456
jenkins
postgres
minecraft
```

AI-powered SSH:

```bash
ssh -p 2223 root@localhost
```

### HTTP

```bash
curl -v http://localhost:8080/
curl http://localhost:8081/
curl http://localhost:8082/
```

### MySQL

```bash
mysql -h localhost -P 3306 -u root -p
```

Any password works.

### TCP

```bash
telnet localhost 3307
# or
nc localhost 3307
```

---

## Monitoring

### RabbitMQ Dashboard

* URL: [http://localhost:15672](http://localhost:15672)
* Username: `guest`
* Password: `guest`

### Prometheus Metrics

```bash
curl http://localhost:2112/metrics
```

### Logs

```bash
docker logs honeypot
docker logs rabbitmq
```

Follow live:

```bash
docker logs -f honeypot
```

---

## Troubleshooting

**Port Conflict**

```bash
netstat -ano | findstr :5672
taskkill /PID <PID> /F
```

**Docker Not Running**

```bash
docker version
```

**Missing Configuration**

```bash
ls configurations/honeypot.yaml
```

**RabbitMQ Check**

```bash
docker exec rabbitmq rabbitmq-diagnostics ping
```

**SSH Debug**

```bash
ssh -v -p 2222 root@localhost
```

---

## Development

### Build Manually

```bash
go build -o honeypot .
go test ./... -v
```

### Makefile Commands

```bash
make honeypot.start
make honeypot.stop
make test.unit
make test.integration
```

---

## Add a Custom Service

Example file under `configurations/services/`:

```yaml
apiVersion: "v1"
protocol: "http"
address: ":8083"
description: "Custom HTTP honeypot"
commands:
  - regex: ".*"
    handler: "Custom Response"
    headers:
      - "Content-Type: text/plain"
    statusCode: 200
```

Restart containers after adding.

---

## Security Practices

* Run only in isolated lab environments.
* Never expose honeypot to production or public networks.
* Inspect logs and metrics regularly.
* Use dedicated VMs or sandbox systems.
* Keep honeypot traffic separated from internal assets.

---

## System Requirements

| Resource | Minimum | Recommended |
| -------- | ------- | ----------- |
| CPU      | 2 cores | 4 cores     |
| RAM      | 2 GB    | 4 GB        |
| Storage  | 10 GB   | 10+ GB      |

---

## Contributing

1. Fork the repository
2. Create a feature branch
3. Commit clearly
4. Include tests
5. Open a pull request

---

## License

Released under the MIT License.
See the `LICENSE` file for terms.

---

## Disclaimer

This project is intended **only for cybersecurity research and education**.
Unauthorized deployment or malicious use is strictly prohibited.
It will render perfectly — vertical directory structure preserved, markdown clean, and fully compliant with GitHub formatting.
```
