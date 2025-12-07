# The netstat Command

netstat is a useful tool for checking your network configuration and activity. It is in fact a collection of several tools lumped together. We discuss each of its functions in the following sections.



## Displaying the Routing Table

When you invoke **netstat** with the `–r` flag, it displays the kernel routing table in the way we've been doing with **route**. 

    netstat -nr

The `–n` option makes **netstat** print addresses as dotted quad IP numbers rather than the symbolic host and network names. This option is especially useful when you want to avoid address lookups over the network (e.g., to a DNS or NIS server).




## Displaying Interface Statistics

When invoked with the `–i` flag, **netstat** displays statistics for the network interfaces currently configured. If the `–a` option is also given, it prints all interfaces present in the kernel, not only those that have been configured currently.




## Displaying Connections

**netstat** supports a set of options to display active or passive sockets. The options `–t`, `–u`, `–w`, and `–x` show active TCP, UDP, RAW, or Unix socket connections. 

If you provide the `–a` flag in addition, sockets that are waiting for a connection (i.e., listening) are displayed as well. This display will give you a list of all servers that are currently running on your system.



`netstat -ltunp | grep portOrName`  

