package main

import (
	"errors"
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/sirupsen/logrus"
	"github.com/yitsushi/go-misskey"
	"github.com/yitsushi/go-misskey/services/admin/moderation"
	"gorm.io/driver/postgres"

	"os"
	"slices"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type UserReport struct {
	gorm.Model
	ID string `gorm:"primaryKey"`
}

func main() {
	db, err := gorm.Open(postgres.Open(fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s TimeZone=%s",
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB"),
		os.Getenv("POSTGRES_TIMEZONE"),
	)))
	if err != nil {
		logrus.Error(err.Error())
		return
	}

	migrator := db.Migrator()

	if !migrator.HasTable(&UserReport{}) {
		if err = migrator.CreateTable(&UserReport{}); err != nil {
			logrus.Error(err.Error())
			return
		}
	}

	serverHost := os.Getenv("MISSKEY_HOST")

	client, err := misskey.NewClientWithOptions(
		misskey.WithAPIToken(os.Getenv("MISSKEY_TOKEN")),
		misskey.WithBaseURL("https", serverHost, ""),
		misskey.WithLogLevel(logrus.DebugLevel),
	)
	if err != nil {
		logrus.Error(err.Error())
		return
	}

	userReportsRequest := moderation.UserReportsRequest{Limit: 2}
	var latestUserReport UserReport

	err = db.
		Order(clause.OrderByColumn{Column: clause.Column{Name: "updated_at"}, Desc: true}).
		Order(clause.OrderByColumn{Column: clause.Column{Name: "id"}, Desc: true}).
		First(&latestUserReport).Error
	if err == nil {
		userReportsRequest.SinceID = latestUserReport.ID
	} else if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		logrus.Error(err.Error())
		return
	}

	discord, err := discordgo.New("")
	if err != nil {
		logrus.Error(err.Error())
		return
	}

	reports, err := client.Admin().Moderation().UserReports(userReportsRequest)
	if err != nil {
		logrus.Error(err.Error())
		return
	}

	slices.Reverse(reports)

	for _, report := range reports {
		hostName := serverHost

		if report.TargetUser.Host != nil {
			hostName = *report.TargetUser.Host
		}

		_, err = discord.WebhookExecute(
			os.Getenv("DISCORD_WEBHOOK_ID"),
			os.Getenv("DISCORD_WEBHOOK_TOKEN"),
			true,
			&discordgo.WebhookParams{Content: fmt.Sprintf(
				"<通報ID>\n"+
					"%s\n"+
					"<対象ユーザー>\n"+
					"https://%s/@%s\n"+
					"<通報内容>\n"+
					"%s",
				report.ID,
				hostName,
				report.TargetUser.Username,
				report.Comment,
			)},
		)
		if err != nil {
			logrus.Error(err.Error())
			return
		}

		if err := db.Create(&UserReport{ID: report.ID}).Error; err != nil {
			logrus.Error(err.Error())
			return
		}
	}
}
