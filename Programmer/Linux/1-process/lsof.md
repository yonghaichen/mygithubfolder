
The **lsof** command stands for **LiSt Open Files** and provides a list of files that are opened. 

Basically, it gives the information to find out the files which are opened by which process. 

With one go it lists out all open files in output console.

Here, you observe there are details of files which are opened. 





Let's say you suspect a file is causing a problem, such as being locked or preventing unmounting. You can use lsof to identify the processes that have that file open:    
`lsof /path/to/file`   




If you want to focus on a specific process, you can filter the output of lsof to display only the files opened by that process. Replace PID in the following command with the actual process ID:   
`lsof -p PID`  



