````markdown
# 🏓 goping — Lightweight ICMP Ping Utility in Go

**goping** is a simple command-line utility written in Go that replicates the core functionality of the classic Unix `ping` tool.  
It sends ICMP Echo Requests, measures round-trip times (RTT), and reports packet loss and latency statistics — all with clean Go code and proper signal handling.

---

## 🚀 Features

- ✅ Send ICMP Echo Requests (`ping`)
- ✅ Measure RTT (Round-Trip Time)
- ✅ Display min / avg / max RTT statistics
- ✅ Track packet loss percentage
- ✅ Graceful shutdown on **Ctrl+C** (prints stats before exit)
- ✅ Configurable packet count and interval via flags
- ✅ Lightweight, cross-platform, and dependency-minimal

---

## ⚙️ Installation

Clone and build the project:

```bash
git clone https://github.com/<your-username>/goping.git
cd goping
go build -o goping ./cmd/app
````

You can now run it directly:

```bash
sudo ./goping [options] <address>
```

---

## 🧩 Usage

```bash
sudo ./goping [options] <address>
```

### Options

| Flag | Description                        | Default |
| ---- | ---------------------------------- | ------- |
| `-c` | Number of packets to send          | `4`     |
| `-i` | Interval between packets (seconds) | `1`     |

### Example

```bash
sudo ./goping -c 5 -i 0.5 8.8.8.8
```

**Output:**

```
64 bytes from 8.8.8.8: ttl=118 time=23.1 ms
64 bytes from 8.8.8.8: ttl=118 time=21.8 ms
Request timed out

--- 8.8.8.8 ping statistics ---
5 packets transmitted, 4 received, 20.0% packet loss
rtt min/avg/max = 21.8/22.5/23.1 ms
```

Press **Ctrl+C** at any time to interrupt the process —
`goping` will print final statistics before exiting.

---

## 🧠 Project Structure

```
goping/
├── cmd
│   └── app
└── internal
    ├── flags
    ├── icmp
    │   ├── reply
    │   └── request
    ├── ping
    ├── print
    └── statistics

```

---

## 🧰 Technical Notes

* Uses Go’s `x/net/icmp` and `x/net/ipv4` packages for raw socket access.
* Requires root privileges (`sudo`) to send ICMP packets.
* Handles `SIGINT` (Ctrl+C) and `SIGTERM` for graceful termination.
* Default payload size: 56 bytes.

---

## 💡 Future Improvements

* Add TTL (`-t`) and timeout (`-W`) flags
* Support for IPv6
* Continuous ping mode (`-c 0` = infinite)
* JSON/CSV output for automation

