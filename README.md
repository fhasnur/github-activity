# GitHub Activity CLI

A command-line interface (CLI) tool to display a user's GitHub activity, including pushes, issues, stars, forks, and more. The sample solution for the [github-user-activity](https://roadmap.sh/projects/github-user-activity) challenge from [roadmap.sh](https://roadmap.sh/)

## Installation

To use the GitHub Activity CLI, you need to have Go installed on your machine.

1. **Clone the Repository:**
```bash
git clone https://github.com/yourusername/github-activity.git
```

2. **Navigate to the project directory:**
```bash
cd github-activity
```

3. **Build the CLI:**
```bash
go build -o github-activity
```

## Usage

Once built, you can run the CLI tool from your terminal. The basic usage is:
```bash
./github-activity <username>

```

Replace <username> with the GitHub username you want to query. For example:
```bash
./github-activity fhasnur
``` 

### Example Output

```bash
-  Created branch in fhasnur/github-activity
-  Created repository in fhasnur/github-activity
-  Pushed 1 commit(s) to fhasnur/fhasnur
```