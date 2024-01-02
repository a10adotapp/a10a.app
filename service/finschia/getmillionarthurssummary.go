package finschia

import (
	"context"
	"database/sql"
	"fmt"
	"log/slog"
	"reflect"
	"strings"
	"time"
	"unicode/utf8"
	"unsafe"

	"github.com/a10adotapp/a10a.app/ent"
	"github.com/a10adotapp/a10a.app/lib/line/message"
)

const sqlStr = `
	SELECT
		properties.gear_rarity AS gear_rarity,
		SUM(
			CASE
				WHEN (activities.activated_at >= DATE_SUB(CURRENT_DATE, INTERVAL 1 DAY))
					AND (activities.activated_at < CURRENT_DATE)
				THEN 1
				ELSE 0
			END
		) AS yesterday,
		SUM(
			CASE
				WHEN (activities.activated_at >= DATE_ADD(DATE_SUB(CURRENT_DATE, INTERVAL 1 DAY), INTERVAL (1 - DAYOFWEEK(DATE_SUB(CURRENT_DATE, INTERVAL 1 DAY))) DAY))
					AND (activities.activated_at < CURRENT_DATE)
				THEN 1
				ELSE 0
			END
		) / (DATEDIFF(DATE_SUB(CURRENT_DATE, INTERVAL 1 DAY), DATE_ADD(DATE_SUB(CURRENT_DATE, INTERVAL 1 DAY), INTERVAL (1 - DAYOFWEEK(DATE_SUB(CURRENT_DATE, INTERVAL 1 DAY))) DAY)) + 1) AS week_avg,
		SUM(
			CASE
				WHEN (activities.activated_at >= DATE_FORMAT(DATE_SUB(CURRENT_DATE, INTERVAL 1 DAY), '%Y-%m-01'))
					AND (activities.activated_at < CURRENT_DATE)
				THEN 1
				ELSE 0
			END
		) / (DATEDIFF(DATE_SUB(CURRENT_DATE, INTERVAL 1 DAY), DATE_FORMAT(DATE_SUB(CURRENT_DATE, INTERVAL 1 DAY), '%Y-%m-01')) + 1) AS month_avg
	FROM
		finschia_item_token_activities AS activities
		JOIN
			finschia_item_token_million_arthurs_properties AS properties
			ON (properties.deleted_at IS NULL)
			AND (properties.finschia_item_token_id = activities.finschia_item_token_id)
	WHERE
		(activities.deleted_at IS NULL)
		AND (activities.activated_at >= DATE_FORMAT(DATE_SUB(CURRENT_DATE, INTERVAL 1 DAY), '%Y-%m-01'))
	GROUP BY
		properties.gear_rarity
	ORDER BY
		properties.gear_rarity ASC
`

type Summary struct {
	GearRarity string
	Yesterday  uint32
	WeekAvg    float32
	MonthAvg   float32
}

func (s FinschiaService) GetMillionArthursSummary(ctx context.Context) error {
	summaries, err := querySummary(s.entClient)
	if err != nil {
		return err
	}

	yesterday := time.Now().Add(-24 * time.Hour).In(time.FixedZone("Asia/Tokyo", 9*60*60))

	messageChunks := []string{
		"Million Arthurs' Summary",
		yesterday.Format("2006-01-02 (Mon)"),
	}

	for _, summary := range summaries {
		gearRarity := summary.GearRarity
		for i := 0; i < 5-(utf8.RuneCountInString(summary.GearRarity)); i++ {
			gearRarity += "☆"
		}

		messageChunks = append(messageChunks, "")
		messageChunks = append(messageChunks, gearRarity)
		messageChunks = append(messageChunks, fmt.Sprintf("[昨日計] %d", summary.Yesterday))
		messageChunks = append(messageChunks, fmt.Sprintf("[週平均] %.1f", summary.WeekAvg))
		messageChunks = append(messageChunks, fmt.Sprintf("[月平均] %.1f", summary.MonthAvg))
	}

	if err := message.Push(strings.Join(messageChunks, "\n")); err != nil {
		return err
	}

	return nil
}

func querySummary(entClient *ent.Client) ([]Summary, error) {
	execQuerierRef := reflect.ValueOf(entClient).Elem().
		FieldByName("config").
		FieldByName("driver").Elem().Elem().
		FieldByName("Conn").
		FieldByName("ExecQuerier")

	db := reflect.NewAt(execQuerierRef.Type(), unsafe.Pointer(execQuerierRef.UnsafeAddr())).Elem().Interface().(*sql.DB)

	rows, err := db.Query(sqlStr)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	summaries := []Summary{}

	for rows.Next() {
		summary := Summary{}

		if err := rows.Scan(&summary.GearRarity, &summary.Yesterday, &summary.WeekAvg, &summary.MonthAvg); err != nil {
			slog.Error("GetMillionArthursSummary", slog.String("step", "querySummary"), slog.Any("error", err))

			continue
		}

		summaries = append(summaries, summary)
	}

	return summaries, nil
}
