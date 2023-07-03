package model

type GG struct {
	CharacterData struct {
		AchievementPoints int64   `json:"AchievementPoints"`
		AchievementRank   int64   `json:"AchievementRank"`
		CharacterImageURL string  `json:"CharacterImageURL"`
		Class             string  `json:"Class"`
		ClassRank         int64   `json:"ClassRank"`
		EXP               int64   `json:"EXP"`
		EXPPercent        float64 `json:"EXPPercent"`
		GlobalRanking     int64   `json:"GlobalRanking"`
		GraphData         []struct {
			CurrentEXP      int64  `json:"CurrentEXP"`
			DateLabel       string `json:"DateLabel"`
			EXPDifference   int64  `json:"EXPDifference"`
			EXPToNextLevel  int64  `json:"EXPToNextLevel"`
			Level           int64  `json:"Level"`
			TotalOverallEXP int64  `json:"TotalOverallEXP"`
		} `json:"GraphData"`
		LegionCoinsPerDay  int64  `json:"LegionCoinsPerDay"`
		LegionLevel        int64  `json:"LegionLevel"`
		LegionPower        int64  `json:"LegionPower"`
		LegionRank         int64  `json:"LegionRank"`
		Level              int64  `json:"Level"`
		Name               string `json:"Name"`
		Server             string `json:"Server"`
		ServerClassRanking int64  `json:"ServerClassRanking"`
		ServerRank         int64  `json:"ServerRank"`
		ServerSlug         string `json:"ServerSlug"`
	} `json:"CharacterData"`
}
