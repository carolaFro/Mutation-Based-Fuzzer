# Mutation-Based Fuzzer

## Overview

This project implements a mutation-based fuzzer that evaluates the robustness of a target program by automatically generating malformed inputs. The fuzzer works by corrupting a default seed input using randomized mutations and executing the target program to observe abnormal behavior.

The project was developed as part of a graduate-level software testing and security course and emphasizes automated robustness testing, reproducibility, and defensive security practices.

---

## Key Features

* Mutation-based fuzzing of a **default seed**
* Configurable fuzzing runs via input files
* Deterministic execution using a PRNG seed
* Automated execution across multiple iterations
* Corruption of arbitrary seed content
* Implemented in **Go (Golang)**

---

## Project Structure

```
Mutation-Based-Fuzzer/
├── fuzzer.go           # Fuzzer implementation
├── fuzzer.exe          # Compiled executable (optional)
├── seed                # Default seed file (arbitrary content)
├── prog_0.txt          # Config file
├── prog_1.txt
├── prog_2.txt
├── prog_3.txt
├── prog_4.txt
├── prog_5.txt
├── prog_6.txt
├── prog_7.txt
├── prog_8.txt
├── prog_9.txt
```

---

## How It Works

### 1. Seed Input

The file named `seed` serves as the base input for fuzzing.

> The seed file may contain any text or data.
> During execution, the fuzzer intentionally corrupts the seed contents to generate malformed inputs.

This allows experimentation with different baseline inputs without modifying the fuzzer logic.

---

### 2. Configuration Files (`prog_x.txt`)

Each configuration file (`prog_0.txt` through `prog_9.txt`) specifies:

* **`prng_seed`** – seed for the pseudo-random number generator (controls reproducibility)
* **`num_of_iters`** – number of mutation iterations to perform

Using fixed PRNG seeds ensures fuzzing runs are **deterministic and repeatable**.

---

### 3. Mutation Phase

For each iteration:

* The seed input is copied
* Random mutations are applied (byte modifications, insertions, deletions)
* The corrupted input is passed to the target program

---

### 4. Execution & Observation

The target program is executed with each mutated input.
The fuzzer monitors execution to detect:

* Crashes
* Abnormal termination
* Unexpected behavior

---

## Running the Fuzzer

### Requirements

* **Go 1.18+**
* Target program executable

### Build

```bash
go build fuzzer.go
```

### Run with a Config File

```bash
./fuzzer prog_3.txt
```

The fuzzer reads `prng_seed` and `num_of_iters` from the specified config file and fuzzes the default `seed` accordingly.

---

## Input Files

### `seed`

* Base input used for fuzzing
* Can contain **any arbitrary text**
* Contents are intentionally corrupted during fuzzing

### `prog_x.txt`

* Configuration files controlling fuzzing behavior
* Enable reproducible experiments

---

## Output

Depending on configuration and target behavior, the fuzzer may:

* Print execution status to the console
* Detect and report crashes
* Allow reproduction of failures using the same config file

---

## Limitations

* Black-box fuzzing (no coverage feedback)
* Random mutation strategy
* Designed for academic experimentation

