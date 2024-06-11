import { WordCard } from "@/app/components/word-card";
import SearchBar from "../../searchBar";

export const words = [
  {
    category: "HSK1",
    type: "HSK",
    image: "/hsk1.jpeg",
    carac_simpl: "还",
    carac_trad: "還",
    pinyin: "hai2, huan2",
    meaning: "still; yet; in addition; even; repay; to return",
  },
  {
    category: "HSK1",
    type: "HSK",
    image: "/hsk1.jpeg",
    carac_simpl: "给",
    carac_trad: "給",
    pinyin: "gei3, ji3",
    meaning: "to give; to grant; (passive particle); provide",
  },
  {
    category: "HSK1",
    type: "HSK",
    image: "/hsk1.jpeg",
    carac_simpl: "还",
    carac_trad: "還",
    pinyin: "hai2, huan2",
    meaning: "still; yet; in addition; even; repay; to return",
  },
  {
    category: "HSK1",
    type: "HSK",
    image: "/hsk1.jpeg",
    carac_simpl: "给",
    carac_trad: "給",
    pinyin: "gei3, ji3",
    meaning: "to give; to grant; (passive particle); provide",
  },
  {
    category: "HSK1",
    type: "HSK",
    image: "/hsk1.jpeg",
    carac_simpl: "还",
    carac_trad: "還",
    pinyin: "hai2, huan2",
    meaning: "still; yet; in addition; even; repay; to return",
  },
  {
    category: "HSK1",
    type: "HSK",
    image: "/hsk1.jpeg",
    carac_simpl: "给",
    carac_trad: "給",
    pinyin: "gei3, ji3",
    meaning: "to give; to grant; (passive particle); provide",
  },
];

async function getSentences() {
  const res = await fetch(`http://localhost:3001/sentences`);
  return res.json();
}

export default async function RadicalsPage() {
  const wordList = await getSentences();

  return (
    <>
      {/*  <Header title='Search a hanzi, a word or a sentence ' /> */}
      <SearchBar />
      <div
        className='          
              pt-24
              grid
              grid-cols-1
              sm:grid-cols-2
              md:grid-cols-3
              lg:grid-cols-4
              xl:grid-cols-5
              2xl:grid-cols-6
              gap-8 '
      >
        {/*           {categories.map((categorie: any) => {
            return (
              <CategorieCard
                title={categorie.label}
                categorieAttribute={categorie.image}
              />
            );
          })} */}
      </div>
      {words.map((word: any) => {
        return (
          <WordCard
            category={word.category}
            type={word.type}
            image={word.image}
            carac_simpl={word.carac_simpl}
            carac_trad={word.carac_simpl}
            pinyin={word.pinyin}
            meaning={word.meaning}
          />
        );
      })}
    </>
  );
}
