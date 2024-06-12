package repository_test

import (
	"context"
	"fmt"
	"os"
	"testing"
	"yoshubackend/db/sqlc"          // Assurez-vous que le chemin est correct
	"yoshubackend/study/repository" // Assurez-vous que le chemin est correct

	"github.com/jackc/pgx/v4/pgxpool"
	_ "github.com/lib/pq"
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
}

func setupDatabase() (*pgxpool.Pool, error) {
    ctx := context.Background()
	connStr := "postgresql://root:root@localhost:5432/yoshu?sslmode=disable"

	// Create a new connection pool
	pool, err := pgxpool.Connect(ctx, connStr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		return nil, fmt.Errorf("failed to parse database URL: %v", err)
	}
	//defer pool.Close()
    return pool, nil
}

func seedDatabase(db *pgxpool.Pool) error {
     _, err := db.Exec(context.Background(), `DELETE FROM character`)
    if err != nil {
        return fmt.Errorf("failed to clean the database: %v", err)
    }

    for i, char := range characters {
        _, err := db.Exec(context.Background(), `INSERT INTO character (id, carac_simpl, carac_trad, pinyin, meaning, category, carac_antonym, carac_similar, radical_list) 
                            VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)`,
            i+1, char.Simplified, char.Traditional, char.Pinyin, char.Meaning, char.Category, char.Antonym, char.Similar, char.RadicalList)
        if err != nil {
            return fmt.Errorf("failed to seed the database: %v", err)
        }
    }
    return nil
}

func TestListCharactersByKeyset(t *testing.T) {
    db, err := setupDatabase()
    if err != nil {
        t.Fatalf("failed to setup the database: %v", err)
    }
    defer db.Close()

    err = seedDatabase(db)
    if err != nil {
        t.Fatalf("failed to seed the database: %v", err)
    }

    queries := sqlc.New(db)

    repo := repository.NewStudyRepository(&repository.StudyRepo{
        Querie: queries,
    })

    t.Run("Test Limit", func(t *testing.T) {
        params := sqlc.ListCharactersByKeysetParams{
            ID:       0,
            Pagesize: 15,
        }

        characters, err := repo.ListCharactersByKeyset(context.Background(), params.Pagesize, params.ID)
        assert.NoError(t, err)
        assert.Len(t, characters, 15) // Vérifiez que 10 caractères sont retournés
        assert.Equal(t, "你", characters[0].CaracSimpl) // Vérifiez le premier caractère
    })
}

func TestFilteredCharactersByKeyset(t *testing.T) {
    db, err := setupDatabase()
    if err != nil {
        t.Fatalf("failed to setup the database: %v", err)
    }
    defer db.Close()

    err = seedDatabase(db)
    if err != nil {
        t.Fatalf("failed to seed the database: %v", err)
    }

    queries := sqlc.New(db)

    repo := repository.NewStudyRepository(&repository.StudyRepo{
        Querie: queries,
    })

    t.Run("Test Limit", func(t *testing.T) {
        params := sqlc.FilteredCharactersByKeysetParams{
            ID:       0,
            Pagesize: 15,
            Carac: "你",

        }


        characters, err := repo.FilteredCharactersByKeyset(context.Background(), params.Pagesize, params.ID, params.Carac)
        assert.NoError(t, err)
        assert.Len(t, characters, 1) // Vérifiez que 10 caractères sont retournés
        assert.Equal(t, "你", characters[0].CaracSimpl) // Vérifiez le premier caractère
    })
}