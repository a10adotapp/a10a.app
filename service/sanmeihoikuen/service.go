package sanmeihoikuen

import (
	"context"
	"log/slog"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/a10adotapp/a10a.app/ent"
	entsanmeihoikuenpost "github.com/a10adotapp/a10a.app/ent/sanmeihoikuenpost"
	"github.com/a10adotapp/a10a.app/lib/line/message"
	"github.com/gocolly/colly"
	"github.com/line/line-bot-sdk-go/v8/linebot/messaging_api"
)

type SanmeiHoikuenService struct {
	entClient *ent.Client
}

func NewSanmeiHoikuenService(entClient *ent.Client) *SanmeiHoikuenService {
	return &SanmeiHoikuenService{
		entClient: entClient,
	}
}

func (s *SanmeiHoikuenService) ExplorePosts(ctx context.Context) error {
	collector := colly.NewCollector()

	collector.OnError(func(r *colly.Response, err error) {
		slog.Error(
			"sanmeihoikuen.SanmeiHoikuenService.ExplorePosts.OnError",
			slog.Any("err", err),
		)
	})

	collector.OnHTML("article", func(article *colly.HTMLElement) {
		article.DOM.Find(".articleBox").Each(func(i int, articleBox *goquery.Selection) {
			href, _ := articleBox.Find("a").Attr("href")

			date := articleBox.Find(".date").Text()

			title := articleBox.Find(".title").Text()

			sanmeiHoikuenPost, err := s.entClient.SanmeiHoikuenPost.Query().
				Where(
					entsanmeihoikuenpost.URL(href),
				).
				First(ctx)
			if ent.MaskNotFound(err) != nil {
				slog.Error(
					"sanmeihoikuen.SanmeiHoikuenService.ExplorePosts.OnHTML",
					slog.Any("err", err),
				)
				return
			}

			if sanmeiHoikuenPost != nil {
				return
			}

			message.Push(
				strings.Join([]string{
					date,
					href,
				}, "\n"),
				message.WithSender(&messaging_api.Sender{
					Name:    "SanmeiHoikuen",
					IconUrl: "https://lh3.googleusercontent.com/fife/ALs6j_GKbLP3NkT1YtWFPvkaYOVWLdFZobJZQe6D_mX8BAf1EuGZYqqSZnOkqfcmh4-M7hluT9y-42fe8vbydDz7E0WUTsFqqCWzUkAM8Je2-VbhyYRcmw4kxapTxcA26QzhO2Be8H3L12yto18pcNOM0RcBiSwWj1oGZ42E2qb6sBNa8aaZOuBshirfXxEldU0KCeEru_A7l1q3t2F2yKmP6rLQGTsajuM0Bh8h9zt7oJUTvrda7ezhdYXoYX2MTyhIDdoDVb71vZEQimSMv59D3kzGqbmSFiFq5Ch64zbzdqMANb6FO2TDWAZxVutNmnp3ZHC-Hf4AnzFLA52Je6Cd1Yx9l6cPJ6xHwwcR__Ud7gzGMsjSefcG-cqsGK-kARXTayUfDz32eAP3joo4Rskf6oshnQrqS3DWa4YGYx0dUNXqeUdye1YjK65cLk3q7QLgPwIVop8uoxd5F3_vuRbJGfGSzpltIV502FRnMgdlcXdtYqoJMTkskutQHGmTCRoCXKab1QFx3E2tERD0BYsrhiFQ0hSvZV-hqXDNGqbF1zADnSAgxvKhzRDOcpmOP-KgTSTCz4kqmzk17ynnsdzMVlw3WXQST5BYT4VkDVBED-TsIIWKrp7Vb7ZOL_jO1SnOiCVGw9NS58jK4xISvnowNgsNJx8U5Tc5-xCXygtuC9PbkoArWc47phcrUTrIk1XdubGSqpapgoWhRxAbwKnMx5xK-1kwatNg6H3v1Uhaha_QR8JIOhHbQWjAKDrUjF2wxNsSqIj_ei0FiwG_gONF5MC1DkzEZtkvTYalzibeZ-JI1ccPuBWCCaTDg6lo-FBK51Sxrg-VDoGJPOEKPls1KxGlyc7SlQ_CHNfI4LiE_9eNx9Yc8TWTsO2fxDM92hG69mR7XkUgQPkwW7jLJySjgTiPaZn2uLAuUVj18hTDfFHkLG3g-edUOyPI8--OFrxjQ_mY1OstRyAkcSE7sBc_8sR8G1RJg9B_vwsBX38Qhpqr47mRez7L5lYWcg53OMhxs7R2jcMnIKZC-4JiUrmmiDSLVfWwCgx-HqX35U8C4iQGylsKZoe_oKyyoBDgmypnDxAV8H6wBm1o35gUP9kZfoX_XJTwAuP_JKu74maGU_ERnZ5o4XktWgIttrfBbV_XeRDjqu-h-qq5K-kSrFA0zoMgtT3Ph4eTi0fXNYaT-UtgJMTgNp-fEZwuFB_mIAJUvgFH_qOevFkZfxnmoNdHCd1WVyYP-uFnzA1lMNgm3W--C0o6wztPSkajYZHw_yP7YtjRE6SS0JBLZOyd2IeFQMKUayAx7P6tTk0oufZ6CpW8eVBkavCJ7_aoqI2yWwKUPjTb79rUhK31muSRcNgtX7ZvJXYgz0BHYApAAsrPzsoyHNgrcXf3m35M8009WqwIgiHSEcPjxYm-CEDuM10IbkRtD9qVKevhyzMpiRK1FLUWnnYQQJCdoHSIcO-GBnKdMD-at1hEr_IQI0vqvzOrNAtjoCdQr28Z_SVu0kE0e-9DdXoM8t8AytUjMWWvBNZiVLDhuK4J4KCVuuFa_Dy-D0n9AMAL-Y3XJyqiuZOrLN6X1XMN1_Z1L34Y7NZ2m1t4zm_WLN3YeC7Z7A3YIXE=w3584-h1996",
				}),
			)

			_, err = s.entClient.SanmeiHoikuenPost.Create().
				SetURL(href).
				SetTitle(title).
				SetDate(date).
				Save(ctx)
			if err != nil {
				slog.Error(
					"sanmeihoikuen.SanmeiHoikuenService.ExplorePosts.OnHTML",
					slog.Any("err", err),
				)
				return
			}
		})
	})

	return collector.Visit("https://sanmeihoikuen.com/post")
}
