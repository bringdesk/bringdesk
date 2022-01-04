package todoist

import (
	"encoding/json"
	"fmt"
	"github.com/bringdesk/bringdesk/ctx"
	"github.com/bringdesk/bringdesk/evt"
	"github.com/bringdesk/bringdesk/widgets"
	"io/ioutil"
	"log"
	"os"
	"path"
	"strings"
	"time"
)

type TodoistTask struct {
	Id          int64  `json:"id"`
	Assigner    int    `json:"assigner"`
	ProjectId   int    `json:"project_id"`
	SectionId   int    `json:"section_id"`
	Order       int    `json:"order"`
	Content     string `json:"content"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
	//	LabelIds     []interface{} `json:"label_ids"`
	Priority     int         `json:"priority"`
	CommentCount int         `json:"comment_count"`
	Creator      interface{} `json:"creator"`
	Created      time.Time   `json:"created"`
	Due          struct {
		Date      string `json:"date"`
		String    string `json:"string"`
		Lang      string `json:"lang"`
		Recurring bool   `json:"recurring"`
	} `json:"due"`
}

type Task struct {
	content string
}

type TodoistWidget struct {
	widgets.BaseWidget
	apiToken string  /* Todoist API token                 */
	tasks    []*Task /* Output tasks                      */
}

func NewTodoistWidget() *TodoistWidget {
	newTodoistWidget := new(TodoistWidget)
	newTodoistWidget.recoverToken()
	go func() {
		for {
			newTodoistWidget.updateData()
			time.Sleep(5 * time.Minute)
		}
	}()
	return newTodoistWidget
}

func (self *TodoistWidget) recoverToken() {

	/* Step 0. Prepare reding user home directory */
	userDirName, err := os.UserHomeDir()
	if err != nil {
		log.Printf("Todoist error API token reading: err = %#v", err)
	}

	/* Step 1. Prepare Todoist token path */
	newTokenPath := path.Join(userDirName, ".todoist")
	log.Printf("Todoist token path: %#v", newTokenPath)

	/* Step 2. Reading content with token */
	content, err := ioutil.ReadFile(newTokenPath)
	if err != nil {
		log.Printf("Todoist error API token reading: err = %#v", err)
	}
	self.apiToken = strings.Trim(string(content), " \r\n\t")

}

func (self *TodoistWidget) updateData() {
	log.Printf("Todoist update start: apiToekn = %#v", self.apiToken)
	if self.apiToken == "" {
		log.Printf("No API token. Skip update.")
		return
	}
	mainNetworkManager := ctx.GetNetworkManager()

	/* Step 1. Create new request */
	req, err1 := mainNetworkManager.MakeRequest("TodoistWidget", "GET", "https://api.todoist.com/rest/v1/tasks", 15)
	if err1 != nil {
		log.Printf("err = %#v", err1)
		return
	}
	newAuthorization := fmt.Sprintf("Bearer %s", self.apiToken)
	req.AddHeader("Authorization", newAuthorization)

	/* Step 2. Perform request */
	resp, err2 := mainNetworkManager.Perform(req)
	if err2 != nil {
		log.Printf("err = %#v", err2)
		return
	}
	defer resp.Close()

	/* Step 2. Processing TASK */
	var tasks []TodoistTask
	out := resp.Bytes()
	err3 := json.Unmarshal(out, &tasks)
	if err3 != nil {
		log.Printf("Unable parser Task issues")
		return
	}

	/* Step 3. Convert issue to local struct */
	self.performTask(tasks)

}

func (self *TodoistWidget) ProcessEvent(e *evt.Event) {
}

func (self *TodoistWidget) Render() {
	self.BaseWidget.Render()

	/* Show task */
	for idx, t := range self.tasks {
		newText := widgets.NewTextWidget("", 21)
		newText.SetRect(self.X, self.Y+20*idx, self.Width, self.Height)
		newText.SetText(t.content)
		newText.Render()
		newText.Destroy()
	}

}

func (self *TodoistWidget) performTask(tasks []TodoistTask) {

	log.Printf("tasks = %#v", tasks)

	self.tasks = nil
	for _, t := range tasks {
		if t.Completed == false {
			newTask := &Task{
				content: t.Content,
			}
			self.tasks = append(self.tasks, newTask)
		}
	}
}
