package todoist

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/bringdesk/bringdesk/evt"
	"github.com/bringdesk/bringdesk/widgets"
	"io"
	"io/ioutil"
	"log"
	"net/http"
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
	apiToken string /* Todoist API token                 */
	tasks    []*Task
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
	if self.apiToken != "" {
		// curl -X GET  -H "Authorization: "

		/* Step 1. Download response */
		client := http.Client{
			Timeout: 15 * time.Second,
		}
		req, _ := http.NewRequest("GET", "https://api.todoist.com/rest/v1/tasks", nil)
		newAuthorization := fmt.Sprintf("Bearer %s", self.apiToken)
		req.Header.Add("Authorization", newAuthorization)

		resp, err := client.Do(req)
		if err != nil {
			log.Printf("Problem with Gismeto result: err = %#v", err)
			return
		}
		defer resp.Body.Close()
		var out bytes.Buffer
		io.Copy(&out, resp.Body)
		log.Printf("out = %s", out.String())

		/* Step 2. Processing TASK */
		var tasks []TodoistTask
		err3 := json.Unmarshal(out.Bytes(), &tasks)
		if err3 != nil {
			log.Printf("Unable parser Task issues")
		}
		/* Step 3. Convert issue to local struct */
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
