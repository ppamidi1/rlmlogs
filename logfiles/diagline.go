package logfiles

const (
	AM = iota
	AS
	CS
	SS
	VS
	AB
	IP
	FI
	FO
	FD
	CPU_Usage
	Up_Time
	Pct_Mem_Used_sys
	Mem_Used_sys
	Mem_Free_sys
	Mem_Used_rlc
	Processes
)

var itemNames = [...]string{
	"AM",
	"AS",
	"CS",
	"SS",
	"VS",
	"AB",
	"IP",
	"FI",
	"FO",
	"FD",
	"CPU_Usage",
	"Up_Time",
	"Pct_Mem_Used_sys",
	"Mem_Used_sys",
	"Mem_Free_sys",
	"Mem_Used_rlc",
	"Processes"}

var itemStatusLine = [...]string{
	"AM=",
	"AS=",
	"CS=",
	"SS=",
	"VS=",
	"AB=",
	"IP=",
	"FI=",
	"FO=",
	"FD=",
	"CPU Usage:",
	"Up Time:",
	"% Mem Used(sys):",
	"Mem Used(sys):",
	"Mem Free(sys):",
	"Mem Used(rlc):",
	"Processes:"}

/*
	Jun 20 23:29:39.289159 RL00122 rlc[694]: RLM RL00122 Connection Status
	                                         AM=2
	                                         AS=STATE_started
	                                         CS=STATE_connected
	                                         SS=STATE_established
	                                         VS=STATE_high_quality
	                                         AB=Ethernet
	                                         IP=192.168.128.37
	                                         FI=146308
	                                         FO=146302
	                                         FD=6
	                                         CPU Usage: 55%
	                                         Up Time: 14496
	                                         % Mem Used(sys):39.07781
	                                         Mem Used(sys):  411009024
	                                         Mem Free(sys):  640761856
	                                         Mem Used(rlc):  48696
	                                         Processes: 130  */
