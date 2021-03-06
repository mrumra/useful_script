package task

import (
	"condition"
	"crypto/sha1"
	"encoding/hex"
	"log"
	"net/url"
	"routine"
	"rut"
	"taskresult"
)

type Task struct {
	Name          string               `json:"name"`
	PreCondition  *condition.Condition `json:"precondition"`
	Routine       *routine.Routine     `json:"routine"`
	PostCondition *condition.Condition `json:"postcondition"`
	Clear         *routine.Routine     `json:"clear"`
	Description   string               `json:"description"`
	ID            string               `json:"id"`
}

func Hash(name []byte) []byte {
	hash := sha1.New()
	return []byte(hex.EncodeToString(hash.Sum([]byte("taskTASK" + string(name)))))
}

func (t *Task) Run(db *rut.DB) *taskresult.Result {
	if res := t.CheckPreCondition(db); !res.Success {
		return res
	}

	if res := t.RunMainRoutine(db); !res.Success {
		return res
	}

	if res := t.CheckPostCondition(db); !res.Success {
		return res
	}

	defer t.RunClearRoutine(db)

	return &taskresult.Result{Name: t.Name, Description: t.Description, Success: true}
}

func (t *Task) CheckPreCondition(db *rut.DB) *taskresult.Result {
	for _, r := range db.DB {
		r.GoInitMode()
	}

	if err := t.PreCondition.Check(db); err != nil {
		return &taskresult.Result{
			Name:        t.Name,
			Description: t.Description,
			Success:     false,
			Message:     err.Error(),
		}
	}

	return &taskresult.Result{
		Name:        t.Name,
		Description: t.Description,
		Success:     true,
	}
}

func (t *Task) CheckPostCondition(db *rut.DB) *taskresult.Result {
	for _, r := range db.DB {
		r.GoInitMode()
	}

	if err := t.PostCondition.Check(db); err != nil {
		return &taskresult.Result{
			Name:        t.Name,
			Description: t.Description,
			Success:     false,
			Message:     err.Error(),
		}
	}

	return &taskresult.Result{
		Name:        t.Name,
		Description: t.Description,
		Success:     true,
	}

}

func (t *Task) RunMainRoutine(db *rut.DB) *taskresult.Result {
	for _, r := range db.DB {
		r.GoInitMode()
	}

	if err := t.Routine.Run(db); err != nil {
		return &taskresult.Result{
			Name:        t.Name,
			Description: t.Description,
			Success:     false,
			Message:     err.Error(),
		}
	}

	return &taskresult.Result{
		Name:        t.Name,
		Description: t.Description,
		Success:     true,
	}

}

func (t *Task) RunClearRoutine(db *rut.DB) *taskresult.Result {
	for _, r := range db.DB {
		r.GoInitMode()
	}
	if err := t.Clear.Run(db); err != nil {
		return &taskresult.Result{
			Name:        t.Name,
			Description: t.Description,
			Success:     false,
			Message:     err.Error(),
		}
	}

	return &taskresult.Result{
		Name:        t.Name,
		Description: t.Description,
		Success:     true,
	}

}

func (t *Task) SetPreCondition(con *condition.Condition) {
	t.PreCondition = con
}

func (t *Task) SetPostCondition(con *condition.Condition) {
	t.PostCondition = con
}

func (t *Task) SetMainRoutine(r *routine.Routine) {
	t.Routine = r
}

func (t *Task) SetClearRoutine(r *routine.Routine) {
	t.Clear = r
}

func (t *Task) SetName(name string) {
	t.Name = name
}

func (t *Task) SetDescription(desc string) {
	t.Description = desc
}

func IsTaskParamsValid(in url.Values) bool {
	for k, v := range in {
		log.Println(k, "----------->", v, len(v))
		if len(v) == 0 {
			return false
		}
	}
	return true
}
