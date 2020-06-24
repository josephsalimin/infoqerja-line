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

	jobs, err := handler.GetData()
	if err != nil {
		return []linebot.SendingMessage{linebot.NewTextMessage(constant.ShowMessageFail)}
	}

	var placeholder []*linebot.BubbleContainer
	for _, job := range jobs {
		placeholder = append(placeholder, getTemplate(job))
	}
	contents := &linebot.CarouselContainer{
		Type:     linebot.FlexContainerTypeCarousel,
		Contents: placeholder,
	}
	return []linebot.SendingMessage{linebot.NewFlexMessage(constant.InvalidMessage, contents)}
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

func getTemplate(job model.Job) *linebot.BubbleContainer {
	return &linebot.BubbleContainer{
		Size: "kilo",
		Type: linebot.FlexContainerTypeBubble,
		Body: &linebot.BoxComponent{
			Type:   linebot.FlexComponentTypeBox,
			Layout: linebot.FlexBoxLayoutTypeVertical,
			Contents: []linebot.FlexComponent{
				&linebot.TextComponent{
					Type:  linebot.FlexComponentTypeText,
					Text:  "InfoQerja",
					Color: "#1DB446",
					Size:  "md",
				},
				&linebot.TextComponent{
					Type:   linebot.FlexComponentTypeText,
					Text:   job.Title,
					Weight: "bold",
					Size:   "xxl",
					Margin: "md",
				},
				&linebot.BoxComponent{
					Type:   linebot.FlexComponentTypeBox,
					Layout: linebot.FlexBoxLayoutTypeHorizontal,
					Contents: []linebot.FlexComponent{
						&linebot.TextComponent{
							Type:  linebot.FlexComponentTypeText,
							Text:  "Job ID : ",
							Size:  "xs",
							Color: "#aaaaaa",
						},
						&linebot.TextComponent{
							Type:  linebot.FlexComponentTypeText,
							Text:  job.ID.Hex(),
							Size:  "xs",
							Color: "#aaaaaa",
							Align: "end",
						},
					},
				},
				&linebot.BoxComponent{
					Type:    linebot.FlexComponentTypeBox,
					Layout:  linebot.FlexBoxLayoutTypeVertical,
					Margin:  "xxl",
					Spacing: "sm",
					Contents: []linebot.FlexComponent{
						&linebot.BoxComponent{
							Type:   linebot.FlexComponentTypeBox,
							Layout: linebot.FlexBoxLayoutTypeHorizontal,
							Contents: []linebot.FlexComponent{
								&linebot.TextComponent{
									Type:  linebot.FlexComponentTypeText,
									Text:  "Deadline Date : ",
									Size:  "sm",
									Color: "#555555",
								},
								&linebot.TextComponent{
									Type:  linebot.FlexComponentTypeText,
									Text:  job.Deadline.Format(constant.DateFormatLayout),
									Size:  "sm",
									Color: "#555555",
									Align: "end",
								},
							},
						},
						&linebot.SeparatorComponent{
							Type:   linebot.FlexComponentTypeSeparator,
							Margin: "xxl",
						},
					},
				},
				&linebot.ButtonComponent{
					Type:   linebot.FlexComponentTypeButton,
					Action: linebot.NewPostbackAction("View Details", constant.JobIDData+"|"+job.ID.Hex(), "", ""),
				},
			},
		},
	}
}
