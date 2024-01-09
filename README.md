# EvoLang Compiler

## Overview

EvoLang is a Domain-Specific Language (DSL) designed to simplify the process of defining data models and their interactions in a web application context. It focuses on reducing boilerplate code and streamlining the development process, particularly for tasks involving data model definition, event handling, access control, and model interaction.

### Design Goals

1. **Simplicity**: Provide a straightforward syntax for defining models, events, and model functions.

2. **Efficiency**: Reduce the amount of code needed to define complex data models and their interactions.

3. **Flexibility**: Allow for custom event handling and complex interactions between models while maintaining a clear and concise syntax.

4. **Integration**: Allow for inlining of Go code, allowing integration with existing Go codebases and leveraging Go's existing ecosystem.

## Syntax

EvoLang uses a concise, readable syntax to define various aspects of web application data handling:

### Models

```evo
model User {
    userId: Int
    name: String
    email: String
    // Relationships
    lists: [TodoList] @linkedBy(creator)
}
```

### Events

```evo
model TodoItem {
    // Fields
    events {
        onStatusChange: Event<OldStatus, NewStatus> {
            <sandbox lang="go">
                // Embedded Go code for implementation details
            </sandbox>
        }
    }
}
```

### API Endpoints

```evo
endpoint "/todoList" {
    body: TodoList
    authenticate: User
    return TodoList.create(body)
}
```

### Server Definition and Main Function

```evo
server MyServer {
    endpoint "/todoList" { ... }
    // Other endpoints
}

main {
    MyServer.Listen(8080)
}
```

## Usage

### Setting Up

Ensure you have Go installed on your system. The EvoLang compiler is written in Go and requires the Go toolchain.

### Compiling EvoLang Code

1. Write your EvoLang code in a `.evo` file.

2. Run the EvoLang compiler with the `.evo` file as input:

   ```sh
   go run main.go your_file.evo
   ```

   This will generate Go code based on your EvoLang definitions.

3. The compiler automatically handles the Go code generation and compiles it into an executable binary.

### Running Your Application

After compiling, run the generated binary to start your web application:

```sh
./.evo-out/output
```

This starts the web server as defined in your EvoLang code, listening on the specified port and setting up the defined API endpoints.

## Contributing

Contributions to EvoLang are welcome! Whether it's improving the syntax, extending features, or fixing bugs, your input is valuable. Please feel free to fork the repository, make your changes, and submit a pull request.

---

EvoLang aims to streamline web application development, making it faster and more intuitive while leveraging the power and efficiency of Go. Whether you're building a small project or a large-scale application, EvoLang offers a unique approach to defining and managing your application's data layer and interactions.

