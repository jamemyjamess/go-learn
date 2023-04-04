package utils

func MapStatusText(status string) string {
	switch status {
	case "2_1":
		return "รออนุมัติ"
	case "2_2":
		return "ผ่านอนุมัติ"
	case "2_3":
		return "แก้ไข"
	case "2_4":
		return "ไม่ผ่านอนุมัติ"
	case "2_5":
		return "แบบร่าง"
	default:
		return status
	}
}
