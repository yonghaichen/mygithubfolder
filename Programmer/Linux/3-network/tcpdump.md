
`tcpdump -i any -vv -n -s 0`

`tcpdump -i eth1 -t -s 0 and dst port ! 22 and src net 192.168.1.0/24`
	

	
`-i eth0` Listen on the eth0 interface. Or get all interfaces with `-i any`

`port 25001`  To filter packets based on the desired service or port, use the `port` filter.

`-s`  Define the snaplength (size) of the capture in bytes. Use `-s0` to get everything, unless you are intentionally capturing less.

`-t`      Give human-readable timestamp output.

`-c 100`  To limit the number of packets captured and stop tcpdump, use the `-c` (for count) option

`-X`      to print content in hex, and ASCII or `-A` to print the content in ASCII.

`-w filename.pcap`  To save packets to a file instead of displaying them on screen

`-vv`  Verbose output (more v’s gives more output).

`-n` Don't convert addresses (i.e., host addresses, port numbers, etc.) to names.