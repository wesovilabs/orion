
trait print debugger {
    prefix = "DEBUG"
    showtimestamp = true
}

scenario "demo of traits" {
    when "prints a message with" {
        debugger {
            msg = "hello my friend"
        }
    }
    
}