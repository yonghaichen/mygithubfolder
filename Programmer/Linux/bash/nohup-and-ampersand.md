
# &

The ampersand, `&`, is used to run a command in the background of a shell. 
It’s useful in situations where you’re starting a long running process in a shell and don’t need to follow any output. 
The problem most people run into with this is that they run a command using `&`, log out of the shell, and then their process is all of a sudden gone. 
This is because running a command with the `&` doesn’t prevent hangup, it just hides the process in the background. 
When you do something like `./start_server.sh &` and then log out of the shell, that command is killed as a sub-command of the current shell session.




# nohup

`nohup` on the other hand, is used to run a command without hanging up at the end of the shell session. 
It’s more ideal for those who want to do something like start a server and then leave. 
It’s not foolproof, however, some processes may reconnect the hangup signal and wind up killing the process later on. 
Think of `nohup` as a temporary fix in a rush as the **pid** can still be re-added to the shell’s job list.