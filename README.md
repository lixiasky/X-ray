# Xray: Full Source Code of Behavior Chain Tracing System

This is a high-privilege malware and backdoor behavior chain auditing tool. It compares the current system state with the original ISO image, automatically identifies modified files, performs decoy-based induction, traces behavior chains, exports diagrams, and supports limited process countermeasures.

---

## Quick Start

Please ensure you have installed the Go programming environment.

```bash
go build -o xray
sudo ./xray
```

After launching, the program will prompt you to enter the path to your ISO file (used as the baseline for comparison), and will then automatically begin scanning the entire system.

---

## Screenshots

Full runtime screenshots can be found in the `screenshot/` folder, including `.dot` file contents, visualization diagrams, and behavior detection records.

Some of the `failed` and `error` messages in the screenshots are expected: the program may attempt to access protected directories or processes and be denied. This is part of the system's normal protection and does not affect the overall functionality.

---

## Project Structure and File Descriptions

### Root Directory

- `main.go`: Program entry point. Initializes logging, module scheduling, and overall control flow.
- `go.mod` & `go.sum`: Go module files declaring dependencies.

### core/ - Main Core Modules

- `autodefense.go`: Auto-defense logic. Attempts to terminate suspicious processes and mark threats.
- `compare.go`: Compares current system files with those in the ISO image to detect modified targets.
- `monitor.go`: Monitors system-level file and directory changes for real-time reactions.
- `procwatcher.go`: Monitors newly created processes, extracts source paths and behaviors.
- `responder.go`: Responds to abnormal behavior in cooperation with the auto-defense module.
- `scanner.go`: Main scanning logic. Traverses directories, hashes files, and records actions.
- `trace.go`: Constructs actions into a behavior chain structure for analysis.
- `tracker.go`: Tracks system behavior paths and assists in graph exports.

### decoy/

- `decoy.go`: Creates decoy files or structures to detect unauthorized access or modifications.

### engine/

- `hash.go`: File hashing module (supports SHA256) for content difference comparison.
- `iso.go`: Parses ISO image contents for use as a comparison baseline.

### export/

- `graph.go`: Exports behavior chains in `.dot` format for visualization via Graphviz.
- `log_export.go`: Attempts to export data as JSON and LOG (not yet successful).

### platform/

- `linux.go`: Linux-specific platform handling.
- `windows.go`: Placeholder for Windows support (not yet implemented).

---

## Current Status

- `.dot` export is **implemented** and can be visualized using Graphviz.
- `.json` and `.log` exports are **not yet functional** — still under development.
- Monitoring and decoy-based triggers are working; some malicious processes can be auto-terminated.

---

## Author Statement and Licensing

The author is currently a first-year student at a regular university in China, not from a CS background, and self-learning Go and system-level programming. This project was born from their curiosity while studying OS and cybersecurity and is still a work in progress.

This project is open-source with permission-based licensing: anyone is free to fork and modify it without submitting PRs, as long as you **retain the line “First version written by TangTian” in your README**.

Please respect this sole requirement and avoid removing the author’s name.

This project is licensed under the Apache License 2.0. See LICENSE for details.

First developed and open-sourced by Tang Tian on 2025-05-17.

Fun fact: During initial testing, the tool terminated VSCode while the author was installing the Graphviz Interactive Preview plugin — showcasing its real-world detection capabilities (author QwQ).

If you'd like to get in touch, feel free to reach me : lixiasky@protonmail.com — I'm always open to feedback or discussion.

---

## Legal Notice

A provisional patent application has been filed with the United States Patent and Trademark Office (USPTO) for this project under the name of the author. All rights reserved.

This is an independent research work, created and maintained by an individual student without institutional backing.

---

## Donations

If you appreciate this project, you’re welcome to donate via the author's QR code.  
(**Limit: 1 CNY = approx. $0.13 USD**)

![Alipay donation QR code](donate_alipay.jpg)

---

## Disclaimer

This project is intended for educational and research purposes only.  
Unauthorized scanning, intrusion, or damage to any system is strictly prohibited.  
The author bears no legal responsibility or liability for any misuse of this tool.

---
