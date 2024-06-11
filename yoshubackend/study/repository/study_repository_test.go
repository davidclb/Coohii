package repository_test

import (
	"context"
	"math/rand"
	"testing"
	"time"
	"yoshubackend/db/sqlc"          // Assurez-vous que le chemin est correct
	"yoshubackend/study/repository" // Assurez-vous que le chemin est correct

	"github.com/pashagolub/pgxmock"
	"github.com/stretchr/testify/assert"
)

type Character struct {
    Simplified    string
    Traditional   string
    Pinyin        string
    Meaning       string
    Category      string
    Antonym       string
    Similar       string
    RadicalList   string
}

var characters = []Character{
    {"你", "你", "nǐ", "you", "pronoun", "我", "None", "⼈"},
    {"好", "好", "hǎo", "good", "adjective", "坏", "None", "⼥"},
    {"学", "學", "xué", "study", "verb", "None", "None", "⼦"},
    {"吃", "吃", "chī", "eat", "verb", "None", "None", "⼝"},
    {"大", "大", "dà", "big", "adjective", "小", "None", "⼤"},
    {"小", "小", "xiǎo", "small", "adjective", "大", "None", "⼩"},
    {"我", "我", "wǒ", "I/me", "pronoun", "你", "None", "⼿"},
    {"爱", "愛", "ài", "love", "verb", "恨", "None", "⼼"},
    {"恨", "恨", "hèn", "hate", "verb", "爱", "None", "⼼"},
    {"快", "快", "kuài", "fast", "adjective", "慢", "None", "⾜"},
    {"慢", "慢", "màn", "slow", "adjective", "快", "None", "⾜"},
    {"高", "高", "gāo", "tall", "adjective", "矮", "None", "高"},
    {"矮", "矮", "ǎi", "short", "adjective", "高", "None", "矮"},
    {"多", "多", "duō", "many", "adjective", "少", "None", "多"},
    {"少", "少", "shǎo", "few", "adjective", "多", "None", "少"},
    {"强", "強", "qiáng", "strong", "adjective", "弱", "None", "弓"},
    {"弱", "弱", "ruò", "weak", "adjective", "强", "None", "弓"},
    {"开", "開", "kāi", "open", "verb", "关", "None", "門"},
    {"关", "關", "guān", "close", "verb", "开", "None", "門"},
    {"热", "熱", "rè", "hot", "adjective", "冷", "None", "火"},
    {"冷", "冷", "lěng", "cold", "adjective", "热", "None", "冫"},
    {"新", "新", "xīn", "new", "adjective", "旧", "None", "斤"},
    {"旧", "舊", "jiù", "old", "adjective", "新", "None", "臼"},
    {"美", "美", "měi", "beautiful", "adjective", "丑", "None", "羊"},
    {"丑", "醜", "chǒu", "ugly", "adjective", "美", "None", "鬼"},
    {"安", "安", "ān", "safe", "adjective", "危", "None", "宀"},
    {"危", "危", "wēi", "dangerous", "adjective", "安", "None", "厄"},
    {"早", "早", "zǎo", "early", "adjective", "晚", "None", "日"},
    {"晚", "晚", "wǎn", "late", "adjective", "早", "None", "日"},
    {"清", "清", "qīng", "clear", "adjective", "浊", "None", "氵"},
    {"浊", "濁", "zhuó", "muddy", "adjective", "清", "None", "氵"},
    {"轻", "輕", "qīng", "light", "adjective", "重", "None", "車"},
    {"重", "重", "zhòng", "heavy", "adjective", "轻", "None", "里"},
    {"强", "強", "qiáng", "strong", "adjective", "弱", "None", "弓"},
    {"弱", "弱", "ruò", "weak", "adjective", "强", "None", "弓"},
    {"高", "高", "gāo", "high", "adjective", "低", "None", "高"},
    {"低", "低", "dī", "low", "adjective", "高", "None", "亻"},
    {"明", "明", "míng", "bright", "adjective", "暗", "None", "日"},
    {"暗", "暗", "àn", "dark", "adjective", "明", "None", "日"},
    {"新", "新", "xīn", "new", "adjective", "旧", "None", "斤"},
    {"旧", "舊", "jiù", "old", "adjective", "新", "None", "臼"},
    {"安", "安", "ān", "safe", "adjective", "危", "None", "宀"},
    {"危", "危", "wēi", "dangerous", "adjective", "安", "None", "厄"},
    {"早", "早", "zǎo", "early", "adjective", "晚", "None", "日"},
    {"晚", "晚", "wǎn", "late", "adjective", "早", "None", "日"},
}

func generateRandomRows(numRows int) *pgxmock.Rows {
    rows := pgxmock.NewRows([]string{"id", "carac_simpl", "carac_trad", "pinyin", "meaning", "category", "carac_antonym", "carac_similar", "radical_list"})
    rand.Seed(time.Now().UnixNano())
    for i := 0; i < numRows; i++ {
        char := characters[rand.Intn(len(characters))]
        rows.AddRow(int32(i+1), char.Simplified, char.Traditional, char.Pinyin, char.Meaning, char.Category, char.Antonym, char.Similar, char.RadicalList)
    }
    return rows
}

func TestListCharactersByKeyset(t *testing.T) {
    mockDB, err := pgxmock.NewPool()
	if err != nil {
		t.Fatalf("failed to create pgx mock: %v", err)
	}
	defer mockDB.Close()

    queries := sqlc.New(mockDB)

    repo := repository.NewStudyRepository(&repository.StudyRepo{
		Querie : queries,
	})

    t.Run("Test Limit", func(t *testing.T) {
        rows := pgxmock.NewRows([]string{"id", "carac_simpl", "carac_trad", "pinyin", "meaning", "category", "carac_antonym", "carac_similar", "radical_list"}).
            AddRow(int32(30), "吃", "吃", "chi", "eat", "verb", "None", "None", "⼝")
            

        mockDB.ExpectQuery(`\QSELECT id, carac_simpl, carac_trad, pinyin, meaning, category, carac_antonym, carac_similar, radical_list
            FROM character
            WHERE (id > $1::int OR $1::int IS NULL)
            ORDER BY id
            LIMIT $2::int\E`).
            WithArgs(int32(29),int32(10)).
            WillReturnRows(rows)

        params := sqlc.ListCharactersByKeysetParams{
            ID:    29,
            Pagesize: 10,
        }

        characters, err := repo.ListCharactersByKeyset(context.Background(), params.Pagesize, params.ID )
        assert.NoError(t, err)
        assert.Equal(t, "吃", (*characters)[0].CaracSimpl)
        //assert.Len(t, characters, 1)
    })


        t.Run("Test Nombre Retour Carac", func(t *testing.T) {
        //rows := pgxmock.NewRows([]string{"id", "carac_simpl", "carac_trad", "pinyin", "meaning", "category", "carac_antonym", "carac_similar", "radical_list"}).
          //  AddRow(int32(30), "吃", "吃", "chi", "eat", "verb", "None", "None", "⼝")
        numRows := 30
        rows := generateRandomRows(numRows) 

        mockDB.ExpectQuery(`\QSELECT id, carac_simpl, carac_trad, pinyin, meaning, category, carac_antonym, carac_similar, radical_list
            FROM character
            WHERE (id > $1::int OR $1::int IS NULL)
            ORDER BY id
            LIMIT $2::int\E`).
            WithArgs(int32(29),int32(10)).
            WillReturnRows(rows)

        params := sqlc.ListCharactersByKeysetParams{
            ID:    29,
            Pagesize: 10,
        }

        characters, err := repo.ListCharactersByKeyset(context.Background(), params.Pagesize, params.ID )
        assert.NoError(t, err)
        //assert.Equal(t, "吃", (*characters)[0].CaracSimpl)
        assert.Len(t, characters, 10)
    })
}

