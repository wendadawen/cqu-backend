package bo

type ClassScheduleBo = []ClassInfo

type ClassInfo struct {
	Title      string `json:"title"`
	Id         string `json:"id"`
	TeachClass string `json:"class"` // 教学班
	Day        int    `json:"day"`
	Weeks      []int  `json:"weeks"`
	Room       string `json:"room"`
	Start      int    `json:"start"`
	Teacher    string `json:"teacher"`
	Num        int    `json:"num"`
	More       string `json:"more"`
	Content    string `json:"content,omitempty"`
	CampusId   string `json:"campusId"`
	Meeting    string `json:"meeting"`
}
