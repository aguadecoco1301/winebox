/*
Copyright © 2026 ADRIEL ULLOA <adrielalejoulloa@gmail.com>
*/
package listing

import "fmt"

func HumanSize(size uint64) string {
	switch {
	case size >= 1<<30:
		return fmt.Sprintf("%.1fG", float64(size)/(1<<30))
	case size >= 1<<20:
		return fmt.Sprintf("%.1fM", float64(size)/(1<<20))
	case size >= 1<<10:
		return fmt.Sprintf("%.1fK", float64(size)/(1<<10))
	default:
		return fmt.Sprintf("%dB", size)
	}
}
