# Lem-in: Digital Ant Farm

## 📖 Description

**Lem-in** is a program that simulates a digital ant farm. The objective is to find the quickest path to move `n` ants from the starting room (`##start`) to the ending room (`##end`) through a network of rooms and tunnels.

The program reads an input file describing the colony and outputs the movements of the ants step-by-step while adhering to specific rules.

---

## ✨ Features

- **Input Parsing**: Reads and validates colony descriptions from the input file.
- **Pathfinding**: Computes the shortest path or paths to move ants efficiently.
- **Ant Movement Simulation**:
  - Ants move one by one to an empty room via tunnels.
  - Tunnels can only be used once per turn.
  - Avoids traffic jams for optimal efficiency.
- **Error Handling**: Displays detailed error messages for invalid input.

---

## 📄 Input Data Format

The input file should follow this structure:

1. **Number of ants** (on the first line).
2. **Rooms**: Defined as `name coord_x coord_y`.
   - Example: `Room1 23 3`.
   - Rules:
     - Room names cannot start with `L` or `#`.
     - Room names must not contain spaces.
3. **Special Commands**:
   - `##start`: Indicates the starting room.
   - `##end`: Indicates the ending room.
4. **Links**: Defined as `name1-name2`.
   - Example: `Room1-Room2`.

---
## ⚙️ Algorithm: Breadth-First Search (BFS)
The program uses the Breadth-First Search (BFS) algorithm to find the shortest path(s) from the ##start room to the ##end room. BFS is well-suited for this problem because it explores all possible paths layer by layer, ensuring the shortest path is found efficiently.

BFS Steps
Initialization: Start with the ##start room as the root node. Add it to a queue.
Exploration: For each room in the queue:
Explore its connected rooms (neighbors) through tunnels.
If a neighbor is unvisited and valid, mark it as visited and add it to the queue.
Termination: Stop when the ##end room is reached or when all paths have been explored.
Path Construction: Backtrack from the ##end room to construct the shortest path(s).
Benefits of BFS
Guarantees the shortest path in an unweighted graph (like this one, where all tunnels are of equal "weight").
Handles complex cases with multiple paths, loops, and disconnected components effectively.

## 🛠 Usage

To run the program, use the following command:

```bash
$ go run . <input_file>
```

## 🧪 Example
Input File (test.txt)

```plaintext
3
##start  
1 23 3  
2 16 7  
##end  
0 9 5  
1-2  
2-0
```
## ⚙️ Rules
The starting room (##start) and ending room (##end) can contain multiple ants.
Each intermediate room can only contain one ant at a time.
Tunnels can only be used once per turn.
Ants must avoid traffic jams for optimal performance.
## ❌ Error Handling
If the input is invalid, the program will display:

```plaintext
ERROR: invalid data format
Examples of invalid input:
Missing or incorrect number of ants.
No ##start or ##end room.
Invalid or duplicate room/link definitions.
No valid path between ##start and ##end.
```
## 📋 Requirements
Written in Go.
Only standard Go packages are allowed.
Includes unit tests for validation.
