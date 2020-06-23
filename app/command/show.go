package command

import (
	model "infoqerja-line/app/model"
	util "infoqerja-line/app/utils"
	constant "infoqerja-line/app/utils/constant"
	"log"
	"time"

	"github.com/line/line-bot-sdk-go/linebot"
	"go.mongodb.org/mongo-driver/bson"
)

// Show : A class to represent the show job command
type Show struct{}

// GetReply : Method service for IncomingHelp instance
func (handler *Show) GetReply() []linebot.SendingMessage {

	// template : carousel
	jobs, err := handler.GetData()
	if err != nil {
		return []linebot.SendingMessage{linebot.NewTextMessage(constant.ShowMessageFail)}
	}

	var holder []*linebot.CarouselColumn
	for _, job := range jobs {
		holder = append(holder, linebot.NewCarouselColumn(
			constant.ResumeImageURL, job.Title, job.Deadline.Format(constant.DateFormatLayout),
			linebot.NewPostbackAction("View Details", constant.JobIDData, job.ID.String(), ""),
		))
	}

	return []linebot.SendingMessage{linebot.NewTemplateMessage(constant.UnavailableMessage, linebot.NewCarouselTemplate(holder...))}
}

// GetData : Get the data necessary for jobs
func (handler *Show) GetData() ([]model.Job, error) {
	jobs, err := (&util.JobReader{}).ReadFiltered(bson.M{
		constant.IsComplete: true,
		constant.Deadline: bson.M{
			"$gt": time.Now(),
		},
	})

	if err != nil {
		log.Print(err)
		return nil, err
	}
	return jobs, nil
}

// GetState : Method to get any state a certain command produce, if present
func (handler *Show) GetState() (model.State, error) {
	return nil, nil
}
