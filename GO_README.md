# Go: Document Processing Engine

## Mission
Build a CLI that imports TXT/JSON documents, extracts metadata, validates and classifies them, persists an index, and exports reports. Use the standard library first.

**Time budget:** 90 consecutive calendar days, approximately 60 minutes per day. Days 1–3 set up the project; days 4–72 cover 23 classical patterns at three days each; days 73–86 cover seven supplemental topics at two days each; days 87–90 are integration and review. This means **one pattern per three days in this language**, not one pattern per day across all three languages. If studying all three simultaneously, budget about three hours per day or stagger the tracks.

## Ground rules
- Write all application and test code yourself. This document provides requirements, not solutions.
- Before implementing each pattern, write a short baseline solution without it or describe the existing pain point.
- After implementation, record why the pattern helps, its cost, two other plausible use cases, and a case where it should not be used.
- Use at least two positive tests and one negative/edge test per assignment; preserve all previous behavior.
- Do not force inheritance-based textbook structures into languages where functions, interfaces, composition, or framework facilities are simpler.
- Commit at the end of every three-day block; use meaningful commit messages.

## Initial scope and setup: days 1–3
Create a Go module, CLI entry point, internal packages for domain/import/process/store/export, fixtures and tests. Start with a single TXT document and an in-memory index; add JSON and disk persistence only when assignments call for them. Use interfaces at consumer boundaries, explicit constructors, context cancellation and table-driven tests.

## Daily rhythm
**First day of each three-day pattern block:** 15 minutes reading, 15 minutes identifying the change pressure, 30 minutes drafting an interface and tests. **Second day:** 10 minutes revisiting the design, 40 minutes implementing, 10 minutes testing. **Third day:** 30 minutes edge cases and refactoring, 20 minutes explaining trade-offs and two additional use cases, 10 minutes committing. For a faster one-day variant, combine the three stages in a 60-minute focused session with a smaller scope.

### Three required written references for every pattern
1. **GoF:** Gamma, Helm, Johnson and Vlissides, *Design Patterns: Elements of Reusable Object-Oriented Software* (1994), chapter for the named pattern. For supplemental patterns use Fowler, *Patterns of Enterprise Application Architecture*, or the official platform documentation instead.
2. **Refactoring.Guru:** https://refactoring.guru/design-patterns — open the named pattern and read Intent, Problem, Solution, Applicability, and Pros/Cons. For supplemental patterns, use https://martinfowler.com/eaaCatalog/ and the platform-specific references below.
3. **Pattern Garden:** https://designpatterns.guru/ — locate the named pattern and compare the illustrated use case. For supplemental patterns, substitute the relevant official documentation below.

Platform-specific written references: Go https://go.dev/doc/effective_go and https://go.dev/doc/ ; Rails https://guides.rubyonrails.org/ ; JavaScript https://developer.mozilla.org/en-US/docs/Web/JavaScript and https://developer.mozilla.org/en-US/docs/Web/API . No videos are required. Read references for understanding; do not copy example implementations into the project.


## Pattern-by-pattern assignments

### Days 4–6: Factory Method
**Essence and purpose:** Create interchangeable products through a creation hook. **Project assignment:** Extend the application with new input/output provider. Preserve the existing external behavior and make the new variant independently testable.

**Day 4 — Design:** Read the three written references for **Factory Method**. Describe the coupling or variation in the current implementation, identify participants and their responsibilities, and sketch two candidate designs (pattern-based and simpler alternative). Write failing tests for the intended change.

**Day 5 — Implement:** Implement the smallest working version in this project's language. Add a second variant to demonstrate that the design is genuinely extensible. Avoid introducing classes, interfaces, or modules solely to mimic a diagram.

**Day 6 — Verify:** Test the two variants and one failure case. Explain how the pattern works, why it is useful here, and its trade-offs. Name at least two other use cases (e.g., alternate providers and alternate formats) and one situation where direct code is clearer. Commit.

**Acceptance:** The new variant can be selected without rewriting its consumer; old tests pass; a short design note identifies the pattern's roles and an explicit reason to use or reject it.

### Days 7–9: Abstract Factory
**Essence and purpose:** Create compatible families of collaborators. **Project assignment:** Extend the application with local versus cloud processing bundle. Preserve the existing external behavior and make the new variant independently testable.

**Day 7 — Design:** Read the three written references for **Abstract Factory**. Describe the coupling or variation in the current implementation, identify participants and their responsibilities, and sketch two candidate designs (pattern-based and simpler alternative). Write failing tests for the intended change.

**Day 8 — Implement:** Implement the smallest working version in this project's language. Add a second variant to demonstrate that the design is genuinely extensible. Avoid introducing classes, interfaces, or modules solely to mimic a diagram.

**Day 9 — Verify:** Test the two variants and one failure case. Explain how the pattern works, why it is useful here, and its trade-offs. Name at least two other use cases (e.g., alternate providers and alternate formats) and one situation where direct code is clearer. Commit.

**Acceptance:** The new variant can be selected without rewriting its consumer; old tests pass; a short design note identifies the pattern's roles and an explicit reason to use or reject it.

### Days 10–12: Builder
**Essence and purpose:** Assemble complex configurations incrementally. **Project assignment:** Extend the application with processing job configuration. Preserve the existing external behavior and make the new variant independently testable.

**Day 10 — Design:** Read the three written references for **Builder**. Describe the coupling or variation in the current implementation, identify participants and their responsibilities, and sketch two candidate designs (pattern-based and simpler alternative). Write failing tests for the intended change.

**Day 11 — Implement:** Implement the smallest working version in this project's language. Add a second variant to demonstrate that the design is genuinely extensible. Avoid introducing classes, interfaces, or modules solely to mimic a diagram.

**Day 12 — Verify:** Test the two variants and one failure case. Explain how the pattern works, why it is useful here, and its trade-offs. Name at least two other use cases (e.g., alternate providers and alternate formats) and one situation where direct code is clearer. Commit.

**Acceptance:** The new variant can be selected without rewriting its consumer; old tests pass; a short design note identifies the pattern's roles and an explicit reason to use or reject it.

### Days 13–15: Prototype
**Essence and purpose:** Copy configured objects without sharing mutable state. **Project assignment:** Extend the application with processing template. Preserve the existing external behavior and make the new variant independently testable.

**Day 13 — Design:** Read the three written references for **Prototype**. Describe the coupling or variation in the current implementation, identify participants and their responsibilities, and sketch two candidate designs (pattern-based and simpler alternative). Write failing tests for the intended change.

**Day 14 — Implement:** Implement the smallest working version in this project's language. Add a second variant to demonstrate that the design is genuinely extensible. Avoid introducing classes, interfaces, or modules solely to mimic a diagram.

**Day 15 — Verify:** Test the two variants and one failure case. Explain how the pattern works, why it is useful here, and its trade-offs. Name at least two other use cases (e.g., alternate providers and alternate formats) and one situation where direct code is clearer. Commit.

**Acceptance:** The new variant can be selected without rewriting its consumer; old tests pass; a short design note identifies the pattern's roles and an explicit reason to use or reject it.

### Days 16–18: Singleton
**Essence and purpose:** Control a single shared instance and examine global-state costs. **Project assignment:** Extend the application with configuration registry. Preserve the existing external behavior and make the new variant independently testable.

**Day 16 — Design:** Read the three written references for **Singleton**. Describe the coupling or variation in the current implementation, identify participants and their responsibilities, and sketch two candidate designs (pattern-based and simpler alternative). Write failing tests for the intended change.

**Day 17 — Implement:** Implement the smallest working version in this project's language. Add a second variant to demonstrate that the design is genuinely extensible. Avoid introducing classes, interfaces, or modules solely to mimic a diagram.

**Day 18 — Verify:** Test the two variants and one failure case. Explain how the pattern works, why it is useful here, and its trade-offs. Name at least two other use cases (e.g., alternate providers and alternate formats) and one situation where direct code is clearer. Commit.

**Acceptance:** The new variant can be selected without rewriting its consumer; old tests pass; a short design note identifies the pattern's roles and an explicit reason to use or reject it.

### Days 19–21: Adapter
**Essence and purpose:** Translate an incompatible interface. **Project assignment:** Extend the application with legacy metadata source. Preserve the existing external behavior and make the new variant independently testable.

**Day 19 — Design:** Read the three written references for **Adapter**. Describe the coupling or variation in the current implementation, identify participants and their responsibilities, and sketch two candidate designs (pattern-based and simpler alternative). Write failing tests for the intended change.

**Day 20 — Implement:** Implement the smallest working version in this project's language. Add a second variant to demonstrate that the design is genuinely extensible. Avoid introducing classes, interfaces, or modules solely to mimic a diagram.

**Day 21 — Verify:** Test the two variants and one failure case. Explain how the pattern works, why it is useful here, and its trade-offs. Name at least two other use cases (e.g., alternate providers and alternate formats) and one situation where direct code is clearer. Commit.

**Acceptance:** The new variant can be selected without rewriting its consumer; old tests pass; a short design note identifies the pattern's roles and an explicit reason to use or reject it.

### Days 22–24: Bridge
**Essence and purpose:** Vary abstraction and implementation independently. **Project assignment:** Extend the application with document type versus storage backend. Preserve the existing external behavior and make the new variant independently testable.

**Day 22 — Design:** Read the three written references for **Bridge**. Describe the coupling or variation in the current implementation, identify participants and their responsibilities, and sketch two candidate designs (pattern-based and simpler alternative). Write failing tests for the intended change.

**Day 23 — Implement:** Implement the smallest working version in this project's language. Add a second variant to demonstrate that the design is genuinely extensible. Avoid introducing classes, interfaces, or modules solely to mimic a diagram.

**Day 24 — Verify:** Test the two variants and one failure case. Explain how the pattern works, why it is useful here, and its trade-offs. Name at least two other use cases (e.g., alternate providers and alternate formats) and one situation where direct code is clearer. Commit.

**Acceptance:** The new variant can be selected without rewriting its consumer; old tests pass; a short design note identifies the pattern's roles and an explicit reason to use or reject it.

### Days 25–27: Composite
**Essence and purpose:** Treat individual objects and groups uniformly. **Project assignment:** Extend the application with nested document collections. Preserve the existing external behavior and make the new variant independently testable.

**Day 25 — Design:** Read the three written references for **Composite**. Describe the coupling or variation in the current implementation, identify participants and their responsibilities, and sketch two candidate designs (pattern-based and simpler alternative). Write failing tests for the intended change.

**Day 26 — Implement:** Implement the smallest working version in this project's language. Add a second variant to demonstrate that the design is genuinely extensible. Avoid introducing classes, interfaces, or modules solely to mimic a diagram.

**Day 27 — Verify:** Test the two variants and one failure case. Explain how the pattern works, why it is useful here, and its trade-offs. Name at least two other use cases (e.g., alternate providers and alternate formats) and one situation where direct code is clearer. Commit.

**Acceptance:** The new variant can be selected without rewriting its consumer; old tests pass; a short design note identifies the pattern's roles and an explicit reason to use or reject it.

### Days 28–30: Decorator
**Essence and purpose:** Add behavior around an existing component. **Project assignment:** Extend the application with logging and caching around importer. Preserve the existing external behavior and make the new variant independently testable.

**Day 28 — Design:** Read the three written references for **Decorator**. Describe the coupling or variation in the current implementation, identify participants and their responsibilities, and sketch two candidate designs (pattern-based and simpler alternative). Write failing tests for the intended change.

**Day 29 — Implement:** Implement the smallest working version in this project's language. Add a second variant to demonstrate that the design is genuinely extensible. Avoid introducing classes, interfaces, or modules solely to mimic a diagram.

**Day 30 — Verify:** Test the two variants and one failure case. Explain how the pattern works, why it is useful here, and its trade-offs. Name at least two other use cases (e.g., alternate providers and alternate formats) and one situation where direct code is clearer. Commit.

**Acceptance:** The new variant can be selected without rewriting its consumer; old tests pass; a short design note identifies the pattern's roles and an explicit reason to use or reject it.

### Days 31–33: Facade
**Essence and purpose:** Expose a small interface over a complex subsystem. **Project assignment:** Extend the application with one-command processing workflow. Preserve the existing external behavior and make the new variant independently testable.

**Day 31 — Design:** Read the three written references for **Facade**. Describe the coupling or variation in the current implementation, identify participants and their responsibilities, and sketch two candidate designs (pattern-based and simpler alternative). Write failing tests for the intended change.

**Day 32 — Implement:** Implement the smallest working version in this project's language. Add a second variant to demonstrate that the design is genuinely extensible. Avoid introducing classes, interfaces, or modules solely to mimic a diagram.

**Day 33 — Verify:** Test the two variants and one failure case. Explain how the pattern works, why it is useful here, and its trade-offs. Name at least two other use cases (e.g., alternate providers and alternate formats) and one situation where direct code is clearer. Commit.

**Acceptance:** The new variant can be selected without rewriting its consumer; old tests pass; a short design note identifies the pattern's roles and an explicit reason to use or reject it.

### Days 34–36: Flyweight
**Essence and purpose:** Share immutable repeated state. **Project assignment:** Extend the application with shared classification rules. Preserve the existing external behavior and make the new variant independently testable.

**Day 34 — Design:** Read the three written references for **Flyweight**. Describe the coupling or variation in the current implementation, identify participants and their responsibilities, and sketch two candidate designs (pattern-based and simpler alternative). Write failing tests for the intended change.

**Day 35 — Implement:** Implement the smallest working version in this project's language. Add a second variant to demonstrate that the design is genuinely extensible. Avoid introducing classes, interfaces, or modules solely to mimic a diagram.

**Day 36 — Verify:** Test the two variants and one failure case. Explain how the pattern works, why it is useful here, and its trade-offs. Name at least two other use cases (e.g., alternate providers and alternate formats) and one situation where direct code is clearer. Commit.

**Acceptance:** The new variant can be selected without rewriting its consumer; old tests pass; a short design note identifies the pattern's roles and an explicit reason to use or reject it.

### Days 37–39: Proxy
**Essence and purpose:** Control access to an underlying service. **Project assignment:** Extend the application with lazy document loader. Preserve the existing external behavior and make the new variant independently testable.

**Day 37 — Design:** Read the three written references for **Proxy**. Describe the coupling or variation in the current implementation, identify participants and their responsibilities, and sketch two candidate designs (pattern-based and simpler alternative). Write failing tests for the intended change.

**Day 38 — Implement:** Implement the smallest working version in this project's language. Add a second variant to demonstrate that the design is genuinely extensible. Avoid introducing classes, interfaces, or modules solely to mimic a diagram.

**Day 39 — Verify:** Test the two variants and one failure case. Explain how the pattern works, why it is useful here, and its trade-offs. Name at least two other use cases (e.g., alternate providers and alternate formats) and one situation where direct code is clearer. Commit.

**Acceptance:** The new variant can be selected without rewriting its consumer; old tests pass; a short design note identifies the pattern's roles and an explicit reason to use or reject it.

### Days 40–42: Chain of Responsibility
**Essence and purpose:** Pass a request through ordered handlers. **Project assignment:** Extend the application with validation and classification chain. Preserve the existing external behavior and make the new variant independently testable.

**Day 40 — Design:** Read the three written references for **Chain of Responsibility**. Describe the coupling or variation in the current implementation, identify participants and their responsibilities, and sketch two candidate designs (pattern-based and simpler alternative). Write failing tests for the intended change.

**Day 41 — Implement:** Implement the smallest working version in this project's language. Add a second variant to demonstrate that the design is genuinely extensible. Avoid introducing classes, interfaces, or modules solely to mimic a diagram.

**Day 42 — Verify:** Test the two variants and one failure case. Explain how the pattern works, why it is useful here, and its trade-offs. Name at least two other use cases (e.g., alternate providers and alternate formats) and one situation where direct code is clearer. Commit.

**Acceptance:** The new variant can be selected without rewriting its consumer; old tests pass; a short design note identifies the pattern's roles and an explicit reason to use or reject it.

### Days 43–45: Command
**Essence and purpose:** Represent actions as objects or values. **Project assignment:** Extend the application with reprocessable import actions. Preserve the existing external behavior and make the new variant independently testable.

**Day 43 — Design:** Read the three written references for **Command**. Describe the coupling or variation in the current implementation, identify participants and their responsibilities, and sketch two candidate designs (pattern-based and simpler alternative). Write failing tests for the intended change.

**Day 44 — Implement:** Implement the smallest working version in this project's language. Add a second variant to demonstrate that the design is genuinely extensible. Avoid introducing classes, interfaces, or modules solely to mimic a diagram.

**Day 45 — Verify:** Test the two variants and one failure case. Explain how the pattern works, why it is useful here, and its trade-offs. Name at least two other use cases (e.g., alternate providers and alternate formats) and one situation where direct code is clearer. Commit.

**Acceptance:** The new variant can be selected without rewriting its consumer; old tests pass; a short design note identifies the pattern's roles and an explicit reason to use or reject it.

### Days 46–48: Iterator
**Essence and purpose:** Traverse a collection without exposing its representation. **Project assignment:** Extend the application with paged document traversal. Preserve the existing external behavior and make the new variant independently testable.

**Day 46 — Design:** Read the three written references for **Iterator**. Describe the coupling or variation in the current implementation, identify participants and their responsibilities, and sketch two candidate designs (pattern-based and simpler alternative). Write failing tests for the intended change.

**Day 47 — Implement:** Implement the smallest working version in this project's language. Add a second variant to demonstrate that the design is genuinely extensible. Avoid introducing classes, interfaces, or modules solely to mimic a diagram.

**Day 48 — Verify:** Test the two variants and one failure case. Explain how the pattern works, why it is useful here, and its trade-offs. Name at least two other use cases (e.g., alternate providers and alternate formats) and one situation where direct code is clearer. Commit.

**Acceptance:** The new variant can be selected without rewriting its consumer; old tests pass; a short design note identifies the pattern's roles and an explicit reason to use or reject it.

### Days 49–51: Mediator
**Essence and purpose:** Coordinate collaborators through a central component. **Project assignment:** Extend the application with pipeline coordinator. Preserve the existing external behavior and make the new variant independently testable.

**Day 49 — Design:** Read the three written references for **Mediator**. Describe the coupling or variation in the current implementation, identify participants and their responsibilities, and sketch two candidate designs (pattern-based and simpler alternative). Write failing tests for the intended change.

**Day 50 — Implement:** Implement the smallest working version in this project's language. Add a second variant to demonstrate that the design is genuinely extensible. Avoid introducing classes, interfaces, or modules solely to mimic a diagram.

**Day 51 — Verify:** Test the two variants and one failure case. Explain how the pattern works, why it is useful here, and its trade-offs. Name at least two other use cases (e.g., alternate providers and alternate formats) and one situation where direct code is clearer. Commit.

**Acceptance:** The new variant can be selected without rewriting its consumer; old tests pass; a short design note identifies the pattern's roles and an explicit reason to use or reject it.

### Days 52–54: Memento
**Essence and purpose:** Capture and restore state safely. **Project assignment:** Extend the application with processing configuration snapshots. Preserve the existing external behavior and make the new variant independently testable.

**Day 52 — Design:** Read the three written references for **Memento**. Describe the coupling or variation in the current implementation, identify participants and their responsibilities, and sketch two candidate designs (pattern-based and simpler alternative). Write failing tests for the intended change.

**Day 53 — Implement:** Implement the smallest working version in this project's language. Add a second variant to demonstrate that the design is genuinely extensible. Avoid introducing classes, interfaces, or modules solely to mimic a diagram.

**Day 54 — Verify:** Test the two variants and one failure case. Explain how the pattern works, why it is useful here, and its trade-offs. Name at least two other use cases (e.g., alternate providers and alternate formats) and one situation where direct code is clearer. Commit.

**Acceptance:** The new variant can be selected without rewriting its consumer; old tests pass; a short design note identifies the pattern's roles and an explicit reason to use or reject it.

### Days 55–57: Observer
**Essence and purpose:** Notify subscribers when events occur. **Project assignment:** Extend the application with processing progress subscribers. Preserve the existing external behavior and make the new variant independently testable.

**Day 55 — Design:** Read the three written references for **Observer**. Describe the coupling or variation in the current implementation, identify participants and their responsibilities, and sketch two candidate designs (pattern-based and simpler alternative). Write failing tests for the intended change.

**Day 56 — Implement:** Implement the smallest working version in this project's language. Add a second variant to demonstrate that the design is genuinely extensible. Avoid introducing classes, interfaces, or modules solely to mimic a diagram.

**Day 57 — Verify:** Test the two variants and one failure case. Explain how the pattern works, why it is useful here, and its trade-offs. Name at least two other use cases (e.g., alternate providers and alternate formats) and one situation where direct code is clearer. Commit.

**Acceptance:** The new variant can be selected without rewriting its consumer; old tests pass; a short design note identifies the pattern's roles and an explicit reason to use or reject it.

### Days 58–60: State
**Essence and purpose:** Change behavior according to explicit lifecycle state. **Project assignment:** Extend the application with queued/running/failed/completed jobs. Preserve the existing external behavior and make the new variant independently testable.

**Day 58 — Design:** Read the three written references for **State**. Describe the coupling or variation in the current implementation, identify participants and their responsibilities, and sketch two candidate designs (pattern-based and simpler alternative). Write failing tests for the intended change.

**Day 59 — Implement:** Implement the smallest working version in this project's language. Add a second variant to demonstrate that the design is genuinely extensible. Avoid introducing classes, interfaces, or modules solely to mimic a diagram.

**Day 60 — Verify:** Test the two variants and one failure case. Explain how the pattern works, why it is useful here, and its trade-offs. Name at least two other use cases (e.g., alternate providers and alternate formats) and one situation where direct code is clearer. Commit.

**Acceptance:** The new variant can be selected without rewriting its consumer; old tests pass; a short design note identifies the pattern's roles and an explicit reason to use or reject it.

### Days 61–63: Strategy
**Essence and purpose:** Swap algorithms behind one contract. **Project assignment:** Extend the application with classification algorithms. Preserve the existing external behavior and make the new variant independently testable.

**Day 61 — Design:** Read the three written references for **Strategy**. Describe the coupling or variation in the current implementation, identify participants and their responsibilities, and sketch two candidate designs (pattern-based and simpler alternative). Write failing tests for the intended change.

**Day 62 — Implement:** Implement the smallest working version in this project's language. Add a second variant to demonstrate that the design is genuinely extensible. Avoid introducing classes, interfaces, or modules solely to mimic a diagram.

**Day 63 — Verify:** Test the two variants and one failure case. Explain how the pattern works, why it is useful here, and its trade-offs. Name at least two other use cases (e.g., alternate providers and alternate formats) and one situation where direct code is clearer. Commit.

**Acceptance:** The new variant can be selected without rewriting its consumer; old tests pass; a short design note identifies the pattern's roles and an explicit reason to use or reject it.

### Days 64–66: Template Method
**Essence and purpose:** Fix workflow skeleton and customize steps. **Project assignment:** Extend the application with document import workflow. Preserve the existing external behavior and make the new variant independently testable.

**Day 64 — Design:** Read the three written references for **Template Method**. Describe the coupling or variation in the current implementation, identify participants and their responsibilities, and sketch two candidate designs (pattern-based and simpler alternative). Write failing tests for the intended change.

**Day 65 — Implement:** Implement the smallest working version in this project's language. Add a second variant to demonstrate that the design is genuinely extensible. Avoid introducing classes, interfaces, or modules solely to mimic a diagram.

**Day 66 — Verify:** Test the two variants and one failure case. Explain how the pattern works, why it is useful here, and its trade-offs. Name at least two other use cases (e.g., alternate providers and alternate formats) and one situation where direct code is clearer. Commit.

**Acceptance:** The new variant can be selected without rewriting its consumer; old tests pass; a short design note identifies the pattern's roles and an explicit reason to use or reject it.

### Days 67–69: Visitor
**Essence and purpose:** Add operations across a stable object structure. **Project assignment:** Extend the application with export and report over document nodes. Preserve the existing external behavior and make the new variant independently testable.

**Day 67 — Design:** Read the three written references for **Visitor**. Describe the coupling or variation in the current implementation, identify participants and their responsibilities, and sketch two candidate designs (pattern-based and simpler alternative). Write failing tests for the intended change.

**Day 68 — Implement:** Implement the smallest working version in this project's language. Add a second variant to demonstrate that the design is genuinely extensible. Avoid introducing classes, interfaces, or modules solely to mimic a diagram.

**Day 69 — Verify:** Test the two variants and one failure case. Explain how the pattern works, why it is useful here, and its trade-offs. Name at least two other use cases (e.g., alternate providers and alternate formats) and one situation where direct code is clearer. Commit.

**Acceptance:** The new variant can be selected without rewriting its consumer; old tests pass; a short design note identifies the pattern's roles and an explicit reason to use or reject it.

### Days 70–72: Interpreter
**Essence and purpose:** Interpret a small domain-specific grammar. **Project assignment:** Extend the application with metadata filter expression. Preserve the existing external behavior and make the new variant independently testable.

**Day 70 — Design:** Read the three written references for **Interpreter**. Describe the coupling or variation in the current implementation, identify participants and their responsibilities, and sketch two candidate designs (pattern-based and simpler alternative). Write failing tests for the intended change.

**Day 71 — Implement:** Implement the smallest working version in this project's language. Add a second variant to demonstrate that the design is genuinely extensible. Avoid introducing classes, interfaces, or modules solely to mimic a diagram.

**Day 72 — Verify:** Test the two variants and one failure case. Explain how the pattern works, why it is useful here, and its trade-offs. Name at least two other use cases (e.g., alternate providers and alternate formats) and one situation where direct code is clearer. Commit.

**Acceptance:** The new variant can be selected without rewriting its consumer; old tests pass; a short design note identifies the pattern's roles and an explicit reason to use or reject it.

### Days 73–74: Dependency Injection
**Purpose:** Pass collaborators explicitly; isolate side effects. **Assignment:** inject filesystem and clock.

**Day 73:** Read the relevant official platform documentation and Fowler's enterprise pattern catalog where applicable; draw the data/control flow and specify the failure case. **Day 74:** Implement and test one normal and one failure path; document the costs and compare to a simpler approach. For this supplemental topic use three written references: Fowler's catalog, official platform documentation, and the relevant MDN/Rails/Go reference listed above.

### Days 75–76: Repository and Unit of Work
**Purpose:** Separate domain operations from persistence and coordinate transactions. **Assignment:** document repository and atomic batch.

**Day 75:** Read the relevant official platform documentation and Fowler's enterprise pattern catalog where applicable; draw the data/control flow and specify the failure case. **Day 76:** Implement and test one normal and one failure path; document the costs and compare to a simpler approach. For this supplemental topic use three written references: Fowler's catalog, official platform documentation, and the relevant MDN/Rails/Go reference listed above.

### Days 77–78: MVC and Presentation Model
**Purpose:** Separate domain state, presentation state, and input handling. **Assignment:** CLI command/view separation.

**Day 77:** Read the relevant official platform documentation and Fowler's enterprise pattern catalog where applicable; draw the data/control flow and specify the failure case. **Day 78:** Implement and test one normal and one failure path; document the costs and compare to a simpler approach. For this supplemental topic use three written references: Fowler's catalog, official platform documentation, and the relevant MDN/Rails/Go reference listed above.

### Days 79–80: Pub/Sub and Event Bus
**Purpose:** Decouple event producers from multiple consumers. **Assignment:** progress event bus.

**Day 79:** Read the relevant official platform documentation and Fowler's enterprise pattern catalog where applicable; draw the data/control flow and specify the failure case. **Day 80:** Implement and test one normal and one failure path; document the costs and compare to a simpler approach. For this supplemental topic use three written references: Fowler's catalog, official platform documentation, and the relevant MDN/Rails/Go reference listed above.

### Days 81–82: Pipeline and Middleware
**Purpose:** Compose ordered processing stages. **Assignment:** bounded document pipeline.

**Day 81:** Read the relevant official platform documentation and Fowler's enterprise pattern catalog where applicable; draw the data/control flow and specify the failure case. **Day 82:** Implement and test one normal and one failure path; document the costs and compare to a simpler approach. For this supplemental topic use three written references: Fowler's catalog, official platform documentation, and the relevant MDN/Rails/Go reference listed above.

### Days 83–84: Concurrency and Async Coordination
**Purpose:** Manage cancellation, ordering, retries and failures. **Assignment:** worker pool with context.

**Day 83:** Read the relevant official platform documentation and Fowler's enterprise pattern catalog where applicable; draw the data/control flow and specify the failure case. **Day 84:** Implement and test one normal and one failure path; document the costs and compare to a simpler approach. For this supplemental topic use three written references: Fowler's catalog, official platform documentation, and the relevant MDN/Rails/Go reference listed above.

### Days 85–86: Capstone and Pattern Audit
**Purpose:** Evaluate architecture and remove unnecessary abstractions. **Assignment:** end-to-end import/export.

**Day 85:** Read the relevant official platform documentation and Fowler's enterprise pattern catalog where applicable; draw the data/control flow and specify the failure case. **Day 86:** Implement and test one normal and one failure path; document the costs and compare to a simpler approach. For this supplemental topic use three written references: Fowler's catalog, official platform documentation, and the relevant MDN/Rails/Go reference listed above.

## Days 87–90: integration and assessment
- **Day 87:** Run an end-to-end user journey; identify missing integration points and write failing integration tests.
- **Day 88:** Fix integration failures and document data flow, module boundaries and error handling.
- **Day 89:** Audit every introduced pattern: keep, simplify or remove it based on observed requirements. Note at least three examples of overengineering avoided.
- **Day 90:** Demonstrate the application, run the full test suite, explain five randomly selected patterns without notes, and submit a short architecture decision record for the most consequential design choice.

## Evaluation rubric
A pattern is complete only if the feature works, tests cover normal and failure behavior, the trainee can state the concrete change pressure it addresses, and can articulate why a simpler implementation might be preferable. Prioritize maintainability and clear behavior over pattern count.
