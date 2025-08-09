# Expense Tracker API

## 📝 Project Description

**Expense Tracker API** is a simple expense tracker application to manage your finances. The application allows users to add, delete, and view their expenses. The application also provide a summary of the expenses.

This project serves as an excellent starting point for Go beginners to solidify fundamental concepts such as:

- Go Modules and Packages: Understanding project structure and dependencies.

- Structs and Slices: Defining and managing custom data collections.

- Functions and Error Handling: Writing modular, robust, and idiomatic Go code.

- File I/O: Reading from and writing to the local file system.

- JSON Serialization/Deserialization: Persisting structured data to disk.

## ✨ Requirements

Application should run from the command line and should have the following features:

- Users can add an expense with a description and amount.
- Users can update an expense.
- Users can delete an expense.
- Users can view all expenses.
- Users can view a summary of all expenses.
- Users can view a summary of expenses for a specific month (of current year).

Here are some additional features that you can add to the application:

- Add expense categories and allow users to filter expenses by category.
- Allow users to set a budget for each month and show a warning when the user exceeds the budget.
- Allow users to export expenses to a CSV file.

- Register User: Interactively add new user with a firstName, lastName, email, and password tags.

- User Login: Interactively login with a email, and password tags.

- Create Notes: Interactively add new notes with a title, and description tags.

- View Notes: Retrieve and display the full details of a specific user notes by user ID.

- Delete Note: Remove a note from your vault using its ID.

- Local Persistence: All notes and users are saved to a local JSON file, ensuring your data is available across sessions.

## 🛠️ Core Technologies Used

- Go (Golang): The primary programming language.

- Standard Library Packages:

  - fmt: For formatted input and output.

  - os: For operating system interactions, including file system operations.

  - strings: For string manipulation (e.g., trimming, splitting, case conversion).

  - encoding/json: For marshaling (saving) and unmarshaling (loading) Go structs to/from JSON.

## 🚀 Installation

To get Go Note Application up and running on your local machine, follow these steps:

1. Ensure Go is Installed:
   Make sure you have Go installed (version 1.18 or higher is recommended). You can download it from go.dev/dl/.
   Verify your installation:

```
go version
```

2. Clone the Repository (or create project manually):
   If you're starting from scratch as part of a learning exercise, you'd create the project structure manually as described in the task instructions. If this were a real repository:

```
git clone https://github.com/philipoyelegbin/golang-training.git/noteapp
cd noteApp
```

3. Initialize Go Module (if not already done):

```
go mod init github.com/philipoyelegbin/golang-training.git/noteapp # Only if you created the project manually
```

4. Build the Executable:
   This command compiles your Go source code into a single executable binary.

```
go build -o noteapp
```

This will create an executable file named noteapp in your project's root directory.

## 💡 Usage

Once built, you can run the CLI commands from your terminal.

**General Usage**

```
./noteapp     # Prompt you interactively to select an action to perform
```

**Commands**

This command will prompt you interactively

1. `1` - Register a new user

2. `2` - Create a new note

3. `3` - View your notes

4. `4` - Delete a notes

5. `5` - Exit the app

## 📂 Project Structure

```
noteapp/
├── main.go # Main entry point and CLI command handling
└── functions/ # Directory for defined logic
    ├── user.go # Defines the User input validation and file persistence logic
    ├── note.go # Defines the Note input validation and file persistence logic
└── schema/ # Directory for defined model
    ├── model.go # Defines the User and Note data model
└── db/ # Directory for data storage
    └── users.json # JSON file where all users are saved
    └── notes.json # JSON file where all notes are saved
```

File Breakdown:

- `main.go`:

  - Contains the main function, which is the entry point of the application.

  - Includes some helper functions to load users from file, load notes from file, and action selector.

  - Includes some global variables

  - Orchestrates calls to the external functions.

- `functions/users.go`:

  - Contains functions to perform certain tasks like;

    - generate a user id for new user

    - validate user input for creating new user and login

    - save a new user to file

    - creating a new user

    - log in as a user

- `functions/notes.go`:

  - Contains functions to perform certain tasks like;

    - generate a note id for newly created note

    - validate note creation input

    - save a new note to file

    - creating a new note

    - viewing all notes for a logged in user

    - delete a note by id for a logged in user

- `schema/model.go`:

  - Contains the user and note type/model

## 💾 Data Persistence

All your code users and notes are stored locally in a users.json and notes.json file respoectively which are located in the db/ directory within your project.

- When the CLI starts, users.json and notes.json is loaded into memory.

- Any add or delete operations modify the in-memory data.

- After a modification, the function to save is called to write the updated data back to respective file, ensuring your changes are persistent.

## 🤝 Contributing

Contributions are welcome! If you'd like to contribute, please:

1. Fork the repository.

2. Create a new branch (git checkout -b feature/your-feature-name).

3. Make your changes.

4. Commit your changes (git commit -m 'feat: Add new feature').

5. Push to the branch (git push origin feature/your-feature-name).

6. Open a Pull Request.
