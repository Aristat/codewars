/*

You're analyzing authentication logs. Each log entry is a string like:

192.168.1.1 LOGIN_FAIL user=admin
192.168.1.1 LOGIN_SUCCESS user=admin
10.0.0.5 LOGIN_FAIL user=root

An IP is suspicious if it has 3 or more consecutive failures without a success in between. Return a list of suspicious IPs, sorted alphabetically.

logs = [
    "192.168.1.1 LOGIN_FAIL user=admin",
    "192.168.1.1 LOGIN_FAIL user=admin",
    "192.168.1.1 LOGIN_FAIL user=root",
    "10.0.0.5 LOGIN_FAIL user=test",
    "10.0.0.5 LOGIN_SUCCESS user=test"
]
detect_brute_force(logs)  # ["192.168.1.1"]

The 10.0.0.5 IP had a failure then a success, so its streak reset. The 192.168.1.1 IP hit 3 failures in a row - busted. Only respond with a list of the suspicious IPs.

A success resets that IP's failure count to zero. Empty list returns empty list.

PS. You do not need to validate the IP addresses.

*/

package kata

import (
	"sort"
	"strings"
)

type Counter struct {
	consecutive int
	total       int
}

func DetectBruteForce(logs []string) []string {
	counters := make(map[string]*Counter)
	for _, log := range logs {
		parts := strings.Split(log, " ")
		if counters[parts[0]] == nil {
			counters[parts[0]] = &Counter{}
		}
		if parts[1] == "LOGIN_SUCCESS" {
			counters[parts[0]].consecutive = 0
		} else {
			counters[parts[0]].consecutive++
			if counters[parts[0]].consecutive >= 3 {
				counters[parts[0]].total++
				counters[parts[0]].consecutive = 0
			}
		}
	}

	var keys []string
	for k, v := range counters {
		if v.total > 0 {
			keys = append(keys, k)
		}
	}

	if len(keys) == 0 {
		return []string{}
	} else {
		sort.Strings(keys)
		return keys
	}
}
