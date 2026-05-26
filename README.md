# 🚀 100 Go (Golang) Programming Challenges & Backend Projects

<p align="">
  <img src="https://blog.suprabha.me/static/1-635fc115b7d799983a3a02926db54bdf.gif" alt="Golang Challenges Demo" width="30%" height="30%">
</p>


Welcome to the **Ultimate Go 100 Challenge**! This repository is a curated roadmap of 100 structured production-grade programming challenges designed to master high-performance backend development using Golang. 

The progression takes you from fundamental algorithmic problem-solving to highly scalable concurrent patterns, advanced memory management, data persistence, and microservices architecture.

---

## 🛠️ Key Learning Objectives & Core Competencies
By exploring these challenges, you will build production-ready backend engineering skills including:
* **System Design & Core Logic:** Mastering native Go data structures (`maps`, `slices`, `structs`).
* **Concurrency Paradigms:** High-performance systems using `Goroutines`, `Channels`, `sync.Mutex`, and `Context`.
* **Standard Library Mastery:** Full control over `net/http`, `os`, `io`, `time`, `encoding/json`, and `crypto`.
* **Clean Code Architecture:** Domain-Driven Design (DDD), interface abstractions, and explicit error handling.

---

## 📋 Repository Progress Matrix


| Project Code | Challenge Title | Advanced Technical Concepts & Go Packages | Level | Status |
| :---: | :--- | :--- | :---: | :---: |
| **`p001`** | Temperature Converter | CLI I/O, `fmt` string formatting, explicit Type Assertions | Beginner | ✅ Done |
| **`p002`** | ATM Security | Multi-layered condition logic, `strings` sanitation, custom control flow | Beginner | ✅ Done |
| **`p003`** | Word Frequency Counter | Native `maps` analytics, string parsing, tokenization | Beginner | ✅ Done |
| **`p004`** | Invoice Generator | Compound `struct` modeling, deterministic math floating calculations | Beginner | ✅ Done |
| **`p005`** | Student Grade Calculator | Dynamic memory `slices`, basic slice analytics & statistical reductions | Beginner | ✅ Done |
| **`p006`** | Bank Account CLI | Object-Oriented patterns via Pointer Receivers and Struct Methods | Beginner | ✅ Done |
| **`p007`** | Todo CLI Tool | Persistent State via JSON Serialization (`encoding/json`), File I/O (`os`) | Intermediate | ✅ Done |
| **`p008`** | Shopping Cart System | Key-value store simulations using complex Maps and structural pointer modifications | Intermediate | ✅ Done |
| **`p009`** | Password Validator | Regular Expressions (`regexp`), custom Error handling architectures | Intermediate | ✅ Done |
| **`p010`** | Expense Tracker CLI | High-accuracy Date parsing (`time.Parse`), structural sorting logic | Intermediate | ✅ Done |
| **`p011`** | Library Manager | Multi-structural relationship models, sort abstractions for slices | Intermediate | ✅ Done |
| **`p012`** | URL Shortener Engine | Fast random generation (`math/rand/v2`), memory-efficient `strings.Builder` | Intermediate | ✅ Done |
| **`p013`** | Concurrent Web Crawler | Multiplexing execution via `Goroutines`, `Channels`, and `sync.WaitGroup` | Advanced | ⏳ Next |

---

## ⚙️ Local Development Guide

### Prerequisites
Make sure you have **Go 1.22 or higher** installed on your system.
```bash
go version
```

### Installation & Execution
1. **Clone the repository:**
   ```bash
   git clone https://github.com.git
   cd go-100-challenge
   ```

2. **Navigate and execute any specific project target:**
   ```bash
   cd p012_url_shortener
   go run main.go
   ```

---

## 🌟 Software Engineering Practices Applied
* **Idiomatic Go:** Avoiding over-engineering; keeping code readable, explicit, and performant.
* **Zero Magic Frameworks:** Relying heavily on the robust Go standard library to understand deep underlying engine mechanics.
* **Memory Optimization:** Utilizing proper pre-allocation structures (`strings.Builder`, explicit slice capacity scaling) to minimize garbage collection (GC) pressure.

---
⭐ **If you find this repository helpful for your Golang journey, please consider giving it a star!** 
