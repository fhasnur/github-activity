# GitHub Activity CLI

A command-line interface (CLI) tool to display a user's GitHub activity, including pushes, issues, stars, forks, and more. The sample solution for the [github-user-activity](https://roadmap.sh/projects/github-user-activity) challenge from [roadmap.sh](https://roadmap.sh/)

## Installation

To use the GitHub Activity CLI, you need to have Go installed on your machine.

1. **Clone the Repository:**
```bash
git clone https://github.com/fhasnur/github-activity.git
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
Github User: fhasnur
Name: Fandi Hasnur
Bio: my codes are funday space 🌞
Public Repos: 52 | Followers: 46

Recent Activities:
-------------------------------------------------
- Pushed 1 commit(s) to fhasnur/github-activity
 (at 2024-09-03T16:52:59Z)
 -------------------------------------------------
- Created branch in fhasnur/github-activity 
 (at 2024-08-24T01:15:41Z)
-------------------------------------------------
- Created repository in fhasnur/github-activity 
 (at 2024-08-24T01:09:37Z)
```