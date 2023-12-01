# Project Requirements

This project is built using Go, so ensure that you have **Go installed** on your machine. Additionally, you'll need to have **Graphviz** installed. If you haven't installed Graphviz yet, you can download it from [here](https://graphviz.org/download/). After installation, make sure to add the Graphviz bin folder to your **PATH** environment variables.

To run the project, execute the following commands in your terminal:

```bash
go run ./main.go
dot -Tsvg -O sp-graph.gv
```
## Bonus Tip
While executing **go run ./main.go** , if you encounter versioning issues, <br>you can resolve them by removing the version from the go.mod file.<br> If you encounter any problems during the process, feel free to open an issue in the repository for assistance.
