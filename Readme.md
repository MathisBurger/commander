<div align="center">
<h1>Commander</h1>
<hr>
<strong>
    A lightweight command handler for discordgo
</strong>
<br>
<hr>
<img src="https://upload.wikimedia.org/wikipedia/commons/thumb/0/05/Go_Logo_Blue.svg/1200px-Go_Logo_Blue.svg.png" height="100">
</div>


# Information

---
The Commander is a lightweight command handler build for discordgo. You can use it
to control the command execution of your commands.
It is easy to setup and use. Test it out yourself!


# Documentation

Install the library with:
```
go get -t github.com/MathisBurger/commander
```

Example:
```go
// Init command handler
handler := commander.New(";;", "826378911444500501")

// add command to handler
handler.Register("example", "Example command", ExampleCommand, 100)

// add handler to discord session
session.AddHandler(handler.Execute)
```

<strong>Note:</strong> The permission system is not working in v0.0.1


# Examples

---
- <a href="https://github.com/MathisBurger/commander/blob/master/examples/default.go">Default example</a>

