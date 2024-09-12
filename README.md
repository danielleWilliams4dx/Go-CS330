# History
Go was created in 2007 by Robert Griesemer, Rob Pike, and Ken Thompson at Google and was launched as an open-source programming language in 2009 (Pike, 2012; Yadav, 2024). According to Rob Pike, Go was written to address the various problems in software infrastructure development that arose from the introduction of networked systems, large-scale computing clusters, multicore processors, and the web-based programming framework (Pike, 2012). At the time, languages such as C++, Java, and Python were being used to circumvent rather than directly address these problems. Go was designed to make the software development process more productive and scalable. Features that make Go desirable to programmers are its clear dependencies, syntax, and semantics; its prioritization of composition over inheritance; its garbage collection; its concurrency; and its simple tooling. Go is primarily used for creating web applications, command-line tools, and scalable network services (Yadav, 2024).

[Pike, R. (2012, October 25). Go at Google: Language Design in the Service of Software Engineering. The Go Programming Language. Retrieved September 11, 2024.](https://go.dev/talks/2012/splash.article)

[Yadav, A. (2024, August 23). Go Programming Language (Introduction). Geeks for Geeks. Retrieved September 11, 2024.](https://www.geeksforgeeks.org/go-programming-language-introduction/)

Information about this language can be found in documentation on [The Go Programming Language website](https://go.dev/doc/). This website also provides Go tutorials as does the [W3Schools website](https://www.w3schools.com/go/index.php).

# Getting Started
## Installation
1. Install Go [here](https://go.dev/dl/) and download the proper version for your computer. A step-by-step guide can be found [here](https://go.dev/doc/install).
2. Any text editor can be used to run Go. VSCode is popular and can be installed [here](https://code.visualstudio.com/Download).
3. Download the Go extension in your text editor. It can be downloaded for VSCode [here](https://marketplace.visualstudio.com/items?itemName=golang.go).
   
## Module Set-up and Running Code
1. A helpful video can be found [here](https://www.youtube.com/watch?v=1MXIGYrMk80).
2. Type **go mod init** in a terminal window followed by the name of your code’s module to enable dependency tracking. This command creates a go.mod file which tracks the modules of all imported packages in your code.
3. Create a .go file in your new module.
4. Write your program. See [hello_example](https://github.com/danielleWilliams4dx/Go-CS330/tree/main/hello_example) for a basic example.
5. To run your file, **cd** to your new module and then type **go run fileName.go**.
   
## Comments
**Comments in Go are the same as in languages like Java and JavaScript.**
-  // - Single line comments
-  /*…*/  - Multi line comments 
