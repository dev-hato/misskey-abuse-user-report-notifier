package main

import (
	"context"
	"fmt"
	"strconv"

	"github.com/bwmarrin/discordgo"
	"github.com/dev-hato/misskey-abuse-user-report-notifier/ent"
	"github.com/dev-hato/misskey-abuse-user-report-notifier/ent/userreport"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/yitsushi/go-misskey"
	"github.com/yitsushi/go-misskey/services/admin/moderation"

	"os"
	"slices"
)

func closeDBClient(db *ent.Client) {
	if err := db.Close(); err != nil {
		logrus.Fatal(err.Error())
	}
}

func main() {
	dbClient, err := ent.Open("postgres", fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s TimeZone=%s",
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB"),
		os.Getenv("POSTGRES_TIMEZONE"),
	))
	if err != nil {
		logrus.Fatal(err.Error())
	}
	defer closeDBClient(dbClient)

	ctx := context.Background()

	if err = dbClient.Schema.Create(ctx); err != nil {
		logrus.Fatal(err.Error())
	}

	serverHost := os.Getenv("MISSKEY_HOST")

	misskeyClient, err := misskey.NewClientWithOptions(
		misskey.WithAPIToken(os.Getenv("MISSKEY_TOKEN")),
		misskey.WithBaseURL("https", serverHost, ""),
		misskey.WithLogLevel(logrus.DebugLevel),
	)
	if err != nil {
		logrus.Fatal(err.Error())
	}

	userReportsLimit, err := strconv.ParseUint(os.Getenv("MISSKEY_USER_REPORTS_LIMIT"), 10, 64)
	if err != nil {
		logrus.Fatal(err.Error())
	}

	userReportsRequest := moderation.UserReportsRequest{Limit: uint(userReportsLimit)}

	latestUserReport, err := dbClient.UserReport.Query().
		Order(ent.Desc(userreport.FieldCreatedAt), ent.Desc(userreport.FieldID)).
		First(ctx)
	if err == nil {
		userReportsRequest.SinceID = latestUserReport.ID
	} else if err != nil && !ent.IsNotFound(err) {
		logrus.Fatal(err.Error())
	}

	discordSession, err := discordgo.New("")
	if err != nil {
		logrus.Fatal(err.Error())
	}

	reports, err := misskeyClient.Admin().Moderation().UserReports(userReportsRequest)
	if err != nil {
		logrus.Fatal(err.Error())
	}

	slices.Reverse(reports)

	for _, report := range reports {
		hostName := serverHost

		if report.TargetUser.Host != nil {
			hostName = *report.TargetUser.Host
		}

		_, err = discordSession.WebhookExecute(
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
			logrus.Fatal(err.Error())
		}

		if _, err := dbClient.UserReport.Create().SetID(report.ID).Save(ctx); err != nil {
			logrus.Fatal(err.Error())
		}
	}
}
