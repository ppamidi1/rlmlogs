package logfiles

import (
	"fmt"
	"regexp"
)

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
	"AM",   /* Application Mode                 */
	"AS",   /* Application Current State        */
	"CS",   /* Connection Manager Current State */
	"SS",   /* Session Manager Current State    */
	"VS",   /* Video Streamer Current State     */
	"AB",   /* Active Bearer                    */
	"IP",   /* Current IP Address               */
	"FI",   /* Input Frames                     */
	"FO",   /* Output Frames                    */
	"FD",   /* Difference of above              */
	"CPU_Usage", 
	"Up_Time",
	"Pct_Mem_Used_sys",
	"Mem_Used_sys",
	"Mem_Free_sys",
	"Mem_Used_rlc",
	"Processes"}
var itemDescriptions = [...]string{
		"Application Mode" ,
		"Application Current State" ,
		"Connection Manager Current State" ,
		"Session Manager Current State" ,
		"Video Streamer Current State" ,
		"Active Bearer" ,
		"Current IP Address" ,
		"Input Frames" ,
		"Output Frames" ,
		"Difference of above" ,
		"CPU_Usage" , 
		"Up_Time" ,
		"Pct_Mem_Used_sys",
		"Mem_Used_sys",
		"Mem_Free_sys",
		"Mem_Used_rlc",
		"Processes"}

		/*
	       "AM=%d\n"
        "AS=%s\n"
        "CS=%s\n"
        "SS=%s\n"
        "VS=%s\n"
        "AB=%s\n"
        "IP=%s\n"
        "FI=%llu\n"
        "FO=%llu\n"
        "FD=%llu\n"
        "CPU Usage: %lu%%\n"
        "Up Time: %-10ld\n"
        "%% Mem Used(sys):%-5.05f\n"
        "Mem Used(sys):  %-10ld\n"
        "Mem Free(sys):  %-10ld\n"
        "Mem Used(rlc):  %-10ld\n"
		"Processes: %-10d\n",
		
        app_get_rlm_serial(),
        app_cm_mode(),
		app_fsm_get_cur_state(),
		
        cm_fsm_get_cur_state(),
        sm_fsm_get_cur_state(),
		vs_fsm_get_cur_state(),
		
        cm_fsm_get_active_bearer_name(),
        cm_get_cur_ip_address(),
        vs_fsm_get_input_frames(),
        vs_fsm_get_output_frames(),
        vs_fsm_get_input_frames() - vs_fsm_get_output_frames(),
        g_rlc_stats.sys_cpu_usage,
        si.uptime,
        mem_usage_percent,
        si.totalram - si.freeram,
        si.freeram,
        ru.ru_maxrss,
        si.procs
	*/

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
var numPlots int

func ValidItem(nm string) bool {
	for _, opt := range itemNames {
		if nm == opt {
			return true
		}
	}
	return false
}
func ShowValidItems() {
	for idx, opt := range itemNames {
		fmt.Printf("%s \t\t\t: %s\n", opt , itemDescriptions[idx])
	}
}
func Index(nm string) int {
	for i, opt := range itemNames {
		if nm == opt {
			return i
		}
	}
	return -1
}

func SetupStats(nm string) {
	if ValidItem(nm) {
		numPlots++
		idx := Index(nm)
		gatheredStats[idx] = New(nm)
		switch idx {
		case FI:
			gatheredStats[FI].Extractor = regexp.MustCompile("([0-9]+)")
		case FO:
			gatheredStats[FO].Extractor = regexp.MustCompile("([0-9]+)")
		case FD:
			gatheredStats[FD].Extractor = regexp.MustCompile("([0-9]+)")
		case CPU_Usage:
			gatheredStats[CPU_Usage].Extractor = regexp.MustCompile("([0-9]+)%")
		case Up_Time:
			gatheredStats[Up_Time].Extractor = regexp.MustCompile("([0-9]+)")
		case Pct_Mem_Used_sys:
			gatheredStats[Pct_Mem_Used_sys].Extractor = regexp.MustCompile("([0-9]+\\.[0-9]+)")
		case Mem_Used_sys:
			gatheredStats[Mem_Used_sys].Extractor = regexp.MustCompile("([0-9]+)")
		case Mem_Free_sys:
			gatheredStats[Mem_Free_sys].Extractor = regexp.MustCompile("([0-9]+)")
		case Mem_Used_rlc:
			gatheredStats[Mem_Used_rlc].Extractor = regexp.MustCompile("([0-9]+)")
		case Processes:
			gatheredStats[Processes].Extractor = regexp.MustCompile("([0-9]+)")
		}
	} else {
		if nm == "CPUTemp" {
			cpuTempStats = New("CPUTemp")
			numPlots++
		}
	}
}
